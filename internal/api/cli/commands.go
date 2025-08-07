package cli

import (
	"fmt"
	"log"
)

import (
	"analysis-question-AI/internal/service"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/model"
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
	var flagConfig = c.flags.GetFlags()
	
	question := model.Question{
		Question: "Придумай 3 интересных факта о языке Go", 
		Answer: "",
	}
	answer, err := c.svc.Send(question)
	if err != nil {
		log.Fatal(err)
	}

	
	


	fmt.Println(answer)

	fmt.Println(flagConfig.FILE_OUTPUT)
	fmt.Println(flagConfig.FILE_INPUT)
}
