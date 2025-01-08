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

func marshalAdvertisements(advertisements []models.UserAdvertisement) (string, error) {
	data, err := json.Marshal(advertisements)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func unmarshalAdvertisements(data string) ([]models.UserAdvertisement, error) {
	var advertisements []models.UserAdvertisement
	err := json.Unmarshal([]byte(data), &advertisements)
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func marshalAdvertisement(advertisement models.UserAdvertisement) (string, error) {
	data, err := json.Marshal(advertisement)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
