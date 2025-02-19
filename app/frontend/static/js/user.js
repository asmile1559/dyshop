import * as router from './router.js'
import modalCtrl from './common.js'

function panelSwitch() {
  const userPanel = document.querySelector('.user-panel')
  const userPanelInfo = document.querySelector('.user-panel .info')
  const userPanelAccount = document.querySelector('.user-panel .account')
  const userPanelAddress = document.querySelector('.user-panel .address')
  const userPanelProduct = document.querySelector('.user-panel .product')

  const userInfoPage = document.querySelector('.user-info-page')
  const accountSettingPage = document.querySelector('.account-setting-page')
  const addressManagementPage = document.querySelector('.address-management-page')
  const productManagementPage = document.querySelector('.product-management-page')

  userPanel.addEventListener('click', (e) => {

    if (e.target.tagName !== 'BUTTON') {
      return
    }

    if (e.target.classList.contains('active')) {
      return
    }
    const p = document.querySelector('.user-panel .active')
    document.querySelector(p.dataset["tpage"]).classList.add('d-none')
    p.classList.remove('active')
    if (e.target.textContent === '个人资料') {
      userPanelInfo.classList.add('active')
      userInfoPage.classList.remove('d-none')
    } else if (e.target.textContent === '账号') {
      userPanelAccount.classList.add('active')
      accountSettingPage.classList.remove('d-none')
    } else if (e.target.textContent === '收货地址') {
      userPanelAddress.classList.add('active')
      addressManagementPage.classList.remove('d-none')
    } else if (e.target.textContent === '商品管理') {
      userPanelProduct.classList.add('active')
      productManagementPage.classList.remove('d-none')
    }
  })
}

