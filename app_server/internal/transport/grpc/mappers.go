package grpc

import (
	"app_server/internal/models"
	client "app_server/pkg/api/swapmeet_grpc"
)

func toGrpcCategory(category *models.Category) *client.Category {
	if category == nil {
		return nil
	}
	return &client.Category{
		Id:       category.ID,
		Name:     category.Name,
		ParentId: category.ParentID,
	}
}

func toGrpcCategories(categories []models.Category) []*client.Category {
	grpcCategories := make([]*client.Category, len(categories))
	for i, catcategory := range categories {
		grpcCategories[i] = toGrpcCategory(&catcategory)
	}
	return grpcCategories
}

////

func toGrpcUserAdvertisements(advertisements []models.UserAdvertisement) []*client.UserAdvertisement {
	grpcAdvertisements := make([]*client.UserAdvertisement, len(advertisements))
	for i, advertisement := range advertisements {
		grpcAdvertisements[i] = toGrpcUserAdvertisement(&advertisement)
	}
	return grpcAdvertisements
}

func toGrpcUserAdvertisement(advertisement *models.UserAdvertisement) *client.UserAdvertisement {
	if advertisement == nil {
		return nil
	}
	return &client.UserAdvertisement{
		Id:           advertisement.ID,
		UserId:       advertisement.UserID,
		UserName:     advertisement.UserName,
		StatusId:     advertisement.StatusID,
		StatusName:   advertisement.StatusName,
		CategoryId:   advertisement.CategoryID,
		CategoryName: advertisement.CategoryName,
		CreatedAt:    advertisement.CreatedAt,
		LastUpd:      advertisement.LastUpd,
		Title:        advertisement.Title,
		Description:  advertisement.Description,
		Price:        advertisement.Price,
		ContactInfo:  advertisement.ContactInfo,
	}
}
