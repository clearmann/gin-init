package v1

import model "gin-init/internal/model/model_type"

type CreatePostRequest struct {
    Title     string `json:"title"`
    Content   string `json:"content"`
    Tags      string `json:"tags"`
    ThumbNum  int    `json:"thumb_num"`
    FavourNum int    `json:"favour_num"`
    UUID      uint64 `json:"uuid"`
}
type DeletePostRequest struct {
    UUID   uint64 `json:"uuid,omitempty"`
    PostID int    `json:"post_id"`
}
type UpdatePostRequest struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    Tags    string `json:"tags"`
    UUID    uint64 `json:"uuid"`
}
type ListPostRequest struct {
    QueryRequest
}
type ListPostResponse struct {
    Total  int64         `json:"total"`
    Offset int           `json:"offset"`
    Limit  int           `json:"limit"`
    Data   []*model.Post `json:"data"`
}
type GetPostRequest struct {
}
type GetPostResponse struct {
}
