package main

import (
	"go-assignment2/configs"
	"go-assignment2/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs.StartDB()

	db := configs.GetDB()
	inDB := &controllers.InDB{DB: db}

	var PORT = ":8080"

	router := gin.Default()
	router.GET("/orders", inDB.GetOrders)
	router.GET("/orders/:id", inDB.GetOrder)
	router.POST("/orders", inDB.CreateOrderAndItem)
	router.PUT("/orders/:id", inDB.UpdateOrderAndItem)
	router.DELETE("/orders/:id", inDB.DeleteOrderAndItem)

	log.Fatal(http.ListenAndServe(PORT, router))

}
