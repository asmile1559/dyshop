package product

import (
	"github.com/asmile1559/dyshop/app/frontend/middleware"
	"github.com/gin-gonic/gin"
)

func _rootMw() []gin.HandlerFunc {
	return nil
}

func _productMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{middleware.Auth()}
}

func _listProductsMw() []gin.HandlerFunc {
	return nil
}

func _getProductMw() []gin.HandlerFunc {
	return nil
}

func _searchProductsMw() []gin.HandlerFunc {
	return nil
}

func _createProductsMw() []gin.HandlerFunc {
	return nil
}

func _modifyProductsMw() []gin.HandlerFunc {
	return nil
}

func _deleteProductsMw() []gin.HandlerFunc {
	return nil
}
