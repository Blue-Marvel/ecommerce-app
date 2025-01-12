package main

import (
	"log"
	"os"

	"github.com/Blue-Marvel/ecommerce-app/controllers"
	"github.com/Blue-Marvel/ecommerce-app/database"
	"github.com/Blue-Marvel/ecommerce-app/middleware"
	"github.com/Blue-Marvel/ecommerce-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/add-to-cart", app.AddToCart())
	router.GET("/remove-item", app.RemoveItem())
	router.GET("/cart-checkout", app.BuyFromCart())
	router.GET("/instant-buy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
