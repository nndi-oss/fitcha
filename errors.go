package fitcha

import "errors"

var ErrFeatureDoesNotExist = errors.New("feature does not exist")

var ErrFeatureIsNotEnabled = errors.New("feature is not enabled")

var ErrEmptyExpression = errors.New("expression to be evaluated cannot be empty string")

var ErrExpression = errors.New("the expression to be evaluated is invalid or cannot be evaluated")
