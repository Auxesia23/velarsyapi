package dto

type ProjectRequest struct {
	Name            string `form:"name"`
	AboutBrand      string `form:"about_brand"`
	DesignExecution string `form:"design_execution"`
}

type ProjectResponse struct {
	ID        uint   `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type ProjectDetailResponse struct {
	ID              uint            `json:"id"`
	Slug            string          `json:"slug"`
	Name            string          `json:"name"`
	Thumbnail       string          `json:"thumbnail"`
	AboutBrand      string          `json:"about_brand"`
	DesignExecution string          `json:"design_execution"`
	Images          []ImageResponse `json:"images"`
}
