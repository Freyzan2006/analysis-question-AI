package cli

import (
	"log"
	"fmt"
)

import (
	"analysis-question-AI/internal/service"
	"analysis-question-AI/internal/repository"
)

type Commands interface {
	Run()
}

type commands struct {
	flags  Flags
	svc    service.QuestionService
	repo   repository.QuestionRepository
}

func NewCommands(flags Flags, svc *service.QuestionService, repo *repository.QuestionRepository) Commands {
	return &commands{
		flags: flags,
		svc:    *svc,
		repo:   *repo,
	}
}

func (c *commands) Run() {
	answer, err := c.svc.Send()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
