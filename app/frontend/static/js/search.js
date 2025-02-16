import * as router from './router.js'
import modalCtrl from './common.js'

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.GETReqRouters['search'] + '?keyword=' + keyword
  })
}()

!function () {
  document.querySelectorAll('.product-item-g').forEach((item) => {
    console.log(item)
    item.addEventListener('click', (e) => {
      const productId = item.dataset['productId']
      window.location.href = router.GETReqRouters['product'] + '?product_id=' + productId
    })
  })
}()