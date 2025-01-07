package repository

import (
	"app_server/internal/models"
	"encoding/json"
)

func unmarshalCategories(data string) ([]models.Category, error) {
	var categories []models.Category
	err := json.Unmarshal([]byte(data), &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func marshalCategories(categories []models.Category) (string, error) {
	data, err := json.Marshal(categories)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func marshalCategory(category models.Category) (string, error) {
	data, err := json.Marshal(category)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func marshalAdvertisements(advertisements []models.PublishedAdvertisement) (string, error) {
	data, err := json.Marshal(advertisements)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func unmarshalAdvertisements(data string) ([]models.PublishedAdvertisement, error) {
	var advertisements []models.PublishedAdvertisement
	err := json.Unmarshal([]byte(data), &advertisements)
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func marshalAdvertisement(advertisement models.PublishedAdvertisement) (string, error) {
	data, err := json.Marshal(advertisement)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
