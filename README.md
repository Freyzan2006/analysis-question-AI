# analysis-question-AI

# Описание
-----------------------------------------
analysis-question-AI - это microservice, который предназначен для анализа вопросов и ответов от ai.
Анализирует корректность ответа и выдает правильный ответ.
-----------------------------------------

# Требования к запуску
-----------------------------------------
* Нужно установить docker и docker-compose или golang
* Нужно иметь свой аккаунт в google cloud
* Нужно иметь свой api key в google cloud
*.env*
```bash 
API_GEMINI_KEY= # вставить свой api key
API_GEMINI_URL= # вставить свой url
```
Это необходимо для работы gemini api
-----------------------------------------

# Запуск
-----------------------------------------

## Через Docker/Docker-compose
1. Установить docker и docker-compose
2. git clone https://github.com/Freyzan2006/analysis-question-AI.git 
3. cd analysis-question-AI
4. Нужно создать .env [Ссылка для получение данных для .env](https://aistudio.google.com/apikey)
```bash
echo "API_GEMINI_KEY=<КЛЮЧ_к_API> API_GEMINI_URL=<URL_к_API>" > .env
```
5. 
```bash
docker-compose build
docker-compose run --rm cli --fileInput test.json --fileOutput result.json
```
* Здесь можно получить ключ - https://aistudio.google.com/apikey
* Здесь можно получить url - https://aistudio.google.com/apikey

## Через golang
1. Установить golang 1.24.4 или выше
2. Скачивание репозитория:
```bash
git clone https://github.com/Freyzan2006/analysis-question-AI.git
```

3. Переход в директиву:
```bash
cd analysis-question-AI
```

4. Создайте .env [Ссылка для получение данных для .env](https://aistudio.google.com/apikey)
```bash
echo "API_GEMINI_KEY=<КЛЮЧ_к_API> API_GEMINI_URL=<URL_к_API>" > .env
```

5. Сборка проект под вашу OS:
```bash
go build -o ./build/analysis-question-AI ./cmd/main.go
```

6. Запуск:
```bash
./build/analysis-question-AI
```

Всё готово !
-----------------------------------------