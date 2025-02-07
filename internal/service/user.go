package service

import (
	"context"
	"fmt"
	"gin-init/api/errcode"
	"gin-init/api/v1"
	"gin-init/internal/model/model_type"
	"gin-init/internal/repository"
	"gin-init/pkg/utils/constants"
	"gin-init/pkg/utils/notify/sms"
	"gin-init/pkg/utils/random"
	"gin-init/pkg/utils/validate"
	"log"
	"time"

	"github.com/silenceper/wechat/v2"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest, resp *v1.LoginResponse) error
	GetProfile(ctx context.Context, req *v1.GetProfileRequest, resp *v1.GetProfileResponse) error
	UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) error
	SendPhoneCode(ctx context.Context, req *v1.SendPhoneCodeRequest) error
	BindWechat(ctx context.Context, req *v1.BindWeChatRequest) error
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo    repository.UserRepository
	settingRepo repository.SettingRepository
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check username
	isExist, err := s.userRepo.ExistUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if !isExist {
		return errcode.ErrEmailAlreadyUse
	}

	isExist, err = s.userRepo.ExistUserByUsername(ctx, req.Email)
	if err != nil {
		return err
	}
	if !isExist {
		return errcode.ErrUsernameAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Logger.Info("Generate password error:", zap.Error(err))
		return err
	}
	// Generate uuid
	uuid, err := s.sid.GenUint64()
	if err != nil {
		log.Println("Generate uuid error:", err)
		return err
	}
	var user = &model_type.User{
		UUID:     uuid,
		Email:    req.Email,
		Password: string(hashedPassword),
		Username: req.Username,
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			log.Println("Create user error:", err)
			return err
		}
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest, resp *v1.LoginResponse) error {
	user, err := s.userRepo.GetByEmailORUsername(ctx, req.Name)
	if err != nil || user == nil {
		return errcode.ErrNoAuth
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		s.logger.Info("")
		return err
	}
	token, err := s.jwt.GenToken(user.UUID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return err
	}
	resp.AccessToken = token
	return nil
}

func (s *userService) GetProfile(ctx context.Context, req *v1.GetProfileRequest, resp *v1.GetProfileResponse) error {
	user, err := s.userRepo.GetByID(ctx, req.UUID)
	if err != nil {
		return err
	}
	resp = &v1.GetProfileResponse{
		Username: user.Username,
		Profile:  user.Profile,
		Email:    user.Email,
		UUID:     user.UUID,
		Avatar:   user.Avatar,
	}
	return nil
}

func (s *userService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) error {
	isExist, err := s.userRepo.ExistUserByUUID(ctx, req.UUID)
	if err != nil {
		return err
	}
	if !isExist {
		return errcode.ErrNotFound
	}
	if err = s.userRepo.Update(ctx, req); err != nil {
		return err
	}

	return nil
}
func (s *userService) SendPhoneCode(ctx context.Context, req *v1.SendPhoneCodeRequest) error {
	// 判断手机号是否正确
	if err := validate.IsPhone(req.Phone); err != nil {
		s.logger.Info("手机号格式错误")
		return errcode.ErrPhoneFormat
	}
	// todo：完成具体业务判断

	// 生成和返回手机验证码
	code := random.Number(6)
	key := constants.RedisKeyPhoneVerifyCode(req.Phone)
	_ = s.userRepo.Set(ctx, key, code, time.Minute*5)
	var cfg sms.SMSConfig
	// todo 去数据库获取短信配置
	if err := s.settingRepo.GetYaml("", "", &cfg); err != nil {
		s.logger.Error("获取短信配置失败")
		return err
	}
	if err := sms.SendPhoneCode(ctx, req.Phone, code, &cfg); err != nil {
		s.logger.Error("发送短信失败")
		return err
	}
	return nil
}

func (s *userService) BindWechat(ctx context.Context, req *v1.BindWeChatRequest) error {
	wApp := wechat.NewWechat()
	var cfg = &offConfig.Config{
		AppID:          "",
		AppSecret:      "",
		Token:          "",
		EncodingAESKey: "",
		Cache:          nil,
	}
	weapp := wApp.GetOfficialAccount(cfg).GetOauth()

	tkn, err := weapp.GetUserAccessToken(req.Request.Code)
	if err != nil {
		return err
	}
	ui, err := weapp.GetUserInfo(tkn.AccessToken, tkn.OpenID, "")
	if err != nil {
		return err
	}
	// todo
	fmt.Println(ui)
	return nil
}
