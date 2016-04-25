package main

import (
	"fmt"
	"time"
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
	 go func () {
	 	time.Sleep(5*time.Second)
	 	zke.Fire("hello world ")
	 }()
	ech,err:=zke.Listen()
	for {
		e:=<- ech
		fmt.Printf("get message %s\n", e.Data)
	}
	// c, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second, 30) //*10)
	// if err != nil {
	// 	panic(err)
	// }
	// _, _, ch, err := c.GetW("/zk/zkEve/aaa")
	// if err != nil {
	//  	panic(err)
	// }
	// // fmt.Printf("%+v %+v\n", children, stat)
	// e := <-ch
	// fmt.Printf("%+v\n", e)
}