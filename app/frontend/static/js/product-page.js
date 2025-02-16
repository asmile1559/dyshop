import * as router from './router.js'

function UpdatePrice() {
  const price = document.querySelector('#price')
  const quantity = document.querySelector('#quantity')
  const valueItme = document.querySelector("input:checked[data-input-type='valueItem']")
  let realPrice = valueItme.dataset.price * quantity.value
  price.textContent = realPrice.toFixed(2)
}

!function () {
  const price = document.querySelector('#price')
  const quantity = document.querySelector('#quantity')
  document.querySelector('.value-item-box').addEventListener('click', (e) => {
    if (e.target.tagName === 'INPUT') {
      UpdatePrice()
    }
  })
}()

!function () {
  const quantity = document.querySelector('#quantity')
  const subQuantity = document.querySelector('#subQuantity')
  const addQuantity = document.querySelector('#addQuantity')
  quantity.addEventListener('blur', (e) => {
    if (quantity.value < 1) {
      quantity.value = 1
    }
    quantity.value = parseInt(quantity.value)
    if (isNaN(quantity.value)) {
      quantity.value = 1
    }
    UpdatePrice()
  })

  subQuantity.addEventListener('click', (e) => {
    if (quantity.value > 1) {
      quantity.value--
    }
    if (isNaN(quantity.value)) {
      quantity.value = 1
    }
    UpdatePrice()
  })

  addQuantity.addEventListener('click', (e) => {
    quantity.value++
    if (isNaN(quantity.value)) {
      quantity.value = 1
    }
    UpdatePrice()
  })
}()

!function () {
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
  axios.defaults.headers.post['Content-Type'] = 'application/json'
  axios.defaults.baseURL = router.DefaultURL
  const buyNow = document.querySelector('#buyNow')
  const addToCart = document.querySelector('#addToCart')
  buyNow.addEventListener('click', (e) => {
    console.log(e.target.dataset.productId)
    const item = document.querySelector("input:checked[data-input-type='valueItem']").nextElementSibling.textContent
    const singlePrice = parseFloat(document.querySelector("input:checked[data-input-type='valueItem']").dataset.price)
    console.log(item)
    axios({
      method: 'post',
      url: router.POSTReqRouters['buyNow'],
      data: {
        product_id: e.target.dataset.productId,
        item: item,
        single_price: singlePrice,
        quantity: parseFloat(document.querySelector('#quantity').value),
        price: parseFloat(document.querySelector('#price').textContent),
      },
    }).then((res) => {
      console.log(res)
      const orderId = res.data['order_id']
      window.location.href = `${router.GETReqRouters['order']}?order_id=${orderId}`
    }
    ).catch((err) => {
      console.log(err)
    })
  })

  addToCart.addEventListener('click', (e) => {
    console.log(e.target.dataset.productId)
    const item = document.querySelector("input:checked[data-input-type='valueItem']").nextElementSibling.textContent
    const singlePrice = parseFloat(document.querySelector("input:checked[data-input-type='valueItem']").dataset.price)
    console.log(item)
    axios({
      method: 'post',
      url: router.POSTReqRouters['addToCart'],
      data: {
        product_id: e.target.dataset.productId,
        item: item,
        single_price: singlePrice,
        quantity: parseFloat(document.querySelector('#quantity').value),
        price: parseFloat(document.querySelector('#price').textContent),
      },
    }).then((res) => {
      console.log(res)
      window.location.href = `${router.GETReqRouters['cart']}` // 使用token替代user id
    }
    ).catch((err) => {
      console.log(err)
    })
  })

  console.log(buyNow)
  console.log(addToCart)
}()

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.GETReqRouters['search'] + '?keyword=' + keyword
  })
}()