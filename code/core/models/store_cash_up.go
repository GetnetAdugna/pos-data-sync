package models

import "time"

type StoreCashUp struct {
	ID                 int        `db:"id"`
	TxDate             time.Time  `db:"txDate"`
	CreatedAt          time.Time  `db:"createdAt"`
	UpdatedAt          time.Time  `db:"updatedAt"`
	BrandUUID          string     `db:"brandUuid"`
	RestaurantUUID     string     `db:"restaurantUuid"`
	OperationalUnitID  string     `db:"operationalUnitId"`
	DeletedAt          *time.Time `db:"deletedAt"`
	TotalFloat         float64    `db:"totalFloat"`
	TotalCashDrop      float64    `db:"totalCashDrop"`
	GrossSales         float64    `db:"grossSales"`
	GrossSalesTax      float64    `db:"grossSalesTax"`
	TotalDiscounts     float64    `db:"totalDiscounts"`
	DiscountTax        float64    `db:"discountTax"`
	TotalOverrings     float64    `db:"totalOverrings"`
	OverringsTax       float64    `db:"overringsTax"`
	TotalRefunds       float64    `db:"totalRefunds"`
	RefundsTax         float64    `db:"refundsTax"`
	TotalDeliveryFees  float64    `db:"totalDeliveryFees"`
	DeliveryFeeTax     float64    `db:"deliveryFeeTax"`
	TotalTips          float64    `db:"totalTips"`
	TipsTax            float64    `db:"tipsTax"`
	TotalDeliveryTips  float64    `db:"totalDeliveryTips"`
	DeliveryTipsTax    float64    `db:"deliveryTipsTax"`
	NetSales           float64    `db:"netSales"`
	TotalCashRecorded  float64    `db:"totalCashRecorded"`
	TotalCardRecorded  float64    `db:"totalCardRecorded"`
	TotalBalance       float64    `db:"totalBalance"`
	ClosedBy           int        `db:"closedBy"`
	EditedBy           int        `db:"editedBy"`
	PaymentTypeDetails string     `db:"paymentTypeDetails"`
	FxRate             float64    `db:"fxRate"`
}
