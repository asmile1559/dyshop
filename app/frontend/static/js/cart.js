import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL


function DeleteACartItemOnPage(itemId) {
  const cartList = document.querySelector('#cart-list')
  const item = document.querySelector(`#cart-item${itemId}`)
  cartList.removeChild(item)
}

function UpdatePrice() {
  const items = document.querySelectorAll('.cart-item input[type=checkbox]:checked')
  let totalPrice = 0
  let totalQuantity = 0
  items.forEach((item) => {
    totalQuantity++
    console.log(item)
    const itemId = item.dataset['itemId']
    const price = document.querySelector('#item-price' + itemId).textContent
    const quantity = document.querySelector('#item-quantity' + itemId).value
    console.log(price)
    console.log(quantity)
    totalPrice += parseFloat(price) * parseFloat(quantity)
  })
  document.querySelector('#totalQuantity').innerText = totalQuantity
  document.querySelector('#totalPrice').innerText = isNaN(totalPrice.toFixed(2)) ? "0.00" : totalPrice.toFixed(2)
}
UpdatePrice()

function UpdateCartTitle() {
  const chooseAll = document.querySelector('#chooseAll')
  const globalDel = document.querySelector('#globalDel')
  const globalCheckout = document.querySelector('#globalCheckout')
  if (document.querySelectorAll('.cart-item').length === 0) {
    chooseAll.checked = false
    chooseAll.disabled = true
    globalDel.disabled = true
    globalCheckout.disabled = true
    checkout.disabled = true
  } else {
    chooseAll.disabled = false
  }
}
UpdateCartTitle()

!function () {
  document.querySelectorAll("[data-role='subQuantity']").forEach((ele) => {
    ele.addEventListener('click', (e) => {
      console.log('sub')
      const quantity = ele.nextElementSibling
      if (quantity.value > 1) {
        quantity.value--
        UpdatePrice()
      }
    })
  })

  document.querySelectorAll("[data-role='addQuantity']").forEach((ele) => {
    ele.addEventListener('click', (e) => {
      console.log('add')
      const quantity = ele.previousElementSibling
      quantity.value++
      UpdatePrice()
    })
  })

  document.querySelectorAll("[data-role='quantity']").forEach((ele) => {
    ele.addEventListener('blur', (e) => {
      const target = e.target
      if (target.value < 1) {
        target.value = 1
      }
      target.value = parseInt(target.value)
      if (isNaN(target.value)) {
        target.value = 1
      }
      UpdatePrice()
    })

    ele.addEventListener('input', (e) => {
      const target = e.target
      target.value = target.value.replace(/\D/g, '')
      UpdatePrice()
    })
  })

  document.querySelectorAll("input[type=checkbox][data-role='choose']").forEach((ele) => {
    ele.addEventListener('change', (e) => {
      UpdatePrice()
      const allCheckedBox = document.querySelectorAll("input[type=checkbox][data-role='choose']:checked")
      const allBox = document.querySelectorAll("input[type=checkbox][data-role='choose']")
      if (allCheckedBox.length === allBox.length) {
        document.querySelector('#chooseAll').checked = true
      } else {
        document.querySelector('#chooseAll').checked = false
      }

      if (allCheckedBox.length === 0) {
        document.querySelector('#globalDel').disabled = true
        document.querySelector('#globalCheckout').disabled = true
      } else {
        document.querySelector('#globalDel').disabled = false
        document.querySelector('#globalCheckout').disabled = false
      }
    })
  })

  document.querySelector('#chooseAll').addEventListener('change', (e) => {
    const allBox = document.querySelectorAll("input[type=checkbox][data-role='choose']")
    if (e.target.checked) {
      allBox.forEach((ele) => {
        ele.checked = true
      })
      document.querySelector('#globalDel').disabled = false
      document.querySelector('#globalCheckout').disabled = false
    } else {
      allBox.forEach((ele) => {
        ele.checked = false
      })
      document.querySelector('#globalDel').disabled = true
      document.querySelector('#globalCheckout').disabled = true
    }
    UpdatePrice()
  })

}()

function getCartItemInfo(itemId) {
  const product_id = document.querySelector("#cart-item" + itemId).dataset['productId']
  const name = document.querySelector('#item-name' + itemId).textContent
  const spec = document.querySelector('#item-spec' + itemId).textContent
  const price = document.querySelector('#item-price' + itemId).textContent
  const quantity = document.querySelector('#item-quantity' + itemId).value
  return {
    item_id: itemId,
    product_id,
    name,
    product_spec: {
      spec_name: spec,
      spec_price: price,
    },
    quantity,
    conrrency: "CNY"
  }
}

function checkout() {
  const items = document.querySelectorAll('.cart-item input[type=checkbox][data-role="choose"]:checked')
  const data = {
    "order_price": document.querySelector('#totalPrice').textContent,
    "cart_items": [],
  }

  items.forEach((item) => {
    const itemId = Number(item.dataset['itemId'])
    data.cart_items.push(getCartItemInfo(itemId))
  })

  console.log(data)

  axios({
    method: 'post',
    url: router.OperationRouters['cartCheckout'],
    data
  }).then((res) => {
    console.log(res)
    const resp = res.data.resp
    for (let item of items) {
      DeleteACartItemOnPage(item.dataset['itemId'])
    }
    UpdatePrice()
    UpdateCartTitle()
    window.location.href = router.OperationRouters['getOrder'] + '?order_id=' + resp["order_id"]
  }).catch((err) => {
    console.log(err)
  })
}


!function () {

  document.querySelector('#globalDel').addEventListener('click', (e) => {
    const items = document.querySelectorAll('.cart-item input[type=checkbox][data-role="choose"]:checked')
    const itemIds = []
    items.forEach((item) => {
      itemIds.push(Number(item.dataset['itemId']))
    })
    console.log(itemIds)
    axios({
      method: 'post',
      url: router.OperationRouters['deleteCartItem'],
      data: {
        item_ids: itemIds
      }
    }).then((res) => {
      console.log(res)
      items.forEach((item) => {
        DeleteACartItemOnPage(item.dataset['itemId'])
      })
      UpdatePrice()
      UpdateCartTitle()
    }).catch((err) => {
      console.log(err)
    })
  })

  document.querySelectorAll('a[data-role="del"]').forEach((ele) => {
    ele.addEventListener('click', (e) => {
      const itemId = ele.dataset['itemId']
      axios({
        method: 'post',
        url: router.OperationRouters['deleteCartItem'],
        data: {
          item_ids: [Number(itemId)]
        }
      }).then((res) => {
        console.log(res)
        DeleteACartItemOnPage(itemId)
        UpdatePrice()
        UpdateCartTitle()
      }).catch((err) => {
        console.log(err)
      })
    })
  })


  document.querySelector('#globalCheckout').addEventListener('click', (e) => {
    checkout()
  })

  document.querySelector('#checkout').addEventListener('click', (e) => {
    if (document.querySelectorAll('.cart-item input[type=checkbox][data-role="choose"]:checked').length === 0) {
      return
    }
    checkout()
  })

  document.querySelectorAll('a[data-role="check"]').forEach((ele) => {
    ele.addEventListener('click', (e) => {
      ele.closest('.cart-item').querySelector('input[type=checkbox][data-role="choose"]').checked = true
      checkout()
    })
  })

}()

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.OperationRouters['search'] + '?keyword=' + keyword
  })
}()