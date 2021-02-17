package etcd_go

import (
	"fmt"
	"github.com/zieckey/etcdsync"
)

func LockWithETCD(){
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil{
		fmt.Printf("etcd sync new failed: %+v\n", err)
		return
	}

	if err := m.Lock(); err != nil{
		fmt.Printf("etcd sync lock failed: %+v\n", err)
		return
	}
	fmt.Printf("etcd lock success.")

	if err := m.Unlock(); err != nil{
		fmt.Printf("etcd unlock failed: %+v\n", err)
	}
	fmt.Printf("etcd unlock success.")
}
