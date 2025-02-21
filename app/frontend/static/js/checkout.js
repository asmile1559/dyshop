import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

async function CancelPayment() {
  try {
    const res = await axios({
      method: 'post',
      url: router.OperationRouters['cancelCheckout'],
      data: {
        transaction_id: window.location.search.match(/transaction_id=(\w+)&*/)[1],
        order_id: window.location.search.match(/order_id=(\w+)&*/)[1],
      }
    })
    console.log(res.data)
  } catch (err) {
    console.log(err)
  }
}

!function () {
  const countdown = document.getElementById('countdown')
  let time = 600
  const timer = setInterval(() => {
    time--
    const minutes = Math.floor(time / 60)
    const seconds = time % 60
    countdown.innerText = `${minutes}:${seconds < 10 ? '0' + seconds : seconds}`
    if (time <= 0) {
      clearInterval(timer)
      CancelPayment()
      common.alertByModal('超时', '支付超时已取消，1s后跳转到首页')
      setTimeout(() => {
        window.location.href = router.OperationRouters['home']
      }, 1000)
    }
  }, 1000)
}()

function TextPasswordChange(ele) {
  ele.addEventListener('focus', () => {
    ele.type = 'text'
  })
  ele.addEventListener('blur', () => {
    ele.type = 'password'
  })
}

function TruncateInput(ele, maxLength) {
  ele.addEventListener('input', () => {
    if (ele.value.length > maxLength) {
      ele.value = ele.value.slice(0, maxLength)
    }
  })
}

!function () {
  const cardNumber = document.getElementById('card-number')
  const cvv = document.getElementById('card-cvv')
  TextPasswordChange(cardNumber)
  TruncateInput(cardNumber, 16)
  TextPasswordChange(cvv)
  TruncateInput(cvv, 3)
}()

!function () {
  const cancelBtn = document.getElementById('cancelBtn')
  const confirmBtn = document.getElementById('confirmBtn')
  cancelBtn.addEventListener('click', () => {
    const res = CancelPayment()
    common.alertByModal('取消', '支付已取消，1s后跳转到首页')
    setTimeout(() => {
      window.location.href = router.OperationRouters['home']
    }, 1000)
  })

  confirmBtn.addEventListener('click', () => {
    const real_price = document.getElementById('realPrice').innerText
    const card_type = document.querySelector('input[type="radio"]:checked').id
    const card_number = document.getElementById('card-number').value
    const card_holder = document.getElementById('card-holder').value
    const cvv = document.getElementById('card-cvv').value
    const expire_month = document.getElementById('card-expiration-month').value
    const expire_year = document.getElementById('card-expiration-year').value
    const transaction_id = window.location.search.match(/transaction_id=(\w+)&*/)[1]
    const order_id = window.location.search.match(/order_id=(\w+)&*/)[1]
    if (
      card_number.length !== 16 ||
      card_holder.length < 1 ||
      cvv.length !== 3 ||
      expire_month === '' ||
      expire_year === '') {
      common.alertByModal('错误', '请检查你的支付信息')
      return
    }

    const data = {
      transaction_id,
      credit_card: {
        card_type,
        card_number,
        card_holder,
        cvv,
        expire_month,
        expire_year,
      },
      final_price: real_price,
    }
    axios({
      method: 'post',
      url: router.OperationRouters['payment'],
      data,
    }).then((res) => {
      console.log(res.data)
      common.alertByModal('成功', '你的支付已经成功。1s后跳转到首页')
      setTimeout(() => {
        window.location.href = router.OperationRouters['home']
      }, 1000)
    }).catch((err) => {
      console.log(err)
      common.alertByModal('失败', '支付失败，请检查你的支付信息')
    })
  })
}()