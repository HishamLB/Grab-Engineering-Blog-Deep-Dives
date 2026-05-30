package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

func getResultFromCommands(cmds []redis.Cmder) error {
	for _, cmd := range cmds {
		if err := cmd.Err(); err != nil && err != redis.Nil {
			return err
		}
	}
	return nil
}

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6001", "127.0.0.1:6002", "127.0.0.1:6003",
			"127.0.0.1:6004", "127.0.0.1:6005", "127.0.0.1:6006",
			"127.0.0.1:6007", "127.0.0.1:6008", "127.0.0.1:6009",
		},
		ReadOnly:       true,
		RouteByLatency: false,
		RouteRandomly:  false,
		MaxRedirects:   3,
	})
	defer client.Close()

	fmt.Println("=== Seeding 200 keys...")
	for i := 1; i <= 200; i++ {
		client.HMSet(fmt.Sprintf("driver:%d", i), map[string]interface{}{
			"id": i, "state": "available", "lat": 1.23 + float64(i)*0.001,
			"lng": 103.45 + float64(i)*0.001, "updated_at": time.Now().Unix(),
		})
	}
	fmt.Println("=== Done. Starting pipeline test...\n")

	fmt.Println("Kill a slave: redis-cli -p 6005 SHUTDOWN")
	fmt.Println("Will auto-stop after 20 iterations (10 seconds).\n")

	timeout := time.After(10 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	var total, failed int

	for {
		select {
		case <-timeout:
			fmt.Printf("\nTotal: %d  Failed: %d (%.1f%%)\n", total, failed, float64(failed)/float64(total)*100)
			return

		case <-ticker.C:
			pipe := client.Pipeline()
			cmds := make([]redis.Cmder, 10)
			for i := 0; i < 10; i++ {
				cmds[i] = pipe.HMGet(fmt.Sprintf("driver:%d", rand.Intn(200)+1), "id", "state", "lat", "lng")
			}
			_, _ = pipe.Exec()
			err := getResultFromCommands(cmds)

			total++
			if err != nil {
				failed++
				fmt.Printf("[FAIL] %v\n", err)
			} else {
				fmt.Printf("[OK]\n")
			}
		}
	}
}
