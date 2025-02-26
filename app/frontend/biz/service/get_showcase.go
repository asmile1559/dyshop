package service

import (
	"github.com/asmile1559/dyshop/pb/frontend/home_page"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type GetShowcaseService struct {
	ctx context.Context
}

func NewGetShowcaseService(c context.Context) *GetShowcaseService {
	return &GetShowcaseService{ctx: c}
}

func (s *GetShowcaseService) Run(req *home_page.GetShowcaseReq) gin.H {
	//listProductsResp, err := rpcclient.ProductClient.ListProducts(s.ctx, &pbproduct.ListProductsReq{
	//	Page:         1,
	//	PageSize:     30,
	//	CategoryName: req.Which,
	//})
	//if err != nil {
	//	logrus.Error(err)
	//	return nil
	//}
	//
	//products := make([]gin.H, 0)
	//for _, product := range listProductsResp.Products {
	//	products = append(products, gin.H{
	//		"Id":      product.GetId(),
	//		"Name":    product.GetName(),
	//		"Picture": product.GetPicture(),
	//		"Price":   product.GetPrice(),
	//		"Sold":    "0",
	//	})
	//}
	//
	//return gin.H{
	//	"resp": products,
	//}
	var products []gin.H
	if req.GetWhich() == "hot" || req.GetWhich() == "discount" {
		products = []gin.H{
			{
				"Id":      "2",
				"Picture": "/static/src/product/bearsweet.webp",
				"Name":    "超级无敌好吃的小熊软糖值得品尝大力推荐",
				"Price":   20.99,
				"Sold":    "1000",
			},
			{
				"Id":      "1",
				"Picture": "/static/src/product/bearcookie.webp",
				"Name":    "超级无敌好吃的小熊饼干",
				"Price":   18.80,
				"Sold":    "200",
			},
		}
	} else {
		products = []gin.H{
			{
				"Id":      "1",
				"Picture": "/static/src/product/bearcookie.webp",
				"Name":    "超级无敌好吃的小熊饼干",
				"Price":   18.80,
				"Sold":    "200",
			},
			{
				"Id":      "2",
				"Picture": "/static/src/product/bearsweet.webp",
				"Name":    "超级无敌好吃的小熊软糖值得品尝大力推荐",
				"Price":   20.99,
				"Sold":    "1000",
			},
		}
	}
	return gin.H{
		"resp": products,
	}
}
