package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"

	"flashflag/conf"
	"flashflag/notice"
)

var (
	fInt = flag.Int("f-int", 1, "flash int flag.")
	fStr = flag.String("f-str", "deadbeaf", "flash string flag.")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	zk := notice.NewConfZK([]string{"127.0.0.1:2181"}, 10*time.Second)
	if zk == nil {
		glog.Fatal("create zookeeper client error.")
	}

	// create(or set) /conf/mysvr.f-int 5  // for all nodes.
	// create(or set) /conf/mysvr.node1.f-int 5  // for node1 only.
	// carete(or set) /conf/mysvr.node1.f-str ok // for node1 only, change f-str to "ok".
	conf.NewConf("/conf", "mysvr", "node1", zk)

	for {
		if *fInt == -1000 {
			break
		}

		fmt.Printf("Hello flash flag, f-int %d, f-str %s\n", *fInt, *fStr)
		time.Sleep(3 * time.Second)
	}
}
