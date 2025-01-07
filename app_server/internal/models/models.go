package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type Category struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	ParentID string `db:"parent_id"`
	// Subcategories []Category
}

type PublishedAdvertisement struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	UserName     string `db:"user_name"`
	StatusID     string `db:"status_id"`
	StatusName   string `db:"status_name"`
	CategoryID   string `db:"category_id"`
	CategoryName string `db:"category_name"`
	CreatedAt    string `db:"created_at"`
	LastUpd      string `db:"last_upd"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	Price        string `db:"price"`
	ContactInfo  string `db:"contact_info"`
}

type UserAdvertisement struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	UserName     string `db:"user_name"`
	StatusID     string `db:"status_id"`
	StatusName   string `db:"status_name"`
	CategoryID   string `db:"category_id"`
	CategoryName string `db:"category_name"`
	CreatedAt    string `db:"created_at"`
	LastUpd      string `db:"last_upd"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	Price        string `db:"price"`
	ContactInfo  string `db:"contact_info"`
}
