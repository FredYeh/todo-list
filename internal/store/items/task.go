package items

type Task struct {
	Name   string `json:"name" redis:"name"`
	Status int    `json:"status" redis:"status"`
}

func (t *Task) Unify() {
	if t.Status > 1 {
		t.Status = 1
	} else if t.Status < 0 {
		t.Status = 0
	}
}
