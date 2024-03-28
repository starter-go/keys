package lib

import (
	"fmt"
	"strings"
	"sync"

	"github.com/starter-go/keys"
)

// DriverManagerImpl ...
type DriverManagerImpl struct {

	//starter:component

	_as func(keys.DriverManager) //starter:as("#")

	RawDrivers []keys.Driver //starter:inject(".")

	table map[string]*keys.DriverRegistration
	mutex sync.Mutex
}

func (inst *DriverManagerImpl) _impl() keys.DriverManager {
	return inst
}

// Find ...
func (inst *DriverManagerImpl) Find(algorithm string, class keys.Class) (keys.Driver, error) {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	table := inst.getDrivers()
	key := inst.keyFor(algorithm, class)
	res := table[key]

	if res == nil {
		return nil, fmt.Errorf("no key driver for [algorithm:%s class:%s]", algorithm, class)
	}
	return res.Driver, nil
}

func (inst *DriverManagerImpl) loadDrivers() map[string]*keys.DriverRegistration {
	src := inst.RawDrivers
	dst := make(map[string]*keys.DriverRegistration)
	for _, dr1 := range src {
		if dr1 == nil {
			continue
		}
		list := dr1.ListRegistrations()
		for _, dr2 := range list {
			if inst.isAvailable(dr2) {
				key := inst.keyFor(dr2.Algorithm, dr2.Class)
				dst[key] = dr2
			}
		}
	}
	return dst
}

func (inst *DriverManagerImpl) getDrivers() map[string]*keys.DriverRegistration {
	t := inst.table
	if t == nil {
		t = inst.loadDrivers()
		inst.table = t
	}
	return t
}

func (inst *DriverManagerImpl) keyFor(algorithm string, class keys.Class) string {
	str := algorithm + "." + string(class)
	return strings.ToLower(str)
}

func (inst *DriverManagerImpl) isAvailable(dr *keys.DriverRegistration) bool {
	if dr == nil {
		return false
	}
	if !dr.Enabled {
		return false
	}
	if dr.Driver == nil {
		return false
	}
	return true
}
