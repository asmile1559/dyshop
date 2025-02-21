import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

!function () {
  document.querySelector("#update-img-btn").addEventListener("click", function () {
    document.querySelector("#user-img-input").click()
  })

  document.querySelector("#user-img-input").addEventListener("change", function (e) {
    const file = e.target.files[0]
    const formData = new FormData()
    formData.append('Img', file)
    axios({
      'method': 'POST',
      'url': router.OperationRouters['updateUserImg'],
      'data': formData,
      'headers': {
        'Content-Type': 'multipart/form-data'
      }
    }).then((res) => {
      console.log(res)
      const resp = res.data.resp
      document.querySelectorAll('.user-img>img').forEach(ele => {
        ele.src = resp["url"]
      });
    }).catch((err) => {
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })
}()


!function () {
  document.querySelector("#update-info-btn").addEventListener("click", function () {
    const userName = document.querySelector("#username").value
    const userSign = document.querySelector("#usersign").value
    const userGender = document.querySelector('#user-gender').value
    const userBirthday = document.querySelector('#user-birthday').value
    console.log(userBirthday)
    if (userName === '' || userSign === '' || userGender === '' || userBirthday === '') {
      common.alertByModal('错误', '请填写完整信息。')
      return
    }
    axios({
      'method': 'POST',
      'url': router.OperationRouters['updateUserInfo'],
      'data': {
        'name': userName,
        'sign': userSign,
        'gender': userGender,
        'birthday': userBirthday
      }
    }).then((res) => {
      const resp = res.data.resp
      console.log(res)
      document.querySelector('#nav-username').textContent = resp["name"]
      document.querySelector('#nav-usersign').textContent = resp["sign"]
    }).catch((err) => {
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })
}()
