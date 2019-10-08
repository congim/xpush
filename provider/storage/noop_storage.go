package storage

type noopStorage struct {
}

var _ Storage = (*noopStorage)(nil)

func newNoopStorage() *noopStorage {
	return &noopStorage{}
}

func (n *noopStorage) Store() error {
	return nil
}