function info(modaler) {
  const updateInfoBtn = document.getElementById('update-info-btn')
  const updateImgBtn = document.getElementById('update-img-btn')

  const usernameInput = document.getElementById('user-name')
  const usersignInput = document.getElementById('user-sign')
  const usergenderInput = document.getElementById('user-gender')
  const userbirthdayInput = document.getElementById('user-birthday')
  const userImgInput = document.getElementById('user-img')
  const allUsernames = document.querySelectorAll('.username')
  const allUsersigns = document.querySelectorAll('.usersign')
  const allUserImgs = document.querySelectorAll('.userimg')

  let lastUsername = usernameInput.value
  let lastUsersign = usersignInput.value
  let lastUsergender = usergenderInput.value
  let lastUserbirthday = userbirthdayInput.value

  updateInfoBtn.addEventListener('click', () => {
    if (usernameInput.value === lastUsername
      && usersignInput.value === lastUsersign
      && usergenderInput.value === lastUsergender
      && userbirthdayInput.value === lastUserbirthday) {
      return
    }
    axios({
      'method': 'post',
      'url': router.POSTReqRouters['updateInfo'],
      'data': {
        'userName': usernameInput.value,
        'userSign': usersignInput.value,
        'userGender': usergenderInput.value,
        'userBirthday': userbirthdayInput.value
      }
    }).then((res) => {
      console.log(res)
      if (lastUsername !== usernameInput.value) {
        allUsernames.forEach(e => {
          e.innerHTML = usernameInput.value
        });
      }

      if (lastUsersign !== usersignInput.value) {
        allUsersigns.forEach(e => {
          e.innerHTML = usersignInput.value
        });
      }
      lastUsername = usernameInput.value
      lastUsersign = usersignInput.value
      lastUsergender = usergenderInput.value
      lastUserbirthday = userbirthdayInput.value
    }).catch((err) => {
      modaler('错误', '一个错误发生了，请稍后再试。')
      if (lastUsername !== usernameInput.value) {
        usernameInput.value = lastUsername
      }
      if (lastUsersign !== usersignInput.value) {
        usersignInput.value = lastUsersign
      }
      console.log(err)
    })
  })

  updateImgBtn.addEventListener('click', () => {
    document.getElementById('user-img').click();
  })

  userImgInput.addEventListener('change', (e) => {
    const file = e.target.files[0]
    const formData = new FormData()
    formData.append('type', 'userImg')
    formData.append('userImg', file)
    axios({
      'method': 'post',
      'url': router.POSTReqRouters['updateImg'],
      'data': formData,
      'headers': {
        'Content-Type': 'multipart/form-data'
      }
    }).then((res) => {
      allUserImgs.forEach(e => {
        e.src = res.data.url
      });
    }).catch((err) => {
      modaler('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })
}

function account(modaler) {
  const accountSettingPage = document.querySelector('.account-setting-page')
  // const updatePhoneBtn = document.getElementById('update-phone-btn')
  // const updateEmailBtn = document.getElementById('update-email-btn')
  // const updatePasswordBtn = document.getElementById('update-password-btn')
  // const delAccountBtn = document.getElementById('del-account-btn')

  const oldPhoneInput = document.getElementById('user-phone')
  const newPhoneInput = document.getElementById('new-user-phone')
  const oldEmailInput = document.getElementById('user-email')
  const newEmailInput = document.getElementById('new-user-email')
  const oldPasswordInput = document.getElementById('user-password')
  const newPasswordInput = document.getElementById('new-user-password')
  const confirmPasswordInput = document.getElementById('confirm-user-password')

  const becomeMerchantLink = document.getElementById('user-role-marchant')

  accountSettingPage.addEventListener('click', (e) => {
    const target = e.target
    console.log(target.tagName)
    if (target.tagName !== 'A' && target.tagName !== 'BUTTON') {
      return
    }

    if (target.id === 'del-account-btn') {
      if (!confirm('确认删除账户？')) {
        return
      }
    }

    console.log(target)
    const data = {
      actionType: target.dataset["actionType"]
    }
    switch (target.dataset["actionType"]) {
      case "role":
        break
      case "phone":
        if (oldPhoneInput.value === newPhoneInput.value || newPhoneInput.value === '') {
          return
        }
        data['phone'] = newPhoneInput.value
        break
      case "email":
        if (oldEmailInput.value === newEmailInput.value || newEmailInput.value === '') {
          return
        }
        data['email'] = newEmailInput.value
        break
      case "password":
        if (newPasswordInput.value === '' || confirmPasswordInput.value === '') {
          return
        }
        data["old_password"] = oldPasswordInput.value
        data["new_password"] = newPasswordInput.value
        data["confirm_password"] = confirmPasswordInput.value
        break
      case "del":
        data["delete_account"] = true
        break
    }

    axios({
      'method': 'post',
      'url': router.POSTReqRouters['updateInfo'],
      'data': data
    }).then((res) => {
      console.log(res)
      switch (target.dataset["actionType"]) {
        case "role":
          break
        case "phone":
          oldPhoneInput.value = newPhoneInput.value
          newPhoneInput.value = ''
          break
        case "email":
          oldEmailInput.value = newEmailInput.value
          newEmailInput.value = ''
          break
        case "password":
          oldPasswordInput.value = ''
          newPasswordInput.value = ''
          confirmPasswordInput.value = ''
          break
        case "del":
          window.location.href = router.GETReqRouters['home']
          document.cookie = ''
          localStorage.clear()
          break
      }
    }).catch((err) => {
      modaler('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })

  })
}

function CreateNewTableRow(id, name, phone, province, city, district, street, detail, isDefault) {
  // could be optimized by using clone template

  const tr = document.createElement('tr')
  tr.dataset['addrId'] = id
  const createTd = (className, content) => {
    const td = document.createElement('td')
    td.classList.add(className)
    td.innerHTML = `<span>${content}</span>`
    return td
  }

  tr.appendChild(createTd('name-column', name))
  tr.appendChild(createTd('phone-column', phone))

  const regionTd = document.createElement('td')
  regionTd.classList.add('region-column')
  regionTd.innerHTML = `
    <span>
      <span class="province">${province} </span>
      <span class="city">${city} </span>
      <span class="district">${district} </span>
      <span class="street">${street}</span>
    </span>`
  tr.appendChild(regionTd)

  tr.appendChild(createTd('detail-column', detail))

  const opTd = document.createElement('td')
  opTd.classList.add('operation-column')
  opTd.innerHTML = `
    <a class="implicit-link underline-link" href="#"><span>修改</span></a>
    <a class="implicit-link underline-link" href="#"><span>删除</span></a>`
  tr.appendChild(opTd)

  const defaultTd = document.createElement('td')
  defaultTd.classList.add('default-column')
  defaultTd.innerHTML = isDefault
    ? `<a class="implicit-link default">默认地址</a>`
    : `<a class="implicit-link underline-link" href="#"><span>设为默认</span></a>`
  tr.appendChild(defaultTd)

  return tr
}

function address(modaler) {
  const myAddresses = document.querySelector('.my-addresses')
  const addressModal = document.querySelector('.modal.address-modal')
  const closeAddressModal = document.querySelector('.modal.address-modal .btn-close')
  const cancelAddressModal = document.querySelector('.modal.address-modal .btn-cancel')
  const confirmAddressModal = document.querySelector('.modal.address-modal .btn-confirm')

  let addressId = ''
  const addressName = document.getElementById('address-name')
  const addressPhone = document.getElementById('address-phone')
  const addressProvince = document.getElementById('address-province')
  const addressCity = document.getElementById('address-city')
  const addressDistrict = document.getElementById('address-district')
  const addressStreet = document.getElementById('address-street')
  const addressDetail = document.getElementById('address-detail')

  closeAddressModal.addEventListener('click', () => {
    addressModal.style.display = 'none'
  })

  cancelAddressModal.addEventListener('click', () => {
    addressModal.style.display = 'none'
  })

  confirmAddressModal.addEventListener('click', () => {

    if (addressName.value === ''
      || addressPhone.value === ''
      || addressProvince.value === ''
      || addressCity.value === ''
      || addressDistrict.value === ''
      || addressStreet.value === ''
      || addressDetail.value === '') {
      addressModal.style.display = 'none'
      modaler('错误', '请填写完整信息。')
      return
    }

    axios({
      'method': 'post',
      'url': router.POSTReqRouters['updateInfo'],
      'data': {
        'address': {
          'id': addressId,
          'name': addressName.value,
          'phone': addressPhone.value,
          'province': addressProvince.value,
          'city': addressCity.value,
          'district': addressDistrict.value,
          'street': addressStreet.value,
          'detail': addressDetail.value
        }
      }
    }).then((res) => {
      addressModal.style.display = 'none'
      if (addressId === '-1') {
        if (myAddresses.querySelector('.no-address')) {
          myAddresses.removeChild(myAddresses.querySelector('.no-address'))
        }
        const lastrow = myAddresses.querySelector('.last-row')
        const data = res.data
        lastrow.parentNode.insertBefore(CreateNewTableRow(
          data["id"],
          addressName.value,
          addressPhone.value,
          addressProvince.value,
          addressCity.value,
          addressDistrict.value,
          addressStreet.value,
          addressDetail.value,
          data["isDefault"]
        ), lastrow)
      } else {
        const row = document.querySelector(`[data-addr-id="${addressId}"]`)
        row.querySelector('.name-column span').textContent = addressName.value
        row.querySelector('.phone-column span').textContent = addressPhone.value
        row.querySelector('.region-column span.province').textContent = addressProvince.value
        row.querySelector('.region-column span.city').textContent = addressCity.value
        row.querySelector('.region-column span.district').textContent = addressDistrict.value
        row.querySelector('.region-column span.street').textContent = addressStreet.value
        row.querySelector('.detail-column span').textContent = addressDetail.value
        console.log(row)
      }
    }).catch((err) => {
      addressModal.style.display = 'none'
      modaler('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })

  })
  myAddresses.addEventListener('click', (e) => {
    const target = e.target
    console.log(target)
    if (target.tagName === 'SPAN') {
      if (target.innerHTML === '修改') {
        addressModal.style.display = 'block'
        addressModal.querySelector('.modal-title').textContent = '修改地址'
        const row = target.parentNode.parentNode.parentNode
        console.log(row)
        addressId = row.dataset["addrId"]
        console.log(addressId)
        addressName.value = row.querySelector('.name-column').textContent.trim()
        addressPhone.value = row.querySelector('.phone-column').textContent.trim()
        addressProvince.value = row.querySelector('.province').textContent.trim()
        addressCity.value = row.querySelector('.city').textContent.trim()
        addressDistrict.value = row.querySelector('.district').textContent.trim()
        addressStreet.value = row.querySelector('.street').textContent.trim()
        addressDetail.value = row.querySelector('.detail-column').textContent.trim()
      } else if (target.innerHTML === '删除') {
        const row = target.parentNode.parentNode.parentNode
        if (row.dataset["addrId"] === myAddresses.dataset["defaultId"]) {
          modaler('错误', '默认地址不能删除。')
          return
        }
        axios({
          'method': 'post',
          'url': router.POSTReqRouters['updateInfo'],
          'data': {
            'address': {
              "id": row.dataset["addrId"],
              "delete": true
            }
          }
        }).then((res) => {
          myAddresses.removeChild(row)
        }).catch((err) => {
          modaler('错误', '一个错误发生了，请稍后再试。')
          console.log(err)
        })
      } else if (target.innerHTML === '设为默认') {
        const row = target.parentNode.parentNode.parentNode
        let defaultRowId = myAddresses.dataset["defaultId"]
        const defaultRow = document.querySelector(`[data-addr-id="${defaultRowId}"]`)
        axios({
          'method': 'post',
          'url': router.POSTReqRouters['updateInfo'],
          'data': {
            'address': {
              'id': row.dataset["addrId"],
              "set_default": true
            }
          }
        }).then((res) => {
          const dc = row.querySelector('.default-column a')
          dc.classList.remove('underline-link')
          dc.classList.add('default')
          dc.innerHTML = '默认地址'
          dc.removeAttribute('href')

          const ddc = defaultRow.querySelector('.default-column a')
          ddc.classList.remove('default')
          ddc.classList.add('underline-link')
          ddc.innerHTML = '<span>设为默认</span>'
          ddc.href = '#'

          const newRow = row.cloneNode(true)
          myAddresses.removeChild(row)
          myAddresses.insertBefore(newRow, defaultRow)

          myAddresses.dataset["defaultId"] = row.dataset["addrId"]
        }).catch((err) => {
          modaler('错误', '一个错误发生了，请稍后再试。')
          console.log(err)
        })
      } else if (target.innerHTML === '添加新地址') {
        addressModal.style.display = 'block'
        addressId = '-1'
        addressModal.querySelector('.modal-title').textContent = '添加新地址'
        addressName.value = ''
        addressPhone.value = ''
        addressProvince.value = ''
        addressCity.value = ''
        addressDistrict.value = ''
        addressStreet.value = ''
        addressDetail.value = ''
      }
      console.log(target)
    }
  })
}

function NewAccordionItem(id, img, name, shortname, desc, categories, price, stock) {
  const accordionItem = document.createElement('div');
  accordionItem.classList.add('accordion-item');
  accordionItem.id = `accordion-item${id}`;
  accordionItem.dataset.productId = id;

  const accordionHeader = document.createElement('h2');
  accordionHeader.classList.add('accordion-header');

  const accordionButton = document.createElement('button');
  accordionButton.classList.add('accordion-button', 'collapsed');
  accordionButton.type = 'button';
  accordionButton.dataset.bsToggle = 'collapse';
  accordionButton.dataset.bsTarget = `#accordion-item${id}-body`;
  accordionButton.textContent = name;

  accordionHeader.appendChild(accordionButton);
  accordionItem.appendChild(accordionHeader);

  const accordionCollapse = document.createElement('div');
  accordionCollapse.id = `accordion-item${id}-body`;
  accordionCollapse.classList.add('accordion-collapse', 'collapse');
  accordionCollapse.dataset.bsParent = '#product-accordion';

  const accordionBody = document.createElement('div');
  accordionBody.classList.add('accordion-body');

  const createInputDiv = (labelText, inputId, inputType, inputValue, maxLength, inputRole) => {
    const div = document.createElement('div');
    div.classList.add('d-flex', 'align-items-center');

    const label = document.createElement('label');
    label.htmlFor = inputId;
    label.classList.add('form-label', 'mb-0');
    label.textContent = labelText;

    const input = document.createElement('input');
    input.id = inputId;
    input.type = inputType;
    input.classList.add('form-control', 'bg-body-tertiary', 'w-75', 'ms-2', 'disabled');
    input.setAttribute('value', inputValue);
    input.dataset.inputRole = inputRole;

    const span = document.createElement('span');
    span.classList.add('ms-2');
    span.textContent = `(${inputValue.length}/${maxLength})`;

    div.appendChild(label);
    div.appendChild(input);
    div.appendChild(span);

    return div;
  };

  const createButton = (btnClass, btnRole, btnText) => {
    const button = document.createElement('button');
    button.classList.add('btn', btnClass, 'btn-sm', 'w-25');
    button.dataset.itemId = `#accordion-item${id}`;
    button.dataset.btnRole = btnRole;
    button.textContent = btnText;
    return button;
  };

  const createCategorySpan = (cat) => {
    const categorySpan = document.createElement('span');
    categorySpan.classList.add('category');

    const categoryNameSpan = document.createElement('span');
    categoryNameSpan.classList.add('category-name');
    categoryNameSpan.textContent = cat;

    const categoryDelLink = document.createElement('a');
    categoryDelLink.href = 'javascript:void(0)';
    categoryDelLink.classList.add('category-del', 'd-none');
    categoryDelLink.dataset.executor = `#categories${id}`;
    categoryDelLink.dataset.aRole = 'categoryDel';
    categoryDelLink.innerHTML = '&otimes;';

    categorySpan.appendChild(categoryNameSpan);
    categorySpan.appendChild(categoryDelLink);
    return categorySpan;
  };

  const row = document.createElement('div');
  row.classList.add('row');

  const col3 = document.createElement('div');
  col3.classList.add('col-3', 'text-center');

  const imgElement = document.createElement('img');
  imgElement.src = img;
  imgElement.alt = 'productImg';
  imgElement.classList.add('img-fluid');
  imgElement.id = `img${id}`;
  imgElement.dataset.status = 'no-change';

  const imgInput = document.createElement('input');
  imgInput.type = 'file';
  imgInput.id = `img-input${id}`;
  imgInput.accept = 'image/*';
  imgInput.style.display = 'none';
  imgInput.dataset.inputRole = 'img';
  imgInput.dataset.productId = id;
  // imgInput.addEventListener('change', (e) => {
  //   const target = e.target
  //   console.log(target.dataset["productId"])

  //   const file = target.files[0]
  //   const request = indexedDB.open('myProduct', 1)
  //   request.onupgradeneeded = (e) => {
  //     const db = e.target.result
  //     if (!db.objectStoreNames.contains('productImgFile')) {
  //       db.createObjectStore('productImgFile', { keyPath: 'id' })
  //     }
  //   }

  //   request.onsuccess = (e) => {
  //     const db = e.target.result
  //     const transaction = db.transaction('productImgFile', 'readwrite')
  //     const store = transaction.objectStore('productImgFile')
  //     console.log(target.dataset["productId"])
  //     console.log(file)
  //     const req = store.put({ id: target.dataset["productId"], data: file })

  //     req.onsuccess = () => {
  //       const reader = new FileReader()
  //       reader.onload = (e) => {
  //         target.previousElementSibling.src = e.target.result
  //         target.previousElementSibling.dataset["status"] = 'changed'
  //       }
  //       reader.readAsDataURL(file)
  //     }

  //     req.onerror = (e) => {
  //       modaler('错误', '上传图片失败, 请稍后再试。')
  //     }
  //   }
  //   request.onerror = (e) => {
  //     console.log(e)
  //   }
  // })

  const imgButton = document.createElement('button');
  imgButton.classList.add('btn', 'btn-success', 'mt-2', 'd-none');
  imgButton.dataset.itemId = `#accordion-item${id}`;
  imgButton.textContent = '修改商品图片';

  col3.appendChild(imgElement);
  col3.appendChild(imgInput);
  col3.appendChild(imgButton);

  const col8 = document.createElement('div');
  col8.classList.add('col-8', 'myproduct');

  const nameDiv = createInputDiv('商品名：', `name${id}`, 'text', name, 30, 'name');
  // nameDiv.querySelector('input').addEventListener('input', (e) => {
  //   const input = e.target;
  //   const span = input.nextElementSibling;
  //   if (input.value.length > 30) {
  //     input.value = input.value.slice(0, 30);
  //   }
  //   span.textContent = `(${input.value.length}/30)`;
  // });

  const shortnameDiv = createInputDiv('简略商品名：', `shortname${id}`, 'text', shortname, 10, 'shortname');
  // shortnameDiv.querySelector('input').addEventListener('input', (e) => {
  //   const input = e.target;
  //   const span = input.nextElementSibling;
  //   if (input.value.length > 10) {
  //     input.value = input.value.slice(0, 10);
  //   }
  //   span.textContent = `(${input.value.length}/10)`;
  // });

  col8.appendChild(nameDiv);
  col8.appendChild(shortnameDiv);
  // col8.appendChild(createInputDiv('商品名：', `name${id}`, 'text', name, 30, 'name'));
  // col8.appendChild(createInputDiv('简略商品名：', `shortname${id}`, 'text', shortname, 10, 'shortname'));

  const descSpan = document.createElement('span');
  descSpan.textContent = '商品的描述信息:';

  const descTextarea = document.createElement('textarea');
  descTextarea.classList.add('form-control', 'bg-body-tertiary', 'mt-1', 'disabled', 'js-desc');
  descTextarea.id = `desc${id}`;
  descTextarea.rows = 3;
  descTextarea.dataset.inputRole = 'desc';
  descTextarea.textContent = desc;

  const descTextEnd = document.createElement('div');
  descTextEnd.classList.add('text-end');
  const descTextSpan = document.createElement('span');
  descTextSpan.textContent = `(${desc.length}/200)`;
  descTextEnd.appendChild(descTextSpan);

  // descTextarea.addEventListener('input', (e) => {
  //   const textarea = e.target;
  //   const span = textarea.nextElementSibling.querySelector('span');
  //   if (textarea.value.length > 200) {
  //     textarea.value = textarea.value.slice(0, 200);
  //   }
  //   span.textContent = `(${textarea.value.length}/200)`;
  // });

  col8.appendChild(descSpan);
  col8.appendChild(descTextarea);
  col8.appendChild(descTextEnd);

  const row2 = document.createElement('div');
  row2.classList.add('row');

  const categoriesDiv = document.createElement('div');
  categoriesDiv.classList.add('mt-2', 'categories');
  categoriesDiv.id = `categories${id}`;

  const categoriesSpan = document.createElement('span');
  categoriesSpan.textContent = `商品类别 (${categories.length}/5): `;
  categoriesDiv.appendChild(categoriesSpan);

  categories.forEach(cat => categoriesDiv.appendChild(createCategorySpan(cat)));

  const categoryAddDiv = document.createElement('div');
  categoryAddDiv.classList.add('d-none');

  const categoryAddInnerDiv = document.createElement('div');
  categoryAddInnerDiv.classList.add('d-flex', 'align-items-center', 'mt-2', 'mb-2', 'w-50', 'gap-4', 'category-add');

  const categorySelect = document.createElement('select');
  categorySelect.classList.add('form-select');
  categorySelect.innerHTML = `
    <option value="0" selected>选择商品类别</option>
    <option value="服装">服装</option>
    <option value="鞋子">鞋子</option>
    <option value="儿童用品">儿童用品</option>
    <option value="家具">家具</option>
    <option value="电子产品">电子产品</option>
    <option value="其他">其他</option>
  `;

  const categoryAddLinkSpan = document.createElement('span');
  categoryAddLinkSpan.classList.add('d-flex', 'align-items-center');
  categoryAddLinkSpan.style.width = '40%';

  const categoryAddLink = document.createElement('a');
  categoryAddLink.href = 'javascript:void(0)';
  categoryAddLink.classList.add('disable-link');
  categoryAddLink.dataset.executor = `#categories${id}`;
  categoryAddLink.dataset.aRole = 'categoryAdd';
  categoryAddLink.textContent = '+ 添加类别';

  categoryAddLinkSpan.appendChild(categoryAddLink);
  categoryAddInnerDiv.appendChild(categorySelect);
  categoryAddInnerDiv.appendChild(categoryAddLinkSpan);
  categoryAddDiv.appendChild(categoryAddInnerDiv);

  row2.appendChild(categoriesDiv);
  row2.appendChild(categoryAddDiv);

  const row3 = document.createElement('div');
  row3.classList.add('row', 'align-items-center', 'mt-2');

  const priceDiv = createInputDiv('价格：', `price${id}`, 'number', price, 0, 'price');
  const stockDiv = createInputDiv('库存：', `stock${id}`, 'number', stock, 0, 'stock');

  priceDiv.classList.add('col-6');
  priceDiv.removeChild(priceDiv.querySelector('span'));
  // priceDiv.querySelector('input').addEventListener('blur', (e) => {
  //   const input = e.target;
  //   if (input.value < 0) {
  //     input.value = 0;
  //   }
  // });
  stockDiv.classList.add('col-6');
  stockDiv.removeChild(stockDiv.querySelector('span'));
  // stockDiv.querySelector('input').addEventListener('blur', (e) => {
  //   const input = e.target;
  //   if (input.value < 0) {
  //     input.value = 0;
  //   }
  //   input.value = Math.round(input.value);
  // });
  row3.appendChild(priceDiv);
  row3.appendChild(stockDiv);

  const textEndDiv1 = document.createElement('div');
  textEndDiv1.classList.add('text-end', 'mt-3', 'onpaper');
  textEndDiv1.appendChild(createButton('btn-outline-primary', 'productEdit', '编辑'));
  textEndDiv1.appendChild(createButton('btn-outline-danger', 'productDel', '删除'));

  const textEndDiv2 = document.createElement('div');
  textEndDiv2.classList.add('text-end', 'mt-3', 'd-none');
  textEndDiv2.appendChild(createButton('btn-outline-success', 'productSave', '保存'));
  textEndDiv2.appendChild(createButton('btn-outline-secondary', 'productDiscard', '放弃'));

  row.appendChild(col3);
  row.appendChild(col8);
  accordionBody.appendChild(row);
  accordionBody.appendChild(row2);
  accordionBody.appendChild(row3);
  accordionBody.appendChild(textEndDiv1);
  accordionBody.appendChild(textEndDiv2);

  accordionCollapse.appendChild(accordionBody);
  accordionItem.appendChild(accordionCollapse);

  return accordionItem;
}

function product(modaler) {
  const newProduct = document.querySelector('.new-product')
  const productAccordion = document.querySelector('#product-accordion')
  const accordionCount = document.querySelector('.myproducts .count>span:nth-child(2)')

  const newProductName = document.querySelector('#new-product-name')
  const newProductShortname = document.querySelector('#new-product-shortname')
  const newProductDesc = document.querySelector('#new-product-desc')
  const newProductCategories = document.querySelector('.new-product .categories')
  const newProductSelect = document.querySelector('#new-product-select')
  const newProductPrice = document.querySelector('#new-product-price')
  const newProductStock = document.querySelector('#new-product-stock')
  const newProductImg = document.querySelector('#new-product-img')

  const uploadProductImg = document.getElementById('upload-product-img')
  const addCategory = document.getElementById('add-category')
  const confirmNewProduct = document.getElementById('confirm-new-product')
  const cancelNewProduct = document.getElementById('cancel-new-product')
  const productImgInput = document.getElementById('new-product-img-input')

  cancelNewProduct.addEventListener('click', () => {
    newProductName.value = ''
    newProductShortname.value = ''
    newProductDesc.value = ''
    newProductPrice.value = ''
    newProductStock.value = ''
    newProductImg.src = '/static/src/basic/noimg.svg'
    newProductCategories.innerHTML = `<span>商品类别 (0/5): </span>`
    newProductSelect.children[0].selected = true
    const request = indexedDB.open('myProduct', 1)
    request.onupgradeneeded = (e) => {
      const db = e.target.result
      if (db.objectStoreNames.contains('productImgFile')) {
        db.deleteObjectStore('productImgFile')
      }
    }

    request.onsuccess = (e) => {
      const db = e.target.result
      const transaction = db.transaction('productImgFile', 'readwrite')
      const store = transaction.objectStore('productImgFile')
      store.delete(-1)
    }
  })

  uploadProductImg.addEventListener('click', () => {
    productImgInput.click()
  })

  productImgInput.addEventListener('change', (e) => {
    const file = e.target.files[0]

    if (file) {
      const request = indexedDB.open('myProduct', 1)
      request.onupgradeneeded = (e) => {
        const db = e.target.result
        if (!db.objectStoreNames.contains('productImgFile')) {
          db.createObjectStore('productImgFile', { keyPath: 'id' })
        }
      }

      request.onsuccess = (e) => {
        const db = e.target.result
        console.log(file)
        const transaction = db.transaction('productImgFile', 'readwrite')
        const store = transaction.objectStore('productImgFile')
        const req = store.put({ id: -1, data: file })

        req.onsuccess = () => {
          console.log("upload new product img success")
        }

        req.onerror = (e) => {
          modaler('错误', '上传图片失败, 请稍后再试。')
          return
        }
      }

      const reader = new FileReader()
      reader.onload = (e) => {
        newProductImg.src = e.target.result
      }
      reader.readAsDataURL(file)
    }
  })

  newProductCategories.addEventListener('click', (e) => {
    if (e.target.tagName === 'A') {
      const target = e.target
      console.log(target)
      console.log(target.parentNode)
      newProductCategories.removeChild(target.parentNode)
      newProductCategories.children[0].textContent = `商品类别 (${newProductCategories.children.length - 1}/5): `
    }
  })


  addCategory.addEventListener('click', () => {
    if (newProductSelect.value === '') {
      modaler('错误', '请选择一个类别。')
      return
    }
    if (newProductCategories.children.length >= 6) {
      modaler('错误', '最多添加5个类别。')
      return
    }
    const newCategory = document.createElement('span')
    newCategory.classList.add('category')
    newCategory.innerHTML =
      `<span class="category-name">${newProductSelect.value}</span>
  <a href="javascript:void(0)" class="category-del">
    &otimes;
  </a>`
    newProductCategories.appendChild(newCategory)
    newProductCategories.children[0].textContent = `商品类别 (${newProductCategories.children.length - 1}/5): `
  })

  newProductName.addEventListener('input', () => {
    if (newProductName.value.length > 30) {
      newProductName.value = newProductName.value.slice(0, 30)
    }
    newProductName.nextElementSibling.textContent = `(${newProductName.value.length}/30)`
  })

  newProductShortname.addEventListener('input', () => {
    if (newProductShortname.value.length > 10) {
      newProductShortname.value = newProductShortname.value.slice(0, 10)
    }
    newProductShortname.nextElementSibling.textContent = `(${newProductShortname.value.length}/10)`
  })

  newProductDesc.addEventListener('input', () => {
    if (newProductDesc.value.length > 200) {
      newProductDesc.value = newProductDesc.value.slice(0, 200)
    }
    console.log(newProductDesc.nextElementSibling)
    newProductDesc.nextElementSibling.textContent = `(${newProductDesc.value.length}/200)`
  })

  newProductPrice.addEventListener('blur', () => {
    if (newProductPrice.value < 0) {
      newProductPrice.value = 0
    }
  })

  newProductStock.addEventListener('blur', () => {
    if (newProductStock.value < 0) {
      newProductStock.value = 0
      return
    }
    newProductStock.value = Math.round(newProductStock.value)
  })

  confirmNewProduct.addEventListener('click', () => {
    const categories = []
    for (let i = 1; i < newProductCategories.children.length; i++) {
      categories.push(newProductCategories.children[i].querySelector('.category-name').textContent)
    }

    if (newProductImg.src === '/static/src/basic/noimg.svg'
      || newProductName.value === ''
      || newProductShortname.value === ''
      || newProductPrice.value === '' || newProductPrice.value == 0
      || newProductStock.value === '' || newProductStock.value == 0
      || categories.length === 0) {
      modaler('错误', '请填写完整信息。')
      return
    }

    const request = indexedDB.open('myProduct', 1)
    request.onupgradeneeded = (e) => {
      const db = e.target.result
      if (!db.objectStoreNames.contains('productImgFile')) {
        db.createObjectStore('productImgFile', { keyPath: 'id' })
      }
    }

    request.onsuccess = (e) => {
      const db = e.target.result
      const transaction = db.transaction('productImgFile', 'readwrite')
      const store = transaction.objectStore('productImgFile')
      const req = store.get(-1)

      req.onsuccess = (e) => {
        const file = e.target.result.data
        const formData = new FormData()
        formData.append('type', 'updateProductWithImg')
        formData.append('productImg', file)
        formData.append('productName', newProductName.value)
        formData.append('productShortname', newProductShortname.value)
        formData.append('productDesc', newProductDesc.value)
        formData.append('productCategories', JSON.stringify(categories))
        formData.append('productPrice', newProductPrice.value)
        formData.append('productStock', newProductStock.value)

        axios({
          'method': 'post',
          'url': router.POSTReqRouters['newProduct'],
          'data': formData,
          'headers': {
            'Content-Type': 'multipart/form-data'
          }
        }).then((res) => {
          console.log(res)
          const data = res.data.res
          const item = NewAccordionItem(
            data.id,
            data.img,
            newProductName.value,
            newProductShortname.value,
            newProductDesc.value,
            categories,
            newProductPrice.value,
            newProductStock.value
          )

          bindInputEvents(modaler, item)
          productAccordion.appendChild(item)
          accordionCount.textContent = parseInt(accordionCount.textContent) + 1
          if (accordionCount.textContent === '1') {
            productAccordion.classList.remove('d-none')
            productAccordion.nextElementSibling.classList.add('d-none')
          }
          newProductName.value = ''
          newProductName.nextElementSibling.textContent = '(0/30)'
          newProductShortname.value = ''
          newProductShortname.nextElementSibling.textContent = '(0/10)'
          newProductDesc.value = ''
          newProductDesc.nextElementSibling.textContent = '(0/200)'
          newProductPrice.value = ''
          newProductStock.value = ''
          newProductImg.src = '/static/src/basic/noimg.svg'
          newProductCategories.innerHTML = `<span>商品类别 (0/5): </span>`
          newProductSelect.children[0].selected = true
        }).catch((err) => {
          modaler('错误', '一个错误发生了，请稍后再试。')
          console.log(err)
        })
      }
    }
  })

  productAccordion.addEventListener('click', (e) => {
    const target = e.target
    if (target.tagName === 'BUTTON') {
      const role = target.dataset['btnRole']
      if (role === 'productDel') {
        if (!confirm('确认删除商品？')) {
          return
        }

        const thisAccordionItem = productAccordion.querySelector(target.dataset["itemId"])
        const productId = thisAccordionItem.dataset["productId"]

        axios({
          'method': 'post',
          'url': router.POSTReqRouters['deleteProduct'],
          'data': {
            'product_id': productId,
            'delete': true
          }
        }).then((res) => {
          productAccordion.removeChild(thisAccordionItem)
          accordionCount.textContent = parseInt(accordionCount.textContent) - 1
          if (accordionCount.textContent === '0') {
            productAccordion.classList.add('d-none')
            productAccordion.nextElementSibling.classList.remove('d-none')
          }
        }).catch((err) => {
          modaler('错误', '一个错误发生了，请稍后再试。')
          console.log(err)
        })
      } else if (role === 'productEdit') {
        const thisAccordionItem = productAccordion.querySelector(target.dataset["itemId"])
        const productId = thisAccordionItem.dataset["productId"]
        console.log(thisAccordionItem)

        const request = indexedDB.open('textBuffer', 1)
        request.onupgradeneeded = (e) => {
          const db = e.target.result
          if (!db.objectStoreNames.contains('accordionText')) {
            db.createObjectStore('accordionText', { keyPath: 'id' })
          }
        }

        thisAccordionItem.querySelectorAll('input[type="text"], input[type="number"]').forEach((input) => {
          input.setAttribute('value', input.value)
        })

        const text = thisAccordionItem.innerHTML
        console.log(text)

        request.onsuccess = (e) => {
          const db = e.target.result
          const transaction = db.transaction('accordionText', 'readwrite')
          const store = transaction.objectStore('accordionText')
          const req = store.put({ id: productId, data: text })

          req.onsuccess = () => {
            console.log("save accordion text success")
            thisAccordionItem.classList.add('edit')
          }

          req.onerror = (e) => {
            modaler('错误', '发生了一个错误，无法修改。')
            return
          }
        }
      } else if (role === 'productDiscard') {
        const thisAccordionItem = productAccordion.querySelector(target.dataset["itemId"])
        const productId = thisAccordionItem.dataset["productId"]

        const request = indexedDB.open('textBuffer', 1)
        request.onupgradeneeded = (e) => {
          const db = e.target.result
          if (!db.objectStoreNames.contains('accordionText')) {
            db.createObjectStore('accordionText', { keyPath: 'id' })
          }
        }

        request.onsuccess = (e) => {
          const db = e.target.result
          const transaction = db.transaction('accordionText', 'readwrite')
          const store = transaction.objectStore('accordionText')
          const req = store.get(productId)

          req.onsuccess = (e) => {
            const text = e.target.result.data
            store.delete(productId)
            console.log(text)
            thisAccordionItem.innerHTML = text
            bindInputEvents(modaler, thisAccordionItem)
            thisAccordionItem.classList.remove('edit')
          }

          req.onerror = (e) => {
            modaler('错误', '发生了一个错误，无法放弃修改。')
            return
          }
        }
      } else if (role === 'productSave') {
        const thisAccordionItem = productAccordion.querySelector(target.dataset["itemId"])
        const productId = thisAccordionItem.dataset["productId"]
        const productCategories = thisAccordionItem.querySelector('.categories')
        const name = thisAccordionItem.querySelector('#name' + productId).value
        const shortname = thisAccordionItem.querySelector('#shortname' + productId).value
        const desc = thisAccordionItem.querySelector('#desc' + productId).value
        const price = thisAccordionItem.querySelector('#price' + productId).value
        const stock = thisAccordionItem.querySelector('#stock' + productId).value
        const imgStatus = thisAccordionItem.querySelector('#img' + productId).dataset["status"]
        const categories = []
        for (let i = 1; i < productCategories.children.length; i++) {
          categories.push(productCategories.children[i].querySelector('.category-name').textContent)
        }

        if (name === '' ||
          shortname === '' ||
          desc === '' ||
          price === '' ||
          stock === '' ||
          categories.length === 0) {
          modaler('错误', '请填写完整信息。')
          return
        }

        console.log(imgStatus)
        const formData = new FormData()
        formData.append('type', 'updateProduct' + (imgStatus === 'no-change' ? 'WithoutImg' : 'WithImg'))
        if (imgStatus === 'changed') {
          const request = indexedDB.open('myProduct', 1)
          request.onupgradeneeded = (e) => {
            const db = e.target.result
            if (!db.objectStoreNames.contains('productImgFile')) {
              db.createObjectStore('productImgFile', { keyPath: 'id' })
            }
          }
          request.onsuccess = (e) => {
            const db = e.target.result
            const transaction = db.transaction('productImgFile', 'readwrite')
            const store = transaction.objectStore('productImgFile')
            const req = store.get(productId)

            req.onsuccess = (e) => {
              const img = e.target.result.data
              console.log(img)
              formData.append('productImg', img)
              formData.append('productId', productId)
              formData.append('productName', name)
              formData.append('productShortname', shortname)
              formData.append('productDesc', desc)
              formData.append('productPrice', price)
              formData.append('productStock', stock)
              formData.append('productCategories', JSON.stringify(categories))
              axios({
                'method': 'post',
                'url': router.POSTReqRouters['updateProduct'],
                'data': formData,
                'headers': {
                  'Content-Type': 'multipart/form-data'
                }
              }).then((res) => {
                console.log(res)
                thisAccordionItem.classList.remove('edit')
                thisAccordionItem.querySelector('.accordion-button').textContent = name
                thisAccordionItem.querySelector('#img' + productId).dataset["status"] = 'no-change'
              }).catch((err) => {
                console.log(err)
                modaler('错误', '发生了一个错误，无法保存修改。')
                return
              })
            }

            req.onerror = (e) => {
              modaler('错误', '发生了一个错误，无法保存修改。')
              return
            }
          }

          request.onerror = (e) => {
            modaler('错误', '发生了一个错误，无法保存修改。')
            return
          }
        } else {
          formData.append('productId', productId)
          formData.append('productName', name)
          formData.append('productShortname', shortname)
          formData.append('productDesc', desc)
          formData.append('productPrice', price)
          formData.append('productStock', stock)
          formData.append('productCategories', JSON.stringify(categories))
          axios({
            'method': 'post',
            'url': router.POSTReqRouters['updateProduct'],
            'data': formData,
            'headers': {
              'Content-Type': 'multipart/form-data'
            }
          }).then((res) => {
            console.log(res)
            thisAccordionItem.classList.remove('edit')
            thisAccordionItem.querySelector('.accordion-button').textContent = name
            thisAccordionItem.querySelector('#img' + productId).dataset["status"] = 'no-change'
          }).catch((err) => {
            console.log(err)
            modaler('错误', '发生了一个错误，无法保存修改。')
            return
          })
        }
      } else if (target.textContent === '修改商品图片') {
        const thisAccordionItem = productAccordion.querySelector(target.dataset["itemId"])
        const productId = thisAccordionItem.dataset["productId"]
        const productImgInput = thisAccordionItem.querySelector('#img-input' + productId)
        productImgInput.click()
      }
    } else if (target.tagName === 'A') {
      const target = e.target
      const role = target.dataset['aRole']
      if (role === "categoryDel") {
        console.log(target)
        console.log(target.dataset['executor'])
        const categories = document.querySelector(target.dataset['executor'])
        categories.removeChild(target.parentNode)
        categories.children[0].textContent = `商品类别 (${categories.children.length - 1}/5): `
      } else if (role === 'categoryAdd') {
        console.log(target)
        console.log(target.dataset['executor'])
        const categories = document.querySelector(target.dataset['executor'])

        if (categories.children.length >= 6) {
          modaler('错误', '最多添加5个类别。')
          return
        }

        const value = target.parentNode.previousElementSibling.value
        if (value === '0') {
          modaler('错误', '请选择一个类别。')
          return
        }
        const newCategory = document.createElement('span')
        newCategory.classList.add('category')
        newCategory.innerHTML =
          `<span class="category-name">${value}</span>
<a href="javascript:void(0)" class="category-del" data-executor="${target.dataset['executor']}" data-a-role="categoryDel">&otimes;</a>`
        categories.appendChild(newCategory)
        categories.children[0].textContent = `商品类别 (${categories.children.length - 1}/5): `
      }
    }
  })
}


function bindInputEvents(modaler, parent) {
  parent.querySelectorAll('[data-input-role="img"]').forEach(ele => {
    ele.addEventListener('change', (e) => {
      const target = e.target
      console.log(target.dataset["productId"])

      const file = target.files[0]
      const request = indexedDB.open('myProduct', 1)
      request.onupgradeneeded = (e) => {
        const db = e.target.result
        if (!db.objectStoreNames.contains('productImgFile')) {
          db.createObjectStore('productImgFile', { keyPath: 'id' })
        }
      }

      request.onsuccess = (e) => {
        const db = e.target.result
        const transaction = db.transaction('productImgFile', 'readwrite')
        const store = transaction.objectStore('productImgFile')
        console.log(target.dataset["productId"])
        console.log(file)
        const req = store.put({ id: target.dataset["productId"], data: file })

        req.onsuccess = () => {
          const reader = new FileReader()
          reader.onload = (e) => {
            target.previousElementSibling.src = e.target.result
            target.previousElementSibling.dataset["status"] = 'changed'
          }
          reader.readAsDataURL(file)
        }

        req.onerror = (e) => {
          modaler('错误', '上传图片失败, 请稍后再试。')
        }
      }
      request.onerror = (e) => {
        console.log(e)
      }
    })
  })

  parent.querySelectorAll('[data-input-role="name"]').forEach(ele => {
    ele.addEventListener('input', (e) => {
      const target = e.target
      if (target.value.length > 30) {
        target.value = target.value.slice(0, 30)
      }
      target.nextElementSibling.textContent = `(${target.value.length}/30)`
    })
  })

  parent.querySelectorAll('[data-input-role="shortname"]').forEach(ele => {
    ele.addEventListener('input', (e) => {
      const target = e.target
      if (target.value.length > 10) {
        target.value = target.value.slice(0, 10)
      }
      target.nextElementSibling.textContent = `(${target.value.length}/10)`
    })
  })

  parent.querySelectorAll('[data-input-role="desc"]').forEach(ele => {
    ele.addEventListener('input', (e) => {
      const target = e.target
      if (target.value.length > 200) {
        target.value = target.value.slice(0, 200)
      }
      target.nextElementSibling.textContent = `(${target.value.length}/200)`
    })
  })

  parent.querySelectorAll('[data-input-role="price"]').forEach(ele => {
    ele.addEventListener('blur', (e) => {
      const target = e.target
      if (target.value < 0) {
        target.value = 0
      }
    })
  })

  parent.querySelectorAll('[data-input-role="stock"]').forEach(ele => {
    ele.addEventListener('blur', (e) => {
      const target = e.target
      if (target.value < 0) {
        target.value = 0
        return
      }
      target.value = Math.round(target.value)
    })
  })
}

!function userProcess() {

  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
  axios.defaults.headers.post['Content-Type'] = 'application/json'
  axios.defaults.baseURL = router.DefaultURL
  const modaler = modalCtrl()

  panelSwitch()
  info(modaler)
  account(modaler)
  address(modaler)
  product(modaler)
  bindInputEvents(modaler, document.querySelector('#product-accordion'))
}()