package driver

import "fmt"

type FilterFn func(Driver) bool

func FilterKind(k Kind) FilterFn {
	return func(d Driver) bool {
		return d.Info().Kind == k
	}
}

// Manager is a singleton to manage multiple drivers and their states
type Manager struct {
	drivers map[string]Driver
}

var manager = &Manager{
	drivers: make(map[string]Driver),
}

// GetManager gets manager singleton instance
func GetManager() *Manager {
	return manager
}

// Register registers adapter to be discoverable by Query
func (m *Manager) Register(a Adapter) error {
	d := wrapAdapter(a)
	if d == nil {
		return fmt.Errorf("adapter has to be either VideoAdapter/AudioAdapter")
	}

	m.drivers[d.ID()] = d
	return nil
}

// Query queries by using f to filter drivers, and simply return the filtered results.
func (m *Manager) Query(f FilterFn) []Driver {
	results := make([]Driver, 0)
	for _, d := range m.drivers {
		if ok := f(d); ok {
			results = append(results, d)
		}
	}

	return results
}