package service 

import (
	"log"
	"strings"
	"fmt"
)

import (
	"analysis-question-AI/internal/model"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/api/http/external"
)

type GoogleDocsService struct {
	api  external.GoogleDocsAPI
	repo repository.QuestionRepository
}

func NewGoogleDocsService(api *external.GoogleDocsAPI, repo *repository.QuestionRepository) *GoogleDocsService {
	return &GoogleDocsService{
		repo: *repo,
		api:  *api,
	}
}

func (g *GoogleDocsService) GetQuestions() ([]model.QuestionWithRow, error) {
	
	api, err := g.api.GetQuestions()
	if err != nil {
		log.Fatal(err)
	}

	
	
	return api, nil
}

func (s *GoogleDocsService) UpdateQuestionRow(sheetName string, rowIndex int, q model.QuestionTable) error {
	values := [][]interface{}{
		{
			q.Question,
			q.Options[0].Text,
			q.Options[0].IsCorrect,
			q.Options[0].Explanation,
			q.Options[1].Text,
			q.Options[1].IsCorrect,
			q.Options[1].Explanation,
			q.Options[2].Text,
			q.Options[2].IsCorrect,
			q.Options[2].Explanation,
			q.Options[3].Text,
			q.Options[3].IsCorrect,
			q.Options[3].Explanation,
			strings.Join(q.Categories, ", "),
		},
	}

	writeRange := fmt.Sprintf("%s!A%d:N%d", sheetName, rowIndex, rowIndex)
	err := s.api.UpdateRange(writeRange, values)

	return err
}


// internal/service/google_docs_service.go
func (s *GoogleDocsService) UpdateQuestionBlock(sheet string, startRow int, qt model.QuestionTable) error {
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
