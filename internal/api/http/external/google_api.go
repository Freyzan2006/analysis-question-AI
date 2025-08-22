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

	var questions []model.QuestionTable
	totalCount := 0

	for _, sheet := range g.cfg.Sheets {
		readRange := sheet
		// например: "'Sheet1'!A161:E"
		if !strings.Contains(sheet, "!") {
			readRange = fmt.Sprintf("'%s'!A:E", sheet)
		}

		resp, err := srv.Spreadsheets.Values.Get(g.cfg.SpreadsheetID, readRange).Do()
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve data from sheet %s: %w", sheet, err)
		}
		if len(resp.Values) == 0 {
			continue
		}

		// ⚡ Берём первые 4 строки (1 вопрос + 3 опции)
		rows := resp.Values
		if len(rows) > 4 {
			rows = rows[:4]
		}

		// // первая строка → сам вопрос (+ категория в E)
		// question := fmt.Sprintf("%v", rows[0][0])
		// categories := []string{}
		// if len(rows[0]) > 4 {
		// 	categories = append(categories, fmt.Sprintf("%v", rows[0][4]))
		// }

		// q := model.QuestionTable{
		// 	Question:   question,
		// 	Options:    []model.AnswerOption{},
		// 	Categories: categories,
		// }

		// // остальные строки → варианты ответа
		// for _, row := range rows[1:] {
		// 	if len(row) < 2 {
		// 		continue
		// 	}

		// 	optionText := fmt.Sprintf("%v", row[1])
		// 	isCorrect := false
		// 	if len(row) > 2 && strings.ToUpper(fmt.Sprintf("%v", row[2])) == "TRUE" {
		// 		isCorrect = true
		// 	}

		// 	explanation := ""
		// 	if len(row) > 3 {
		// 		explanation = fmt.Sprintf("%v", row[3])
		// 	}

		// 	q.Options = append(q.Options, model.AnswerOption{
		// 		Text:        optionText,
		// 		IsCorrect:   isCorrect,
		// 		Explanation: explanation,
		// 	})
		// }

		// собираем временный массив опций
		var options []model.AnswerOption
		question := fmt.Sprintf("%v", rows[0][0]) // дефолтный вопрос (первая строка)

		// категории
		categories := []string{}
		if len(rows[0]) > 4 {
			categories = append(categories, fmt.Sprintf("%v", rows[0][4]))
		}

		for _, row := range rows {
			if len(row) < 2 {
				continue
			}

			optionText := fmt.Sprintf("%v", row[1])
			isCorrect := false
			if len(row) > 2 && strings.ToUpper(fmt.Sprintf("%v", row[2])) == "TRUE" {
				isCorrect = true
				// если нашли правильный вариант → берём его строку как "вопрос"
				question = fmt.Sprintf("%v", row[0])
			}

			explanation := ""
			if len(row) > 3 {
				explanation = fmt.Sprintf("%v", row[3])
			}

			options = append(options, model.AnswerOption{
				Text:        optionText,
				IsCorrect:   isCorrect,
				Explanation: explanation,
			})
		}

		q := model.QuestionTable{
			Question:   question,
			Options:    options,
			Categories: categories,
		}



		if g.cfg.Limit > 0 && totalCount >= g.cfg.Limit {
			return questions, nil
		}
		questions = append(questions, q)
		totalCount++
	}

	return questions, nil
}




// func (g *GoogleDocsAPI) GetQuestions() ([]model.QuestionTable, error) {
// 	ctx := context.Background()

// 	fmt.Println(g.cfg)

// 	b, err := os.ReadFile(g.cfg.ServiceAccountFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to read service account file: %w", err)
// 	}

// 	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to parse service account key: %w", err)
// 	}

	


// 	client := config.Client(ctx)

// 	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to retrieve Sheets client: %w", err)
// 	}

// 	var questions []model.QuestionTable
// 	totalCount := 0

	
// 	for _, sheet := range g.cfg.Sheets {
// 		// readRange := fmt.Sprintf("'%s'!A:E", sheet)
// 		readRange := sheet
// 		if !strings.Contains(sheet, "!") {
// 			readRange = fmt.Sprintf("'%s'!A:Z", sheet)
// 		}

		

