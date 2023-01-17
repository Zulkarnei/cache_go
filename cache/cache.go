package cache

type Cache map[string]interface{}

func New() Cache {
	return make(map[string]interface{})
}

func (c Cache) Set(key string, value interface{}) {
	c[key] = value
}

func (c Cache) Get(key string) interface{} {
	return c[key]
}

func (c Cache) Delete(key string) {
	delete(c, key)
}
