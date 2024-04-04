package items

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Status int

const (
	Todo Status = iota
	Done
)

func (s Status) String() string {
	return map[int]string{0: "Todo", 1: "Done"}[int(s)]
}

func (s *Status) MarshalJSON() ([]byte, error) {
	if *s != 0 && *s != 1 {
		return nil, fmt.Errorf("%v is not a valid status", *s)
	}
	return json.Marshal(s.String())
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var status string
	err := json.Unmarshal(data, &status)
	if err != nil {
		return err
	}
	*s, err = ParseStatus(status)
	if err != nil {
		return err
	}
	return nil
}

func ParseStatus(s string) (Status, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := map[string]int{"todo": 0, "done": 1}[s]
	if !ok {
		return Status(0), fmt.Errorf("%v is not a valid status", s)
	} else {
		return Status(value), nil
	}
}

type Task struct {
	Name   string `json:"name" redis:"name"`
	Status Status `json:"status" redis:"status"`
}
