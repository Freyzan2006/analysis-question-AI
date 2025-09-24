package analysis

import (
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/entity"
	"fmt"
	"strings"
)


func extractJSON(s string) string {
    start := strings.Index(s, "{")
    end := strings.LastIndex(s, "}")
    if start >= 0 && end > start {
        return s[start : end+1]
    }
    return ""
}


func formatOptions(options []entity.AnswerOption) string {
    var result string
    for i, opt := range options {
        result += fmt.Sprintf("%d. %s (correct: %v, explanation: %s)\n",
            i+1, opt.Text, opt.IsCorrect, opt.Explanation)
    }
    return result
}


func getPromptFromFileUtil(path string) string {
    promptTemplate, err := core.LoadPrompt(path)
	if err != nil {
		panic("Не смог прочитать промпт")
	}

    return promptTemplate
}