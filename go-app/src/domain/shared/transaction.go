package shared

import (
	"context"
)

type Transaction interface {
	DoInTx(context.Context, func(context.Context) error) error
}
