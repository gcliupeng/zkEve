package main

import (
	"strconv"
	zkeve "zkEve/lib"
)
func main() {
	zke,err:=zkeve.NewZkEve("127.0.0.1:2181")
	if(err!=nil){
		panic(err)
	}
	err=zke.SetUp("/zk/zkEve/aaa")
	if(err!=nil){
		panic(err)
	}
	for i:=0;;i++ {
		zke.Fire(strconv.Itoa(i))
	}
} 