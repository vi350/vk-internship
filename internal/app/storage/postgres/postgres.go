package postgres

import (
	"github.com/go-pg/pg"
	"github.com/vi350/vk-internship/internal/app/e"
)

type Storage struct {
	DB *pg.DB
}

func New() (ps *Storage, err error) {
	defer func() { err = e.WrapIfErr("error performing request: ", err) }()
	ps = &Storage{}

	if ps.DB, err = connect(); err != nil {
		return nil, err
	}

	return ps, nil
}

func (ps *Storage) Close() error {
	return ps.DB.Close()
}
