package todaTag

import (
	"context"
)

type ITodaTagService interface {

	// basic crud
	Get(context.Context, uint) (*TodaTag, error)

	Save(context.Context, *TodaTag) (*TodaTag, error)

	List(context.Context, *TodaTagQuerier) ([]*TodaTag, error)

	Delete(context.Context, uint) (uint, error)

}