package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
    "strings"
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


// func (a *GeminiAPI) GenerateText(q model.QuestionTable) (*model.QuestionTable, error) {
//     var API = a.URL + a.KEY


// 	prompt := fmt.Sprintf(a.PromptTemplate, q.Question, formatOptions(q.Options))
    

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

    

//     var result string
//     for _, candidate := range geminiResp.Candidates {
//         for _, part := range candidate.Content.Parts {
            
//             if part.Text != "" {
//                 result += part.Text + "\n"
//             }
//         }
//     }

//     // Можно положить результат в Explanation или отдельное поле
//     q.Options = append(q.Options, model.AnswerOption{
//         Text:        "Анализ от Gemini",
//         Explanation: result,
//     })


//     // if len(generatedJSON) <= 2 { // {} 
//     //     // Нет изменений
//     // }

//     return &q, nil
// }

func (a *GeminiAPI) GenerateText(q model.QuestionTable) (*model.QuestionTable, error) {
    API := a.URL + a.KEY

    // Формируем промпт
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

    fmt.Println("Raw response:", string(respBody))

    var geminiResp ResponseDTO
    if err := json.Unmarshal(respBody, &geminiResp); err != nil {
        return nil, err
    }

    // Ищем JSON в ответе модели
    var jsonResult string
    for _, candidate := range geminiResp.Candidates {
        for _, part := range candidate.Content.Parts {
            partText := strings.TrimSpace(part.Text)
            if partText != "" {
                jsonResult = partText
                break
            }
        }
        if jsonResult != "" {
            break
        }
    }

    // Если JSON пустой или пустой объект {}, возвращаем исходный вопрос
    if jsonResult == "" || jsonResult == "{}" {
        return &q, nil
    }

    jsonResult = strings.TrimSpace(jsonResult)
    jsonResult = strings.TrimPrefix(jsonResult, "```json")
    jsonResult = strings.TrimPrefix(jsonResult, "```")
    jsonResult = strings.TrimSuffix(jsonResult, "```")
    jsonResult = strings.TrimSpace(jsonResult)

    // Парсим JSON в структуру QuestionTable
    var updatedQuestion model.QuestionTable
    if err := json.Unmarshal([]byte(jsonResult), &updatedQuestion); err != nil {
        return nil, fmt.Errorf("ошибка парсинга JSON от Gemini: %w", err)
    }

    return &updatedQuestion, nil
}


func formatOptions(options []model.AnswerOption) string {
    var result string
    for i, opt := range options {
        result += fmt.Sprintf("%d. %s (правильный: %v, пояснение: %s)\n", i+1, opt.Text, opt.IsCorrect, opt.Explanation)
    }
    return result
}
