package cli

import (
	"log"
	"fmt"
)

import (
	"analysis-question-AI/internal/service"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/core"
	// "analysis-question-AI/internal/model"
)

type Commands interface {
	Run()
}

type commands struct {
	flags  Flags
	svc    service.QuestionService
	repo   repository.QuestionRepository
	env    core.Environment
}

func NewCommands(flags Flags, svc *service.QuestionService, repo *repository.QuestionRepository, env *core.Environment) Commands {
	return &commands{
		flags: flags,
		svc:    *svc,
		repo:   *repo,
		env:    *env,
	}
}

func (c *commands) Run() {
	
	
	// question := model.Question{
	// 	Question: "Придумай 3 интересных факта о языке Go", 
	// 	Answer: "",
	// }
	answer, err := c.svc.Send()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)

    
	// for i := 0; i < len(answer); i++ {
	// 	fmt.Println("Вопрос: ", answer[i].Question)
	// 	fmt.Println(" ")

	// 	for j := 0; j < len(answer[i].Options); j++ {
	// 		fmt.Println("\t", j+1, " вариант: ", answer[i].Options[j].Text)
	// 		fmt.Println(" ")
	// 		fmt.Println("\tВерно ли ?: ", answer[i].Options[j].IsCorrect)
	// 		fmt.Println(" ")
	// 		fmt.Println("\tПояснение: ", answer[i].Options[j].Explanation)
	// 		fmt.Println(" ")
	// 	}
		

	// 	fmt.Println("-------------------")
	// }
	


	

}
