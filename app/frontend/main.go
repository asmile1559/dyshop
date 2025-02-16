package main

import (
	"fmt"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/jwt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"

	bizrouter "github.com/asmile1559/dyshop/app/frontend/biz/router"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

func IndexMod(idx int, m int) int {
	return idx % m
}

func rangeSlice(s, l int) []int {
	r := make([]int, l)
	for i := 0; i < l; i++ {
		r[i] = s + i
	}
	return r
}

func showPages(curPage, totalPage interface{}) []int {

	var ci int
	var ti int
	var ok bool
	if ci, ok = curPage.(int); !ok {
		ci, _ = strconv.Atoi(curPage.(string))
	}

	if ti, ok = totalPage.(int); !ok {
		ti, _ = strconv.Atoi(totalPage.(string))
	}

	// 1, 2, 3, 4, 5, 6, 7, 8
	// 1*, 2, 3, 4, 5
	// 1, 2*, 3, 4, 5
	// 1, 2, 3*, 4, 5
	// 2, 3, 4*, 5, 6
	// 3, 4, 5*, 6, 7
	// 4, 5, 6*, 7, 8
	// 4, 5, 6, 7*, 8
	// 4, 5, 6, 7, 8*

	pageSlice := make([]int, min(ti, 5))
	if ti == 1 {
		return []int{1}
	}

	if ti < 5 {
		for i := 0; i < ti; i++ {
			pageSlice[i] = i + 1
		}
		return pageSlice
	}

	var startPage = ci - min(ci-1, max(4+ci-ti, 2))
	//if ci == ti {
	//	startPage = ci - 4
	//} else if ci+1 == ti {
	//	startPage = ci - 3
	//} else {
	//	startPage = ci - min(ci-1, 2)
	//}

	for i := 0; i < 5; i++ {
		pageSlice[i] = startPage + i
	}

	return pageSlice
}

func subOne(n int) int {
	return n - 1
}

func addOne(n int) int {
	return n + 1
}

func main() {
	rpcclient.InitRPCClient()

	router := gin.Default()
	router.Use(cors.Default())
	router.SetFuncMap(template.FuncMap{
		"realLen":    utf8.RuneCountInString,
		"iMod":       IndexMod,
		"rangeSlice": rangeSlice,
		"showPages":  showPages,
		"subOne":     subOne,
		"addOne":     addOne,
	})
	router.LoadHTMLGlob("templates/**")

	router.StaticFS("/static", http.Dir("static"))
	router.StaticFS("/test/static", http.Dir("static"))

	router.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pong.html", gin.H{
			"code": http.StatusOK,
			"host": "192.168.191.130:10166",
			"ping": "pong",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404 page not found",
		})
	})

	bizrouter.RegisterRouters(router)
	registerTestRouter(router)

	if err := router.Run(":" + viper.GetString("server.port")); err != nil {
		logrus.Fatal(err)
	}
	//router.Run(":10166")
}

