package trino

import (
	"context"
	"errors"
	"log"

	"github.com/bjornm82/saga/pkg/reqbin"
	saga "github.com/itimofeev/go-saga"
)

func UpdateTrino(id int, fail bool) Trino {
	return Trino{Id: id, Fail: fail}
}

type Trino struct {
	Id   int  `json:"id"`
	Fail bool `json:"fail"`
}

func (s *Trino) GetStep() *saga.Step {
	var state = Trino{}
	rb := reqbin.New()
	return &saga.Step{
		Name: "update-trino-step",
		Func: func(context.Context) error {
			if s.Fail {
				rb.ReturnError = errors.New("failed")
			}
			i, err := rb.Get()
			if err != nil {
				return err
			}
			state.Id = i

			log.Println("trino: set new version with ID: ", s.Id)
			log.Println("trino: previous version has ID: ", i)
			return nil
		},
		CompensateFunc: func(ctx context.Context) error {
			i, err := rb.Post(state.Id)
			if err != nil {
				return err
			}

			log.Println("trino rollback: set back to previous version with ID: ", i)
			return nil
		},
	}
}
