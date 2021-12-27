package pipeline

import (
	"context"

	"github.com/bjornm82/saga/internal/registry"
	"github.com/bjornm82/saga/internal/trino"
	saga "github.com/itimofeev/go-saga"
)

func UpdateSource() error {
	s := saga.NewSaga("process")

	ur := registry.UpdateRegistry(1, false)
	s.AddStep(ur.GetStep())
	ut := trino.UpdateTrino(3, false)
	s.AddStep(ut.GetStep())
	ut2 := trino.UpdateTrino(5, true)
	s.AddStep(ut2.GetStep())

	store := saga.New()
	c := saga.NewCoordinator(
		context.Background(),
		context.Background(),
		s,
		store,
	)

	res := c.Play()
	return res.ExecutionError
}
