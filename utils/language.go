package utils

import "Rabbit-OJ-Backend/models"

var (
	SupportLanguage []models.SupportLanguage
	CompileObject   map[string]CompileInfo
)

func InitLanguage() {
	SupportLanguage = []models.SupportLanguage{
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

	CompileObject = map[string]CompileInfo{
		"cpp11": {
			Args:        "g++ -std=c++11 /submit/code.cpp -Wall -lm --static -O2 -o /compile/code.o",
			Image:       "gcc:9.2.0",
			CompileTime: 5,
		},
		"cpp14": {
			Args:        "g++ -std=c++14 /submit/code.cpp -Wall -lm --static -O2 -o /compile/code.o",
			Image:       "gcc:9.2.0",
			CompileTime: 5,
		},
		"cpp17": {
			Args:        "g++ -std=c++17 /submit/code.cpp -Wall -lm --static -O2 -o /compile/code.o",
			Image:       "gcc:9.2.0",
			CompileTime: 5,
		},
		"cpp20": {
			Args:        "g++ -std=c++2a /submit/code.cpp -Wall -lm --static -O2 -o /compile/code.o",
			Image:       "gcc:9.2.0",
			CompileTime: 5,
		},
	}
}

type CompileInfo struct {
	Args        string
	Image       string
	CompileTime int // unit: seconds
}
