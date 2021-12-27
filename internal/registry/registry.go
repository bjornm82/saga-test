package registry

import (
	"context"
	"errors"
	"log"

	"github.com/bjornm82/saga/pkg/reqbin"
	saga "github.com/itimofeev/go-saga"
)

func UpdateRegistry(id int, fail bool) Registry {
	return Registry{Id: id, Fail: fail}
}

type Registry struct {
	Id   int  `json:"id"`
	Fail bool `json:"fail"`
}

func (s *Registry) GetStep() *saga.Step {
	var state = Registry{}
	rb := reqbin.New()
	return &saga.Step{
		Name: "update-registry-step",
		Func: func(context.Context) error {
			if s.Fail {
				rb.ReturnError = errors.New("failed")
			}
			i, err := rb.Get()
			if err != nil {
				return err
			}
			state.Id = i

			log.Println("registry: set new version with ID: ", s.Id)
			log.Println("registry: previous version has ID: ", i)
			return nil
		},
		CompensateFunc: func(ctx context.Context) error {
			i, err := rb.Post(state.Id)
			if err != nil {
				return err
			}

			log.Println("rollback registry: set back to previous version with ID: ", i)
			return nil
		},
	}
}
