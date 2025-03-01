package user

import (
	"net/http"

	u "github.com/asmile1559/dyshop/app/frontend/biz/handler/user"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	/* auth := func(c *gin.Context) {
		logrus.Infof("Method: %v, URI: %v", c.Request.Method, c.Request.RequestURI)
		// 一般在auth中鉴权, 并将用户id加入到context中
		// 所以在执行用户操作时,不需要传递用户参数
		c.Set("user_id", int64(5163785232846848))
		c.Next()
	} */
	root := e.Group("/",_rootMw()...)
	{
		_user := root.Group("/user", _userMw()...)
		//_user := root.Group("/user", auth)
		_user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})
		_user.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{})
		})
		_user.POST("/login", append(_loginMw(), u.Login)...)
		_user.POST("/register", append(_registerMw(), u.Register)...)

		// GET: /user/info
		// 获取用户信息
		_user.GET("/info",append(_getuserinfoMw(),u.GetUserInfo)...)
		
		// POST /user/info
		// 修改用户文字信息
		_user.POST("/info",append(_updateuserinfoMw(), u.UpdateUserInfo)...)
		
		// POST /user/info/upload
		// 修改用户图片信息
		_user.POST("/info/upload",append(_uploadavatarMw(), u.UploadAvatar)...)
		
		// GET: /user/account
		// 获取用户账户
		_user.GET("/account",append(_getaccountinfoMw(), u.GetAccountInfo)...)
		
		// POST /user/account
		// 修改用户账户信息
		_user.POST("/account",append(_updateaccountinfoMw(), u.UpdateAccount)...)
		
		// GET /user/account/delete
		// 删除账户
		_user.GET("/account/delete",append(_deleteMw(), u.Delete)...)
		
		// GET /user/role/merchant
		// 注册成为商户
		_user.GET("/role/merchant",append(_registermerchantMw(), u.RegisterMerchant)...)
	}
}
