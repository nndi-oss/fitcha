package fitcha

import "context"

type FeatureManager interface {
	// IsEnabled(context.Context, string) (bool, error)

	// Disable(context.Context, string) error

	// Enable(context.Context, string) error

	// AddCondition(context.Context, string, string) error

	Store() Storage

	Evaluate(context.Context, string) (bool, error)

	EvaluateExpr(context.Context, string, string) (bool, error)

	IsEnabled(context.Context, string) (bool, error)

	Disable(context.Context, string) error
}

type simpleFeatureManager struct {
	store Storage
}

func NewFeatureManager(store Storage) FeatureManager {
	return &simpleFeatureManager{
		store: store,
	}
}

func (sfm *simpleFeatureManager) Store() Storage {
	return sfm.store
}

func (sfm *simpleFeatureManager) IsEnabled(ctx context.Context, featureName string) (bool, error) {
	feature, _ := sfm.Store().Find(featureName)
	if !feature.IsEnabled {
		return false, ErrFeatureIsNotEnabled
	}
	return true, nil
}

func (sfm *simpleFeatureManager) Disable(ctx context.Context, featureName string) error {
	feature, err := sfm.Store().Find(featureName)
	if err != nil {
		return err
	}
	feature.IsEnabled = false
	sfm.Store().Add(feature)
	return nil
}

func (sfm *simpleFeatureManager) Evaluate(ctx context.Context, featureName string) (bool, error) {
	feature, _ := sfm.Store().Find(featureName)
	if !feature.IsEnabled {
		return false, ErrFeatureIsNotEnabled
	}
	return evaluateCondition(ctx, feature.Expr), nil
}

func (sfm *simpleFeatureManager) EvaluateExpr(ctx context.Context, featureName string, cond string) (bool, error) {
	feature, _ := sfm.Store().Find(featureName)
	if !feature.IsEnabled {
		return false, ErrFeatureIsNotEnabled
	}
	if cond == "" {
		return false, ErrEmptyExpression
	}
	return evaluateCondition(ctx, cond), nil
}