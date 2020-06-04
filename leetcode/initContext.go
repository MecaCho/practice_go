package leetcode

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func InitContext() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("############### waited for 1 sec")
	case <-time.After(2 * time.Second):
		fmt.Println("############### waited for 2 sec")
	case <-time.After(3 * time.Second):
		fmt.Println("############## waited for 3 sec")
	case <-ctx.Done():
		fmt.Println("############## time out.")
		fmt.Println(ctx.Err())
	}
}

func InitContext1() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch1 := make(chan int, 1)

	ch2 := make(chan int, 1)

	go func() {
		for {
			ch1 <- 1
		}

	}()

	go func() {
		for {
			ch2 <- 1
		}

	}()

	select {
	case <-ch1:
		fmt.Println("select 1")
	case <-ch2:
		fmt.Println("select 2")
	case <-time.After(3 * time.Second):
		fmt.Println("############## waited for 3 sec")
	case <-ctx.Done():
		fmt.Println("############## time out.")
		fmt.Println(ctx.Err())
	}
}

func FallThrough() {

	// techstack := "coding,Golang,Node.js,Python,Kubernetes,Serverless"
	techstack := "Golang"
	passions := "I can coding " + techstack

	fmt.Println(strings.Contains(passions, "coding"))
	if strings.Contains(passions, "coding") {

		switch techstack {

		case "Golang":
			fmt.Println("Golang")
			fallthrough

		case "Node.js":
			fmt.Println("node")
			fallthrough

		case "Python":
			fmt.Println("python")
			fallthrough

		case "Kubernetes":
			fmt.Println("k8s")
			fallthrough

		case "Serverless":
			fmt.Println("serverless")
			fmt.Println("xxxxxxx")

		}

	}
}
