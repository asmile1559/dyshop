// router map
const DefaultURL = "http://192.168.191.130:10166"

const GETReqRouters = {
  "home": "/test",
  "user": "/test/user",
  "login": "/test/user/login",
  "register": "/test/user/register",
  "cart": "/test/cart",
  "checkout": "/test/checkout",
  "product": "/test/product",
  "payment": "/test/payment",
  "order": "/test/order",
  "search": "/test/product/search",
}

const POSTReqRouters = {
  "login": "/test/user/login",
  "register": "/test/user/register",
  "cart": "/test/cart",
  "checkout": "/test/checkout",
  "product": "/test/product",
  "payment": "/test/payment",
  "order": "/test/order",
  "updateInfo": "/test/user",
  "updateImg": "/test/upload/img",
  "newProduct": "/test/upload/product",
  "updateProduct": "/test/upload/product",
  "deleteProduct": "/test/product/delete",
  "buyNow": "/test/product/buyNow",
  "addToCart": "/test/product/add2cart",
  "orderCancel": "/test/order/cancel",
  "orderSubmit": "/test/order/submit",
  "paymentCancel": "/test/payment/cancel",
  "paymentSubmit": "/test/payment/",
  "cartDelete": "/test/cart/delete",
  "checkout": "/test/checkout",
}

export { DefaultURL, GETReqRouters, POSTReqRouters }