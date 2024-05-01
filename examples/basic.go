package main

import (
	"fmt"
	"time"

	"github.com/nndi-oss/fitcha"
)

func main() {
	newFeature := fitcha.Feature{
		ID:        "test",
		Name:      "test",
		IsEnabled: true,
		Expr:      `"typist" in roles`,
		CreatedAt: time.Now(),
	}

	ctx := fitcha.NewContext("administrator", "default", map[string]any{
		"roles": []string{"typist"},
		"age":   21,
	})

	featureManager := fitcha.NewFeatureManager(fitcha.NewInMemoryStorage())

	featureManager.Store().Add(&newFeature)

	// check if the feature is enabled..
	ok, err := featureManager.IsEnabled(ctx, "test")
	if ok && err == nil {
		fmt.Println("FEATURE is enabled")
	}

	// use the expression stored with the feature
	ok, _ = featureManager.Evaluate(ctx, "test")
	if ok {
		fmt.Println("user can use the feature")
	} else {
		fmt.Println("user cannot use this feature.")
	}

	// override the existing expression
	ok, _ = featureManager.EvaluateExpr(ctx, "test", `"typist" in roles && age >= 20`)
	if ok {
		fmt.Println("user can use the feature")
	} else {
		fmt.Println("user cannot use this feature.")
	}

	err = featureManager.Disable(ctx, "test")
	if err != nil {
		panic(err)
	}

	ok, _ = featureManager.IsEnabled(ctx, "test")
	if !ok {
		fmt.Println("FEATURE is disabled")
	}
}
