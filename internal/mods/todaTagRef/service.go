package todaTagRef

import (
	"context"
)

type ITodaTagRefService interface {

	// basic crud
	Get(context.Context, uint) (*TodaTagRef, error)

	First(context.Context, *TodaTagRefQuerier) (*TodaTagRef, error)

	Save(context.Context, *TodaTagRef) (*TodaTagRef, error)

	List(context.Context, *TodaTagRefQuerier) ([]*TodaTagRef, error)

	Delete(context.Context, uint) (uint, error)

}