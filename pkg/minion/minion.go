package minion

type Minion interface {
	GetWork() ([]*Work, error)
}

type Work struct {
}
