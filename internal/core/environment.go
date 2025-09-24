package core 

import (
	"fmt"
	"log"
	"os"
)

import (
	"github.com/joho/godotenv"
)

type Environment struct {}

func NewEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	return &Environment{}
}

func (env *Environment) Get(key string) string {
	var value = os.Getenv(key)
	
	if value == "" {
		panic(fmt.Sprintf("Отсутствует обязательная переменная окружения: %s", key))
	}
	return value
}