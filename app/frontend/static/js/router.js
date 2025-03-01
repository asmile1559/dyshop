// router map
const DefaultURL = "http://localhost:10166"

const OperationRouters = {
  "home": "/", // GET
  "switchShowcase": "/showcase/", // GET
  "verify": "/verify/", // POST
  "updateUserInfo": "/user/info/", // POST
  "updateUserImg": "/user/info/upload/", // POST
  "registerMerchant": "/user/role/merchant/", // GET
  "updateUserAccount": "/user/account/", // POST
  "deleteUserAccount": "/user/account/delete/", // POST
  "updateAddress": "/example/user/address/", // POST
  "deleteAddress": "/example/user/address/delete/", // POST
  "setDefAddress": "/example/user/address/setDefault/", // POST
  "updateProduct": "/example/user/product/", // POST
  "deleteProduct": "/example/user/product/delete/", // POST
  "getProduct": "/example/product/", // GET
  "buy": "/example/product/buy/", // POST
  // "addToCart": "/example/product/add2cart/", // POST
  "addToCart": "/cart/add", // POST
  "getCart": "/cart/", // GET
  "deleteCartItem": "/cart/delete/", // POST
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