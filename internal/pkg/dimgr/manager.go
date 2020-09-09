package dimgr

import "sync"

var dimgr *Manager
var once sync.Once

// The Manager contains the information related to the dependencies can be injected.
type Manager struct {
	deps map[string]interface{}
}

// GetManager gets dependency injection manager instance.
func GetManager() *Manager {
	once.Do(func() {
		dimgr = &Manager{
			deps: make(map[string]interface{}),
		}
	})
	return dimgr
}

// AddDependency add a dependency instance.
func (mgr *Manager) AddDependency(name string, dep interface{}) {
	mgr.deps[name] = dep
}

// GetDependency gets a dependency instance.
func (mgr Manager) GetDependency(name string) interface{} {
	return mgr.deps[name]
}
