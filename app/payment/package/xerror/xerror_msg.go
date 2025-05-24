package xerror

import (
	"context"
	"payment/locale"
)

type ErrorMsg string

func (e ErrorMsg) String(ctx context.Context, params ...interface{}) string {
	return locale.Get(ctx, string(e), params...)
}

var (
	ErrMsgFieldRequired ErrorMsg = "ErrMsgFieldRequired"
	ErrMsgFieldInvalid  ErrorMsg = "ErrMsgFieldInvalid"
)
