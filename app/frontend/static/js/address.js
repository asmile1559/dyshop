import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL


function newAddressTableRow(address, is_default) {
  const tr = document.createElement('tr');
  tr.id = `address${address['address_id']}`;
  tr.dataset.addrId = address['address_id'];

  tr.innerHTML = `
    <td class="name-column">
      <span>${address["recipient"]}</span>
    </td>
    <td class="phone-column">
      <span>${address["phone"]}</span>
    </td>
    <td class="region-column">
      <span>
        <span class="province">${address["province"]}</span>
        <span class="city">${address["city"]}</span>
        <span class="district">${address["district"]}</span>
        <span class="street">${address["street"]}</span>
      </span>
    </td>
    <td class="detail-column">
      <span>${address["full_address"]}</span>
    </td>
    <td class="operation-column">
      <a class="implicit-link underline-link" href="javascript:void(0)">
        <span data-target="#address${address['address_id']}">修改</span>
      </a>
      <a class="implicit-link underline-link" href="javascript:void(0)">
        <span data-target="#address${address['address_id']}">删除</span>
      </a>
    </td>
    <td class="default-column">
      ${is_default ? `
        <a class="implicit-link default">
          <span data-target="#address${address['address_id']}">默认地址</span>
        </a>
      ` : `
        <a href="javascript:void(0)" class="implicit-link underline-link">
          <span data-target="#address${address['address_id']}">设为默认</span>
        </a>
      `}
    </td>
  `;
  return tr;
}


!function () {
  document.querySelector('#address-modal .btn-confirm').addEventListener('click', function (e) {
    const addrTbody = document.querySelector('#address-tbody')
    const addressModal = document.querySelector('#address-modal')
    const targetRow = document.querySelector(addressModal.getAttribute('data-target'))
    const addressName = document.querySelector('#address-name')
    const addressPhone = document.querySelector('#address-phone')
    const addressProvince = document.querySelector('#address-province')
    const addressCity = document.querySelector('#address-city')
    const addressDistrict = document.querySelector('#address-district')
    const addressStreet = document.querySelector('#address-street')
    const addressDetail = document.querySelector('#address-detail')

    const addressId = targetRow.dataset['addrId']
    if (addressName.value === ''
      || addressPhone.value === ''
      || addressProvince.value === ''
      || addressCity.value === ''
      || addressDistrict.value === ''
      || addressStreet.value === ''
      || addressDetail.value === '') {
      addressModal.style.display = 'none'
      common.alertByModal('错误', '请填写完整信息。')
      return
    }
    axios({
      'method': 'post',
      'url': router.OperationRouters['updateAddress'],
      'data': {
        'address_id': addressId,
        'recipient': addressName.value,
        'phone': addressPhone.value,
        'province': addressProvince.value,
        'city': addressCity.value,
        'district': addressDistrict.value,
        'street': addressStreet.value,
        'full_address': addressDetail.value
      }
    }).then((res) => {
      const resp = res.data.resp
      console.log(resp)
      if (addressId === '-1') {
        // add new address
        if (addrTbody.querySelector('.no-address')) {
          addrTbody.removeChild(myAddresses.querySelector('.no-address'))
        }
        const lastrow = addrTbody.querySelector('#last-row')
        lastrow.parentNode.insertBefore(newAddressTableRow(resp['address'], resp['is_default']), lastrow)
      } else {
        targetRow.querySelector('.name-column span').textContent = addressName.value.trim()
        targetRow.querySelector('.phone-column span').textContent = addressPhone.value.trim()
        targetRow.querySelector('.region-column span.province').textContent = addressProvince.value.trim()
        targetRow.querySelector('.region-column span.city').textContent = addressCity.value.trim()
        targetRow.querySelector('.region-column span.district').textContent = addressDistrict.value.trim()
        targetRow.querySelector('.region-column span.street').textContent = addressStreet.value.trim()
        targetRow.querySelector('.detail-column span').textContent = addressDetail.value.trim()
      }
      addressModal.style.display = 'none'
    }).catch((err) => {
      addressModal.style.display = 'none'
      common.alertByModal('错误', '一个错误发生了，请稍后再试。')
      console.log(err)
    })
  })
}()

