package models

import "time"

type SalesHeader struct {
	TxDateTime                time.Time `db:"txDateTime"`
	CreatedAt                 time.Time `db:"createdAt"`
	UpdatedAt                 time.Time `db:"updatedAt"`
	BrandUUID                 string    `db:"brandUuid"`
	RestaurantUUID            string    `db:"restaurantUuid"`
	OperationalUnitID         string    `db:"operationalUnitId"`
	OrderNumber               int       `db:"orderNumber"`
	Status                    string    `db:"status"`
	ErrorCode                 string    `db:"errorCode"`
	ErrorMessage              string    `db:"errorMessage"`
	OrderWasted               bool      `db:"orderWasted"`
	Notes                     string    `db:"notes"`
	ScheduledTime             time.Time `db:"scheduledTime"`
	OrderSubtotal             float64   `db:"orderSubtotal"`
	OrderSubtotalTax          float64   `db:"orderSubtotalTax"`
	DeliveryFee               float64   `db:"deliveryFee"`
	DeliveryFeeTax            float64   `db:"deliveryFeeTax"`
	Tip                       float64   `db:"tip"`
	TipTax                    float64   `db:"tipTax"`
	TotalDeliveryTip          float64   `db:"totalDeliveryTip"`
	DeliveryTipTax            float64   `db:"deliveryTipTax"`
	TotalDiscounts            float64   `db:"totalDiscounts"`
	DiscountTax               float64   `db:"discountTax"`
	DiscountCode              string    `db:"discountCode"`
	DiscountDescription       string    `db:"discountDescription"`
	Total                     float64   `db:"total"`
	OrderUUID                 string    `db:"orderUuid"`
	ExternalReference         string    `db:"externalReference"`
	CreatedShiftNumber        int       `db:"createdShiftNumber"`
	CreatedUserID             int       `db:"createdUserId"`
	AuthorisedUserID          int       `db:"authorisedUserId"`
	CompletedShiftNumber      int       `db:"completedShiftNumber"`
	CompletedUserID           int       `db:"completedUserId"`
	CompletedAuthorisedUserID int       `db:"completedAuthorisedUserID"`
	DeliveryAddress           string    `db:"deliveryAddress"`
	Tendered                  float64   `db:"tendered"`
	ChangeGiven               float64   `db:"changeGiven"`
	TableNumber               string    `db:"tableNumber"`
	GuestCount                int       `db:"guestCount"`
	CustomerUUID              string    `db:"customerUuid"`
	IsManualCapture           bool      `db:"isManualCapture"`
	InvoiceNumber             string    `db:"invoiceNumber"`
	MarketingSource           string    `db:"marketingSource"`
	MarketingMedium           string    `db:"marketingMedium"`
	MarketingCampaign         string    `db:"marketingCampaign"`
	ChannelCode               string    `db:"channelCode"`
	ChannelGroupCode          string    `db:"channelGroupCode"`
	ServiceTypeCode           string    `db:"serviceTypeCode"`
	FoodCost                  float64   `db:"foodCost"`
	FxRate                    float64   `db:"fxRate"`
}
