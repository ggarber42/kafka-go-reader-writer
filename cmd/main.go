package main

import (
	"fmt"
	"kafkago/internal/kafka"
	"os"
)

func main() {
	fmt.Println("vai amyzona")
	brokers := []string{"localhost:9092"}
	p, err := kafka.NewProducer(brokers)
	if err != nil {
		fmt.Println("corram para as montanhas")
		os.Exit(1)
	}
	err = kafka.Publish(p, "aaaa", "faaaaa")
	if err != nil {
		fmt.Println("corram para as montanhas")
		os.Exit(1)
	}
}