package cli

import (
	"flag"
)

type flagsConfig struct {
	FILE_INPUT string
	FILE_OUTPUT string
}

type Flags interface {
	GetFlags() *flagsConfig
}

type flags struct {}

func NewFlags() *flags {
	return &flags{}
}

// Получение флагов
func(f *flags) GetFlags() *flagsConfig {
	var (
		fileInput string
		fileOutput string
	)


	flag.StringVar(&fileInput, "fileInput", "./test.json", "Файл с вопросами")
	flag.StringVar(&fileOutput, "fileOutput", "./result.json", "Файл с ответами")

	flag.Parse()


	return &flagsConfig{
		FILE_INPUT: fileInput,
		FILE_OUTPUT: fileOutput,
	}
}