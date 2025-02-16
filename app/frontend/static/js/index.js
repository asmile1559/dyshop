import * as router from './router.js'
import modalCtrl from './common.js'
import { setCookie, delCookie, getCookie } from './common.js'

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

let currentIndex = 0;
let intervalId;

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.GETReqRouters['search'] + '?keyword=' + keyword
  })
}()

!function () {
  function showSlide(index) {
    const slides = document.querySelectorAll('.mycarousel-item');
    const dots = document.querySelectorAll('.dot');
    const totalSlides = slides.length;

    if (index >= totalSlides) {
      currentIndex = 0;
    } else if (index < 0) {
      currentIndex = totalSlides - 1;
    } else {
      currentIndex = index;
    }

    const offset = -currentIndex * 100;
    document.querySelector('.mycarousel-inner').style.transform = `translateX(${offset}%)`;

    slides.forEach((slide, idx) => {
      slide.classList.toggle('active', idx === currentIndex);
    });

    dots.forEach((dot, idx) => {
      dot.classList.toggle('active', idx === currentIndex);
    });
  }

  document.querySelector('.mycarousel-control.prev').addEventListener('click', function () {
    showSlide(currentIndex - 1);
  });

  document.querySelector('.mycarousel-control.next').addEventListener('click', function () {
    showSlide(currentIndex + 1);
  });

  document.querySelectorAll('.dot').forEach((dot, index) => {
    dot.addEventListener('click', function () {
      showSlide(index);
    });
  });

  intervalId = setInterval(function () {
    showSlide(currentIndex + 1);
  }, 3000);

  document.querySelectorAll('.mycarousel-item').forEach(ele => {
    ele.addEventListener('click', (e) => {
      window.location.href = router.GETReqRouters['search'] + '?category=' + ele.getAttribute('href')
    })
  })

}()

!function () {
  document.querySelectorAll('.product-item-g').forEach((item) => {
    console.log(item)
    item.addEventListener('click', (e) => {
      const productId = item.dataset['productId']
      window.location.href = router.GETReqRouters['product'] + '?product_id=' + productId
    })
  })
}()