package crontask

import (
	"fmt"
	"time"
)

func process() {

	panic("error: test")

}

func CornTask() {
	ticker := time.NewTicker(1 * time.Second)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			go func() {
	// 				defer func() {
	// 					if re := recover(); re != nil {
	// 						fmt.Println("recover", re)
	// 					}
	// 				}()
	// 				process()
	// 			}()
	// 		}
	// 	}
	// }()
	//
	// select {}

	for {
		select {
		case <-ticker.C:
			go func() {
				defer func() {
					if re := recover(); re != nil {
						fmt.Println("recover", re)
					}
				}()
				process()
			}()
		}
	}
}
