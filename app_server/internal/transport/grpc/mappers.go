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

func toGrpcPublishedAdvertisement(advertisement *models.PublishedAdvertisement) *client.PublishedAdvertisement {
	if advertisement == nil {
		return nil
	}
	return &client.PublishedAdvertisement{
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

func toGrpcPublishedAdvertisements(advertisements []models.PublishedAdvertisement) []*client.PublishedAdvertisement {
	grpcAdvertisements := make([]*client.PublishedAdvertisement, len(advertisements))
	for i, advertisement := range advertisements {
		grpcAdvertisements[i] = toGrpcPublishedAdvertisement(&advertisement)
	}
	return grpcAdvertisements
}

func modelToGrpcUserAdvertisements(advertisements []models.UserAdvertisement) []*client.UserAdvertisement {
	var grpcAdvertisements []*client.UserAdvertisement
	for _, advertisement := range advertisements {
		grpcAdvertisements = append(grpcAdvertisements, &client.UserAdvertisement{
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
		})
	}
	return grpcAdvertisements
}

func modelToGrpcUserAdvertisement(advertisement *models.UserAdvertisement) *client.UserAdvertisement {
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