func registerTestRouter(e *gin.Engine) {
	type user struct {
		UserID uint32 `json:"user_id"`
	}

	mid := func(c *gin.Context) {
		logrus.Infof("Method: %v, URI: %v", c.Request.Method, c.Request.RequestURI)
		c.Next()
	}

	_test := e.Group("/test", mid)

	_test.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": "lixiaoming",
			"img":      "/static/src/user/snake.svg",
			"categoryList": [][]string{
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
				{"食品", "酒类", "生鲜", "特产"},
				{"图书", "文娱", "教育", "电子书"},
			},
			"carousel": []map[string]string{
				{"category": "电子设备", "img": "/static/src/carousel/carousel1.jpg"},
				{"category": "彩电", "img": "/static/src/carousel/carousel2.jpg"},
				{"category": "空调", "img": "/static/src/carousel/carousel3.jpg"},
			},
			"products": []map[string]string{
				{
					"productId": "1",
					"img":       "/static/src/product/bearcookie.webp",
					"name":      "超级无敌好吃的小熊饼干",
					"price":     "18.80",
					"sold":      "200",
				},
				{
					"productId": "2",
					"img":       "/static/src/product/bearsweet.webp",
					"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"price":     "20.99",
					"sold":      "1000",
				},
				{
					"productId": "3",
					"img":       "/static/src/product/bearsweet.webp",
					"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"price":     "20.99",
					"sold":      "1000",
				},
				{
					"productId": "4",
					"img":       "/static/src/product/bearsweet.webp",
					"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"price":     "20.99",
					"sold":      "1000",
				},
				{
					"productId": "5",
					"img":       "/static/src/product/bearsweet.webp",
					"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"price":     "20.99",
					"sold":      "1000",
				},
				{
					"productId": "6",
					"img":       "/static/src/product/bearsweet.webp",
					"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"price":     "20.99",
					"sold":      "1000",
				},
				{
					"productId": "7",
					"img":       "/static/src/product/bearsweet.webp",
					"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"price":     "20.99",
					"sold":      "1000",
				},
			},
		})
	})

	{
		_upload := _test.Group("/upload")
		_upload.POST("/img", func(c *gin.Context) {
			t := c.PostForm("type")
			switch t {
			case "userImg":
				file, err := c.FormFile("userImg")
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    http.StatusBadRequest,
						"message": "upload file error.",
					})
					return
				}

				filepath := "/static/src/user/" + file.Filename
				if err := c.SaveUploadedFile(file, "."+filepath); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    http.StatusInternalServerError,
						"message": "something went wrong, please try it later.",
						"error":   err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "upload ok!",
					"url":     filepath,
				})
			}
		})
		_upload.POST("/product", func(c *gin.Context) {
			t := c.PostForm("type")
			fmt.Println(t)
			if t == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "upload file error.",
				})
				return
			}
			fmt.Printf("productName:%v\nproductShortname:%v\nproductDesc:%v\nproductCategories:%v\nproductPrice:%v\nproductStock:%v\n",
				c.PostForm("productName"),
				c.PostForm("productShortname"),
				c.PostForm("productDesc"),
				c.PostForm("productCategories"),
				c.PostForm("productPrice"),
				c.PostForm("productStock"),
			)

			if t == "updateProductWithImg" {
				file, err := c.FormFile("productImg")
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    http.StatusBadRequest,
						"message": "upload file error.",
						"error":   err.Error(),
					})
					return
				}
				filepath := "/static/src/product/" + file.Filename
				if err := c.SaveUploadedFile(file, "."+filepath); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    http.StatusInternalServerError,
						"message": "something went wrong, please try it later.",
						"error":   err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "upload ok!",
					"res": gin.H{
						"id":  "2",
						"img": filepath,
					},
				})
				return
			}
			fmt.Printf("productName:%v\nproductShortname:%v\nproductDesc:%v\nproductCategories:%v\nproductPrice:%v\nproductStock:%v\n",
				c.PostForm("productName"),
				c.PostForm("productShortname"),
				c.PostForm("productDesc"),
				c.PostForm("productCategories"),
				c.PostForm("productPrice"),
				c.PostForm("productStock"),
			)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "upload ok!",
			})
		})
	}

	{
		_user := _test.Group("/user")
		_user.GET("/", func(c *gin.Context) {
			t, _ := time.Parse("2006年1月2日", "1970年1月1日")
			_ = map[string]interface{}{
				"id":    "1",
				"name":  "张三李四",
				"phone": "86-12345678901",
				"region": map[string]string{
					"province": "北京",
					"city":     "北京市",
					"district": "海淀区",
					"street":   "知春路",
				},
				"detail": "北京北京市海淀区知春路甲48号抖音视界（北京）有限公司",
			}
			product := map[string]interface{}{
				"id":          "1",
				"name":        "超级无敌好吃的小熊饼干值得品尝大力推荐",
				"shortname":   "小熊饼干",
				"description": "真的超级无敌好吃啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
				"picture":     "/static/src/product/bearcookie.webp",
				"categories": []string{
					"食品", "玩具", "儿童",
				},
				"price": 20.99,
				"stock": 200,
			}
			c.HTML(http.StatusOK, "user.html", gin.H{
				"username":      "李小明",
				"signature":     "hi 这是李小明的主页",
				"headImg":       "/static/src/user/snake.svg",
				"gender":        "man",
				"birthday":      t,
				"userID":        "100001",
				"userRole":      "merchant",
				"userPhone":     "13245679090",
				"userEmail":     "123@abc.com",
				"defaultAddrID": "1",
				"addresses":     []map[string]interface{}{},
				"products":      []map[string]interface{}{product},
			})
		})

		_user.POST("/", func(c *gin.Context) {
			info := struct {
				Username        string `json:"userName"`
				UserSign        string `json:"userSign"`
				UserGender      string `json:"userGender"`
				UserBirthday    string `json:"userBirthday"`
				Phone           string `json:"phone"`
				Email           string `json:"email"`
				OldPassword     string `json:"old_password"`
				NewPassword     string `json:"new_password"`
				ConfirmPassword string `json:"confirm_password"`
				DeleteAccount   bool   `json:"delete_account"`
				Address         struct {
					Id         string `json:"id"`
					Name       string `json:"name"`
					Telephone  string `json:"telephone"`
					Province   string `json:"province"`
					City       string `json:"city"`
					District   string `json:"district"`
					Street     string `json:"street"`
					Detail     string `json:"detail"`
					SetDefault bool   `json:"set_default"`
					Delete     bool   `json:"delete"`
				}
			}{}

			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": err.Error(),
				})
				return
			}

			fmt.Println(info)
			if info.Address.Id == "-1" {
				c.JSON(http.StatusOK, gin.H{
					"code":      http.StatusOK,
					"message":   "new/update address ok",
					"id":        "1",
					"isDefault": true,
				})
			}

		})
		_user.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{})
		})
		_user.POST("/register", func(c *gin.Context) {
			u := struct {
				Email           string `json:"email"`
				Password        string `json:"password"`
				ConfirmPassword string `json:"confirm_password"`
			}{}
			err := c.BindJSON(&u)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "Expect json format register information",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "register ok!",
				"token":   "111",
			})
		})

		_user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})
		_user.POST("/login", func(c *gin.Context) {
			u := struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}{}
			err := c.BindJSON(&u)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "Expect json format login information",
				})
				return
			}
			token, err := jwt.GenerateJWT(uint32(1))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "Something went wrong. Please try again later.",
				})
				logrus.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "login ok!",
				"token":   token,
			})
		})

	}

	{
		_product := _test.Group("/product")
		_product.GET("/", func(c *gin.Context) {
			id := c.Query("product_id")
			if id == "2" {
				c.HTML(http.StatusOK, "product-page.html", gin.H{
					"username": "李晓明",
					"id":       id,
					"name":     "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"img":      "/static/src/product/bearsweet.webp",
					"price":    "20.99",
					"sales":    3213,
					"deliverDetail": gin.H{
						"deliver":     "顺丰速递",
						"deliverCost": "免运费",
					},
					"insureDetail": gin.H{
						"insureName": "退货宝",
						"insureDesc": "退货免邮费",
					},
					"valueItems": []map[string]interface{}{
						{"itemName": "1分软", "itemValue": "20.99"},
						{"itemName": "2分软", "itemValue": "20.99"},
						{"itemName": "3分软", "itemValue": "20.99"},
						{"itemName": "4分软", "itemValue": "20.99"},
						{"itemName": "5分软", "itemValue": "20.99"},
						{"itemName": "6分软", "itemValue": "20.99"},
						{"itemName": "7分软", "itemValue": "20.99"},
						{"itemName": "8分软", "itemValue": "20.99"},
						{"itemName": "9分软", "itemValue": "20.99"},
						{"itemName": "10分软", "itemValue": "20.99"},
					},
					"productDetails": []map[string]string{
						{"name": "味道", "value": "香橙味"},
						{"name": "材料", "value": "软糖"},
						{"name": "颜色", "value": "橘黄色"},
						{"name": "适宜人群", "value": "所有年龄人群"},
						{"name": "定位", "value": "定位"},
					},
				})
			} else if id == "1" {
				c.HTML(http.StatusOK, "product-page.html", gin.H{
					"username": "李小明",
					"id":       id,
					"name":     "超级无敌好吃的小熊饼干",
					"img":      "/static/src/product/bearcookie.webp",
					"price":    "100.00",
					"sales":    3213,
					"deliverDetail": gin.H{
						"deliver":     "顺丰速递",
						"deliverCost": "￥20",
					},
					"insureDetail": gin.H{
						"insureName": "运送宝",
						"insureDesc": "运送期间出问题赔付",
					},
					"valueItems": []map[string]string{
						{"itemName": "500g装", "itemValue": "18.8"},
						{"itemName": "1000g装", "itemValue": "36.8"},
						{"itemName": "200g装", "itemValue": "72.8"},
					},
					"productDetails": []map[string]string{
						{"name": "味道", "value": "盲盒味"},
						{"name": "材料", "value": "面粉，小麦"},
						{"name": "颜色", "value": "棕黄色"},
						{"name": "定位", "value": "高端产品"},
					},
				})
			}

		})
		_product.POST("/delete", func(c *gin.Context) {
			info := struct {
				ProductId string `json:"product_id"`
				Delete    bool   `json:"delete"`
			}{}

			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}
			if info.Delete {
				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "ok",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "not delete",
			})

		})

		_product.POST("/buyNow", func(c *gin.Context) {
			info := struct {
				ProductId   string  `json:"product_id"`
				Item        string  `json:"item"`
				SinglePrice float64 `json:"single_price"`
				Quantity    float64 `json:"quantity"`
				Price       float64 `json:"price"`
			}{}

			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":     http.StatusOK,
				"message":  "buy now ok",
				"order_id": info.ProductId,
			})
		})

		_product.POST("/add2cart", func(c *gin.Context) {
			info := struct {
				ProductId   string  `json:"product_id"`
				Item        string  `json:"item"`
				SinglePrice float64 `json:"single_price"`
				Quantity    float64 `json:"quantity"`
				Price       float64 `json:"price"`
			}{}

			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "add to cart now ok",
			})
		})

		_product.GET("/search", func(c *gin.Context) {
			kw := c.Query("keyword")
			category := c.Query("category")
			if kw == "" && category == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "must have a keyword or a category",
				})
				return
			}
			pg := c.Query("pg")
			if pg == "" {
				pg = "1"
			}
			sort := c.Query("sort")
			if sort == "" {
				sort = "comprehensive"
			}
			curPage, err := strconv.Atoi(pg)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "page must be a number",
					"error":   err.Error(),
				})
				return
			}
			if curPage <= 0 {
				curPage = 1
			}
			fmt.Println(kw)
			c.HTML(http.StatusOK, "search.html", gin.H{
				"username": "lixiaoming",
				"products": []map[string]string{
					{
						"productId": "1",
						"img":       "/static/src/product/bearcookie.webp",
						"name":      "超级无敌好吃的小熊饼干",
						"price":     "18.80",
						"sold":      "200",
					},
					{
						"productId": "2",
						"img":       "/static/src/product/bearsweet.webp",
						"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"price":     "20.99",
						"sold":      "1000",
					},
				},
				"keyword":   kw,
				"sort":      sort,
				"curPage":   curPage,
				"totalPage": 8,
			})
		})
	}

	{
		_order := _test.Group("/order")
		_order.GET("/", func(c *gin.Context) {
			id := c.Query("order_id")
			fmt.Println(id)
			c.HTML(http.StatusOK, "order.html", gin.H{
				"username": "lixiaoming",
				"addressList": map[string]interface{}{
					"defaultAddrId": "1",
					"addresses": []map[string]string{
						{
							"addrId":   "1",
							"province": "北京",
							"city":     "北京市",
							"district": "海淀区",
							"street":   "知春路",
							"detail":   "北京北京市海淀区知春路甲48号抖音视界（北京）有限公司",
							"name":     "张三李四",
							"phone":    "12345678901",
						},
						{
							"addrId":   "2",
							"province": "广东省",
							"city":     "深圳市",
							"district": "南山区",
							"street":   "海天二路",
							"detail":   "广东省深圳市南山区海天二路33号腾讯滨海大厦",
							"name":     "张三",
							"phone":    "12345678901",
						},
					},
				},
				"orderItems": []map[string]interface{}{
					{
						"productId":   "1",
						"img":         "/static/src/product/bearcookie.webp",
						"name":        "超级无敌好吃的小熊饼干",
						"valueItem":   "500g装",
						"singlePrice": "18.80",
						"quantity":    2,
						"price":       "37.60",
						"deliverCost": "10",
					},
					{
						"productId":   "2",
						"img":         "/static/src/product/bearsweet.webp",
						"name":        "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"valueItem":   "9分软",
						"singlePrice": "20.99",
						"quantity":    1,
						"price":       "20.99",
						"deliverCost": "0",
					},
				},
				"totalPrice":       "58.59",
				"totalDeliverCost": "10",
				"discount":         "0",
				"realPrice":        "68.59",
			})
		})
		_order.GET("/myorder", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/test/order/?order_id=123")
		})
		_order.POST("/cancel", func(c *gin.Context) {
			info := struct {
				OrderId string `json:"order_id"`
			}{}

			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": fmt.Sprintf("cancel order [id=%s] now ok", info.OrderId),
			})
		})

		_order.POST("/submit", func(c *gin.Context) {
			info := struct {
				OrderId          string              `json:"order_id"`
				RealPrice        float64             `json:"real_price"`
				Discount         float64             `json:"discount"`
				TotalPrice       float64             `json:"total_price"`
				TotalDeliverCost float64             `json:"total_deliver_cost"`
				Address          map[string]string   `json:"address"`
				OrderItems       []map[string]string `json:"items"`
			}{}
			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}
			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":       http.StatusOK,
				"message":    fmt.Sprintf("submit order [id=%s] now ok", info.OrderId),
				"payment_id": "123",
			})
		})
	}

	{
		_cart := _test.Group("/cart")
		_cart.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "cart.html", gin.H{
				"username": "lihua",
				"cartItems": []map[string]string{
					{
						"itemId":    "1",
						"productId": "1",
						"img":       "/static/src/product/bearcookie.webp",
						"name":      "超级无敌好吃的小熊饼干",
						"spec":      "500g装",
						"price":     "18.80",
						"quantity":  "2",
					},
					{
						"itemId":    "2",
						"productId": "2",
						"img":       "/static/src/product/bearsweet.webp",
						"name":      "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"spec":      "9分软",
						"price":     "20.99",
						"quantity":  "1",
					},
				},
			})
		})

		_cart.POST("/delete", func(c *gin.Context) {
			info := struct {
				CartItemIds []string `json:"item_ids"`
			}{}
			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}
			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": fmt.Sprintf("delete cart item [%v] now ok", info),
			})

		})
	}

	{
		_payment := _test.Group("/payment")
		_payment.GET("/", func(c *gin.Context) {
			id := c.Query("payment_id")
			c.HTML(http.StatusOK, "payment.html", gin.H{
				"paymentId": id,
				"username":  "李华",
				"orderId":   "123",
				"recipient": "张三李四",
				"phone":     "12345678901",
				"detail":    "北京北京市海淀区知春路甲48号抖音视界（北京）有限公司",
				"items": []map[string]string{
					{
						"productId":   "1",
						"name":        "超级无敌好吃的小熊饼干",
						"singlePrice": "18.80",
						"quantity":    "2",
						"price":       "47.60",
						"deliverCost": "10",
					},
					{
						"productId":   "2",
						"img":         "/static/src/product/bearsweet.webp",
						"name":        "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"valueItem":   "9分软",
						"singlePrice": "20.99",
						"quantity":    "1",
						"price":       "20.99",
						"deliverCost": "0",
					},
				},
				"totalQuantity":    "3",
				"totalDeliverCost": "10",
				"realPrice":        "68.59",
			})
		})

		_payment.POST("/", func(c *gin.Context) {
			info := struct {
				PaymentId   string `json:"payment_id"`
				CardType    string `json:"card_type"`
				CardNumber  string `json:"card_number"`
				CardHolder  string `json:"card_holder"`
				CVV         string `json:"cvv"`
				ExpireMonth string `json:"expire_month"`
				ExpireYear  string `json:"expire_year"`
				RealPrice   string `json:"real_price"`
			}{}

			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}
			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": fmt.Sprintf("submit order [id=%s] now ok", info.PaymentId),
			})
		})

		_payment.POST("/cancel", func(c *gin.Context) {
			info := struct {
				PaymentId string `json:"payment_id"`
			}{}
			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}
			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": fmt.Sprintf("cancel order [id=%s] now ok", info.PaymentId),
			})
		})
	}

	{
		_checkout := _test.Group("/checkout")
		_checkout.POST("/", func(c *gin.Context) {
			info := struct {
				TotalPrice string              `json:"total_price"`
				Items      []map[string]string `json:"items"`
			}{}
			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "failed",
					"err":     err.Error(),
				})
				return
			}
			fmt.Println(info)
			orderID := 123
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": fmt.Sprintf("checkout cart item now ok"),
				"orderId": orderID,
			})
		})
	}
}
