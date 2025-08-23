package external

import (
	"fmt"
	"context"
    "os"
	"strings"
	"regexp"
	"strconv"
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

// external/google_docs_api.go
func (g *GoogleDocsAPI) GetQuestions() ([]model.QuestionWithRow, error) {
    ctx := context.Background()

    b, err := os.ReadFile(g.cfg.ServiceAccountFile)
    if err != nil { return nil, fmt.Errorf("unable to read service account file: %w", err) }

    config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
    if err != nil { return nil, fmt.Errorf("unable to parse service account key: %w", err) }

    srv, err := sheets.NewService(ctx, option.WithHTTPClient(config.Client(ctx)))
    if err != nil { return nil, fmt.Errorf("unable to retrieve Sheets client: %w", err) }

    var out []model.QuestionWithRow
    total := 0

    for _, sheet := range g.cfg.Sheets {
        readRange := sheet
        if !strings.Contains(sheet, "!") {
            readRange = fmt.Sprintf("'%s'!A:E", sheet)
        }

        // вытащим стартовую строку из A1-нотации (например 'Expected value'!A161:E → 161)
        startRow := 1
        if m := regexp.MustCompile(`![A-Z]+(\d+)`).FindStringSubmatch(readRange); len(m) == 2 {
            if v, _ := strconv.Atoi(m[1]); v > 0 { startRow = v }
        }

        resp, err := srv.Spreadsheets.Values.Get(g.cfg.SpreadsheetID, readRange).Do()
        if err != nil { return nil, fmt.Errorf("unable to retrieve data from %s: %w", sheet, err) }
        if len(resp.Values) == 0 { continue }

        // берём только первый блок 4 строк (под твой use-case)
        rows := resp.Values
        if len(rows) > 4 { rows = rows[:4] }

        question := fmt.Sprintf("%v", rows[0][0])
        categories := []string{}
        if len(rows[0]) > 4 {
            categories = append(categories, fmt.Sprintf("%v", rows[0][4]))
        }

        opts := make([]model.AnswerOption, 0, 4)
        for _, r := range rows {
            if len(r) < 2 { continue }
            option := fmt.Sprintf("%v", r[1])
            isCorrect := len(r) > 2 && strings.EqualFold(fmt.Sprintf("%v", r[2]), "TRUE")
            expl := ""
            if len(r) > 3 { expl = fmt.Sprintf("%v", r[3]) }
            opts = append(opts, model.AnswerOption{ Text: option, IsCorrect: isCorrect, Explanation: expl })
        }

        // если правильный вариант не в первой строке — вопрос берём из его строки
        for _, r := range rows {
            if len(r) > 2 && strings.EqualFold(fmt.Sprintf("%v", r[2]), "TRUE") {
                question = fmt.Sprintf("%v", r[0])
                break
            }
        }

        q := model.QuestionWithRow{
            QuestionTable: model.QuestionTable{
                Question:   question,
                Options:    opts,
                Categories: categories,
            },
            SheetName: sheetNameOnly(sheet), // см. helper ниже
            StartRow:  startRow,
        }

        out = append(out, q)
        total++
        if g.cfg.Limit > 0 && total >= g.cfg.Limit { break }
    }

    return out, nil
}

func sheetNameOnly(a1 string) string {
    // "'Expected value'!A161:E" → Expected value; "Sheet1!A:E" → Sheet1; "Sheet1" → Sheet1
    s := a1
    if i := strings.Index(s, "!"); i >= 0 { s = s[:i] }
    s = strings.Trim(s, "'")
    return s
}



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

// 	var questions []model.QuestionTable
// 	totalCount := 0

// 	for _, sheet := range g.cfg.Sheets {
// 		readRange := sheet
// 		// например: "'Sheet1'!A161:E"
// 		if !strings.Contains(sheet, "!") {
// 			readRange = fmt.Sprintf("'%s'!A:E", sheet)
// 		}

// 		resp, err := srv.Spreadsheets.Values.Get(g.cfg.SpreadsheetID, readRange).Do()
// 		if err != nil {
// 			return nil, fmt.Errorf("unable to retrieve data from sheet %s: %w", sheet, err)
// 		}
// 		if len(resp.Values) == 0 {
// 			continue
// 		}

// 		// ⚡ Берём первые 4 строки (1 вопрос + 3 опции)
// 		rows := resp.Values
// 		if len(rows) > 4 {
// 			rows = rows[:4]
// 		}


// 		// собираем временный массив опций
// 		var options []model.AnswerOption
// 		question := fmt.Sprintf("%v", rows[0][0]) // дефолтный вопрос (первая строка)

// 		// категории
// 		categories := []string{}
// 		if len(rows[0]) > 4 {
// 			categories = append(categories, fmt.Sprintf("%v", rows[0][4]))
// 		}

// 		for _, row := range rows {
// 			if len(row) < 2 {
// 				continue
// 			}

// 			optionText := fmt.Sprintf("%v", row[1])
// 			isCorrect := false
// 			if len(row) > 2 && strings.ToUpper(fmt.Sprintf("%v", row[2])) == "TRUE" {
// 				isCorrect = true
// 				// если нашли правильный вариант → берём его строку как "вопрос"
// 				question = fmt.Sprintf("%v", row[0])
// 			}

// 			explanation := ""
// 			if len(row) > 3 {
// 				explanation = fmt.Sprintf("%v", row[3])
// 			}

// 			options = append(options, model.AnswerOption{
// 				Text:        optionText,
// 				IsCorrect:   isCorrect,
// 				Explanation: explanation,
// 			})
// 		}

// 		q := model.QuestionTable{
// 			Question:   question,
// 			Options:    options,
// 			Categories: categories,
// 		}



// 		if g.cfg.Limit > 0 && totalCount >= g.cfg.Limit {
// 			return questions, nil
// 		}
// 		questions = append(questions, q)
// 		totalCount++
// 	}

// 	return questions, nil
// }



// external/google_docs_api.go
func (g *GoogleDocsAPI) UpdateRange(a1 string, values [][]interface{}) error {
    ctx := context.Background()

    b, err := os.ReadFile(g.cfg.ServiceAccountFile)
    if err != nil { return fmt.Errorf("unable to read service account file: %w", err) }

    config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
    if err != nil { return fmt.Errorf("unable to parse service account key: %w", err) }

    srv, err := sheets.NewService(ctx, option.WithHTTPClient(config.Client(ctx)))
    if err != nil { return fmt.Errorf("unable to retrieve Sheets client: %w", err) }

    _, err = srv.Spreadsheets.Values.Update(
        g.cfg.SpreadsheetID,
        a1,
        &sheets.ValueRange{ Values: values },
    ).
        ValueInputOption("USER_ENTERED").
        Do()
    if err != nil {
        return fmt.Errorf("unable to update data in range %s: %w", a1, err)
    }
    return nil
}

