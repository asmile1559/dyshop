import * as router from './router.js'
import * as common from './common.js'

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.OperationRouters['search'] + '?keyword=' + keyword
  })
}()

// !function () {
//   document.querySelectorAll('.product-item-g').forEach((item) => {
//     console.log(item)
//     item.addEventListener('click', (e) => {
//       const productId = item.dataset['productId']
//       window.location.href = router.OperationRouters['product'] + '?product_id=' + productId
//     })
//   })
// }()

!function () {
  document.querySelector('.product-list').addEventListener('click', (e) => {
    const target = e.target
    const product_item = target.closest('.product-item-g')
    if (product_item !== null) {
      window.location.href = router.OperationRouters['getProduct'] + '?product_id=' + product_item.dataset['productId']
    }
  })
}()