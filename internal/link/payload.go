package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}
