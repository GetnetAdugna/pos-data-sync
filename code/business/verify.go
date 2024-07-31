package business

import (
	"database/sql"
	"encoding/json"
	"log"
	"serveos-datasync/code/db"
	"time"
)

type License struct {
	ActivationDate    string `json:"activationDate"`
	ExpiryDate        string `json:"expiryDate"`
	SalesTransactions bool   `json:"salesTransactions"`
	StockTransactions bool   `json:"stockTransactions"`
	KDS               bool   `json:"KDS"`
	Shifts            bool   `json:"shifts"`
	TillCashUp        bool   `json:"tillCashUp"`
}

func VerifyLicense() (*License, error) {
	var licenseData string
	err := db.DB.Get(&licenseData, "SELECT license FROM stores LIMIT 1")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No license found in the database")
		}
		return nil, err
	}

	var license License
	err = json.Unmarshal([]byte(licenseData), &license)
	if err != nil {
		return nil, err
	}

	// Parse dates and verify
	activationDate, err := time.Parse("02/01/2006", license.ActivationDate)
	if err != nil {
		log.Fatalf("Invalid activation date format in license")
	}

	expiryDate, err := time.Parse("02/01/2006", license.ExpiryDate)
	if err != nil {
		log.Fatalf("Invalid expiry date format in license")
	}

	if time.Now().Before(activationDate) {
		log.Fatalf("License activation date is in the future")
	}

	if time.Now().After(expiryDate) {
		log.Fatalf("License has expired")
	}

	log.Println("License verification successful")
	return &license, nil
}