!function () {
  document.querySelector('#address-tbody').addEventListener('click', function (e) {
    const target = e.target;
    console.log(target)
    if (target.tagName === 'SPAN') {
      const operation = target.textContent;
      const tr = document.querySelector(target.dataset['target'])
      if (operation === '修改') {
        const addressModal = document.querySelector('#address-modal');
        console.log(addressModal)
        addressModal.setAttribute('data-target', target.dataset['target'])
        addressModal.querySelector('.modal-title').textContent = '修改地址'
        addressModal.querySelector('#address-name').value = tr.querySelector(`.name-column span`).textContent.trim()
        addressModal.querySelector('#address-phone').value = tr.querySelector(`.phone-column span`).textContent.trim()
        addressModal.querySelector('#address-province').value = tr.querySelector(`.province`).textContent.trim()
        addressModal.querySelector('#address-city').value = tr.querySelector(`.city`).textContent.trim()
        addressModal.querySelector('#address-district').value = tr.querySelector(`.district`).textContent.trim()
        addressModal.querySelector('#address-street').value = tr.querySelector(`.street`).textContent.trim()
        addressModal.querySelector('#address-detail').value = tr.querySelector(`.detail-column span`).textContent.trim()
        addressModal.style.display = 'block'
      } else if (operation === '添加新地址') {
        const addressModal = document.querySelector('#address-modal');
        addressModal.setAttribute('data-target', target.dataset['target'])
        addressModal.querySelector('.modal-title').textContent = '添加新地址'
        addressModal.querySelector('#address-name').value = ''
        addressModal.querySelector('#address-phone').value = ''
        addressModal.querySelector('#address-province').value = ''
        addressModal.querySelector('#address-city').value = ''
        addressModal.querySelector('#address-district').value = ''
        addressModal.querySelector('#address-street').value = ''
        addressModal.querySelector('#address-detail').value = ''
        addressModal.style.display = 'block'
      } else if (operation === '删除') {
        if (tr.dataset['addrId'] === this.dataset['defaultId']) {
          common.alertByModal('错误', '不能删除默认地址。')
          return
        }
        axios({
          'method': 'post',
          'url': router.OperationRouters['deleteAddress'],
          'data': {
            'address_id': tr.dataset['addrId'],
          }
        }).then((res) => {
          const resp = res.data.resp
          console.log(resp)
          this.removeChild(tr)
        }).catch((err) => {
          common.alertByModal('错误', '一个错误发生了，请稍后再试。')
          console.log(err)
        })
      } else if (operation === '设为默认') {
        axios({
          'method': 'post',
          'url': router.OperationRouters['setDefAddress'],
          'data': {
            address_id: tr.dataset['addrId'],
          }
        }).then((res) => {
          const resp = res.data.resp
          console.log(resp)
          console.log(this)
          const oldTr = this.querySelector(`#address${this.dataset['defaultId']}`)
          const oldDefault = this.querySelector('.default-column a.default')
          console.log(oldDefault.children)
          oldDefault.children[0].textContent = '设为默认'
          oldDefault.classList.remove('default')
          oldDefault.classList.add('underline-link')
          oldDefault.href = 'javascript:void(0)'

          this.dataset['defaultId'] = tr.dataset['addrId']
          tr.querySelector('.default-column').innerHTML = `
            <a class="implicit-link default">
              <span data-target="#${tr.id}">默认地址</span>
            </a>
          `
          const newTr = tr.cloneNode(true)
          this.insertBefore(newTr, oldTr)
          this.removeChild(tr)
        }).catch((err) => {
          common.alertByModal('错误', '一个错误发生了，请稍后再试。')
          console.log(err)
        })
      }
    }
  })

}()