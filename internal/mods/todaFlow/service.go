package todaFlow

import (
	"context"
)

type ITodaFlowService interface {

	// basic crud
	Get(context.Context, uint) (*TodaFlow, error)

	Save(context.Context, *TodaFlow) (*TodaFlow, error)

	List(context.Context, *TodaFlowQuerier) ([]*TodaFlow, error)

	Delete(context.Context, uint) (uint, error)

}