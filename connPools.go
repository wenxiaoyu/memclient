package gmc

import (
	"fmt"
	"hash/crc32"
	"net"
	"strconv"
	"sync"
)

const virConnNum int = 3
var pools = make([]pool, 0, 10)

type McConn struct {
	mutex sync.Mutex
	conn  net.Conn
	alive bool
	free  bool
}

type pool struct {
	circle int
	conn   *McConn
}

func init() {
	for i := 0; i < virConnNum; i++ {
		for _, host := range McConfig.Host {
			key := string(host.Ip) + "->" + strconv.Itoa(int(host.Port)) + "->" + strconv.Itoa(i)
			crc := int(crc32.ChecksumIEEE([]byte(key)))
			c, err := net.Dial("tcp", string(host.Ip)+":"+strconv.Itoa(int(host.Port)))
			if err != nil {
				fmt.Println(err)
				continue
			}
			mcconn := &McConn{conn: c, alive: true, free: true}
			p := &pool{circle: crc, conn: mcconn}
			pools = append(pools, *p)
		}
	}
	err := shellSort(pools)
	if err != nil {
		fmt.Errorf("init connection error->", err)
	}
}


func shellSort(pools []pool) error {
	n := len(pools)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			tmp := pools[i]
			j := i
			for j >= gap && pools[j-gap].circle > tmp.circle {
				pools[j] = pools[j-gap]
				j = j - gap
			}
			pools[j] = tmp
		}
		gap = gap / 2
	}
	return nil
}

func GetMcConn(key string)(conn McConn,err error){
	
}
