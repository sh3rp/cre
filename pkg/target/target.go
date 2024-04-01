package target

import "github.com/sh3rp/cre/internal/model"

type Target interface {
	ReadConfig() (*model.ConfigObject, error)
}
