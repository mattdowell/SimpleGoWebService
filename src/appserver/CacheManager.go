package appserver

import "sync"

type data struct {
	name    string
	address string
	active  bool
	zip     int32
}

var singleton *data
var once sync.Once

// In Go, a name is exported if it begins with a capital letter.
func GetManager() *data {
	once.Do(func() {
		singleton = &data{
			name: "Matheson, Richard",
			address: "123 Main St. Los Angeles, CA",
			active:  true,
			zip:     55678}
	})
	return singleton
}

func (sm *data) GetName() string {
	return sm.name
}

func (sm *data) getAddress() string {
	return sm.address
}

func (sm *data) getActive() bool {
	return sm.active
}

func (sm *data) getZip() int32 {
	return sm.zip
}