// 		resp, err := srv.Spreadsheets.Values.Get(g.cfg.SpreadsheetID, readRange).Do()
// 		if err != nil {
// 			return nil, fmt.Errorf("unable to retrieve data from sheet %s: %w", sheet, err)
// 		}

// 		if len(resp.Values) == 0 {
// 			continue
// 		}

// 		for i := 0; i < len(resp.Values); i += 4 {
// 			if g.cfg.Limit > 0 && totalCount >= g.cfg.Limit {
// 				return questions, nil
// 			}

// 			rows := resp.Values[i : i+4]
// 			if len(rows) == 0 || len(rows[0]) < 1 {
// 				continue
// 			}

// 			q := model.QuestionTable{
// 				Question: fmt.Sprintf("%v", rows[0][0]), // A — вопрос
// 				Options:  []model.AnswerOption{},
// 			}

// 			for j := 0; j < 4; j++ {
// 				if j >= len(rows) || len(rows[j]) < 2 {
// 					continue
// 				}

// 				optionText := fmt.Sprintf("%v", rows[j][1]) // B — вариант
// 				isCorrect := false
// 				if len(rows[j]) > 2 && strings.ToUpper(fmt.Sprintf("%v", rows[j][2])) == "TRUE" {
// 					isCorrect = true
// 				}

// 				explanation := ""
// 				if j == 0 && len(rows[j]) > 4 {
// 					explanation = fmt.Sprintf("%v %v", rows[j][3], rows[j][4])
// 				}

// 				q.Options = append(q.Options, model.AnswerOption{
// 					Text:        optionText,
// 					IsCorrect:   isCorrect,
// 					Explanation: explanation,
// 				})
// 			}

// 			questions = append(questions, q)
// 			totalCount++
// 		}

// 	}

	


// 	return questions, nil
// }






// func (g *GoogleDocsAPI) GetQuestions() ([]model.QuestionTable, error) {
// 	ctx := context.Background()

// 	b, err := os.ReadFile(g.cfg.ServiceAccountFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to read service account file: %w", err)
// 	}

// 	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to parse service account key: %w", err)
// 	}

// 	client := config.Client(ctx)

// 	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to retrieve Sheets client: %w", err)
// 	}

// 	spreadsheetId := g.cfg.SpreadsheetID
// 	readRange := g.cfg.ReadRange

// 	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to retrieve data: %w", err)
// 	}

// 	if len(resp.Values) == 0 {
// 		return nil, fmt.Errorf("no data found")
// 	}

// 	var questions []model.QuestionTable

// 	// идём по 4 строки на каждый вопрос
// 	for i := 0; i < len(resp.Values); i += 4 {
// 		rows := resp.Values[i : i+4] // берем блок 4 строки
// 		if len(rows) == 0 || len(rows[0]) < 2 {
// 			continue // пропуск, если данных нет
// 		}

// 		q := model.QuestionTable{
// 			Question: rows[0][0].(string),
// 			Options:  []model.AnswerOption{},
// 		}

// 		for j := 0; j < 4; j++ {
// 			if j >= len(rows) || len(rows[j]) < 2 {
// 				continue
// 			}

// 			optionText := fmt.Sprintf("%v", rows[j][1])
// 			isCorrect := false
// 			if len(rows[j]) > 2 && fmt.Sprintf("%v", rows[j][2]) == "TRUE" {
// 				isCorrect = true
// 			}

// 			explanation := ""
// 			if len(rows[j]) > 3 {
// 				// объединим все оставшиеся колонки в пояснение
// 				explanationParts := []string{}
// 				for k := 3; k < len(rows[j]); k++ {
// 					if val := fmt.Sprintf("%v", rows[j][k]); val != "" {
// 						explanationParts = append(explanationParts, val)
// 					}
// 				}
// 				explanation = strings.Join(explanationParts, " ")
// 			}

// 			q.Options = append(q.Options, model.AnswerOption{
// 				Text:        optionText,
// 				IsCorrect:   isCorrect,
// 				Explanation: explanation,
// 			})
// 		}

// 		questions = append(questions, q)
// 	}

// 	return questions, nil
// }