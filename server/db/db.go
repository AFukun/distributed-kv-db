package db

type Database struct {
	values map[string]int
}

func (db *Database) Init() {
	db.values = make(map[string]int)
}

func (db *Database) Get(key string) int {
	return db.values[key]
}

func (db *Database) Put(key string, value int) {
	db.values[key] = value
}

func (db *Database) Delete(key string) {
	delete(db.values, key)
}
