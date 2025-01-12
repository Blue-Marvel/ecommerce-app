package routes

import (
	"github.com/Blue-Marvel/ecommerce-app/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/users")
	userRoutes.POST("/sign-up", controllers.SignUp())
	userRoutes.POST("/login", controllers.Login())
	userRoutes.GET("/product-view", controllers.SearchProduct())
	userRoutes.GET("/search", controllers.SearchProductByQuery())

	incomingRoutes.POST("/admin/add-product", controllers.ProductViewerAdmin())
}
