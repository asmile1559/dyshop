package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

func registerExampleRouter(e *gin.Engine) {
	auth := func(c *gin.Context) {
		logrus.Infof("Method: %v, URI: %v", c.Request.Method, c.Request.RequestURI)
		// 一般在auth中鉴权, 并将用户id加入到context中
		// 所以在执行用户操作时,不需要传递用户参数
		c.Set("user_id", int64(1))
		c.Next()
	}

	_example := e.Group("/example", auth)

	// Router: /example
	{
		// GET /example
		//
		//	获取主页
		_example.GET("/", func(c *gin.Context) {
			// 根据是否有user_id判断是否登录
			// 提供UserInfo时, 是已登录页面, 否则为未登录页面
			resp := gin.H{
				"PageRouter":   PageRouter,
				"CategoryList": CategoryList,
				"Carousels":    Carousels,
				"Products": []gin.H{
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
				},
			}
			c.HTML(http.StatusOK, "index.html", &resp)
		})

		// POST /example/verify
		_example.POST("/verify", func(c *gin.Context) {
			// TODO:
			//
			//	获得前端token, 验证并返回结果
			req := struct {
				Token string `json:"token"`
			}{}

			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "Expect json format register information",
				})
				return
			}

			if req.Token == "123" {
				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "verify ok!",
					"resp": gin.H{
						"ok":   true,
						"Id":   "1",
						"Name": "lixiaoming",
						"Img":  "/static/src/user/snake.svg",
					},
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "verify no!",
					"resp": gin.H{
						"ok": false,
					},
				})
			}

		})
		// GET /example/showcase
		//
		//	获取展示橱窗的内容
		_example.GET("/showcase", func(c *gin.Context) {
			sub := c.DefaultQuery("sub", "hot")
			var products []gin.H
			if sub == "new" || sub == "seckill" {
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
			c.JSON(http.StatusOK, gin.H{
				"code":     http.StatusOK,
				"message":  fmt.Sprintf("获取新橱窗成功, 新的橱窗类别为: %v", sub),
				"products": products,
			})
		})
	}

	// Router: /example/user
	{
		_user := _example.Group("/user")

		// GET: /example/user/register
		// 获取注册界面
		_user.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"PageRouter": PageRouter,
			})
		})

		// POST: /example/user/register
		// 用户注册
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
			})
		})

		// GET: /example/user/login
		// 获取登录界面
		_user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"PageRouter": PageRouter,
			})
		})

		// POST: /example/user/login
		// 用户登录
		_user.POST("/login", func(c *gin.Context) {
			u := struct {
				Email    string `json:"email"`
				Password string `json:"password"`
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
				"message": "login ok!",
				"token":   "123",
			})
		})

		// GET: /example/user/info
		// 获取用户信息
		_user.GET("/info", func(c *gin.Context) {
			if userId, ok := c.Get("user_id"); ok {
				go func(id int64) {
					fmt.Println("模拟通过user id进行查询或者其他的操作, user_id =", id)
				}(userId.(int64))
				t, _ := time.Parse("2006年1月2日", "1970年1月1日")
				c.HTML(http.StatusOK, "info.html", gin.H{
					"PageRouter": PageRouter,
					"UserInfo": gin.H{
						"Id":       userId,
						"Name":     "lixiaoming",
						"Sign":     "许仙给老婆买了一顶帽子，白娘子戴上之后就死了，因为那是顶鸭（压）舌（蛇）帽。",
						"Img":      "/static/src/user/snake.svg",
						"Role":     []string{"user", "merchant"},
						"Gender":   "male",
						"Birthday": t,
					},
				})
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "no user id error",
			})
		})

		// POST /example/user/info
		// 修改用户文字信息
		_user.POST("/info", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			info := struct {
				Name     string `json:"name"`
				Sign     string `json:"sign"`
				Gender   string `json:"gender"`
				Birthday string `json:"birthday"`
			}{}
			err := c.BindJSON(&info)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "update bad info",
					"err":     err.Error(),
				})
				return
			}
			fmt.Println(info)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "update info ok",
				"resp": gin.H{
					"id":   userId,
					"name": info.Name,
					"sign": info.Sign,
				},
			})
		})

		// POST /example/user/info/upload
		// 修改用户图片信息
		_user.POST("/info/upload", func(c *gin.Context) {
			if userId, ok := c.Get("user_id"); ok {
				file, err := c.FormFile("Img")
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    http.StatusBadRequest,
						"message": "upload file error.",
						"error":   err.Error(),
					})
					return
				}

				fileDir := fmt.Sprintf("/static/src/user/%v/", userId)
				saveDir := "." + fileDir
				if _, err := os.Stat(saveDir); os.IsNotExist(err) {
					err = os.Mkdir(saveDir, 0755)
					if err != nil {
						logrus.Error(err)
						c.JSON(http.StatusInternalServerError, gin.H{
							"code":    http.StatusInternalServerError,
							"message": "something went wrong, please try it later.",
							"error":   err.Error(),
						})
						return
					}
				}

				filePath := fileDir + file.Filename
				savePath := saveDir + file.Filename
				if err := c.SaveUploadedFile(file, savePath); err != nil {
					logrus.Error(err)
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
					"resp": gin.H{
						"user_id": userId,
						"url":     filePath,
					},
				})
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "no user id error",
			})
		})

		// GET: /example/user/account
		// 获取用户账户
		_user.GET("/account", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "Unauthorized",
				})
			}
			c.HTML(http.StatusOK, "account.html", gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Id":    userId.(int64),
					"Name":  "lixiaoming",
					"Sign":  "许仙给老婆买了一顶帽子，白娘子戴上之后就死了，因为那是顶鸭（压）舌（蛇）帽。",
					"Img":   "/static/src/user/snake.svg",
					"Role":  []string{"user"},
					"Phone": "13245678901",
					"Email": "123@abc.com",
				},
			})
			return
		})

		// POST /example/user/account
		// 修改用户账户信息
		_user.POST("/account", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				Phone           string `json:"phone"`
				Email           string `json:"email"`
				Password        string `json:"password"`
				NewPassword     string `json:"new_password"`
				ConfirmPassword string `json:"confirm_password"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "update bad info",
					"err":     err.Error(),
				})
				return
			}
			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "update info ok",
				"resp": gin.H{
					"id":    userId,
					"phone": req.Phone,
					"email": req.Email,
				},
			})
		})

		// GET /example/user/account/delete
		// 删除账户
		_user.GET("/account/delete", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id": userId,
				},
			})
		})

		// GET /example/user/role/merchant
		// 注册成为商户
		_user.GET("/role/merchant", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			fmt.Println(userId)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "update info ok",
				"resp": gin.H{
					"user_id": userId,
				},
			})
		})

		// GET /example/user/address
		// 获取用户的地址信息
		_user.GET("/address", func(c *gin.Context) {
			if userId, ok := c.Get("user_id"); ok {
				go func(id int64) {
					fmt.Println("模拟通过user id进行查询或者其他的操作, user_id =", id)
				}(userId.(int64))

				resp := gin.H{
					"PageRouter": PageRouter,
					"UserInfo": gin.H{
						"Id":   userId,
						"Name": "lixiaoming",
						"Sign": "许仙给老婆买了一顶帽子，白娘子戴上之后就死了，因为那是顶鸭（压）舌（蛇）帽。",
						"Img":  "/static/src/user/snake.svg",
						"Role": []string{"user", "merchant"},
					},
					"AddressInfo": gin.H{
						"Default": "1",
						"Addresses": []gin.H{
							{
								"AddressId":   "1",
								"Recipient":   "张三李四",
								"Phone":       "12345678901",
								"Province":    "中国",
								"City":        "北京市",
								"District":    "海淀区",
								"Street":      "知春路",
								"FullAddress": "北京北京市海淀区知春路甲48号抖音视界"},
							{
								"AddressId":   "2",
								"Recipient":   "张三",
								"Phone":       "12345678901",
								"Province":    "广东省",
								"City":        "深圳市",
								"District":    "南山区",
								"Street":      "海天二路",
								"FullAddress": "广东省深圳市南山区海天二路33号腾讯滨海大厦"},
						},
					},
				}
				c.HTML(http.StatusOK, "address.html", resp)
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "no user id error",
			})
		})

		// POST /example/user/address
		// 修改用户的地址信息 添加一条地址
		_user.POST("/address", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				AddressId   string `json:"address_id"`
				Recipient   string `json:"recipient"`
				Phone       string `json:"phone"`
				Province    string `json:"province"`
				City        string `json:"city"`
				District    string `json:"district"`
				Street      string `json:"street"`
				FullAddress string `json:"full_address"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "update bad info",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			if req.AddressId == "-1" {
				req.AddressId = "3"
				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "del account ok",
					"resp": gin.H{
						"user_id":    userId,
						"is_default": false,
						"address":    req,
					},
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "del account ok",
					"resp": gin.H{
						"user_id":    userId,
						"address_id": req.AddressId,
					},
				})
			}
			return
		})

		// POST /example/user/address/setDefault
		//
		//	将一个地址设置为默认地址
		_user.POST("/address/setDefault", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}

			req := struct {
				AddressId string `json:"address_id"`
			}{}

			if err := c.BindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "update bad info",
					"err":     err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":    userId,
					"address_id": req.AddressId,
				},
			})
		})

		// POST /example/user/address/delete
		//
		//	删除一个地址
		_user.POST("/address/delete", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}

			req := struct {
				AddressId string `json:"address_id"`
			}{}

			if err := c.BindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "update bad info",
					"err":     err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":    userId,
					"address_id": req.AddressId,
				},
			})
		})

		// GET /example/user/product
		// 获取用户发布的商品
		_user.GET("/product", func(c *gin.Context) {
			userId, _ := c.Get("user_id")
			go func(id int64) {
				fmt.Println("模拟通过user id进行查询或者其他的操作, user_id =", id)
			}(userId.(int64))

			resp := gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Id":   userId,
					"Name": "lixiaoming",
					"Sign": "许仙给老婆买了一顶帽子，白娘子戴上之后就死了，因为那是顶鸭（压）舌（蛇）帽。",
					"Img":  "/static/src/user/snake.svg",
					"Role": []string{"user", "merchant"},
				},
				"Products": []gin.H{
					{
						"ProductId":   "1",
						"ProductImg":  "/static/src/product/bearcookie.webp",
						"ProductName": "超级无敌好吃的小熊饼干",
						"ProductDesc": "体验前所未有的美味享受，尽在超级无敌好吃的小熊饼干！这款饼干不仅外形可爱，口感更是令人难以抗拒。每一口都充满了浓郁的香甜味道，让你在忙碌的生活中找到一丝甜蜜的慰藉。",
						"ProductSpecs": []gin.H{
							{"SpecName": "500g装", "SpecPrice": "18.80", "SpecStock": "120"},
							{"SpecName": "1000g装", "SpecPrice": "36.80", "SpecStock": "130"},
							{"SpecName": "2000g装", "SpecPrice": "72.80", "SpecStock": "140"},
						},
						"ProductCategories": []string{"食品", "儿童", "饼干"},
						"ProductParams": []gin.H{
							{"ParamName": "味道", "ParamValue": "盲盒味"},
							{"ParamName": "材料", "ParamValue": "面粉，小麦"},
							{"ParamName": "颜色", "ParamValue": "棕黄色"},
							{"ParamName": "适宜人群", "ParamValue": "所有年龄人群"},
							{"ParamName": "定位", "ParamValue": "高端产品"},
						},
						"ProductInsurance": "退货险",
						"ProductExpress":   "包邮",
					},
					{
						"ProductId":   "2",
						"ProductImg":  "/static/src/product/bearsweet.webp",
						"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"ProductDesc": "如果你正在寻找一款既美味又有趣的零食，那么超级无敌好吃的小熊软糖绝对是你的不二之选！这款软糖不仅外形可爱，口感更是令人难以忘怀。每一颗小熊软糖都充满了浓郁的果味，让你在每一次品尝中都能感受到甜蜜的幸福。",
						"ProductSpecs": []gin.H{
							{"SpecName": "1分软", "SpecPrice": "20.99", "SpecStock": "2000"},
							{"SpecName": "3分软", "SpecPrice": "20.99", "SpecStock": "1000"},
							{"SpecName": "5分软", "SpecPrice": "20.99", "SpecStock": "5000"},
							{"SpecName": "7分软", "SpecPrice": "20.99", "SpecStock": "670"},
							{"SpecName": "9分软", "SpecPrice": "20.99", "SpecStock": "280"},
						},
						"ProductCategories": []string{"食品", "儿童", "软糖"},
						"ProductParams": []gin.H{
							{"ParamName": "味道", "ParamValue": "香橙味"},
							{"ParamName": "材料", "ParamValue": "软糖"},
							{"ParamName": "颜色", "ParamValue": "橘黄色"},
							{"ParamName": "适宜人群", "ParamValue": "所有年龄人群"},
							{"ParamName": "定位", "ParamValue": "定位"},
						},
						"ProductInsurance": "运输险",
						"ProductExpress":   "到付",
					},
				},
				"NoImg":             "/static/src/basic/noimg.svg",
				"CategoriesOptions": []string{"服装", "鞋子", "儿童", "食品", "饼干", "软糖"},
			}

			c.HTML(http.StatusOK, "product-mana.html", &resp)

		})

		// POST /example/user/product
		// 修改用户发布的商品
		_user.POST("/product", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}

			req := struct {
				ProductId         string           `json:"product_id"`
				ProductImg        string           `json:"product_img"`
				ProductName       string           `json:"product_name"`
				ProductDesc       string           `json:"product_desc"`
				ProductSold       string           `json:"product_sold"`
				ProductSpecs      []map[string]any `json:"product_specs"`
				ProductCategories []string         `json:"product_categories"`
				ProductParams     []map[string]any `json:"product_params"`
				ProductInsurance  string           `json:"product_insurance"`
				ProductExpress    string           `json:"product_express"`
			}{}

			err := json.Unmarshal([]byte(c.PostForm("product")), &req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "upload file error. 1",
					"error":   err.Error(),
				})
				return
			}
			if req.ProductImg == "" {
				file, err := c.FormFile("image")
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    http.StatusBadRequest,
						"message": "upload file error. 2",
						"error":   err.Error(),
					})
					return
				}

				fileDir := fmt.Sprintf("/static/src/product/%v/", req.ProductId)
				saveDir := "." + fileDir
				if _, err := os.Stat(saveDir); os.IsNotExist(err) {
					err = os.Mkdir(saveDir, 0755)
					if err != nil {
						logrus.Error(err)
						c.JSON(http.StatusInternalServerError, gin.H{
							"code":    http.StatusInternalServerError,
							"message": "something went wrong, please try it later.",
							"error":   err.Error(),
						})
						return
					}
				}
				filePath := fileDir + file.Filename
				savePath := saveDir + file.Filename
				if err := c.SaveUploadedFile(file, savePath); err != nil {
					logrus.Error(err)
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    http.StatusInternalServerError,
						"message": "something went wrong, please try it later.",
						"error":   err.Error(),
					})
					return
				}
				req.ProductImg = filePath
			}

			if req.ProductId == "0" {
				req.ProductId = "3"
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "upload ok!",
				"resp": gin.H{
					"user_id": userId,
					"product": req,
				},
			})
			return
		})

		// POST /example/user/product/delete
		// 删除一个用户发布的商品
		_user.POST("/product/delete", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				ProductId string `json:"product_id"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "update bad info",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":    userId,
					"product_id": &req.ProductId,
				},
			})
		})
	}

	// Router: /example/product
	{
		_product := _example.Group("/product")

		// GET /example/product
		// 获取一个商品页
		_product.GET("/", func(c *gin.Context) {
			userId, _ := c.Get("user_id")
			productId := c.Query("product_id")
			resp := gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Id":   userId,
					"Name": "lixiaoming",
					"Sign": "许仙给老婆买了一顶帽子，白娘子戴上之后就死了，因为那是顶鸭（压）舌（蛇）帽。",
					"Img":  "/static/src/user/snake.svg",
					"Role": []string{"user", "merchant"},
				},
			}

			if productId == "1" {
				resp["Product"] = gin.H{
					"ProductId":   "1",
					"ProductImg":  "/static/src/product/bearcookie.webp",
					"ProductName": "超级无敌好吃的小熊饼干",
					"ProductDesc": "体验前所未有的美味享受，尽在超级无敌好吃的小熊饼干！这款饼干不仅外形可爱，口感更是令人难以抗拒。每一口都充满了浓郁的香甜味道，让你在忙碌的生活中找到一丝甜蜜的慰藉。",
					"ProductSpecs": []gin.H{
						{"SpecName": "500g装", "SpecPrice": "18.80", "SpecStock": "120"},
						{"SpecName": "1000g装", "SpecPrice": "36.80", "SpecStock": "130"},
						{"SpecName": "2000g装", "SpecPrice": "72.80", "SpecStock": "140"},
					},
					"ProductCategories": []string{"食品", "儿童", "饼干"},
					"ProductParams": []gin.H{
						{"ParamName": "味道", "ParamValue": "盲盒味"},
						{"ParamName": "材料", "ParamValue": "面粉，小麦"},
						{"ParamName": "颜色", "ParamValue": "棕黄色"},
						{"ParamName": "适宜人群", "ParamValue": "所有年龄人群"},
						{"ParamName": "定位", "ParamValue": "高端产品"},
					},
					"ProductInsurance": "退货险",
					"ProductExpress":   "包邮",
				}

			} else if productId == "2" {
				resp["Product"] = gin.H{
					"ProductId":   "2",
					"ProductImg":  "/static/src/product/bearsweet.webp",
					"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
					"ProductDesc": "如果你正在寻找一款既美味又有趣的零食，那么超级无敌好吃的小熊软糖绝对是你的不二之选！这款软糖不仅外形可爱，口感更是令人难以忘怀。每一颗小熊软糖都充满了浓郁的果味，让你在每一次品尝中都能感受到甜蜜的幸福。",
					"ProductSpecs": []gin.H{
						{"SpecName": "1分软", "SpecPrice": "20.99", "SpecStock": "2000"},
						{"SpecName": "3分软", "SpecPrice": "20.99", "SpecStock": "1000"},
						{"SpecName": "5分软", "SpecPrice": "20.99", "SpecStock": "5000"},
						{"SpecName": "7分软", "SpecPrice": "20.99", "SpecStock": "670"},
						{"SpecName": "9分软", "SpecPrice": "20.99", "SpecStock": "280"},
					},
					"ProductCategories": []string{"食品", "儿童", "软糖"},
					"ProductParams": []gin.H{
						{"ParamName": "味道", "ParamValue": "香橙味"},
						{"ParamName": "材料", "ParamValue": "软糖"},
						{"ParamName": "颜色", "ParamValue": "橘黄色"},
						{"ParamName": "适宜人群", "ParamValue": "所有年龄人群"},
						{"ParamName": "定位", "ParamValue": "定位"},
					},
					"ProductInsurance": "运输险",
					"ProductExpress":   "到付",
				}
			}
			fmt.Println(resp)
			c.HTML(http.StatusOK, "product-page.html", resp)
		})

		// POST /example/product/buy
		// 购买商品
		_product.POST("/buy", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				ProductId   string `json:"product_id"`
				ProductSpec gin.H  `json:"product_spec"`
				Quantity    string `json:"quantity"`
				OrderPrice  string `json:"order_price"`
				Postage     string `json:"postage"`
				Currency    string `json:"currency"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":  userId,
					"order_id": 123,
				},
			})
		})

		// POST /example/product/add2cart
		//
		//	添加购物车
		_product.POST("/add2cart", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				ProductId   string `json:"product_id"`
				ProductSpec gin.H  `json:"product_spec"`
				Quantity    string `json:"quantity"`
				Postage     string `json:"postage"`
				Currency    string `json:"currency"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id": userId,
				},
			})
		})
	}

	// Router: /example/cart
	{
		_cart := _example.Group("/cart")

		// GET /example/cart
		// 获取用户的购物车
		_cart.GET("/", func(c *gin.Context) {
			userId, _ := c.Get("user_id")
			resp := gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Id":   userId,
					"Name": "lixiaoming",
				},
				"CartItems": []gin.H{
					{
						"ItemId":      "1",
						"ProductId":   "1",
						"ProductImg":  "/static/src/product/bearcookie.webp",
						"ProductName": "超级无敌好吃的小熊饼干",
						"ProductSpec": gin.H{
							"Name":  "500g装",
							"Price": "18.80",
						},
						"Quantity": "2",
						"Postage":  "10",
					},
					{
						"ItemId":      "2",
						"productId":   "2",
						"ProductImg":  "/static/src/product/bearsweet.webp",
						"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"ProductSpec": gin.H{
							"Name":  "9分软",
							"Price": "20.99",
						},
						"Quantity": "1",
						"Postage":  "0",
					},
				},
			}
			c.HTML(http.StatusOK, "cart.html", resp)
		})

		// POST /example/cart/delete
		// 删除购物车
		_cart.POST("/delete", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				ItemIds []string `json:"item_ids"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id": userId,
				},
			})
		})

		// POST /example/cart/checkout
		// 结算商品, 生成订单
		_cart.POST("/checkout", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				OrderPrice string  `json:"order_price"`
				CartItems  []gin.H `json:"cart_items"`
			}{}

			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":  userId,
					"order_id": 123,
					"item_ids": func() []interface{} {
						s := make([]interface{}, 0)
						for _, o := range req.CartItems {
							s = append(s, o["item_id"])
						}
						return s
					}(),
				},
			})
		})
	}

	// Router: order
	{
		_order := _example.Group("/order")

		// GET /example/order
		// 获得订单
		_order.GET("/", func(c *gin.Context) {
			orderId := c.Query("order_id")
			resp := gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Name": "lixiaoming",
				},
				"AddressInfo": gin.H{
					"Default": "1",
					"Addresses": []gin.H{
						{
							"AddressId":   "1",
							"Recipient":   "张三李四",
							"Phone":       "12345678901",
							"Province":    "中国",
							"City":        "北京市",
							"District":    "海淀区",
							"Street":      "知春路",
							"FullAddress": "北京北京市海淀区知春路甲48号抖音视界",
						},
						{
							"AddressId":   "2",
							"Recipient":   "张三",
							"Phone":       "12345678901",
							"Province":    "广东省",
							"City":        "深圳市",
							"District":    "南山区",
							"Street":      "海天二路",
							"FullAddress": "广东省深圳市南山区海天二路33号腾讯滨海大厦"},
					},
				},
				"Products": []gin.H{
					{
						"ProductId":   "1",
						"ProductImg":  "/static/src/product/bearcookie.webp",
						"ProductName": "超级无敌好吃的小熊饼干",
						"ProductSpec": gin.H{
							"Name":  "500g装",
							"Price": "18.80",
						},
						"Quantity": "2",
						"Currency": "CNY",
						"Postage":  "10.00",
					},
					{
						"ItemId":      "2",
						"productId":   "2",
						"ProductImg":  "/static/src/product/bearsweet.webp",
						"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"ProductSpec": gin.H{
							"Name":  "9分软",
							"Price": "20.99",
						},
						"Quantity": "1",
						"Postage":  "0",
					},
				},
				"OrderPrice":      "58.59",
				"OrderPostage":    "10.00",
				"OrderDiscount":   "0",
				"OrderFinalPrice": "68.59",
			}

			fmt.Println(orderId)
			fmt.Println(resp)
			c.HTML(http.StatusOK, "order.html", resp)
		})

		// POST /example/order/submit
		// 用户提交订单
		_order.POST("/submit", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				OrderId         string  `json:"order_id"`
				Address         gin.H   `json:"address"`
				Products        []gin.H `json:"products"`
				Discount        string  `json:"discount"`
				OrderPrice      string  `json:"order_price"`
				OrderPostage    string  `json:"order_postage"`
				OrderFinalPrice string  `json:"order_final_price"`
			}{}

			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":        userId,
					"order_id":       req.OrderId,
					"transaction_id": 123,
				},
			})
		})

		// POST /example/order/cancel
		// 用户取消订单
		_order.POST("/cancel", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				OrderId string `json:"order_id"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":  userId,
					"order_id": req.OrderId,
				},
			})
		})
	}
	// Router /example/checkout
	{
		_checkout := _example.Group("/checkout")
		// GET /example/checkout?transaction_id=&order_id=
		// 进行结算
		_checkout.GET("/", func(c *gin.Context) {
			transactionId := c.Query("transaction_id")
			orderId := c.Query("order_id")

			fmt.Println(transactionId, ",", orderId)

			resp := gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Name": "lixiaoming",
				},
				"OrderId": orderId,
				"Address": gin.H{
					"Recipient":   "张三李四",
					"Phone":       "12345678901",
					"Province":    "北京",
					"City":        "北京市",
					"District":    "海淀区",
					"Street":      "知春路",
					"FullAddress": "北京北京市海淀区知春路甲48号抖音视界",
				},
				"Products": []gin.H{
					{
						"ProductId":   "1",
						"ProductImg":  "/static/src/product/bearcookie.webp",
						"ProductName": "超级无敌好吃的小熊饼干",
						"ProductSpec": gin.H{
							"Name":  "500g装",
							"Price": "18.80",
						},
						"Quantity": "2",
						"Currency": "CNY",
						"Postage":  "10.00",
					},
					{
						"ItemId":      "2",
						"productId":   "2",
						"ProductImg":  "/static/src/product/bearsweet.webp",
						"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"ProductSpec": gin.H{
							"Name":  "9分软",
							"Price": "20.99",
						},
						"Quantity": "1",
						"Postage":  "0",
					},
				},
				"OrderQuantity":   "3",
				"OrderPostage":    "10.00",
				"OrderPrice":      "58.59",
				"OrderFinalPrice": "68.59",
			}
			c.HTML(http.StatusOK, "checkout.html", resp)
		})

		// GET /example/checkout/cancel
		// 取消结算
		_checkout.POST("/cancel", func(c *gin.Context) {
			userId, ok := c.Get("user_id")
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "no user id error",
				})
				return
			}
			req := struct {
				OrderId       string `json:"order_id"`
				TransactionId string `json:"transaction_id"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"user_id":        userId,
					"order_id":       req.OrderId,
					"transaction_id": req.TransactionId,
				},
			})
		})
	}

	// Router /example/payment
	{
		_payment := _example.Group("/payment")

		// POST /example/payment
		// 进行支付
		_payment.POST("/", func(c *gin.Context) {
			req := struct {
				TransactionId string `json:"transaction_id"`
				CreditCard    gin.H  `json:"credit_card"`
				FinalPrice    string `json:"final_price"`
			}{}
			err := c.BindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "StatusBadRequest",
					"err":     err.Error(),
				})
				return
			}

			fmt.Println(req)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "del account ok",
				"resp": gin.H{
					"transaction_id": req.TransactionId,
				},
			})

		})
	}

	// Router: /example/search
	{
		_search := _example.Group("/search")

		// GET /example/search
		// 查询商品
		_search.GET("/", func(c *gin.Context) {
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
			ps := c.Query("ps")
			if ps == "" {
				ps = "30"
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
			pgSize, err := strconv.Atoi(ps)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": "page must be a number",
					"error":   err.Error(),
				})
				return
			}
			totalPage := 8
			resp := gin.H{
				"PageRouter": PageRouter,
				"UserInfo": gin.H{
					"Name": "lixiaoming",
				},
				"Products": []gin.H{
					{
						"ProductId":   "1",
						"ProductImg":  "/static/src/product/bearcookie.webp",
						"ProductName": "超级无敌好吃的小熊饼干",
						"ProductSpec": gin.H{
							"Name":  "500g装",
							"Price": "18.80",
						},
						"Quantity": "2",
						"Currency": "CNY",
						"Postage":  "10.00",
						"Sold":     "123",
					},
					{
						"ItemId":      "2",
						"ProductId":   "2",
						"ProductImg":  "/static/src/product/bearsweet.webp",
						"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
						"ProductSpec": gin.H{
							"Name":  "9分软",
							"Price": "20.99",
						},
						"Quantity": "1",
						"Postage":  "0",
						"Sold":     "456",
					},
				},
				"Keyword":   kw,
				"Sort":      sort,
				"CurPage":   curPage,
				"pageSize":  pgSize,
				"TotalPage": totalPage,
    "Category":  category,
			}
			fmt.Println(kw)
			c.HTML(http.StatusOK, "search.html", &resp)
		})
	}
}
