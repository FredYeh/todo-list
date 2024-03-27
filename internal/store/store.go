package store

type TaskStorage interface {
	Create(t any) error
	Read() []map[string]string
	Update(id string, t any) error
	Delete(id string) error
}
