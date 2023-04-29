package presenter

type PresenterDevice interface {
	SendMessage(id string, msg string) error
}
