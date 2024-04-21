package main

import (
	"os"
	"restaurant-management/database"
	"restaurant-management/middleware"
	"restaurant-management/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := gin.New()
	app.Use(gin.Logger())
	routes.UserRoutes(app)
	app.Use(middleware.Authentication())

	routes.FoodRoutes(app)
	routes.OrderRoutes(app)
	routes.OrderItemsRoutes(app)
	routes.InvoiceRoutes(app)
	routes.TableRoutes(app)
	routes.MenuRoutes(app)

	app.Run(":" + port)

}
