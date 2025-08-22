// package cli

// import (
// 	"analysis-question-AI/internal/core"
// 	"flag"
// 	// "fmt"
// 	"log"
// )

// type flagsConfig struct {
// 	FileOutput             string
// 	GoogleSpreadsheetID    string
// 	GoogleReadRange        string
// 	GoogleServiceAccountFile string
// 	GooglePromptsPath      string
// }

// type Flags interface {
// 	GetFlags() (*flagsConfig)
// }

// type flags struct {}

// func NewFlags() *flags {
// 	return &flags{}
// }

// func (f *flags) GetFlags() (*flagsConfig,) {
// 	var (
// 		fileOutput         string
// 		spreadsheetID      string
// 		readRange          string
// 		serviceAccountFile string
// 		promptsPath        string
// 		configPath         string
// 	)

// 	// Определяем флаги
// 	flag.StringVar(&fileOutput, "fileOutput", "./result.json", "Файл с ответами")
// 	flag.StringVar(&spreadsheetID, "spreadsheetId", "", "Google Spreadsheet ID")
// 	flag.StringVar(&readRange, "readRange", "", "Диапазон ячеек (например: Лист1!A1:C10)")
// 	flag.StringVar(&serviceAccountFile, "serviceAccountFile", "", "Путь к JSON ключу сервисного аккаунта")
// 	flag.StringVar(&promptsPath, "promptsPath", "", "Путь к папке с prompts")
// 	flag.StringVar(&configPath, "config", "config.json", "Путь к config.json")

// 	flag.Parse()

// 	// Загружаем конфиг
// 	cfg, err := core.LoadConfig(configPath)
// 	if err != nil {
// 		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
// 	}

// 	// Подменяем значениями из флагов, если они переданы
// 	if spreadsheetID != "" {
// 		cfg.SpreadsheetID = spreadsheetID
// 	}
// 	if readRange != "" {
// 		cfg.ReadRange = readRange
// 	}
// 	if serviceAccountFile != "" {
// 		cfg.ServiceAccountFile = serviceAccountFile
// 	}
// 	if promptsPath != "" {
// 		cfg.PromptsPath = promptsPath
// 	}

// 	// Проверка обязательных параметров
// 	if cfg.SpreadsheetID == "" || cfg.ReadRange == "" || cfg.ServiceAccountFile == "" {
// 		log.Fatal("Обязательные параметры отсутствуют: spreadsheetId, readRange, serviceAccountFile")
// 	}

	

// 	return &flagsConfig{
// 		FileOutput:             fileOutput,
// 		GoogleSpreadsheetID:    cfg.SpreadsheetID,
// 		GoogleReadRange:        cfg.ReadRange,
// 		GoogleServiceAccountFile: cfg.ServiceAccountFile,
// 		GooglePromptsPath:      cfg.PromptsPath,
// 	}
// }



package cli

import (
	"analysis-question-AI/internal/core"
	"flag"
	"log"
	"analysis-question-AI/internal/core/types"
)

type flagsConfig struct {
	FileOutput              string
	GoogleSpreadsheetID     string
	GoogleReadRange         string
	GoogleServiceAccountFile string
	GooglePromptsPath       string
	GoogleDocsLimit         int
	GoogleDocsSheets        []string
}

type Flags interface {
	GetFlags() *flagsConfig
}

type flags struct{}

func NewFlags() *flags {
	return &flags{}
}

func (f *flags) GetFlags() *flagsConfig {
	// Значения по умолчанию
	var (
		fileOutput         string
		spreadsheetID      string
		readRange          string
		serviceAccountFile string
		promptsPath        string
		configPath         string
		limit              int
		sheets             types.StringSliceFlag
	)

	// Определяем флаги
	flag.StringVar(&fileOutput, "fileOutput", "./result.json", "Файл с ответами")
	flag.StringVar(&spreadsheetID, "spreadsheetId", "", "Google Spreadsheet ID")
	flag.StringVar(&readRange, "readRange", "", "Диапазон ячеек (например: 'Лист1'!A1:C10)")
	flag.StringVar(&serviceAccountFile, "serviceAccountFile", "", "Путь к JSON ключу сервисного аккаунта")
	flag.StringVar(&promptsPath, "promptsPath", "", "Путь к папке с prompts")
	flag.StringVar(&configPath, "config", "config.json", "Путь к config.json")
	flag.IntVar(&limit, "limit", 0, "Максимальное количество вопросов (0 = без лимита)")
	flag.Var(&sheets, "sheets", "Список листов (можно указывать несколько раз)")
	flag.Parse()

	// Загружаем конфиг
	cfg, err := core.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Приоритет флагов над конфигом
	if spreadsheetID != "" {
		cfg.SpreadsheetID = spreadsheetID
	}
	if readRange != "" {
		cfg.ReadRange = readRange
	}
	if serviceAccountFile != "" {
		cfg.ServiceAccountFile = serviceAccountFile
	}
	if promptsPath != "" {
		cfg.PromptsPath = promptsPath
	}

	if limit != 0 {
		cfg.Limit = limit
	}

	if len(sheets) > 0 {
		cfg.Sheets = sheets
	}

	// Проверка обязательных параметров
	if cfg.SpreadsheetID == "" || len(cfg.Sheets) == 0 || cfg.ServiceAccountFile == "" {
		log.Fatal("Обязательные параметры отсутствуют: spreadsheetId, Sheets, serviceAccountFile")
	}

	return &flagsConfig{
		FileOutput:              fileOutput,
		GoogleSpreadsheetID:     cfg.SpreadsheetID,
		GoogleReadRange:         cfg.ReadRange,
		GoogleServiceAccountFile: cfg.ServiceAccountFile,
		GooglePromptsPath:       cfg.PromptsPath,
		GoogleDocsLimit:         cfg.Limit,
		GoogleDocsSheets:        cfg.Sheets,
	}
}
