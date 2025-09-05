package sheet

import (
	"fmt"
	"context"
    "os"
	"strings"
	"regexp"
	"strconv"
)

import (
    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

import (
    "analysis-question-AI/internal/analysis"
)


type sheetApi struct {}

func newSheetApi() *sheetApi {
	return &sheetApi{}
}


func (a *analysisApi) getQuestions() ([]QuestionWithRow, error) {
	ctx := context.Background()

    b, err := os.ReadFile(a.cfg.ServiceAccountFile)
    if err != nil { return nil, fmt.Errorf("unable to read service account file: %w", err) }

    config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
    if err != nil { return nil, fmt.Errorf("unable to parse service account key: %w", err) }

    srv, err := sheets.NewService(ctx, option.WithHTTPClient(config.Client(ctx)))
    if err != nil { return nil, fmt.Errorf("unable to retrieve Sheets client: %w", err) }

    var out []QuestionWithRow
    total := 0

    for _, sheet := range a.cfg.Sheets {
        readRange := sheet
        if !strings.Contains(sheet, "!") {
            readRange = fmt.Sprintf("'%s'!A:E", sheet)
        }

        // вытащим стартовую строку из A1-нотации (например 'Expected value'!A161:E → 161)
        startRow := 1
        if m := regexp.MustCompile(`![A-Z]+(\d+)`).FindStringSubmatch(readRange); len(m) == 2 {
            if v, _ := strconv.Atoi(m[1]); v > 0 { startRow = v }
        }

        resp, err := srv.Spreadsheets.Values.Get(a.cfg.SpreadsheetID, readRange).Do()
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

        opts := make([]analysis.AnswerOption, 0, 4)
        for _, r := range rows {
            if len(r) < 2 { continue }
            option := fmt.Sprintf("%v", r[1])
            isCorrect := len(r) > 2 && strings.EqualFold(fmt.Sprintf("%v", r[2]), "TRUE")
            expl := ""
            if len(r) > 3 { expl = fmt.Sprintf("%v", r[3]) }
            opts = append(opts, analysis.AnswerOption{ Text: option, IsCorrect: isCorrect, Explanation: expl })
        }

        // если правильный вариант не в первой строке — вопрос берём из его строки
        for _, r := range rows {
            if len(r) > 2 && strings.EqualFold(fmt.Sprintf("%v", r[2]), "TRUE") {
                question = fmt.Sprintf("%v", r[0])
                break
            }
        }

        q := QuestionWithRow{
            QuestionTable: analysis.QuestionTable{
                Question:   question,
                Options:    opts,
                Categories: categories,
            },
            SheetName: sheetNameOnly(sheet), // см. helper ниже
            StartRow:  startRow,
        }

        out = append(out, q)
        total++
        if a.cfg.Limit > 0 && total >= a.cfg.Limit { break }
    }

    return out, nil
}



func (a *analysisApi) updateRange(a1 string, values [][]interface{}) error {
    ctx := context.Background()

    b, err := os.ReadFile(a.cfg.ServiceAccountFile)
    if err != nil { return fmt.Errorf("unable to read service account file: %w", err) }

    config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
    if err != nil { return fmt.Errorf("unable to parse service account key: %w", err) }

    srv, err := sheets.NewService(ctx, option.WithHTTPClient(config.Client(ctx)))
    if err != nil { return fmt.Errorf("unable to retrieve Sheets client: %w", err) }

    _, err = srv.Spreadsheets.Values.Update(
        a.cfg.SpreadsheetID,
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

