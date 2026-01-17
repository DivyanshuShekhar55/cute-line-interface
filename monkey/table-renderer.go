package monkey

import (
	"cute-line-interface/utils"
	"errors"
)

func (t *Table) Render() {
	if t.header == nil {
		err := errors.New("Header missing")
		utils.LogError(err)
		return
	}

	

}