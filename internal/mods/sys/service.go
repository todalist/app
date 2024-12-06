package sys

import (
	"context"
	"github.com/todalist/app/internal/models/dto"
)

type ISysService interface {
	PasswordLogin(context.Context, *dto.PasswordLoginDTO) (*string, error)
}
