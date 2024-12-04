package toda

import (
	"context"
)

type ITodaService interface {

	// basic crud
	Get(context.Context, uint) (*Toda, error)

	Save(context.Context, *Toda) (*Toda, error)

	List(context.Context, *TodaQuerier) ([]*Toda, error)

	Delete(context.Context, uint) (uint, error)

}