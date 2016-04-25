package zkEve
import (
	"time"
	"path"
	"github.com/wandoulabs/go-zookeeper/zk"
)
type ZkEve struct{
	Host string
	Path string
	Conn *zk.Conn
}
type Event struct{
	Path string
	Time string
	Data string
}
func NewZkEve(host string) (zke *ZkEve,err error) {
	c, _, err := zk.Connect([]string{host}, time.Second, 30) //*10)
	if err != nil {
		return nil,err
	}
	return &ZkEve{Host:host,Conn:c},nil
}
func (zke * ZkEve) SetUp(path string) error {
	exist,_,err:=zke.Conn.Exists(path)
	if(err!=nil){
		return err
	}
	if(exist){
		zke.Path=path
		return nil
	}else{
		path,err:=zke.Conn.Create(path,[]byte(""),0,zk.WorldACL(zk.PermAll))
		zke.Path=path
		return err
	}
}
func (zke *ZkEve)Fire(data string) error {
	path:=path.Join(path.Dir(zke.Path),"lock")
	l := zk.NewLock(zke.Conn, path, zk.WorldACL(zk.PermAll))
	defer l.Unlock()
	if err := l.Lock(); err != nil {
		return err
	}
	_,err:=zke.Conn.Set(zke.Path,[]byte(data),-1)
	return err
}
func (zke *ZkEve)Listen() (<-chan *Event,error){
	var ech chan *Event
	ech = make(chan *Event)
	_, _, ch, err := zke.Conn.GetW(zke.Path)
	if err != nil {
	 	return nil,err
	}
	go func () {
		for{
			<-ch
			data,state,err:=zke.Conn.Get(zke.Path)
			if(err != nil ){
				return
			}  
			eve:=Event{
				Data:string(data),
				Path:zke.Path,
				Time:state.MTime().Format("2006-01-02 15:04:05"),
			}
			ech <-&eve
			// rewatch
			_, _, ch, err = zke.Conn.GetW(zke.Path)
			if err != nil {
	 			close(ech)
	 			return 
			}
		}
	}()
	return ech,nil
}