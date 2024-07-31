package mqtt

import (
	"log"
	"sync/atomic"
	"time"

	"serveos-datasync/code/db"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	log.Println("Connected to MQTT broker")
	atomic.StoreInt32(&MqttConnected, 1)
	subscribeToTopics()
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	log.Printf("Connection lost: %v", err)
	atomic.StoreInt32(&MqttConnected, 0)
	incrementDisconnectCounters()
}

func subscribeToTopics() {
	Subscribe("your/subscribe/topic")
}

var messageSubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("Received confirmation message: %s from topic: %s", msg.Payload(), msg.Topic())
	updateRecordStatus(msg.Payload())
}

func updateRecordStatus(payload []byte) {
	recordID := string(payload)
	_, err := db.DB.Exec("UPDATE store_cash_ups SET synced = ? WHERE id = ?", true, recordID)
	if err != nil {
		log.Printf("Failed to update record status: %v", err)
	} else {
		log.Printf("Record status updated successfully for ID: %s", recordID)
	}
}

func incrementDisconnectCounters() {
	atomic.AddInt32(&MqttDisconnectsLastHour, 1)
	atomic.AddInt32(&MqttDisconnectsLast24Hours, 1)

	time.AfterFunc(time.Hour, func() {
		atomic.AddInt32(&MqttDisconnectsLastHour, -1)
	})
	time.AfterFunc(24*time.Hour, func() {
		atomic.AddInt32(&MqttDisconnectsLast24Hours, -1)
	})
}

func GetMQTTConnectionStatus() string {
	if atomic.LoadInt32(&MqttConnected) == 1 {
		return "Connected"
	}
	return "Disconnected"
}

func GetMQTTDisconnectsLastHour() int {
	return int(atomic.LoadInt32(&MqttDisconnectsLastHour))
}

func GetMQTTDisconnectsLast24Hours() int {
	return int(atomic.LoadInt32(&MqttDisconnectsLast24Hours))
}
