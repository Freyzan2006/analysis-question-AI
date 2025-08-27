package core

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(filePath string) *Logger {
	logger := logrus.New()

	// 🔹 Создаём директорию для логов, если её нет
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if mkErr := os.MkdirAll(dir, 0755); mkErr != nil {
			logger.Warnf("не удалось создать директорию логов: %v", mkErr)
		}
	}

	// 🔹 Открываем файл для записи логов
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file) // пишем и в файл, и в консоль
		logger.SetOutput(mw)
	} else {
		logger.Warnf("не удалось открыть файл логов, пишем только в консоль: %v", err)
		logger.SetOutput(os.Stdout)
	}

	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &Logger{logger}
}
