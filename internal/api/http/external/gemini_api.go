package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

import (
	"analysis-question-AI/internal/model"
)


type GeminiAPI struct {
	URL string
	KEY string
	PromptTemplate string
}

func NewGeminiAPI(url string, key string, prompt string) *GeminiAPI {
	return &GeminiAPI{
		URL: url,
		KEY: key,
		PromptTemplate: prompt,
	}
}


func (a *GeminiAPI) GenerateText(q model.QuestionTable) (*model.QuestionTable, error) {
    var API = a.URL + a.KEY


	prompt := fmt.Sprintf(a.PromptTemplate, q.Question, formatOptions(q.Options))
    

    input := ContentDTO{
        Contents: []PartDTO{
            {
                Parts: []MessageDTO{
                    {Text: prompt},
                },
            },
        },
    }

    body, err := json.Marshal(input)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", API, bytes.NewReader(body))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    respBody, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    var geminiResp ResponseDTO
    if err := json.Unmarshal(respBody, &geminiResp); err != nil {
        return nil, err
    }

    var result string
    for _, candidate := range geminiResp.Candidates {
        for _, part := range candidate.Content.Parts {
            result += part.Text + "\n"
        }
    }

    // Можно положить результат в Explanation или отдельное поле
    q.Options = append(q.Options, model.AnswerOption{
        Text:        "Анализ от Gemini",
        Explanation: result,
    })

    return &q, nil
}

func formatOptions(options []model.AnswerOption) string {
    var result string
    for i, opt := range options {
        result += fmt.Sprintf("%d. %s (правильный: %v, пояснение: %s)\n", i+1, opt.Text, opt.IsCorrect, opt.Explanation)
    }
    return result
}
