package service

import (
	rpcclient "github.com/asmile1559/dyshop/app/user/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/pb/frontend/home_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type GetHomepageService struct {
	ctx context.Context
}

func NewGetHomepageService(c context.Context) *GetHomepageService {
	return &GetHomepageService{ctx: c}
}

func (s *GetHomepageService) Run(_ *home_page.GetHomepageReq) gin.H {
	listProductsResp, err := rpcclient.ProductClient.ListProducts(s.ctx, &pbproduct.ListProductsReq{
		Page:         1,
		PageSize:     30,
		CategoryName: "",
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}

	products := make([]gin.H, 0)
	for _, product := range listProductsResp.Products {
		products = append(products, gin.H{
			"Id":      product.GetId(),
			"Name":    product.GetName(),
			"Picture": product.GetPicture(),
			"Price":   product.GetPrice(),
			"Sold":    "0",
		})
	}

	return gin.H{
		"resp": products,
	}

	//return gin.H{
	//	"resp": []gin.H{
	//		{
	//			"Id":      "1",
	//			"Picture": "/static/src/product/bearcookie.webp",
	//			"Name":    "超级无敌好吃的小熊饼干",
	//			"Price":   18.80,
	//			"Sold":    "200",
	//		},
	//		{
	//			"Id":      "2",
	//			"Picture": "/static/src/product/bearsweet.webp",
	//			"Name":    "超级无敌好吃的小熊软糖值得品尝大力推荐",
	//			"Price":   20.99,
	//			"Sold":    "1000",
	//		},
	//	},
	//}
}
