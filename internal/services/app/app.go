package app

type App struct{}

func (a *App) Run() {

}

func (a *App) Close() {

}

func CreateNewApp() *App {
	return &App{}
}
