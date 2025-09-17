package sheet

import (
	"fmt"
	"strings"

	"analysis-question-AI/internal/entity"
)

type sheetRepository struct {
	api *sheetApi
}

func newSheetRepository(api *sheetApi) *sheetRepository {
	return &sheetRepository{
        api: api,
    }
}

func (s *sheetRepository) findAll() ([]entity.QuestionWithRow, error) {
	return s.api.getQuestions()
}

func (s *sheetRepository) UpdateQuestionBlock(sheet string, startRow int, qt entity.QuestionTable) error {
    // собираем 4 строки A:E
    values := make([][]interface{}, 4)

    // helper: метка TRUE/пусто
    mark := func(b bool) interface{} { if b { return "TRUE" } ; return "" }

    // категории одной строкой
    cats := strings.Join(qt.Categories, ", ")

    for i := 0; i < 4; i++ {
        var text string
        var isCorr bool
        var expl string
        if i < len(qt.Options) {
            text  = qt.Options[i].Text
            isCorr = qt.Options[i].IsCorrect
            if isCorr { expl = qt.Options[i].Explanation }
        }

        // только в первой строке пишем вопрос и категории
        qCell := ""
        catCell := ""
        if i == 0 {
            qCell = qt.Question
            catCell = cats
        }

        values[i] = []interface{}{
            qCell,           // A — вопрос (только в первой строке)
            text,            // B — вариант
            mark(isCorr),    // C — TRUE или пусто
            expl,            // D — пояснение только у правильного
            catCell,         // E — категории (только в первой строке)
        }
    }

    a1 := fmt.Sprintf("'%s'!A%d:E%d", sheet, startRow, startRow+3)
    return s.api.UpdateRange(a1, values)
}
