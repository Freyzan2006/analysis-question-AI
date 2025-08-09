package test_features

import (
    "context"
    "fmt"
    "log"
    "os"

    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

func Test() {
    ctx := context.Background()

    // Подключаемся по ключу сервисного аккаунта
    b, err := os.ReadFile("analysis-question-ai-230c2feec375.json")
    if err != nil {
        log.Fatalf("Unable to read service account file: %v", err)
    }

    config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
    if err != nil {
        log.Fatalf("Unable to parse service account key: %v", err)
    }

    client := config.Client(ctx)

    srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    // ID документа (в ссылке Google Sheets)
    spreadsheetId := "1R7xfAJ6RyytICa5VFYF7nhVwilWlfW2VvAvRUruJIus"
    readRange := "Лист1!A1:C10" // диапазон

    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data: %v", err)
    }

    if len(resp.Values) == 0 {
        fmt.Println("No data found.")
    } else {
        for _, row := range resp.Values {
            fmt.Println(row)
        }
    }
}
