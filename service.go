package cang

type Service interface {
	Start(app *App) error
	Stop()
	GetName() string
}
