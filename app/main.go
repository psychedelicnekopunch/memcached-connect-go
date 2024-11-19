
package main


import (
	"fmt"

	"github.com/psychedelicnekopunch/memcached-connect-go/src"
)


func main() {

	memcached := src.NewMemcached()
	// fmt.Print("\n=====\n", memcached, "\n=====\n")
	key := "foo"
	// memcached.Set(key, []byte("test"))
	memcached.Set(key, []byte("my value"))
	v, err := memcached.Get(key)
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}
	fmt.Printf("%s is %s\n", key, string(v))
}
