package models

var Cache *LocalCache = NewCache()

type LocalCache struct {
	data map[string]interface{}
}

func NewCache() *LocalCache {
	c := &LocalCache{}
	c.data = make(map[string]interface{})
	return c
}

func (this *LocalCache) Delete(key string) error {
	delete(this.data,key)
	return nil
}

func (this *LocalCache) IsExist(key string) bool {
	if _, ok := this.data[key]; ok {
		return true
	}
	return false

}
