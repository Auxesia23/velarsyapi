package dto

type ServiceRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type ServiceResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
