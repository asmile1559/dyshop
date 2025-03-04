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
  "updateProduct": "/product/update", // POST
  "deleteProduct": "/product/delete/", // POST
  "getProduct": "/product/", // GET
  "buy": "/product/buy/", // POST
  "addToCart": "/cart/add", // POST
  "getCart": "/cart/", // GET
  "deleteCartItem": "/cart/delete/", // POST
  "cartCheckout": "/example/cart/checkout/", // POST
  "getOrder": "/order/", // GET
  "cancelOrder": "/order/cancel/", // POST
  "submitOrder": "/order/submit/", // POST
  "checkout": "/checkout/", // GET
  "cancelCheckout": "/example/checkout/cancel/", // POST
  "payment": "/payment/", // POST
  "search": "/product/search/", // GET
  "register": "/user/register/", // GET|POST
  "login": "/user/login/", // GET|POST
}
export { DefaultURL, OperationRouters }