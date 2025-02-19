import * as router from './router.js'
import { setCookie, delCookie, getCookie } from './common.js'

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

let currentIndex = 0;
let intervalId;

!function () {
  document.querySelector('#search').addEventListener('click', (e) => {
    const keyword = document.querySelector('#search-input').value
    window.location.href = router.OperationRouters['search'] + '?keyword=' + keyword
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
      window.location.href = router.OperationRouters['search'] + '?category=' + ele.getAttribute('href')
    })
  })

}()

!function () {
  document.querySelector('#product-showcase').addEventListener('click', (e) => {
    const target = e.target
    const product_item = target.closest('.product-item-g')
    if (product_item !== null) {
      window.location.href = router.OperationRouters['getProduct'] + '?product_id=' + product_item.dataset['productId']
    }
  })
}()

function NewProductItem(product) {
  const item = document.createElement('div')
  item.classList.add('col-2', 'product-item-g')
  item.dataset['productId'] = product.Id

  function divElement(className, innerHTML) {
    const div = document.createElement('div')
    div.classList.add(className)
    div.innerHTML = innerHTML
    return div
  }

  const img = divElement('item-img', `<img src="${product.Picture}">`)
  const name = divElement('item-name', product.Name)
  const price = divElement('item-price', `<span class="currency">￥</span><span class="price">${product.Price}</span>`)
  const sold = divElement('item-sold', `<span>${product.Sold}</span><span>+人购买</span>`)
  item.appendChild(img)
  item.appendChild(name)
  item.appendChild(price)
  item.appendChild(sold)
  return item
}

!function () {
  document.querySelectorAll('.product-list-tabtitle li a').forEach((item) => {
    item.addEventListener('click', (e) => {
      e.preventDefault()
      console.log(item.dataset['sub'])
      axios({
        'method': 'GET',
        'url': router.OperationRouters['switchShowcase'] + '?sub=' + item.dataset['sub']
      }).then((res) => {
        console.log(res)
        document.querySelector('.product-list-tabtitle a.active').classList.remove('active')
        item.classList.add('active')
        document.querySelector('#product-showcase').innerHTML = ''
        res.data.products.forEach((product) => {
          const item = NewProductItem(product)
          document.querySelector('#product-showcase').appendChild(item)
        })
      }).catch((err) => {
        console.log(err)
      })
    })
  })
}()

!function () {
  const localStoragetoken = localStorage.getItem('token')
  const cookieToken = getCookie('token')

  if (!cookieToken || !localStoragetoken) {
    return
  }

  axios({
    'method': 'POST',
    'url': router.OperationRouters['verify'],
    data: {
      token: localStoragetoken
    }
  }).then((res) => {
    console.log(res)
    const resp = res.data.resp
    if (!resp.ok) {
      return
    }

    document.querySelector("#wel-bar").classList.add("d-none")
    document.querySelector("#user-bar").classList.remove("d-none")
    document.querySelector("#nav-user-name").textContent = resp["Name"]
    document.querySelector("#need-login").classList.add("d-none")
    document.querySelector("#logined").classList.remove("d-none")
    document.querySelector("#logout").classList.remove("d-none")
    document.querySelector("#card-user-name").textContent = resp["Name"]
    document.querySelector("#user-img").src = resp["Img"]
    // document.querySelector('#login').style.display = 'none'
    // document.querySelector('#register').style.display = 'none'
    // document.querySelector('#user-info').style.display = 'block'
    // document.querySelector('#user-name').innerText = res.data.username
  }).catch((err) => {
    console.error(err)
  })
}()