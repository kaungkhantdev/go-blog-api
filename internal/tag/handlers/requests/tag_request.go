package requests

type TagCreateRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	ParentId *int   `json:"parent_id"`
	IconId   int    `json:"icon_id"`
}

type TagUpdateRequest struct {
	TagCreateRequest
}
