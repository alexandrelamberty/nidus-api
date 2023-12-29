package main

import (
	"fmt"
	"log"
	"nidus-server/internal/config"
	"os"
	"os/signal"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

/*
Nidus API
https://docs.gofiber.io/
*/
func main() {
	// Environments variables
	log.Println("[Config] Checking configuration ...")
	err := config.CheckConfig()
	if err != nil {
		log.Fatal("[CheckConfig] ", err)
	}

	// MQTT
	var broker = "localhost"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("mqtt://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client_23423423")
	// opts.SetUsername("")
	// opts.SetPassword("")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub(client)
	// publish(client)

	// Set up a signal handler to gracefully exit the script
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Received interrupt signal. Closing MQTT client...")
		client.Disconnect(250) // Allow 250 milliseconds for the disconnect to complete
		os.Exit(0)
	}()

	// Run a loop to keep the script alive
	for {
		time.Sleep(time.Second) // Adjust the sleep duration as needed

		// Check if the client is still connected, and reconnect if necessary
		if !client.IsConnected() {
			fmt.Println("Reconnecting to MQTT broker...")
			if token := client.Connect(); token.Wait() && token.Error() != nil {
				fmt.Printf("Error reconnecting: %s\n", token.Error())
			} else {
				fmt.Println("Reconnected to MQTT broker.")
				// Resubscribe to the topic after reconnecting if needed
				// if token := client.Subscribe(topic, 0, messagePubHandler); token.Wait() && token.Error() != nil {
				//     fmt.Printf("Error resubscribing: %s\n", token.Error())
				// }
			}
		}
	}
}
