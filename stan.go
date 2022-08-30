package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
)

func (a *application) startSubNats() {
	clientID := "test-client-1"
	clusterID := "test-cluster"
	url := "nats://0.0.0.0:4222"

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("\nCan't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, url)
	}

	log.Println("Connected Nats")

	go a.publishNats(sc)

	_, err = sc.Subscribe("sub", func(msg *stan.Msg) {
		a.parseJsonFile(msg.Data)
	}, stan.DurableName("durable-1"))

	if err != nil {
		fmt.Println(err.Error())
	}
}

func (a *application) publishNats(sc stan.Conn) {
	byteValue, err := ioutil.ReadFile("model1.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = sc.Publish("sub", byteValue)
	if err != nil {
		fmt.Println(err.Error())
	}

	//ach := func(s string, err2 error) {}
	//_, err := Sc.PublishAsync(channel, data, ach)
	//if err != nil {
	//	log.Fatalf("Error during async publish: %v\n", err)
	//}
}
