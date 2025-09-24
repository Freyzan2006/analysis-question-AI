package sheet 

import (
	"analysis-question-AI/internal/core"
)

type SheetModule struct {
	commands *sheetCommand
}

func NewSheetModule(log *core.Logger, cfg *core.Config) *SheetModule {
	api := newSheetApi(cfg);
	repo := newSheetRepository(api);
	svc := newSheetService(repo, log);
	commands := newSheetCommand(log, cfg, svc)

	return &SheetModule{
		commands: commands,
	}
}

func (s *SheetModule) Commands() *sheetCommand {
	return s.commands;
}
