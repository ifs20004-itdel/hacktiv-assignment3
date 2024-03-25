package main

import (
	"assignment3/models"
	"assignment3/routers"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	go func() {
		for range time.Tick(time.Second * 15) {
			writeRandNumber()
		}
	}()

	routers.StartServer().Run(":8080")
}

func writeRandNumber() {
	data := models.Status{
		Status: map[string]int{
			"water": rand.Intn(100),
			"wind":  rand.Intn(100),
		},
	}
	newData, _ := json.MarshalIndent(data, "", "\t")
	err := os.WriteFile("status.json", newData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
