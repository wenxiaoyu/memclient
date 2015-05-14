// go memcache client
package gmc

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
)

type MIN int
type MAX int
type IP string
type PORT int

type Root struct {
	Host []McHost
	Pool McPool
}
type McHost struct {
	Ip   IP
	Port PORT
}

type McPool struct {
	Min MIN
	Max MAX
}

var McConfig = Root{}

func init() {
	fmt.Println("--init memcached client config---")
	bxml, err := ioutil.ReadFile("memcache.xml")
	if err != nil {
		panic(err)
	}	 
    fmt.Println(string(bxml))
	err = xml.Unmarshal(bxml, &McConfig)
   
	if err != nil {
	}
}
