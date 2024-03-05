package usecases

import (
	"context"
	"orders-service/pkg/entities"
)

type AmendOrderRequest struct {
	Status       string  `json:"status" binding:"required"`
	ItemID       string  `json:"itemID" binding:"required" `
	OrderID      string  `json:"orderID" binding:"required" `
	Quantity     int     `json:"quantity" binding:"required"`
	PricePerUnit float64 `json:"pricePerUnit" binding:"required" `
}

type AmendOrderResponse struct {
	Order      entities.Order       `json:"order"`
	OrderItems []entities.OrderItem `json:"orderItem"`
}

type AmendOrderRepository interface {
	AmendOrder(context.Context, entities.Order) error
}
