import * as router from './router.js'
import modalCtrl from './common.js'

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

const modal = modalCtrl()

function CancelPayment() {
  axios({
    method: 'post',
    url: router.POSTReqRouters['paymentCancel'],
    data: {
      payment_id: window.location.search.match(/payment_id=(\w+)&*/)[1],
    }
  }).then((res) => {
    console.log(res.data)
    return { "status": "success", "message": "Payment cancelled", resp: res.data }
  }).catch((err) => {
    console.log(err)
    return { status: "error", message: "Payment cancel failed", resp: err }
  })
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
      modal('超时', '支付超时已取消，1s后跳转到首页')
      setTimeout(() => {
        window.location.href = router.GETReqRouters['home']
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
    modal('取消', '支付已取消，1s后跳转到首页')
    setTimeout(() => {
      window.location.href = router.GETReqRouters['home']
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
    const payment_id = window.location.search.match(/payment_id=(\w+)&*/)[1]

    if (
      card_number.length !== 16 ||
      card_holder.length < 1 ||
      cvv.length !== 3 ||
      expire_month === '' ||
      expire_year === '') {
      modal('错误', '请检查你的支付信息')
      return
    }

    axios({
      method: 'post',
      url: router.POSTReqRouters['paymentSubmit'],
      data: {
        payment_id,
        card_type,
        card_number,
        card_holder,
        cvv,
        expire_month,
        expire_year,
        real_price,
      }
    }).then((res) => {
      console.log(res.data)
      modal('成功', '你的支付已经成功。1s后跳转到首页')
      setTimeout(() => {
        window.location.href = router.GETReqRouters['home']
      }, 1000)
    }).catch((err) => {
      console.log(err)
      modal('失败', '支付失败，请检查你的支付信息')
    })
  })
}()