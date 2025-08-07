package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

import (
	"analysis-question-AI/internal/model"
)


type GeminiAPI struct {
	URL string
	KEY string
}

func NewGeminiAPI(url string, key string) *GeminiAPI {
	return &GeminiAPI{
		URL: url,
		KEY: key,
	}
}

func(a *GeminiAPI) GenerateText(text string) (*model.Question, error) {
	var API = a.URL + a.KEY 

	input := ContentDTO{
		Contents: []PartDTO{
			{
				Parts: []MessageDTO{
					{Text: text},
				},
			},
		},
	}

	body, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}


	req, err := http.NewRequest("POST", API, bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")


	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Raw response:", string(respBody))


	var geminiResp ResponseDTO
	if err := json.Unmarshal(respBody, &geminiResp); err != nil {
		log.Fatal(err)
	}

	// Выводим ответ
	var result string
	for _, candidate := range geminiResp.Candidates {
		for _, part := range candidate.Content.Parts {
			fmt.Println(part.Text)
			result += part.Text
		}
	}

	return &model.Question{
		Question: text,
		Answer: result,
	}, nil
}


// func main() {
	


	// // Запрос
	// input := ContentDTO{
	// 	Parts: []PartDTO{
	// 		{
	// 			Parts: []MessageDTO{
	// 				{Text: "Придумай 3 интересных факта о языке Go"},
	// 			},
	// 		},
	// 	},
	// }

	// body, err := json.Marshal(input)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req, err := http.NewRequest("POST", geminiURL+apiKey, bytes.NewReader(body))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Set("Content-Type", "application/json")


	// client := &http.Client{}
	// res, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()



	// respBody, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var geminiResp Response
	// if err := json.Unmarshal(respBody, &geminiResp); err != nil {
	// 	log.Fatal(err)
	// }

	// // Выводим ответ
	// for _, candidate := range geminiResp.Candidates {
	// 	for _, part := range candidate.Content.Parts {
	// 		fmt.Println(part.Text)
	// 	}
	// }
// }
