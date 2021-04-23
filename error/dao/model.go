package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type ErrNoRows struct {
	message string
}

func (e *ErrNoRows) Error() string {
	return e.message
}

func NewModel() *Select {
	return &Select{}
}

type Select struct{}

var ErrNoRowsType = &ErrNoRows{"不存在的结果集"}

func (s *Select) Select() (*Select, error) {
	//假如报错了
	isErr := true
	if isErr {
		return nil, errors.Wrap(ErrNoRowsType, "查询失败")
	}
	return s, nil
}
