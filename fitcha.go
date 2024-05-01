package fitcha

import (
	"time"
)

type FeatureName string

type Feature struct {
	ID        string      `json:"id"`
	Name      FeatureName `json:"name"`
	IsEnabled bool        `json:"isEnabled"`
	Expr      string      `json:"expr"`
	CreatedAt time.Time   `json:"createdAt"`
}
