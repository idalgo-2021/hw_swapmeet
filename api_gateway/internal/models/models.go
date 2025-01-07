package models

type GenerateTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type CreateCategoryRequest struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type CreateAdvertisementRequest struct {
	CategoryId  string `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ContactInfo string `json:"contact_info"`
}

type UpdateAdvertisementRequest struct {
	AdvertisementId string `json:"advertisement_id,"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Price           string `json:"price"`
	ContactInfo     string `json:"contact_info"`
}
