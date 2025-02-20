import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

!function () {
  updatePrice()
}()

function updatePrice() {
  const spec = document.querySelector("input:checked[data-role='spec']:checked")
  const quantity = document.querySelector('#quantity')
  document.querySelector('#show-price').textContent = parseFloat(parseFloat(spec.dataset["price"]) * parseInt(quantity.value)).toFixed(2)
}

!function () {
  document.querySelector('#quantity').addEventListener('blur', (e) => {
    if (e.target.value < 1) {
      e.target.value = 1
    }
    e.target.value = parseInt(e.target.value)
    if (isNaN(e.target.value)) {
      e.target.value = 1
    }
    updatePrice()
  })

  document.querySelector('#subQuantity').addEventListener('click', (e) => {
    const quantity = document.querySelector('#quantity')
    if (quantity.value > 1) {
      quantity.value--
    }
    if (isNaN(quantity.value)) {
      quantity.value = 1
    }
    updatePrice()
  })

  document.querySelector('#addQuantity').addEventListener('click', (e) => {
    const quantity = document.querySelector('#quantity')
    quantity.value++
    if (isNaN(quantity.value)) {
      quantity.value = 1
    }
    updatePrice()
  })

  document.querySelectorAll(".spec-box input").forEach((input) => {
    input.addEventListener('click', (e) => {
      updatePrice()
    })
  })
}()

!function () {
  const buyNow = document.querySelector('#buyNow')
  const addToCart = document.querySelector('#addToCart')
  buyNow.addEventListener('click', (e) => {
    const target = e.target
    const spec = document.querySelector("input:checked[data-role='spec']:checked")
    axios({
      method: 'post',
      url: router.OperationRouters['buy'],
      data: {
        product_id: target.dataset.productId,
        product_spec: {
          name: spec.nextElementSibling.textContent,
          price: spec.dataset.price,
        },
        quantity: document.querySelector('#quantity').value,
        order_price: parseFloat(document.querySelector('#show-price').textContent).toFixed(2),
        postage: parseInt(document.querySelector('#postage').textContent).toFixed(2),
        currency: "CNY",
      },
    }).then((res) => {
      const resp = res.data.resp
      const orderId = resp['order_id']
      window.location.href = `${router.OperationRouters['getOrder']}?order_id=${orderId}`
    }
    ).catch((err) => {
      console.log(err)
    })
  })

  addToCart.addEventListener('click', (e) => {
    const target = e.target
    const spec = document.querySelector("input:checked[data-role='spec']:checked")
    axios({
      method: 'post',
      url: router.OperationRouters['addToCart'],
      data: {
        product_id: target.dataset.productId,
        product_spec: {
          name: spec.nextElementSibling.textContent,
          price: spec.dataset.price,
        },
        quantity: document.querySelector('#quantity').value,
        postage: parseInt(document.querySelector('#postage').textContent).toFixed(2),
        currency: "CNY",
      },
    }).then((res) => {
      window.location.href = `${router.OperationRouters['getCart']}` // 使用token替代user id
    }
    ).catch((err) => {
      console.log(err)
    })
  })

}()

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.OperationRouters['search'] + '?keyword=' + keyword
  })
}()