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
			BuildArgs:      "g++ -std=c++11 /home/code.cpp -Wall -lm --static -O2 -o /home/code.o",
			BuildTime:      5,
			BuildImage:     "gcc:9.2.0",
			BuildSource:    "/home/code.cpp",
			BuildTarget:    "/home/code.o",
			RunArgs:        "/home/code.o",
			SourceFileName: "code.cpp",
			ExecFilePath:   "/home",
			RunImage:       "alpine_tester:latest",
			RunMaxTimeout:  120,
		},
		"cpp14": {
			BuildArgs:      "g++ -std=c++14 /home/code.cpp -Wall -lm --static -O2 -o /home/code.o",
			BuildTime:      5,
			BuildImage:     "gcc:9.2.0",
			BuildSource:    "/home/code.cpp",
			BuildTarget:    "/home/code.o",
			RunArgs:        "/home/code.o",
			SourceFileName: "code.cpp",
			ExecFilePath:   "/home",
			RunImage:       "alpine_tester:latest",
			RunMaxTimeout:  120,
		},
		"cpp17": {
			BuildArgs:      "g++ -std=c++17 /home/code.cpp -Wall -lm --static -O2 -o /home/code.o",
			BuildTime:      5,
			BuildImage:     "gcc:9.2.0",
			BuildSource:    "/home/code.cpp",
			BuildTarget:    "/home/code.o",
			RunArgs:        "/home/code.o",
			SourceFileName: "code.cpp",
			ExecFilePath:   "/home",
			RunImage:       "alpine_tester:latest",
			RunMaxTimeout:  120,
		},
		"cpp20": {
			BuildArgs:      "g++ -std=c++2a /home/code.cpp -Wall -lm --static -O2 -o /home/code.o",
			BuildTime:      5,
			BuildImage:     "gcc:9.2.0",
			BuildSource:    "/home/code.cpp",
			BuildTarget:    "/home/code.o",
			RunArgs:        "/home/code.o",
			SourceFileName: "code.cpp",
			ExecFilePath:   "/home",
			RunImage:       "alpine_tester:latest",
			RunMaxTimeout:  120,
		},
	}
}

type CompileInfo struct {
	BuildArgs      string
	BuildSource    string
	BuildTarget    string
	BuildImage     string
	BuildTime      int // unit: seconds
	RunArgs        string
	RunImage       string
	RunMaxTimeout  int // unit: seconds
	SourceFileName string
	ExecFilePath   string
}
