!function loginProcess() {

    axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
    axios.defaults.headers.post['Content-Type'] = 'application/json'
    axios.defaults.baseURL = 'http://192.168.191.130:10166'

    const loginTabtitle = document.querySelector('.login-tabtitle')
    const emailTabtitle = document.getElementById('email-tabtitle')
    const phoneTabtitle = document.getElementById('phone-tabtitle')

    const accountInput = document.getElementById('accountInput')
    const accountInputLabelImg = document.querySelector('label[for="accountInput"] img')
    const accountInputPrompt = document.querySelector('.input-box.account .prompt-area')

    const passwordInput = document.getElementById('passwordInput')
    const passwordInputLabelImg = document.querySelector('label[for="passwordInput"] img')
    const passwordInputPrompt = document.querySelector('.input-box.password .prompt-area')

    const messageBox = document.querySelector('.message-box')
    const jumpLink = document.querySelector('.jump a')

    const modal = document.querySelector('.modal')
    const modalBody = document.querySelector('.modal-body p')
    const jumpCancel = document.getElementById('jump-cancel')
    const jumpNow = document.getElementById('jump-now')

    let accountInputOk = false
    let passwordInputOk = false
    let loginWay = 'email'

    jumpLink.addEventListener('click', function (e) {
        e.preventDefault()
        window.location.href = '/test/register'
    })

    loginTabtitle.addEventListener('click', (e) => {
        const target = e.target
        if (e.target.tagName === 'A') {
            if (target.classList.contains('active')) {
                return
            }
            loginWay = target.dataset['loginWay']
            emailTabtitle.classList.remove('active')
            phoneTabtitle.classList.remove('active')
            accountInput.className = 'form-control'
            passwordInput.className = 'form-control'
            accountInput.value = ''
            passwordInput.value = ''
            target.classList.add('active')
            accountInput.placeholder = loginWay === 'email' ? '请输入邮箱' : '请输入手机号'
        }
    })

    accountInput.addEventListener('blur', function () {
        accountInputOk = false
        if (this.value === '') {
            accountInputLabelImg.style.display = 'none'
            accountInputPrompt.style.display = 'none'
            return
        }
        accountInputLabelImg.style.display = 'block'
        accountInputLabelImg.src = '/static/src/basic/wrong.svg'
        accountInputPrompt.style.display = 'block'
        if (loginWay === 'email') {
            if (!/^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/.test(this.value)) {
                accountInputPrompt.innerText = '你输入的邮箱格式不正确'
                return
            }
        } else if (loginWay === 'phone') {
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
        passwordInputLabelImg.style.display = 'block'
        passwordInputLabelImg.src = '/static/src/basic/wrong.svg'
        passwordInputPrompt.style.display = 'block'
        if (!/^[a-zA-Z0-9-_!@#$%^&*~+=]{6,16}$/.test(this.value)) {
            passwordInputPrompt.innerText = '你输入的密码格式不正确'
            return
        }
        passwordInputOk = true
        passwordInputLabelImg.src = '/static/src/basic/correct.svg'
        passwordInputPrompt.style.display = 'none'
        return
    })

    const loginBtn = document.getElementById('login-btn')
    loginBtn.addEventListener('click', function () {
        messageBox.style.display = 'none'
        if (!accountInputOk) {
            messageBox.style.display = 'block'
            messageBox.innerText = '请检查邮箱或手机号格式是否正确。'
            return
        }
        if (!passwordInputOk) {
            messageBox.style.display = 'block'
            messageBox.innerText = '密码不能为空且格式需要正确。'
            return
        }

        axios({
            method: 'post',
            url: '/test/login',
            data: {
                email: accountInput.value,
                password: passwordInput.value
            }
        }).then(res => {
            if (res.data.code === 200) {
                localStorage.setItem('token', res.data.token)
                modal.style.display = 'block'
                let count = 3
                modalBody.innerText = `登录成功，${count}秒后跳转到首页。`
                let timer = setInterval(() => {
                    count--
                    modalBody.innerText = `登录成功，${count}秒后跳转到首页。`
                    if (count === 0) {
                        clearInterval(timer)
                        window.location.href = '/index.html'
                    }
                }, 1000)

                jumpCancel.addEventListener('click', function () {
                    clearInterval(timer)
                    modal.style.display = 'none'
                })

                jumpNow.addEventListener('click', function () {
                    clearInterval(timer)
                    window.location.href = '/index.html'
                })
                console.log(res.data)
            } else {
                messageBox.style.display = 'block'
                messageBox.innerText = '登录失败，账号或密码错误。请检查你的账号和密码。'
            }
        }).catch(err => {
            console.log(err)
            messageBox.style.display = 'block'
            messageBox.innerText = '一个错误发生了，请稍后再试。'
        })
    })
}()

// // File: public/js/login.js
// document.getElementById('login-form').addEventListener('submit', async (e) => {
//     e.preventDefault();

//     const formData = new FormData(e.target);
//     const data = {
//         username: formData.get('username'),
//         password: formData.get('password')
//     };

//     try {
//         const response = await fetch('/api/v1/login', {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify(data)
//         });

//         const result = await response.json();

//         if (!response.ok) {
//             throw new Error(result.error || 'Login failed');
//         }

//         // Store token and redirect
//         localStorage.setItem('token', result.token);
//         window.location.href = '/dashboard';  // Redirect to dashboard page
//     } catch (error) {
//         const errorDiv = document.getElementById('error-message');
//         errorDiv.textContent = error.message;
//         errorDiv.classList.remove('hidden');
//     }
// });

// // File: public/js/register.js
// document.getElementById('register-form').addEventListener('submit', async (e) => {
//     e.preventDefault();

//     const formData = new FormData(e.target);
//     const data = {
//         username: formData.get('username'),
//         password: formData.get('password'),
//         email: formData.get('email'),
//         phone: formData.get('phone')
//     };

//     try {
//         const response = await fetch('/api/v1/register', {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify(data)
//         });

//         const result = await response.json();

//         if (!response.ok) {
//             throw new Error(result.error || 'Registration failed');
//         }

//         // Redirect to login page on success
//         window.location.href = '/login';
//     } catch (error) {
//         const errorDiv = document.getElementById('error-message');
//         errorDiv.textContent = error.message;
//         errorDiv.classList.remove('hidden');
//     }
// });