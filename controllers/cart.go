package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Blue-Marvel/ecommerce-app/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection

	userCollection *mongo.Collection
}

func NewApplication(prodCollection *mongo.Collection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Study what this checks are for
		productQueryId := c.Query("id")
		if productQueryId == "" {
			// c.Header("Content-Type", "application/json")
			_ = c.AbortWithError(http.StatusNotFound, errors.New("product id is required"))
			return
		}
		userQueryId := c.Query("user_id")
		if userQueryId == "" {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is required"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryId)

		if err != nil {
			fmt.Print(err)
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, userQueryId, productID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(http.StatusOK, "Successfully added to cart")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Study what this checks are for
		productQueryId := c.Query("id")
		if productQueryId == "" {
			// c.Header("Content-Type", "application/json")
			_ = c.AbortWithError(http.StatusNotFound, errors.New("product id is required"))
			return
		}
		userQueryId := c.Query("user_id")
		if userQueryId == "" {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is required"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryId)

		if err != nil {
			fmt.Print(err)
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = database.RemoveProductFromCart(ctx, app.prodCollection, app.userCollection, productID, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully Removed Item from cart")
	}
}
func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryId := c.Query("id")
		if userQueryId == "" {
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user_id is empty"))
		}

		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		err := database.BuyFromCart(ctx, app.userCollection, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "Successfully placed order")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Study what this checks are for
		productQueryId := c.Query("id")
		if productQueryId == "" {
			// c.Header("Content-Type", "application/json")
			_ = c.AbortWithError(http.StatusNotFound, errors.New("product id is required"))
			return
		}
		userQueryId := c.Query("user_id")
		if userQueryId == "" {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is required"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryId)

		if err != nil {
			fmt.Print(err)
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = database.InstantBuy(ctx, app.prodCollection, app.userCollection, productID, userQueryId)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully placed order")
	}
}
