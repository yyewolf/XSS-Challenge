package database

import "sync"

var dbmutex = sync.RWMutex{}
var db map[string][]interface{}

func init() {
	db = make(map[string][]interface{})
}

func GetAll(key string) (val []interface{}, found bool) {
	dbmutex.RLock()
	defer dbmutex.RUnlock()
	val, found = db[key]
	return
}

func GetOne(key string, id int) (val interface{}, found bool) {
	dbmutex.RLock()
	defer dbmutex.RUnlock()
	if id < len(db[key]) {
		val, found = db[key][id], true
	}
	return
}

func Add(key string, value interface{}) int {
	dbmutex.Lock()
	defer dbmutex.Unlock()
	db[key] = append(db[key], value)
	return len(db[key]) - 1
}

func Delete(key string, id int) {
	dbmutex.Lock()
	defer dbmutex.Unlock()
	db[key] = append(db[key][:id], db[key][id+1:]...)
}
