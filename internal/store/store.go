package store

type TaskStorage interface {
	Create(t any) (string, error)
	Read() ([]map[string]string, error)
	Update(id string, t any) error
	Delete(id string) error
}
