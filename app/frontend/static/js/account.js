import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

!function () {
  document.querySelector("#register-merchant").addEventListener("click", (e) => {
    const target = e.target
    const user_id = target.dataset["userId"]
    axios({
      'method': 'GET',
      'url': router.OperationRouters['registerMerchant'],
    }).then((res) => {
      const resp = res.data.resp
      document.querySelector(".account-info .user-role>span:nth-child(2)").classList.add("active")
      target.classList.add('d-none')
      target.nextElementSibling.classList.remove('d-none')
      document.querySelector(".user-content .panel .product").classList.remove('d-none')

    }).catch((err) => {
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })
}()

!function () {
  document.querySelector('#update-phone-btn').addEventListener('click', () => {
    const user_phone = document.querySelector('#new-user-phone').value
    if (user_phone === '') {
      common.alertByModal('错误', '请填写完整信息。')
      return
    }
    axios({
      'method': 'POST',
      'url': router.OperationRouters['updateUserAccount'],
      'data': {
        'phone': user_phone
      }
    }).then((res) => {
      const resp = res.data.resp
      console.log(resp)
      document.querySelector('#user-phone').value = user_phone
    }).catch((err) => {
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })

  document.querySelector('#update-email-btn').addEventListener('click', () => {
    const user_email = document.querySelector('#new-user-email').value
    if (user_email === '') {
      common.alertByModal('错误', '请填写完整信息。')
      return
    }
    axios({
      'method': 'POST',
      'url': router.OperationRouters['updateUserAccount'],
      'data': {
        'email': user_email
      }
    }).then((res) => {
      const resp = res.data.resp
      console.log(resp)
      document.querySelector('#user-email').value = user_email
    }).catch((err) => {
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })

  document.querySelector('#update-password-btn').addEventListener('click', () => {
    const user_password = document.querySelector('#user-password').value
    const user_new_password = document.querySelector('#new-user-password').value
    const user_confirm_password = document.querySelector('#confirm-user-password').value
    if (user_new_password === '' || user_confirm_password === '' || user_new_password !== user_confirm_password) {
      common.alertByModal('错误', '请填写完整信息。')
      return
    }
    axios({
      'method': 'POST',
      'url': router.OperationRouters['updateUserAccount'],
      'data': {
        password: user_password,
        new_password: user_new_password,
        confirm_password: user_confirm_password
      }
    }).then((res) => {
      const resp = res.data.resp
      console.log(resp)
      document.querySelector('#user-password').value = ""
      document.querySelector('#new-user-password').value = ""
      document.querySelector('#confirm-user-password').value = ""
    }).catch((err) => {
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })
}()

!function () {
  document.querySelector('#del-account-btn').addEventListener('click', (e) => {
    if (!confirm('确定要注销账户吗？')) {
      return
    }
    const user_id = e.target.dataset['userId']
    axios({
      'method': 'GET',
      'url': router.OperationRouters['deleteUserAccount'],
    }).then((res) => {
      const resp = res.data.resp
      console.log(resp)
      common.alertByModal('成功', '账户已注销。3s后返回首页。')
      localStorage.removeItem('token')
      common.deleteAllCookies()
      setTimeout(() => {
        window.location.href = router.OperationRouters['home']
      }, 3000)
    }).catch((err) => {
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })
}()