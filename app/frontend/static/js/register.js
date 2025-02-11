!function registerProcess() {
    axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
    axios.defaults.headers.post['Content-Type'] = 'application/json'
    axios.defaults.baseURL = 'http://192.168.191.130:10166'

    const registerTabTitle = document.querySelector('.register-tabtitle')
    const emailTabTitle = document.getElementById('email-tabtitle')
    const phoneTabTitle = document.getElementById('phone-tabtitle')

    const accountInput = document.getElementById('accountInput')
    const passwordInput = document.getElementById('passwordInput')
    const confirmPasswordInput = document.getElementById('confirmPasswordInput')

    const accountInputLabelImg = document.querySelector('.account label[for="accountInput"] img')
    const passwordInputLabelImg = document.querySelector('.password label[for="passwordInput"] img')
    const confirmPasswordInputLabelImg = document.querySelector('.confirm-password label[for="confirmPasswordInput"] img')

    const accountInputPrompt = document.querySelector('.account .prompt-area')
    const passwordInputPrompt = document.querySelector('.password .prompt-area')
    const confirmPasswordInputPrompt = document.querySelector('.confirm-password .prompt-area')

    const messageBox = document.querySelector('.message-box')
    const agreement = document.getElementById('invalidCheck')

    const registerBtn = document.getElementById('register-btn')
    const jumpLink = document.querySelector('.jump a')

    const modal = document.querySelector('.modal')
    const modalBody = document.querySelector('.modal-body p')
    const jumpCancel = document.getElementById('jump-cancel')
    const jumpNow = document.getElementById('jump-now')

    let registerWay = 'email'
    let accountInputOk = false
    let passwordInputOk = false
    let confirmPasswordInputOk = false
    jumpLink.addEventListener('click', function (e) {
        e.preventDefault()
        window.location.href = '/test/login'
    })

    registerTabTitle.addEventListener('click', (e) => {
        const target = e.target
        if (target.tagName === 'A') {
            if (target.classList.contains('active')) {
                return
            }
            console.log(target.dataset)
            registerWay = target.dataset['registerWay']
            emailTabTitle.classList.remove('active')
            phoneTabTitle.classList.remove('active')
            accountInput.classList = 'form-control'
            passwordInput.classList = 'form-control'
            confirmPasswordInput.classList = 'form-control'
            accountInput.value = ''
            passwordInput.value = ''
            confirmPasswordInput.value = ''

            target.classList.add('active')
            if (registerWay === 'email') {
                accountInput.setAttribute('placeholder', '请输入邮箱')
                confirmPasswordInput.setAttribute('placeholder', '再次输入登陆密码')
            } else if (registerWay === 'phone') {
                accountInput.setAttribute('placeholder', '请输入手机号')
                confirmPasswordInput.setAttribute('placeholder', '请输入验证码(123456)')
            }
        }
    })

    accountInput.addEventListener('blur', function () {
        accountInputOk = false
        if (this.value === '') {
            accountInputLabelImg.style.display = 'none'
            accountInputPrompt.style.display = 'none'
            return
        }
        accountInputLabelImg.style.display = 'inline-block'
        accountInputLabelImg.src = '/static/src/basic/wrong.svg'
        accountInputPrompt.style.display = 'block'
        if (registerWay === 'email') {
            if (!/^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/.test(this.value)) {
                accountInputPrompt.innerText = '你输入的邮箱格式不正确'
                return
            }
        } else if (registerWay === 'phone') {
            if (!/^1[3456789]\d{9}$/.test(this.value)) {
                accountInputPrompt.innerText = '你输入的手机号格式不正确'
                return
            }
        }
        accountInputOk = true
        accountInputLabelImg.src = '/static/src/basic/correct.svg'
        accountInputPrompt.style.display = 'none'
        return
    })

    passwordInput.addEventListener('blur', function () {
        passwordInputOk = false
        if (this.value === '') {
            passwordInputLabelImg.style.display = 'none'
            passwordInputPrompt.style.display = 'none'
            return
        }
        passwordInputLabelImg.style.display = 'inline-block'
        passwordInputLabelImg.src = '/static/src/basic/wrong.svg'
        passwordInputPrompt.style.display = 'block'
        if (!/^[a-zA-Z0-9-_!@#$%^&*~+=]{6,16}$/.test(this.value)) {
            passwordInputPrompt.innerText = '你输入的密码格式不正确'
            return
        }
        passwordInputOk = true
        passwordInputLabelImg.src = '/static/src/basic/correct.svg'
        passwordInputPrompt.style.display = 'none'

        if (confirmPasswordInput.value !== '' && confirmPasswordInput.value === this.value) {
            confirmPasswordInputOk = true
            confirmPasswordInputLabelImg.src = '/static/src/basic/correct.svg'
            confirmPasswordInputPrompt.style.display = 'none'
        }
        return
    })

    confirmPasswordInput.addEventListener('blur', function () {
        confirmPasswordInputOk = false
        if (this.value === '') {
            confirmPasswordInputLabelImg.style.display = 'none'
            confirmPasswordInputPrompt.style.display = 'none'
            return
        }
        confirmPasswordInputLabelImg.style.display = 'inline-block'
        confirmPasswordInputLabelImg.src = '/static/src/basic/wrong.svg'
        confirmPasswordInputPrompt.style.display = 'block'
        if (!passwordInputOk || this.value !== passwordInput.value) {
            confirmPasswordInputPrompt.innerText = '登录密码输入格式不正确或两次输入的密码不一致'
            return
        }
        confirmPasswordInputOk = true
        confirmPasswordInputLabelImg.src = '/static/src/basic/correct.svg'
        confirmPasswordInputPrompt.style.display = 'none'
        return
    })

    registerBtn.addEventListener('click', (e) => {
        if (!accountInputOk) {
            messageBox.style.display = 'block'
            messageBox.innerText = '请检查邮箱或手机号格式是否正确。'
            return
        }

        if (!passwordInputOk || !confirmPasswordInputOk) {
            messageBox.style.display = 'block'
            messageBox.innerText = '密码不能为空且格式需要正确。'
            return
        }

        if (!agreement.checked) {
            messageBox.style.display = 'block'
            messageBox.innerText = '请先阅读并同意用户协议、隐私政策、产品服务协议。'
            return
        }
        modal.style.display = 'block'
        axios({
            method: 'post',
            url: '/test/register',
            data: {
                email: accountInput.value,
                password: passwordInput.value,
                confirm_password: confirmPasswordInput.value
            }
        }).then(res => {
            if (res.data.code === 200) {
                modal.style.display = 'block'
                let count = 3
                modalBody.innerText = `注册成功，${count}秒后跳转到登录页面。`
                let timer = setInterval(() => {
                    count--
                    modalBody.innerText = `注册成功，${count}秒后跳转到登录页面。`
                    if (count === 0) {
                        clearInterval(timer)
                        window.location.href = '/test/login'
                    }
                }, 1000)

                jumpCancel.addEventListener('click', () => {
                    clearInterval(timer)
                    modal.style.display = 'none'
                })

                jumpNow.addEventListener('click', () => {
                    clearInterval(timer)
                    window.location.href = '/test/login'
                })
            } else {
                messageBox.style.display = 'block'
                messageBox.innerText = '注册失败，账号已存在。'
            }
        }).catch(err => {
            console.log(err)
            messageBox.style.display = 'block'
            messageBox.innerText = '一个错误发生了，请稍后再试。'
        })
    })

}()