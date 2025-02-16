package product

import (
	p "github.com/asmile1559/dyshop/app/frontend/biz/handler/product"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	root := e.Group("/", _rootMw()...)
	{
		_product := root.Group("/product", _productMw()...)
		//
		//_product.GET("/", )
		_product.GET("/:id", append(_getProductMw(), p.GetProduct)...)
		_product.GET("/search", append(_searchProductsMw(), p.SearchProduct)...)
		_product.POST("/create", append(_createProductsMw(), p.CreateProduct)...)
		_product.POST("/update", append(_modifyProductsMw(), p.ModifyProduct)...)
		_product.GET("/delete", append(_deleteProductsMw(), p.DeleteProduct)...)
	}
}
