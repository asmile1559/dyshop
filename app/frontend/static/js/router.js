// router map
const DefaultURL = "http://192.168.191.130:10166"

const OperationRouters = {
  "home": "/", // GET
  "switchShowcase": "/showcase/", // GET
  "verify": "/verify/", // POST
  "updateUserInfo": "/example/user/info/", // POST
  "updateUserImg": "/example/user/info/upload/", // POST
  "registerMerchant": "/example/user/role/merchant/", // GET
  "updateUserAccount": "/example/user/account/", // POST
  "deleteUserAccount": "/example/user/account/delete/", // POST
  "updateAddress": "/example/user/address/", // POST
  "deleteAddress": "/example/user/address/delete/", // POST
  "setDefAddress": "/example/user/address/setDefault/", // POST
  "updateProduct": "/example/user/product/", // POST
  "deleteProduct": "/example/user/product/delete/", // POST
  "getProduct": "/example/product/", // GET
  "buy": "/example/product/buy/", // POST
  "addToCart": "/example/product/add2cart/", // POST
  "getCart": "/example/cart/", // GET
  "deleteCartItem": "/example/cart/delete/", // POST
  "cartCheckout": "/example/cart/checkout/", // POST
  "getOrder": "/example/order/", // GET
  "cancelOrder": "/example/order/cancel/", // POST
  "submitOrder": "/example/order/submit/", // POST
  "checkout": "/example/checkout/", // GET
  "cancelCheckout": "/example/checkout/cancel/", // POST
  "payment": "/example/payment/", // POST
  "search": "/example/search/", // GET
  "register": "/user/register/", // GET|POST
  "login": "/user/login/", // GET|POST

}
export { DefaultURL, OperationRouters }