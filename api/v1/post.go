package v1

type CreatePostRequest struct {
    Title     string `json:"title"`
    Content   string `json:"content"`
    Tags      string `json:"tags"`
    ThumbNum  int    `json:"thumb_num"`
    FavourNum int    `json:"favour_num"`
    UUID      uint64 `json:"uuid"`
}
type DeletePostRequest struct {
    UUID   uint64 `json:"uuid"`
    PostID int    `json:"post_id"`
}
type DeletePostResponse struct {
}
type UpdatePostRequest struct {
    Title     string `json:"title"`
    Content   string `json:"content"`
    Tags      string `json:"tags"`
    ThumbNum  int    `json:"thumb_num"`
    FavourNum int    `json:"favour_num"`
    UUID      uint64 `json:"uuid"`
}
type UpdatePostResponse struct {
}
type ListPostRequest struct {
}
type ListPostResponse struct {
}
type GetPostRequest struct {
}
type GetPostResponse struct {
}
