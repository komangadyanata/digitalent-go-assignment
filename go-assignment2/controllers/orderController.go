package controllers

import (
	"fmt"
	"go-assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrder(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Preload("Items").First(&order, "order_id = ?", id).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"result": order,
			"count":  1,
		}
		c.JSON(http.StatusOK, result)
	}

}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []models.Order
		result gin.H
	)

	err := idb.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
		c.JSON(http.StatusInternalServerError, result)
	} else {
		if len(orders) <= 0 {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": orders,
				"count":  len(orders),
			}
		}
		c.JSON(http.StatusOK, result)
	}
}

func (idb *InDB) CreateOrderAndItem(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)
	if err := c.BindJSON(&order); err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
	} else {
		err := idb.DB.Create(&order).Error
		if err != nil {
			result = gin.H{
				"result": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, result)
		} else {
			result = gin.H{
				"result": order,
			}
			c.JSON(http.StatusOK, result)
		}
	}

}

func (idb *InDB) UpdateOrderAndItem(c *gin.Context) {
	var (
		order    models.Order
		newOrder models.Order
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&order, "order_id = ?", id).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusOK, result)
	} else {
		if err := c.BindJSON(&newOrder); err != nil {
			result = gin.H{
				"result": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, result)
		} else {
			err1 := idb.DB.Save(&newOrder).Error
			err2 := idb.DB.Save(newOrder.Items).Error
			if err1 != nil || err2 != nil {
				result = gin.H{
					"result": fmt.Sprintf("update failed, error table item %s error table order %s ",
						err2.Error(), err1.Error()),
				}
				c.JSON(http.StatusInternalServerError, result)
			} else {
				result = gin.H{
					"result": fmt.Sprintf("%s %s", "update successfully", newOrder),
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}
}

func (idb *InDB) DeleteOrderAndItem(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&order, "order_id = ?", id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	} else {
		err1 := idb.DB.Where("order_id = ?", id).Delete(&models.Item{}).Error
		err2 := idb.DB.Where("order_id = ?", id).Delete(&order).Error
		if err1 != nil || err2 != nil {
			result = gin.H{
				"result": fmt.Sprintf("delete failed, error table item %s error table order %s ",
					err1.Error(), err2.Error()),
			}
		} else {
			result = gin.H{
				"result": fmt.Sprintf("delete successfully, id %s", id),
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
