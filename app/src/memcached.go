
package src


import (
	"github.com/bradfitz/gomemcache/memcache"
)


type Memcached struct {
	Connect *memcache.Client
}


func NewMemcached() *Memcached {
	return &Memcached{
		Connect: memcache.New("memcached-connect-go-memcached:11211"),
	}
}


func (memcached *Memcached) Get(key string) (res []byte, err error) {
	// mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")
	// mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	// it, err := mc.Get("foo")
	// ...
	// mc := memcache.New("memcached-connect-go-memcached:11211")
	// mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	// item, err := mc.Get("foo")
	// if err != nil {
	// 	c.JSON(200, gin.H{"message": "error"})
	// 	return
	// }

	// fmt.Print("\n=====\n", item, "\n", string(item.Value), "\n=====\n")

	item, err := memcached.Connect.Get(key)
	if err != nil {
		return []byte{}, err
	}

	return item.Value, nil
}


func (memcached *Memcached) Set(key string, value []byte) {
	// mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	memcached.Connect.Set(&memcache.Item{Key: key, Value: value})
}
