package types

import (
	"github.com/branila/shortener/utils"
)

type Db map[string]string

func (d Db) String() string {
	return utils.PrettifyObject(d)
}
