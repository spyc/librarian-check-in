package modals

import "fmt"

func (r Rank) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", r.String())), nil
}
