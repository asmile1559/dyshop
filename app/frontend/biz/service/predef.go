package service

import "github.com/gin-gonic/gin"

var (
	PageRouter = gin.H{
		"HomePage":     "/",
		"LoginPage":    "/user/login/",
		"RegisterPage": "/user/register/",
		"UserPage":     "/user/",
		"ProductPage":  "/product/",
		"OrderPage":    "/order/",
		"CartPage":     "/cart/",
		"SearchPage":   "/search/",
		"PaymentPage":  "/payment/",
	}
	CategoryList = [][]string{
		{"家用电器"},
		{"手机", "运营商", "数码"},
		{"电脑", "办公", "文具用品"},
		{"家居", "家具", "家装", "厨具"},
		{"男装", "女装", "童装", "内衣"},
		{"美妆", "个护清洁", "宠物"},
		{"女鞋", "箱包", "钟表", "珠宝"},
		{"男鞋", "运动", "户外"},
		{"房产", "汽车", "汽车用品"},
		{"母婴", "玩具乐器"},
		{"母婴", "玩具乐器"},
		{"食品", "酒类", "生鲜", "特产"},
		{"图书", "文娱", "教育", "电子书"},
	}
	Carousels = []gin.H{
		{"Category": "电子设备", "Img": "/static/src/carousel/carousel1.jpg"},
		{"Category": "彩电", "Img": "/static/src/carousel/carousel2.jpg"},
		{"Category": "空调", "Img": "/static/src/carousel/carousel3.jpg"},
	}
)
