# analysis-question-AI

![Intro](./public/intro.png)

## 📌 Описание
`analysis-question-AI` — это микросервис для анализа вопросов и ответов, полученных от AI.  
Он проверяет корректность ответа и возвращает правильный.

---

## 🚀 Требования
1. Установить:
   - [Docker / Docker Compose](https://docs.docker.com/get-started/get-docker/) **или**
   - [Golang](https://go.dev/doc/install)
2. Аккаунт в [Google Cloud](https://console.cloud.google.com/)
3. Создать проект в Google Cloud → [ссылка](https://console.cloud.google.com/welcome/new)
4. Получить [API ключ Gemini](https://aistudio.google.com/apikey) и добавить его в `.env`:
   ```bash
   API_GEMINI_KEY=ваш_api_key
   ```
5. Создать **service account** в Google Cloud → [инструкция](https://console.cloud.google.com/iam-admin/serviceaccounts)
6. Скачать JSON-ключ, сохранить его в корень проекта и указать в настройках.

---

## ⚙️ Настройка

### Вариант 1: `config.json`
Пример: `./analysis-question-AI/config.json`

| Поле                 | Тип       | Пример значения                                  | Обязателен | Описание |
|----------------------|----------|--------------------------------------------------|------------|----------|
| `spreadsheetId`      | string   | `"1B-OgvMNFt8pApbpwbabVG5rHqp29ztLMOK9yqsagd1Q"` | ✅ | ID Google Spreadsheet |
| `sheets`             | string[] | `[ "Copy of Expected value!A165:E" ]`            | ✅ | Список листов и диапазонов |
| `limit`              | int      | `1`                                              | ❌ | Ограничение количества вопросов (`0` = без лимита) |
| `serviceAccountFile` | string   | `"analysis-question-ai-230c2feec375.json"`       | ✅ | JSON-ключ сервисного аккаунта |
| `promptsPath`        | string   | `"./prompts.md"`                                 | ❌ | Путь к файлу с prompts |

---

### Вариант 2: Флаги запуска
| Флаг                  | Тип       | По умолчанию         | Обязателен | Описание |
|-----------------------|-----------|----------------------|------------|----------|
| `-spreadsheetId`      | string    | `""`                 | ✅ | ID Google Spreadsheet |
| `-readRange`          | string    | `""`                 | ❌ | Диапазон (`Лист1!A1:C10`) |
| `-serviceAccountFile` | string    | `""`                 | ✅ | Путь к JSON-ключу |
| `-promptsPath`        | string    | `""`                 | ❌ | Папка/файл с prompts |
| `-config`             | string    | `config.json`        | ❌ | Путь к config.json |
| `-limit`              | int       | `0`                  | ❌ | Лимит вопросов |
| `-sheets`             | string[]  | `[]`                 | ✅ | Листы (`-sheets Лист1 -sheets Лист2`) |
| `-logPath`            | string    | `logs/app.log`       | ❌ | Файл логов |

👉 Флаги приоритетнее `config.json`, можно комбинировать.

---

## ▶️ Запуск

### Вариант 1: Docker/Docker Compose
```bash
git clone https://github.com/Freyzan2006/analysis-question-AI.git
cd analysis-question-AI
echo "API_GEMINI_KEY=<API_KEY>" > .env
docker-compose build
docker-compose run --rm cli --fileInput test.json --fileOutput result.json
```

---

### Вариант 2: Golang
1. Установить Go **>=1.24.4**
2. Клонировать проект:
   ```bash
   git clone https://github.com/Freyzan2006/analysis-question-AI.git
   cd analysis-question-AI
   ```
3. Создать `.env`:
   ```bash
   echo "API_GEMINI_KEY=<API_KEY>" > .env
   ```
4. Собрать:
   ```bash
   go build -o ./build/analysis-question-AI ./cmd/main.go
   ```
5. Запустить:
   ```bash
   ./build/analysis-question-AI
   ```

---

✅ Всё готово к использованию!
