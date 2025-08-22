package repository

import (
	"fmt"
	"os"
	"strings"

	"analysis-question-AI/internal/model"
	"analysis-question-AI/internal/core"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"context"
)

import (
    "encoding/json"
)

type QuestionRepository struct {
	cfg *core.Config
}

func NewQuestionRepository(cfg *core.Config) *QuestionRepository {
	return &QuestionRepository{
		cfg: cfg,
	}
}

func (r *QuestionRepository) Save(questions []model.QuestionTable, formatSave string) error {
	var sb strings.Builder

	for i, q := range questions {
		sb.WriteString(fmt.Sprintf("### Вопрос %d\n", i+1))
		// sb.WriteString(fmt.Sprintf("%s\n\n", q.Question))

		for j, opt := range q.Options {
			sb.WriteString(fmt.Sprintf("%d) %s\n", j+1, opt.Text))
			// if opt.IsCorrect {
			// 	sb.WriteString("✅ Правильный ответ\n")
			// } else {
			// 	sb.WriteString("❌ Неправильный ответ\n")
			// }
			if opt.Explanation != "" {
				sb.WriteString(fmt.Sprintf("_Пояснение_: %s\n", opt.Explanation))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n---\n\n")
	}

	// Записываем в файл
	if err := os.WriteFile(formatSave, []byte(sb.String()), 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	fmt.Println("Всё успешно сохранено в", formatSave)
	return nil
}



func (r *QuestionRepository) SaveJSON(questions []model.QuestionTable, fileName string) error {
    data, err := json.MarshalIndent(questions, "", "  ")
    if err != nil {
        return fmt.Errorf("ошибка маршалинга JSON: %w", err)
    }

    if err := os.WriteFile(fileName, data, 0644); err != nil {
        return fmt.Errorf("ошибка записи в файл: %w", err)
    }

    fmt.Println("JSON успешно сохранён в", fileName)
    return nil
}

func (r *QuestionRepository) SaveToSheets(results []model.QuestionTable, sheetName string) error {
	ctx := context.Background()

	// читаем ключ сервисного аккаунта
	b, err := os.ReadFile(r.cfg.ServiceAccountFile)
	if err != nil {
		return fmt.Errorf("ошибка чтения service account файла: %w", err)
	}

	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		return fmt.Errorf("ошибка парсинга service account: %w", err)
	}

	client := config.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("ошибка создания Sheets API клиента: %w", err)
	}

	// получаем spreadsheet
	ss, err := srv.Spreadsheets.Get(r.cfg.SpreadsheetID).IncludeGridData(false).Do()
	if err != nil {
		return fmt.Errorf("ошибка получения spreadsheet: %w", err)
	}

	// проверяем, есть ли лист
	sheetExists := false
	for _, s := range ss.Sheets {
		if s.Properties.Title == sheetName {
			sheetExists = true
			break
		}
	}

	// если листа нет — создаём
	if !sheetExists {
		_, err := srv.Spreadsheets.BatchUpdate(r.cfg.SpreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				{
					AddSheet: &sheets.AddSheetRequest{
						Properties: &sheets.SheetProperties{
							Title: sheetName,
						},
					},
				},
			},
		}).Do()
		if err != nil {
			return fmt.Errorf("ошибка создания листа %s: %w", sheetName, err)
		}
	}

	// читаем уже существующие данные (теперь 5 колонок)
	readRange := fmt.Sprintf("%s!A:E", sheetName)
	resp, err := srv.Spreadsheets.Values.Get(r.cfg.SpreadsheetID, readRange).Do()
	if err != nil {
		return fmt.Errorf("ошибка чтения данных из %s: %w", sheetName, err)
	}

	existing := map[string]int{} // Question → row index
	for i, row := range resp.Values {
		if len(row) > 0 {
			existing[fmt.Sprintf("%v", row[0])] = i + 1 // строка начинается с 1
		}
	}

	var data []*sheets.ValueRange

	for _, q := range results {
		categories := strings.Join(q.Categories, ", ")

		for i, opt := range q.Options {
			question := ""
			category := ""
			isCorrect := ""
			explanation := ""

			// вопрос и категория только в первой строке блока
			if i == 0 {
				question = q.Question
				category = categories
			}

			// если это правильный ответ
			if opt.IsCorrect {
				isCorrect = "TRUE"
				explanation = opt.Explanation
			}

			row := []interface{}{question, opt.Text, isCorrect, explanation, category}

			if idx, ok := existing[q.Question]; ok {
				// обновляем только если раньше было неверно
				if len(resp.Values) > idx-1 && len(resp.Values[idx-1]) > 2 {
					oldCorrect := fmt.Sprintf("%v", resp.Values[idx-1][2])
					if strings.ToUpper(oldCorrect) != "TRUE" {
						rng := fmt.Sprintf("%s!A%d:E%d", sheetName, idx, idx)
						data = append(data, &sheets.ValueRange{
							Range:  rng,
							Values: [][]interface{}{row},
						})
					}
				}
			} else {
				// добавляем в конец
				data = append(data, &sheets.ValueRange{
					Range:  fmt.Sprintf("%s!A%d", sheetName, len(resp.Values)+1),
					Values: [][]interface{}{row},
				})
				resp.Values = append(resp.Values, row)
			}
		}
	}


	if len(data) == 0 {
		return nil // ничего не обновляли
	}

	// batch update
	_, err = srv.Spreadsheets.Values.BatchUpdate(r.cfg.SpreadsheetID, &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "RAW",
		Data:             data,
	}).Do()
	if err != nil {
		return fmt.Errorf("ошибка записи в Google Sheets: %w", err)
	}

	return nil
}




