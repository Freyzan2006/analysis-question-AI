package sheet 

import (
	"analysis-question-AI/internal/module/analysis"
)

type QuestionWithRow struct {
    analysis.QuestionTable
    SheetName string // из какого листа
    StartRow  int    // 1-based: первая строка 4-строчного блока
}