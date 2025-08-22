// package external

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
//     "strings"
// )

// import (
// 	"analysis-question-AI/internal/model"
// )


// type GeminiAPI struct {
// 	URL string
// 	KEY string
// 	PromptTemplate string
// }

// func NewGeminiAPI(url string, key string, prompt string) *GeminiAPI {
// 	return &GeminiAPI{
// 		URL: url,
// 		KEY: key,
// 		PromptTemplate: prompt,
// 	}
// }


// func (a *GeminiAPI) GenerateText(q model.QuestionTable) (*model.QuestionTable, error) {
//     API := a.URL + a.KEY

//     // Формируем промпт
//     prompt := fmt.Sprintf(a.PromptTemplate, q.Question, formatOptions(q.Options))

//     input := ContentDTO{
//         Contents: []PartDTO{
//             {
//                 Parts: []MessageDTO{
//                     {Text: prompt},
//                 },
//             },
//         },
//     }

//     body, err := json.Marshal(input)
//     if err != nil {
//         return nil, err
//     }

//     req, err := http.NewRequest("POST", API, bytes.NewReader(body))
//     if err != nil {
//         return nil, err
//     }
//     req.Header.Set("Content-Type", "application/json")

//     client := &http.Client{}
//     res, err := client.Do(req)
//     if err != nil {
//         return nil, err
//     }
//     defer res.Body.Close()

//     respBody, err := io.ReadAll(res.Body)
//     if err != nil {
//         return nil, err
//     }

//     fmt.Println("Raw response:", string(respBody))

//     var geminiResp ResponseDTO
//     if err := json.Unmarshal(respBody, &geminiResp); err != nil {
//         return nil, err
//     }

//     // Ищем JSON в ответе модели
//     var jsonResult string
//     for _, candidate := range geminiResp.Candidates {
//         for _, part := range candidate.Content.Parts {
//             partText := strings.TrimSpace(part.Text)
//             if partText != "" {
//                 jsonResult = partText
//                 break
//             }
//         }
//         if jsonResult != "" {
//             break
//         }
//     }

//     // Если JSON пустой или пустой объект {}, возвращаем исходный вопрос
//     if jsonResult == "" || jsonResult == "{}" {
//         return &q, nil
//     }

//     jsonResult = strings.TrimSpace(jsonResult)
//     jsonResult = strings.TrimPrefix(jsonResult, "```json")
//     jsonResult = strings.TrimPrefix(jsonResult, "```")
//     jsonResult = strings.TrimSuffix(jsonResult, "```")
//     jsonResult = strings.TrimSpace(jsonResult)

//     // Парсим JSON в структуру QuestionTable
//     var updatedQuestion model.QuestionTable
//     if err := json.Unmarshal([]byte(jsonResult), &updatedQuestion); err != nil {
//         return nil, fmt.Errorf("ошибка парсинга JSON от Gemini: %w", err)
//     }

//     return &updatedQuestion, nil
// }


// func formatOptions(options []model.AnswerOption) string {
//     var result string
//     for i, opt := range options {
//         result += fmt.Sprintf("%d. %s (правильный: %v, пояснение: %s)\n", i+1, opt.Text, opt.IsCorrect, opt.Explanation)
//     }
//     return result
// }


package external

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "strings"

    "google.golang.org/genai"

    "analysis-question-AI/internal/model"
)

type GeminiAPI struct {
    PromptTemplate string
    Model          string // например: "gemini-2.5-flash"
    client         *genai.Client
}

func NewGeminiAPI(apiKey string, model string, promptTemplate string) *GeminiAPI {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  apiKey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        log.Fatal("Не удалось создать Gemini client:", err)
    }



    return &GeminiAPI{
        PromptTemplate: promptTemplate,
        Model:          model,
        client:         client,
    }
}

func (a *GeminiAPI) GenerateText(q model.QuestionTable) (*model.QuestionTable, error) {
    ctx := context.Background()

    // Формируем промпт
    prompt := fmt.Sprintf(a.PromptTemplate, q.Question, formatOptions(q.Options))

    // Запрос в Gemini
    resp, err := a.client.Models.GenerateContent(
        ctx,
        a.Model,
        genai.Text(prompt),
        nil,
    )
    if err != nil {
        return nil, fmt.Errorf("ошибка при обращении к Gemini: %w", err)
    }

    // Ответ модели в текстовом виде
    raw := resp.Text()
    fmt.Println("Raw response:", raw)

    // Очистка от ```json ... ```
    jsonResult := strings.TrimSpace(raw)
    jsonResult = strings.TrimPrefix(jsonResult, "```json")
    jsonResult = strings.TrimPrefix(jsonResult, "```")
    jsonResult = strings.TrimSuffix(jsonResult, "```")
    jsonResult = strings.TrimSpace(jsonResult)

    if jsonResult == "" || jsonResult == "{}" {
        return &q, nil
    }

    // Парсим JSON в структуру QuestionTable
    var updated model.QuestionTable
    if err := json.Unmarshal([]byte(jsonResult), &updated); err != nil {
        return nil, fmt.Errorf("ошибка парсинга JSON от Gemini: %w", err)
    }

    return &updated, nil
}

func formatOptions(options []model.AnswerOption) string {
    var result string
    for i, opt := range options {
        result += fmt.Sprintf("%d. %s (correct: %v, explanation: %s)\n",
            i+1, opt.Text, opt.IsCorrect, opt.Explanation)
    }
    return result
}
