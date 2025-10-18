package dto

type WorkResponse struct {
	ID    uint   `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type WorkDetailResponse struct {
	ID       uint              `json:"id"`
	Slug     string            `json:"slug"`
	Title    string            `json:"title"`
	Image    string            `json:"image"`
	Projects []ProjectResponse `json:"projects"`
}
