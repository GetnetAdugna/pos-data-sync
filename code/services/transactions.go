package services

import (
	"log"
	"serveos-datasync/code/core/models"
	"serveos-datasync/code/db"
	"serveos-datasync/code/services/mqtt"
)

func FetchAndPublishTransactions() {
	// Fetch transactions from the database
	var transactions []models.StoreCashUp
	err := db.DB.Select(&transactions, "SELECT * FROM store_cash_ups WHERE synced = ?", false)
	if err != nil {
		log.Printf("Error fetching transactions: %v", err)
		return
	}

	for _, transaction := range transactions {
		// Publish each transaction to the MQTT broker
		mqtt.Publish("your/publish/topic", transaction)
	}
}
