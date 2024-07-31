package mqtt

import (
	"fmt"
	"log"

	"serveos-datasync/config"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var MqttClient MQTT.Client
var MqttConnected int32
var MqttDisconnectsLastHour int32
var MqttDisconnectsLast24Hours int32

func InitMQTT(cfg config.Config) {
	opts := MQTT.NewClientOptions().
		AddBroker(fmt.Sprintf("%s://%s:%d", cfg.ServerMQTTProtocol, cfg.ServerMQTTHost, cfg.ServerMQTTPort)).
		SetClientID("serveos-datasync").
		SetUsername(cfg.ServerMQTTUsername).
		SetPassword(cfg.ServerMQTTPassword).
		SetCleanSession(false).
		SetOrderMatters(false).
		SetKeepAlive(30).
		SetDefaultPublishHandler(messagePubHandler).
		SetOnConnectHandler(connectHandler).
		SetConnectionLostHandler(connectLostHandler)

	if cfg.ServerMQTTEnableTLS {
		opts.SetTLSConfig(NewTLSConfig(cfg))
	}

	MqttClient = MQTT.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %s", token.Error())
	}
}

func Publish(topic string, payload interface{}) {
	token := MqttClient.Publish(topic, 2, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Failed to publish message to topic %s: %v", topic, token.Error())
	} else {
		log.Printf("Published message to topic %s", topic)
	}
}

func Subscribe(topic string) {
	token := MqttClient.Subscribe(topic, 2, messageSubHandler)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Failed to subscribe to topic: %v", token.Error())
	} else {
		log.Printf("Subscribed to topic: %s", topic)
	}
}
