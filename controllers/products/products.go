package products

import (
	"gin_todo/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetProduct - Get product by product id
func GetProduct(c *gin.Context) {

	// c.String(http.StatusNotImplemented, "Not yet implement!")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "eda",
	//	"data" : "data",
	//})

	restaurant := pkg.MasterRestaurant()
	c.JSON(200, restaurant)

}

// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implement!")
}