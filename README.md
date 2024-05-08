fitcha
======

Small feature management library for Go developers. Slightly opionated, highly experimental. YMMV.

Fitcha enables you to check features available to users per-request or globally. It stores 
features in a `Store`, and uses `Manager` to check availability of a feature. Features can
be added to the store dynamically or pre-loaded. One of the things that makes fitcha special
is it's support for dynamic expressions which enable you to implement feature validation with
additional logic via expressions written with [expr](https://github.com/expr-lang/expr) - hopefully this will give you more power for rolling out features.

The API is unstable, the project is experimental - but we are dogfooding it so if doesn't work
out, we have a lot of refactoring to do on our internal projects :D


## Usage

## Example

See [./examples/basic.go](examples/basic.go)


### Fiber Middleware

```go
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nndi-oss/fitcha"
)

const (
	FeatureTest = "test"
)

func main() {
	app := fiber.New()
	newFeature := fitcha.Feature{
		ID:        "test",
		Name:      "test",
		IsEnabled: false,
		Expr:      `"typist" in roles`,
		CreatedAt: time.Now(),
	}
	featureManager := fitcha.NewFeatureManager(fitcha.NewInMemoryStorage())
	featureManager.Store().Add(&newFeature)

	app.Use(useFeatureMiddleware(featureManager, "test"))

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	err := app.Listen(":3006")
	if err != nil {
		panic(err)
	}
}

func getUserContext(c *fiber.Ctx) context.Context {
    // ctx := fitcha.NewContext("administrator", "default", map[string]any{
    // 	"roles": []string{"typist"},
    // 	"age":   21,
    // })
    // return ctx
    return context.Background()
}

func useFeatureMiddleware(featureManager fitcha.FeatureManager, featureName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		ok, err := featureManager.IsEnabled(getUserContext(c), featureName)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "feature not enabled for your account",
			})
		}

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":   err.Error(),
				"message": "failed to check feature, internal error",
			})
		}

		return c.Next()
	}
}

```