package sheet 


type sheetModule struct {
	commands *sheetCommands
}

func NewSheetModule(log *core.Logger, flags *cli.flagsConfig) *sheetModule {
	api := newSheetApi();
	repo := newSheetRepository(api);
	svc := newSheetService(repo);
	commands := newSheetCommands(log, flags, svc)

	return &sheetModule{
		commands: commands,
	}
}

func (s *sheetModule) Commands() *sheetCommands {
	return s.commands;
}
