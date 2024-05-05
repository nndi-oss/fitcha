package fitcha

import "context"

type FitchaKey string

const (
	UserCtx  FitchaKey = "fitcha:user"
	OrgCtx   FitchaKey = "fitcha:org"
	ExtraCtx FitchaKey = "fitcha:_extra"
)

func NewContext(user, org string, extra map[string]any) context.Context {
	fctx := context.WithValue(context.Background(), UserCtx, user)
	fctx = context.WithValue(fctx, OrgCtx, org)
	fctx = context.WithValue(fctx, ExtraCtx, extra)
	return fctx
}
