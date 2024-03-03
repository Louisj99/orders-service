package usecases

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"orders-service/pkg/entities"
)

type OrderCreator interface {
	CreateOrder(ctx context.Context, order entities.Order) error
	CreateItems(ctx context.Context, items []entities.OrderItem) error
	GetItems(ctx context.Context, ItemID []string) ([]entities.OrderItem, error)
}

func CreateOrder(orderCreator OrderCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order entities.Order
		err := c.BindJSON(&order)
		if err != nil {
			c.JSON(400, "error bad request")
			return
		}
		orderID := uuid.New()
		order.ID = orderID.String()
		err = orderCreator.CreateOrder(c, order)
		if err != nil {
			c.JSON(500, "error creating order")
			return
		}
		c.JSON(200, gin.H{"message": "Order Created"})
	}
}
