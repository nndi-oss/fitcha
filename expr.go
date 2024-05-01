package fitcha

import (
	"context"
	"fmt"

	"github.com/expr-lang/expr"
)

func evaluateCondition(ctx context.Context, condition string) bool {
	env := make(map[string]any)

	env["fitcha.user"] = ctx.Value(fitchaUserKey)
	env["fitcha.org"] = ctx.Value(fitchaOrgKey)

	data := ctx.Value(fitchaExtraKey)
	if data != nil {
		dataMap := data.(map[string]any)
		for key, val := range dataMap {
			env[key] = val
		}
	}

	program, err := expr.Compile(condition, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(output) == "true"
}
