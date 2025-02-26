import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

function UpdatePrice() {
  const orderListItems = document.querySelector('.order-list-items')
  let totalPrice = 0
  let totalDeliverCost = 0
  for (let item of orderListItems.children) {
    // console.log(item)
    const product_id = item.dataset.productId
    const q = item.querySelector('#quantity' + product_id).value
    const sp = item.querySelector('#single-price' + product_id).textContent
    const select = item.querySelector("#deliver-select" + product_id)
    const dc = parseFloat(select.options[select.selectedIndex].dataset['deliverCost']).toFixed(2)
    const sum = q * sp
    item.querySelector('#price' + product_id).textContent = sum.toFixed(2)
    item.querySelector('#deliver-cost' + product_id).textContent = dc
    totalPrice += parseFloat(sum.toFixed(2))
    totalDeliverCost += parseFloat(dc)
  }

  const discount = document.querySelector('#discount').textContent
  document.querySelector('#total-price').textContent = totalPrice.toFixed(2)
  document.querySelector('#total-deliver-cost').textContent = totalDeliverCost.toFixed(2)
  document.querySelector('#real-price').textContent = (+totalPrice + totalDeliverCost - discount).toFixed(2)
}

!function () {
  document.querySelector('.order-list-items').addEventListener('click', (e) => {
    const target = e.target
    if (target.tagName === 'BUTTON') {
      if (target.dataset['btnRole'] === 'subQuantity') {
        const quantity = target.nextElementSibling
        if (quantity.value > 1) {
          quantity.value--
          UpdatePrice()
        }

      } else if (target.dataset['btnRole'] === 'addQuantity') {
        const quantity = target.previousElementSibling
        quantity.value++
        UpdatePrice()
      }
    }
  })

  document.querySelectorAll('input[data-input-role="quantity"]').forEach((ele) => {
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
  })

  document.querySelectorAll('select[data-select-role="deliver"]').forEach((ele) => {
    ele.addEventListener('change', (e) => {
      UpdatePrice()
    })
  })
}()


!function () {
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
  axios.defaults.headers.post['Content-Type'] = 'application/json'
  axios.defaults.baseURL = router.DefaultURL
  document.querySelector('#goBack').addEventListener('click', (e) => {
    console.log('go back')
    e.preventDefault()
    console.log(window.location.search)
    const order_id = window.location.search.match(/order_id=(\w+)&*/)[1]

    axios({
      method: 'post',
      url: router.OperationRouters['cancelOrder'],
      data: {
        order_id: order_id,
      }
    }).then((res) => {
      console.log(res.data)
      window.location.href = '/'
    }).catch((err) => {
      console.log(err)
      window.history.back()
    })
  })

  document.querySelector('#submitOrder').addEventListener('click', (e) => {
    console.log('submit order')
    e.preventDefault()
    // 方法1：直接使用addressId，在后端查找对应的address
    const addressId = document.querySelector('input[type="radio"][name="address"]:checked').dataset['addrId']
    // 方法2：直接使用address信息
    console.log(addressId)
    const addressItem = document.querySelector('.address-list-item[data-addr-id="' + addressId + '"]')
    const address = {
      address_id: addressId,
      name: addressItem.querySelector('.name').textContent.trim(),
      phone: addressItem.querySelector('.phone').textContent.trim(),
      detail: addressItem.querySelector('.detail').textContent.trim(),
      province: addressItem.querySelector('.province').textContent.trim(),
      city: addressItem.querySelector('.city').textContent.trim(),
      district: addressItem.querySelector('.district').textContent.trim(),
      street: addressItem.querySelector('.street').textContent.trim(),
    }

    console.log(address)
    const orderListItems = document.querySelector('.order-list-items')
    const orderItems = []
    for (let item of orderListItems.children) {
      const product_id = item.dataset.productId
      const q = item.querySelector('#quantity' + product_id).value
      const sp = item.querySelector('#single-price' + product_id).textContent
      const select = item.querySelector("#deliver-select" + product_id)
      const dc = parseFloat(select.options[select.selectedIndex].dataset['deliverCost']).toFixed(2)
      const p = item.querySelector('#price' + product_id).textContent
      const remark = item.querySelector('#order-remark' + product_id).textContent
      orderItems.push({
        product_id: product_id,
        quantity: q,
        single_price: sp,
        price: p,
        deliver_cost: dc,
        remark: remark,
      })
    }

    console.log(orderItems)
    const order_id = window.location.search.match(/order_id=(\w+)&*/)[1]
    const data = {
      order_id: order_id,
      address: address,
      products: orderItems,
      discount: parseFloat(document.querySelector('#discount').textContent).toFixed(2),
      order_price: parseFloat(document.querySelector('#total-price').textContent).toFixed(2),
      order_postage: parseFloat(document.querySelector('#total-deliver-cost').textContent).toFixed(2),
      order_final_price: parseFloat(document.querySelector('#real-price').textContent).toFixed(2),
    }

    console.log(data)
    axios({
      method: 'post',
      url: router.OperationRouters['submitOrder'],
      data: data,
    }).then((res) => {
      const resp = res.data['resp']
      const transaction_id = resp['transaction_id']
      const order_id = resp['order_id']
      console.log(transaction_id)
      window.location.href = `${router.OperationRouters['checkout']}?transaction_id=${transaction_id}&order_id=${order_id}`
    }).catch((err) => {
      console.log(err)
    })
  })
}()

UpdatePrice()

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.OperationRouters['search'] + '?keyword=' + keyword
  })
}()