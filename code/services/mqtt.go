package services

// import (
// 	"crypto/tls"
// 	"crypto/x509"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sync/atomic"
// 	"time"

// 	"serveos-datasync/code/db"
// 	"serveos-datasync/config"

// 	MQTT "github.com/eclipse/paho.mqtt.golang"
// )

// var MqttClient MQTT.Client
// var MqttConnected int32
// var MqttDisconnectsLastHour int32
// var MqttDisconnectsLast24Hours int32

// func InitMQTT(cfg config.Config) {
// 	opts := MQTT.NewClientOptions().
// 		AddBroker(fmt.Sprintf("%s://%s:%d", cfg.ServerMQTTProtocol, cfg.ServerMQTTHost, cfg.ServerMQTTPort)).
// 		SetClientID("serveos-datasync").
// 		SetUsername(cfg.ServerMQTTUsername).
// 		SetPassword(cfg.ServerMQTTPassword).
// 		SetCleanSession(false).
// 		SetOrderMatters(false).
// 		SetKeepAlive(30).
// 		SetDefaultPublishHandler(messagePubHandler).
// 		SetOnConnectHandler(connectHandler).
// 		SetConnectionLostHandler(connectLostHandler)

// 	if cfg.ServerMQTTEnableTLS {
// 		opts.SetTLSConfig(NewTLSConfig(cfg))
// 	}

// 	MqttClient = MQTT.NewClient(opts)
// 	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
// 		log.Fatalf("Error connecting to MQTT broker: %s", token.Error())
// 	}
// }

// var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
// 	log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
// }

// var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
// 	log.Println("Connected to MQTT broker")
// 	atomic.StoreInt32(&MqttConnected, 1)
// 	subscribeToTopics()
// }

// var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
// 	log.Printf("Connection lost: %v", err)
// 	atomic.StoreInt32(&MqttConnected, 0)
// 	incrementDisconnectCounters()
// }

// func subscribeToTopics() {
// 	token := MqttClient.Subscribe("your/subscribe/topic", 2, messageSubHandler)
// 	token.Wait()
// 	if token.Error() != nil {
// 		log.Printf("Failed to subscribe to topic: %v", token.Error())
// 	} else {
// 		log.Println("Subscribed to topic: your/subscribe/topic")
// 	}
// }

// var messageSubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
// 	log.Printf("Received confirmation message: %s from topic: %s", msg.Payload(), msg.Topic())
// 	// Update the local database to mark records as processed
// 	updateRecordStatus(msg.Payload())
// }

// func updateRecordStatus(payload []byte) {
// 	// Assuming payload contains the ID of the record to be updated
// 	recordID := string(payload)
// 	_, err := db.DB.Exec("UPDATE store_cash_ups SET synced = ? WHERE id = ?", true, recordID)
// 	if err != nil {
// 		log.Printf("Failed to update record status: %v", err)
// 	} else {
// 		log.Printf("Record status updated successfully for ID: %s", recordID)
// 	}
// }

// func Publish(topic string, payload interface{}) {
// 	token := MqttClient.Publish(topic, 2, false, payload)
// 	token.Wait()
// 	if token.Error() != nil {
// 		log.Printf("Failed to publish message to topic %s: %v", topic, token.Error())
// 	} else {
// 		log.Printf("Published message to topic %s", topic)
// 	}
// }

// func NewTLSConfig(cfg config.Config) *tls.Config {
// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: !cfg.ServerMQTTValidateCert,
// 	}
// 	if cfg.ServerMQTTEnableTLS {
// 		certpool := x509.NewCertPool()
// 		pemCerts, err := os.ReadFile(cfg.ServerMQTTCAPath)
// 		if err == nil {
// 			certpool.AppendCertsFromPEM(pemCerts)
// 		}
// 		tlsConfig.RootCAs = certpool
// 	}
// 	return tlsConfig
// }

// func incrementDisconnectCounters() {
// 	atomic.AddInt32(&MqttDisconnectsLastHour, 1)
// 	atomic.AddInt32(&MqttDisconnectsLast24Hours, 1)

// 	time.AfterFunc(time.Hour, func() {
// 		atomic.AddInt32(&MqttDisconnectsLastHour, -1)
// 	})
// 	time.AfterFunc(24*time.Hour, func() {
// 		atomic.AddInt32(&MqttDisconnectsLast24Hours, -1)
// 	})
// }

// func GetMQTTConnectionStatus() string {
// 	if atomic.LoadInt32(&MqttConnected) == 1 {
// 		return "Connected"
// 	}
// 	return "Disconnected"
// }

// func GetMQTTDisconnectsLastHour() int {
// 	return int(atomic.LoadInt32(&MqttDisconnectsLastHour))
// }

// func GetMQTTDisconnectsLast24Hours() int {
// 	return int(atomic.LoadInt32(&MqttDisconnectsLast24Hours))
// }
