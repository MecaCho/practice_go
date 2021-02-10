package redis_go

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func Incr() {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	lockKey := "counter_lock"
	counterKey := "counter"

	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Printf("Lock failed: %+v, %+v\n", lockSuccess, err)
		return
	}

	counterResp := client.Get(counterKey)
	counterValue, err := counterResp.Int64()
	if err != nil {
		counterValue++
		resp := client.Set(counterKey, counterValue, 0)
		if _, err := resp.Result(); err != nil {
			fmt.Printf("Set counter value error: %+v\n", err)
		}
	}
	fmt.Printf("Current counter value is: %+v\n", counterValue)

	delResp := client.Del(lockKey)
	unlockResult, err := delResp.Result()
	if err == nil && unlockResult > 0 {
		fmt.Printf("Unlock success.\n")
	} else {
		fmt.Printf("unlock failed: %+v\n", err)
	}

}

func CounterWithRedisLock() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Incr()
		}()
	}
	wg.Wait()
}
