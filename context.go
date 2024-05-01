package fitcha

import "context"

type FitchaKey string

const (
	fitchaUserKey  FitchaKey = "fitcha:user"
	fitchaOrgKey   FitchaKey = "fitcha:org"
	fitchaExtraKey FitchaKey = "fitcha:_extra"
)

func NewContext(user, org string, extra map[string]any) context.Context {
	fctx := context.WithValue(context.Background(), fitchaUserKey, user)
	fctx = context.WithValue(fctx, fitchaOrgKey, org)
	fctx = context.WithValue(fctx, fitchaExtraKey, extra)
	return fctx
}
