package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}
type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash,omitempty"`
}
