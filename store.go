package fitcha

import "sync"

type Storage interface {
	Add(feature *Feature) error

	Find(featureName string) (*Feature, error)

	Exists(featureName string) (bool, error)
}

type inmemoryStore struct {
	mu       sync.Mutex
	features map[string]*Feature
}

func NewInMemoryStorage() Storage {
	return &inmemoryStore{
		features: make(map[string]*Feature),
	}
}

func (m *inmemoryStore) Add(feature *Feature) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.features[string(feature.Name)] = feature
	return nil
}

func (m *inmemoryStore) Find(featureName string) (*Feature, error) {
	return m.features[featureName], nil
}

func (m *inmemoryStore) Exists(featureName string) (bool, error) {
	_, ok := m.features[featureName]
	return ok, nil
}
