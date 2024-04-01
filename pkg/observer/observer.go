package observer

import "github.com/sh3rp/cre/internal/model"

type Observer interface {
	Observe() (*model.ConfigObject, error)
}
