function modalCtrl() {
  const modal = document.querySelector('.modal.message-modal')
  const modalTitle = modal.querySelector('.modal-title')
  const modalMsg = modal.querySelector('.modal-body p')
  const modalClose = modal.querySelector('.modal .btn-close')
  const modalConfirm = modal.querySelector('.modal-footer button')
  modalClose.addEventListener('click', () => {
    modal.style.display = 'none'
  })

  modalConfirm.addEventListener('click', () => {
    modal.style.display = 'none'
  })

  return function (title, msg) {
    modal.style.display = 'block'
    modalTitle.textContent = title
    modalMsg.textContent = msg
  }
}

function setCookie(key, value, expires = 7, path = "/") {
  const cookieExpire = "expires=" + new Date(Date.now() + expires * 24 * 60 * 60 * 1000).toUTCString()
  const cookiePath = `path=${path}`
  document.cookie = `${key}=${value}; ${cookieExpire}; ${cookiePath};`
}

function delCookie(key) {
  document.cookie = `${key}=; expires=Thu, 01 Jan 1970 00:00:00 GMT`
}

function getCookie(key) {
  var arr, reg = new RegExp("(^| )" + key + "=([^;]*)(;|$)");
  if (arr = document.cookie.match(reg))
    return arr[2]
  else
    return '';
}
export default modalCtrl;
export { setCookie, delCookie, getCookie };