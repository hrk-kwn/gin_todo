package app

import (
	"gin_todo/controllers/products"
)

func mapUrls() {

	// controllers products 呼び出し
	r.GET("/products/:product_id", products.GetProduct)
	r.POST("/products", products.CreateProduct)


}