package utils

import "Rabbit-OJ-Backend/models"

func SupportLanguage() []models.SupportLanguage {
	return []models.SupportLanguage{
		{
			Name:  "C++/11",
			Value: "cpp11",
		},
		{
			Name:  "C++/14",
			Value: "cpp14",
		},
		{
			Name:  "C++/17",
			Value: "cpp17",
		},
		{
			Name:  "C++/2a",
			Value: "cpp20",
		},
	}
}
