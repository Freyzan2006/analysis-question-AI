// package core 

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Config struct {
// 	SpreadsheetID      string `json:"spreadsheetId"`
// 	ReadRange          string `json:"readRange"`
// 	ServiceAccountFile string `json:"serviceAccountFile"`
// 	PromptsPath        string `json:"promptsPath"`
// }


// func LoadConfig(path string) (*Config, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка открытия config.json: %w", err)
// 	}
// 	defer file.Close()

// 	var cfg Config
// 	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
// 		return nil, fmt.Errorf("ошибка парсинга config.json: %w", err)
// 	}

// 	return &cfg, nil
// }


package core

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	SpreadsheetID      string `json:"spreadsheetId"`
	ReadRange          string `json:"readRange"`
	ServiceAccountFile string `json:"serviceAccountFile"`
	PromptsPath        string `json:"promptsPath"`
	Limit              int    `json:"limit"`
	Sheets             []string `json:"sheets"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия config.json: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка парсинга config.json: %w", err)
	}

	return &cfg, nil
}


func LoadPrompt(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", fmt.Errorf("не удалось прочитать промпт: %w", err)
    }
    return string(data), nil
}
