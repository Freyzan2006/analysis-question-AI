package external

import (
	"fmt"
	"context"
    "os"
	"strings"
)

import (
	"analysis-question-AI/internal/model"
	"analysis-question-AI/internal/core"
)

import (
    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

type GoogleDocsAPI struct {
	cfg *core.Config
}

func NewGoogleDocsAPI(cfg *core.Config) *GoogleDocsAPI {
	return &GoogleDocsAPI{
		cfg: cfg,
	}
}


func (g *GoogleDocsAPI) GetQuestions() ([]model.QuestionTable, error) {
	ctx := context.Background()

	b, err := os.ReadFile(g.cfg.ServiceAccountFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read service account file: %w", err)
	}

	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse service account key: %w", err)
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %w", err)
	}

	spreadsheetId := g.cfg.SpreadsheetID
	readRange := g.cfg.ReadRange

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data: %w", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	var questions []model.QuestionTable

	// идём по 4 строки на каждый вопрос
	for i := 0; i < len(resp.Values); i += 4 {
		rows := resp.Values[i : i+4] // берем блок 4 строки
		if len(rows) == 0 || len(rows[0]) < 2 {
			continue // пропуск, если данных нет
		}

		q := model.QuestionTable{
			Question: rows[0][0].(string),
			Options:  []model.AnswerOption{},
		}

		for j := 0; j < 4; j++ {
			if j >= len(rows) || len(rows[j]) < 2 {
				continue
			}

			optionText := fmt.Sprintf("%v", rows[j][1])
			isCorrect := false
			if len(rows[j]) > 2 && fmt.Sprintf("%v", rows[j][2]) == "TRUE" {
				isCorrect = true
			}

			explanation := ""
			if len(rows[j]) > 3 {
				// объединим все оставшиеся колонки в пояснение
				explanationParts := []string{}
				for k := 3; k < len(rows[j]); k++ {
					if val := fmt.Sprintf("%v", rows[j][k]); val != "" {
						explanationParts = append(explanationParts, val)
					}
				}
				explanation = strings.Join(explanationParts, " ")
			}

			q.Options = append(q.Options, model.AnswerOption{
				Text:        optionText,
				IsCorrect:   isCorrect,
				Explanation: explanation,
			})
		}

		questions = append(questions, q)
	}

	return questions, nil
}
