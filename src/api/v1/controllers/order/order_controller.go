package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-bpjs/src/api/v1/models"
	"test-bpjs/src/api/v1/usecase"
)

type OrderController struct {
	uscOrder usecase.UscOrderInterfaces
}

func OrderControllerImpl(r *gin.RouterGroup, uscOrder usecase.UscOrderInterfaces) {
	handler := &OrderController{uscOrder}

	r.POST("/", handler.InsertOrder)
	r.GET("/", handler.FindAllOrder)
	r.PUT("/:id", handler.UpdateOrder)
	r.DELETE("/:id", handler.DeleteOrder)
}

func (i *OrderController) InsertOrder(c *gin.Context) {
	var orderReq models.ReqOrder

	if err := c.BindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := i.uscOrder.CreateOrder(orderReq.RequestID, orderReq.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders created successfully"})
}

func (i *OrderController) UpdateOrder(c *gin.Context) {
	var orderReq models.ReqOrder

	if err := c.BindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := i.uscOrder.UpdateOrders(c.Param("id"), orderReq.RequestID, orderReq.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders saved successfully"})
}

func (i *OrderController) FindAllOrder(c *gin.Context) {
	orders, err := i.uscOrder.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (i *OrderController) DeleteOrder(c *gin.Context) {

	err := i.uscOrder.DeleteOrder(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders deleted successfully"})
}
