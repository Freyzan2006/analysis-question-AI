package analysis

import (
    "context"
    "encoding/json"
    "fmt"
    "strings"

    "google.golang.org/genai"

    "analysis-question-AI/internal/core"
    "analysis-question-AI/internal/entity"
)

type analysisApi struct {
	PromptTemplate string
    Model          string // например: "gemini-2.5-flash"
    client         *genai.Client
    log            *core.Logger
}


func newAnalysisApi(log *core.Logger, apiKey string, model string, promptTemplate string) *analysisApi {
	ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  apiKey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        log.Fatal("Не удалось создать Gemini client:", err)
    }


	return &analysisApi{
		PromptTemplate: promptTemplate,
        Model:          model,
        client:         client,
        log:            log,
	}
}


func (a *analysisApi) analyzeQuestionsApi(q entity.QuestionTable) (*entity.QuestionTable, bool, error) {
    ctx := context.Background()


    prompt := fmt.Sprintf(a.PromptTemplate, q.Question, formatOptions(q.Options))

    resp, err := a.client.Models.GenerateContent(
        ctx,
        a.Model,
        genai.Text(prompt),
        nil,
    )
    if err != nil {
        return nil, false, fmt.Errorf("ошибка при обращении к Gemini: %w", err)
    }

    raw := strings.TrimSpace(resp.Text())
    clean := extractJSON(raw)

    // Если Gemini вернул пустой JSON или ничего
    if clean == "" || clean == "{}" {
        return &q, false, nil
    }

    // Парсим JSON
    var updated entity.QuestionTable
    if err := json.Unmarshal([]byte(clean), &updated); err != nil {
        return nil, false, fmt.Errorf("ошибка парсинга JSON от Gemini: %w\nraw response: %s", err, raw)
    }


    // 🔹 если Gemini не вернул categories — берём старые
    if len(updated.Categories) == 0 {
        updated.Categories = q.Categories
    }

    return &updated, true, nil
}





