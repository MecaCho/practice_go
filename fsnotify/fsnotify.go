package fsnotify

import (
	"gopkg.in/fsnotify.v1"
	"log"
)

func FsNotify(fileName string) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("New watcher failed: %+v.\n", err)
	}
	defer watcher.Close()

	done := make(chan int)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event get: ", event.String())
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error: ", err.Error())
			}
		}
	}()

	err = watcher.Add(fileName)
	if err != nil {
		log.Fatalf("Add file (%s) error %+v.", fileName, err)
	}
	<-done
}
