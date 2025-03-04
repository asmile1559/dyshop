package service

import (
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/pb/frontend/home_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type GetShowcaseService struct {
	ctx context.Context
}

func NewGetShowcaseService(c context.Context) *GetShowcaseService {
	return &GetShowcaseService{ctx: c}
}

func (s *GetShowcaseService) Run(req *home_page.GetShowcaseReq) []gin.H {
	productClient, conn, err := rpcclient.GetProductClient()
	if err != nil {
		logrus.WithError(err).Error("failed to create product client")
		return nil
	}
	defer conn.Close()
	listProductsResp, err := productClient.ListProducts(s.ctx, &pbproduct.ListProductsReq{
		Page:         1,
		PageSize:     30,
		CategoryName: req.Which,
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}

	products := make([]gin.H, 0)
	for _, product := range listProductsResp.Products {
		products = append(products, gin.H{
			"Id":      product.Id,
			"Name":    product.Name,
			"Picture": product.Picture,
			"Price":   product.Price,
			"Sold":    "1k",
		})
	}

	return products
}