// func (r *QuestionRepository) SaveToSheets(results []model.QuestionTable, sheetName string) error {
// 	ctx := context.Background()

// 	// читаем ключ сервисного аккаунта
// 	b, err := os.ReadFile(r.cfg.ServiceAccountFile)
// 	if err != nil {
// 		return fmt.Errorf("ошибка чтения service account файла: %w", err)
// 	}

// 	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
// 	if err != nil {
// 		return fmt.Errorf("ошибка парсинга service account: %w", err)
// 	}

// 	client := config.Client(ctx)
// 	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		return fmt.Errorf("ошибка создания Sheets API клиента: %w", err)
// 	}

// 	// проверяем, есть ли лист
// 	ss, err := srv.Spreadsheets.Get(r.cfg.SpreadsheetID).Do()
// 	if err != nil {
// 		return fmt.Errorf("ошибка получения spreadsheet: %w", err)
// 	}

// 	sheetExists := false
// 	for _, s := range ss.Sheets {
// 		if s.Properties.Title == sheetName {
// 			sheetExists = true
// 			break
// 		}
// 	}

// 	// если листа нет — создаём
// 	if !sheetExists {
// 		_, err := srv.Spreadsheets.BatchUpdate(r.cfg.SpreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
// 			Requests: []*sheets.Request{
// 				{
// 					AddSheet: &sheets.AddSheetRequest{
// 						Properties: &sheets.SheetProperties{
// 							Title: sheetName,
// 						},
// 					},
// 				},
// 			},
// 		}).Do()
// 		if err != nil {
// 			return fmt.Errorf("ошибка создания листа %s: %w", sheetName, err)
// 		}
// 	}

// 	// формируем данные для записи
// 	var values [][]interface{}
// 	values = append(values, []interface{}{"Question", "Option", "IsCorrect", "Explanation"})

// 	for _, q := range results {
// 		for _, opt := range q.Options {
// 			values = append(values, []interface{}{
// 				q.Question,
// 				opt.Text,
// 				opt.IsCorrect,
// 				opt.Explanation,
// 			})
// 		}
// 	}

// 	writeRange := fmt.Sprintf("%s!A1", sheetName)

// 	_, err = srv.Spreadsheets.Values.Update(
// 		r.cfg.SpreadsheetID,
// 		writeRange,
// 		&sheets.ValueRange{
// 			Values: values,
// 		},
// 	).ValueInputOption("RAW").Do()

// 	if err != nil {
// 		return fmt.Errorf("ошибка записи в Google Sheets: %w", err)
// 	}

// 	return nil
// }
