package usecase

import (
	"github.com/amiame/go-websocket-server/pkg/usecase/presenter"
)

type UsecaseDevice interface {
	SendMessage(id string, msg string) error
}

type device struct {
	presenter presenter.PresenterDevice
}

func NewDevice(presenter presenter.PresenterDevice) *device {
	return &device{
		presenter: presenter,
	}
}

func (d *device) SendMessage(id string, msg string) error {
	return d.presenter.SendMessage(id, msg)
}
