import * as router from './router.js';
import * as common from './common.js';

axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.baseURL = router.DefaultURL

!function () {
  takeSnapshot("#public-product")
  document.querySelectorAll('.accordion-item').forEach(ele => {
    accordionItemInputBind(`#${ele.id}`)
  })
}()


// function takeSnapshot(accordionId, save = true) {
//   const accordionItem = document.querySelector(accordionId)
//   const id = accordionItem.dataset['productId']
//   const snapshot = {
//     id: id,
//     name: accordionItem.querySelector('#name' + id).value,
//     desc: accordionItem.querySelector('#desc' + id).value,
//     categories: function () {
//       const res = []
//       const categories = accordionItem.querySelector('.categories')
//       for (let i = 1; i < categories.children.length; i++) {
//         // console.log(categories.children[i])
//         res.push(categories.children[i].querySelector('.category-name').textContent)
//       }
//       return res
//     }(),
//     insure: accordionItem.querySelector('#insurance-select' + id).selectedIndex,
//     express: accordionItem.querySelector('#express-select' + id).selectedIndex,
//     specs: function () {
//       const specs = []
//       const tbody = accordionItem.querySelector('.spec-table tbody')
//       const rnxidx = tbody.dataset['rowNxtIdx']
//       const cnxidx = tbody.dataset['colNxtIdx']
//       for (let i = 0; i < rnxidx - 1; i++) {
//         const cells = tbody.children[i].children
//         console.log(cells)
//         specs.push({
//           name: cells[0].textContent,
//           price: cells[1].textContent,
//           stock: cells[2].textContent,
//         })
//         specs.push({
//           name: cells[4].textContent,
//           price: cells[5].textContent,
//           stock: cells[6].textContent,
//         })
//       }
//       if (cnxidx === "2") {
//         const cells = tbody.children[rnxidx - 1].children
//         specs.push({
//           name: cells[0].textContent,
//           price: cells[1].textContent,
//           stock: cells[2].textContent,
//         })
//       }
//       return specs
//     }(),
//     params: function () {
//       const params = []
//       const tbody = accordionItem.querySelector('.params-table tbody')
//       const rnxidx = tbody.dataset['rowNxtIdx']
//       const cnxidx = tbody.dataset['colNxtIdx']
//       for (let i = 0; i < rnxidx - 1; i++) {
//         const cells = tbody.children[i].children
//         params.push({
//           name: cells[0].textContent,
//           value: cells[1].textContent,
//         })
//         params.push({
//           name: cells[3].textContent,
//           value: cells[4].textContent,
//         })
//       }
//       if (cnxidx === "2") {
//         const cells = tbody.children[rnxidx - 1].children
//         params.push({
//           name: cells[0].textContent,
//           value: cells[1].textContent,
//         })
//       }
//       return params
//     }(),
//     image: function () {
//       const image = accordionItem.querySelector('img')
//       return image.src
//     }(),
//   }
//   if (save) {
//     common.writeDB("productmana", "published", id, snapshot).then((res) => {
//       console.log(res)
//       console.log("Snapshot saved.")
//     }).catch((err) => {
//       console.log(err)
//     })
//   }
//   return snapshot
// }

// async function recoverSnapshot(accordionId) {
//   const accordionItem = document.querySelector(accordionId)
//   const id = accordionItem.dataset['productId']

//   try {
//     const res = await common.readDB("productmana", "published", id)
//     const snapshot = res.data
//     console.log(snapshot)
//     accordionItem.querySelector('#name' + id).value = snapshot.name
//     accordionItem.querySelector('#desc' + id).value = snapshot.desc
//     accordionItem.querySelector('#select' + id).selectedIndex = 0
//     const categories = accordionItem.querySelector('.categories')
//     categories.innerHTML = `<span>商品类别 (${snapshot.categories.length}/5): </span>`
//     for (let i = 0; i < snapshot.categories.length; i++) {
//       const category = document.createElement('span')
//       category.className = 'category'
//       category.innerHTML = `<span class="category-name">${snapshot.categories[i]}</span><a href="javascript:void(0)" class="category-del d-none"
//                             data-executor="#categories${id}" data-a-role="categoryDel">&otimes;</a>`
//       categories.appendChild(category)
//     }
//     accordionItem.querySelector('#insurance-select' + id).selectedIndex = snapshot.insure
//     accordionItem.querySelector('#express-select' + id).selectedIndex = snapshot.express
//     const spec_tbody = accordionItem.querySelector('.spec-table tbody')
//     spec_tbody.innerHTML = ''
//     for (let i = 0, c = 0; i < parseInt(snapshot.specs.length / 2) + 1 && c < snapshot.specs.length; i++) {
//       const row = document.createElement('tr')
//       spec_tbody.appendChild(row)
//       for (let j = 0; j < 8; j++) {
//         const cell = document.createElement('td')
//         row.appendChild(cell)
//         if (j % 4 === 3) {
//           if (c !== snapshot.specs.length) {
//             cell.innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
//                               <span data-role="modifyRow" data-domain="spec" data-target="#${spec_tbody.id}">修改</span>
//                             </a>
//                             <a href="javascript:void(0)" class="implicit-link change-danger disabled-white" >
//                               <span data-role="deleteRow">删除</span>
//                             </a>`
//             cell.dataset['rowIdx'] = i + 1
//             cell.dataset['colIdx'] = parseInt(j / 4) + 1
//           }
//           else {
//             cell.innerHTML = ``
//             cell.dataset['rowIdx'] = i + 1
//             cell.dataset['colIdx'] = parseInt(j / 4) + 1
//           }
//           c++
//         } else {
//           if (c === snapshot.specs.length) {
//             continue
//           }
//           if (j % 4 === 0) {
//             cell.innerHTML = snapshot.specs[c].name
//           } else if (j % 4 === 1) {
//             cell.innerHTML = snapshot.specs[c].price
//           } else if (j % 4 === 2) {
//             cell.innerHTML = snapshot.specs[c].stock
//           }
//         }
//       }
//     }
//     spec_tbody.dataset['rowNxtIdx'] = parseInt(snapshot.specs.length / 2) + 1
//     spec_tbody.dataset['colNxtIdx'] = parseInt(snapshot.specs.length % 2) + 1

//     const params_tbody = accordionItem.querySelector('.params-table tbody')
//     params_tbody.innerHTML = ''
//     for (let i = 0, c = 0; i < parseInt(snapshot.params.length / 2) + 1 && c < snapshot.params.length; i++) {
//       const row = document.createElement('tr')
//       params_tbody.appendChild(row)
//       for (let j = 0; j < 6; j++) {
//         const cell = document.createElement('td')
//         row.appendChild(cell)
//         if (j % 3 === 2) {
//           if (c !== snapshot.params.length) {
//             cell.innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
//                               <span data-role="modifyRow" data-domain="param" data-target="#${params_tbody.id}">修改</span>
//                             </a>
//                             <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
//                               <span data-role="deleteRow">删除</span>
//                             </a>`
//             cell.dataset['rowIdx'] = i + 1
//             cell.dataset['colIdx'] = parseInt(j / 3) + 1
//           } else {
//             cell.innerHTML = ``
//             cell.dataset['rowIdx'] = i + 1
//             cell.dataset['colIdx'] = parseInt(j / 3) + 1
//           }
//           c++
//         } else {
//           if (c === snapshot.params.length) {
//             continue
//           }
//           if (j % 3 === 0) {
//             cell.innerHTML = snapshot.params[i].name
//           } else if (j % 3 === 1) {
//             cell.innerHTML = snapshot.params[i].value
//           }
//         }
//       }
//     }
//     params_tbody.dataset['rowNxtIdx'] = parseInt(snapshot.params.length / 2) + 1
//     params_tbody.dataset['colNxtIdx'] = parseInt(snapshot.params.length % 2) + 1

//     accordionItem.querySelector('img').src = snapshot.image
//     updateIndicators(accordionId)
//   } catch (err) {
//     console.log(err)
//   }
// }

function serializeAccordionItem(accordionId, save = true) {
  const accordionItem = document.querySelector(accordionId)
  const id = accordionItem.dataset['productId']
  const snapshot = {
    product_id: id,
    product_img: accordionItem.querySelector('img').src,
    product_name: accordionItem.querySelector('#name' + id).value,
    product_desc: accordionItem.querySelector('#desc' + id).value,
    product_categories: function () {
      const res = []
      const categories = accordionItem.querySelector('.categories')
      for (let i = 1; i < categories.children.length; i++) {
        res.push(categories.children[i].querySelector('.category-name').textContent)
      }
      return res
    }(),
    product_insurance: accordionItem.querySelector('#insurance-select' + id).value,
    product_express: accordionItem.querySelector('#express-select' + id).value,
    product_specs: function () {
      const specs = []
      const tbody = accordionItem.querySelector('.spec-table tbody')
      const rnxidx = tbody.dataset['rowNxtIdx']
      const cnxidx = tbody.dataset['colNxtIdx']
      for (let i = 0; i < rnxidx - 1; i++) {
        const cells = tbody.children[i].children
        console.log(cells)
        specs.push({
          spec_name: cells[0].textContent,
          spec_price: cells[1].textContent,
          spec_stock: cells[2].textContent,
        })
        specs.push({
          spec_name: cells[4].textContent,
          spec_price: cells[5].textContent,
          spec_stock: cells[6].textContent,
        })
      }
      if (cnxidx === "2") {
        const cells = tbody.children[rnxidx - 1].children
        specs.push({
          spec_name: cells[0].textContent,
          spec_price: cells[1].textContent,
          spec_stock: cells[2].textContent,
        })
      }
      return specs
    }(),
    product_params: function () {
      const params = []
      const tbody = accordionItem.querySelector('.params-table tbody')
      const rnxidx = tbody.dataset['rowNxtIdx']
      const cnxidx = tbody.dataset['colNxtIdx']
      for (let i = 0; i < rnxidx - 1; i++) {
        const cells = tbody.children[i].children
        params.push({
          param_name: cells[0].textContent,
          param_value: cells[1].textContent,
        })
        params.push({
          param_name: cells[3].textContent,
          param_value: cells[4].textContent,
        })
      }
      if (cnxidx === "2") {
        const cells = tbody.children[rnxidx - 1].children
        params.push({
          param_name: cells[0].textContent,
          param_value: cells[1].textContent,
        })
      }
      return params
    }()
  }
  return snapshot
}

function deserializeAccordionItem(accordionId, snapshot) {
  const accordionItem = document.querySelector(accordionId)
  const id = accordionItem.dataset['productId']
  accordionItem.querySelector('.accordion-header button').textContent = snapshot.product_name
  accordionItem.querySelector('#name' + id).value = snapshot.product_name
  accordionItem.querySelector('#desc' + id).value = snapshot.product_desc
  accordionItem.querySelector('#select' + id).selectedIndex = 0
  const categories = accordionItem.querySelector('.categories')
  categories.innerHTML = `<span>商品类别 (${snapshot.product_categories.length}/5): </span>`
  for (let i = 0; i < snapshot.product_categories.length; i++) {
    const category = document.createElement('span')
    category.className = 'category'
    category.innerHTML = `<span class="category-name">${snapshot.product_categories[i]}</span><a href="javascript:void(0)" class="category-del d-none"
                            data-executor="#categories${id}" data-a-role="categoryDel">&otimes;</a>`
    categories.appendChild(category)
  }
  accordionItem.querySelector('#insurance-select' + id).selectedIndex = function (insurance) {
    switch (insurance) {
      case "无保险":
        return 0
      case "退货险":
        return 1
      case "运费险":
        return 2
    }
  }(snapshot.product_insurance)
  accordionItem.querySelector('#express-select' + id).selectedIndex = function (express) {
    switch (express) {
      case "包邮":
        return 1
      case "到付":
        return 0
    }
  }(snapshot.product_express)
  const spec_tbody = accordionItem.querySelector('.spec-table tbody')
  spec_tbody.innerHTML = ''
  for (let i = 0, c = 0; i < parseInt(snapshot.product_specs.length / 2) + 1 && c < snapshot.product_specs.length; i++) {
    const row = document.createElement('tr')
    spec_tbody.appendChild(row)
    for (let j = 0; j < 8; j++) {
      const cell = document.createElement('td')
      row.appendChild(cell)
      if (j % 4 === 3) {
        if (c !== snapshot.product_specs.length) {
          cell.innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
                              <span data-role="modifyRow" data-domain="spec" data-target="#${spec_tbody.id}">修改</span>
                            </a>
                            <a href="javascript:void(0)" class="implicit-link change-danger disabled-white" >
                              <span data-role="deleteRow">删除</span>
                            </a>`
          cell.dataset['rowIdx'] = i + 1
          cell.dataset['colIdx'] = parseInt(j / 4) + 1
        }
        else {
          cell.innerHTML = ``
          cell.dataset['rowIdx'] = i + 1
          cell.dataset['colIdx'] = parseInt(j / 4) + 1
        }
        c++
      } else {
        if (c === snapshot.product_specs.length) {
          continue
        }
        if (j % 4 === 0) {
          cell.innerHTML = snapshot.product_specs[c].spec_name
        } else if (j % 4 === 1) {
          cell.innerHTML = snapshot.product_specs[c].spec_price
        } else if (j % 4 === 2) {
          cell.innerHTML = snapshot.product_specs[c].spec_stock
        }
      }
    }
  }
  spec_tbody.dataset['rowNxtIdx'] = parseInt(snapshot.product_specs.length / 2) + 1
  spec_tbody.dataset['colNxtIdx'] = parseInt(snapshot.product_specs.length % 2) + 1

  const params_tbody = accordionItem.querySelector('.params-table tbody')
  params_tbody.innerHTML = ''
  for (let i = 0, c = 0; i < parseInt(snapshot.product_params.length / 2) + 1 && c < snapshot.product_params.length; i++) {
    const row = document.createElement('tr')
    params_tbody.appendChild(row)
    for (let j = 0; j < 6; j++) {
      const cell = document.createElement('td')
      row.appendChild(cell)
      if (j % 3 === 2) {
        if (c !== snapshot.product_params.length) {
          cell.innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
                              <span data-role="modifyRow" data-domain="param" data-target="#${params_tbody.id}">修改</span>
                            </a>
                            <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
                              <span data-role="deleteRow">删除</span>
                            </a>`
          cell.dataset['rowIdx'] = i + 1
          cell.dataset['colIdx'] = parseInt(j / 3) + 1
        } else {
          cell.innerHTML = ``
          cell.dataset['rowIdx'] = i + 1
          cell.dataset['colIdx'] = parseInt(j / 3) + 1
        }
        c++
      } else {
        if (c === snapshot.product_params.length) {
          continue
        }
        if (j % 3 === 0) {
          cell.innerHTML = snapshot.product_params[i].param_name
        } else if (j % 3 === 1) {
          cell.innerHTML = snapshot.product_params[i].param_value
        }
      }
    }
  }
  params_tbody.dataset['rowNxtIdx'] = parseInt(snapshot.product_params.length / 2) + 1
  params_tbody.dataset['colNxtIdx'] = parseInt(snapshot.product_params.length % 2) + 1

  accordionItem.querySelector('img').src = snapshot.product_img
  accordionItem.querySelector('img').dataset["status"] = "no-change"
  updateIndicators(accordionId)
}

async function takeSnapshot(accordionId) {
  try {
    const snapshot = serializeAccordionItem(accordionId)
    const res = common.writeDB("productmana", "published", snapshot.product_id, snapshot)
    console.log('take snapshot success')
    console.log(res)
  } catch (err) {
    console.log(err)
    common.alertByModal('错误', '什么地方出错了')
  }
}

async function recoverSnapshot(accordionId) {
  try {
    const id = document.querySelector(accordionId).dataset['productId']
    const res = await common.readDB("productmana", "published", id)
    const snapshot = res.data
    deserializeAccordionItem(accordionId, snapshot)
  } catch (err) {
    console.log(err)
    common.alertByModal('错误', '什么地方出错了')
  }
}

!function () {
  document.querySelector('#product-accordion').addEventListener('click', (e) => {
    const target = e.target
    const role = target.dataset['role']
    if (role === undefined) {
      return
    }
    if (role === "productEdit") {
      takeSnapshot(target.dataset['accordionId'])
      document.querySelector(target.dataset['accordionId']).classList.add('edit')
    } else if (role === "productDiscard") {
      document.querySelector(target.dataset['accordionId']).classList.remove('edit')
      recoverSnapshot(target.dataset['accordionId'])
    } else if (role === "categoryAdd") {
      console.log(target)
      const executor = document.querySelector(target.dataset['executor'])
      const selector = document.querySelector(target.dataset['selector'])
      if (executor.children.length >= 6) {
        common.alertByModal('错误', '最多添加5个类别。')
        return
      }
      if (selector.value === '') {
        common.alertByModal('错误', '请选择一个类别。')
        return
      }
      const category = document.createElement('span')
      category.classList.add('category')
      category.innerHTML =
        `<span class="category-name">${selector.value}</span>
        <a href="javascript:void(0)" class="category-del d-none"
        data-executor="${target.dataset['executor']}" data-role="categoryDel">&otimes;</a>`
      executor.appendChild(category)
      executor.children[0].textContent = `商品类别 (${executor.children.length - 1}/5): `
    } else if (role === "modifyRow") {
      const domain = target.dataset['domain']
      if (domain === "spec") {
        const specModal = document.querySelector('#spec-modal')
        const td = target.closest('td')
        const tr = target.closest('tr')
        specModal.setAttribute('data-target-row', td.dataset['rowIdx'])
        specModal.setAttribute('data-target-col', td.dataset['colIdx'])
        specModal.setAttribute('data-action', 'modify')
        specModal.setAttribute('data-target', target.dataset['target'])
        specModal.querySelector('.modal-title').textContent = '修改规格'
        specModal.querySelector('#spec-name').value = tr.children[4 * td.dataset['colIdx'] - 4].textContent
        specModal.querySelector('#spec-price').value = tr.children[4 * td.dataset['colIdx'] - 3].textContent
        specModal.querySelector('#spec-stock').value = tr.children[4 * td.dataset['colIdx'] - 2].textContent
        specModal.style.display = 'block'
      } else if (domain === "param") {
        const tr = target.closest('tr')
        const td = target.closest('td')
        const param_name = tr.children[3 * td.dataset['colIdx'] - 3].textContent
        const value = prompt(`修改参数: ${param_name}`, tr.children[3 * td.dataset['colIdx'] - 2].textContent)
        if (!value) {
          return
        }
        tr.children[3 * td.dataset['colIdx'] - 2].textContent = value
      }
    } else if (role === "categoryDel") {
      const target = e.target
      console.log(target)
      console.log(target.parentNode)
      document.querySelector(target.dataset['executor']).removeChild(target.parentNode)
      document.querySelector(target.dataset['executor']).children[0].textContent = `商品类别 (${document.querySelector(target.dataset['executor']).children.length - 1}/5): `
    } else if (role === "deleteRow") {
      const td = target.closest('td')

      const tr = target.closest('tr')
      const rowIdx = parseInt(td.dataset['rowIdx'])
      const colIdx = parseInt(td.dataset['colIdx'])
      const no = rowIdx * 2 + colIdx - 2

      const radius = parseInt(tr.children.length / 2) - 1

      const tbody = target.closest('tbody')
      const rnxidx = tbody.dataset['rowNxtIdx']
      const cnxidx = tbody.dataset['colNxtIdx']

      const data = []
      let count = 1
      for (let i = 0; i < rnxidx - 1; i++) {
        const cells = tbody.children[i].children
        if (count !== no) {
          const cell = []
          for (let j = 0; j < radius; j++) {
            cell.push(cells[j].textContent)
          }
          data.push(cell)
        }
        count++
        if (count !== no) {
          const cell = []
          for (let j = 0; j < radius; j++) {
            cell.push(cells[radius + 1 + j].textContent)
          }
          data.push(cell)
        }
        count++
      }
      if (cnxidx === "2") {
        const cells = tbody.children[rnxidx - 1].children
        if (count !== no) {
          const cell = []
          for (let j = 0; j < radius; j++) {
            cell.push(cells[j].textContent)
          }
          data.push(cell)
        }
      }
      console.log(data)

      tbody.innerHTML = ''
      for (let i = 0, c = 0; i < parseInt(data.length / 2) + 1 && c < data.length; i++) {
        const row = document.createElement('tr')
        tbody.appendChild(row)

        for (let k = 0; k < 2; k++) {
          if (c !== data.length) {
            for (let j = 0; j < radius; j++) {
              const cell = document.createElement('td')
              row.appendChild(cell)
              cell.textContent = data[c][j]
            }
            const cell = document.createElement('td')
            row.appendChild(cell)
            cell.innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
                            <span data-role="modifyRow" data-domain="${radius === 3 ? "spec" : "param"}" data-target="#${tbody.id}">修改</span>
                          </a>
                          <a href="javascript:void(0)" class="implicit-link change-danger disabled-white" >
                            <span data-role="deleteRow">删除</span>
                          </a>`
            cell.dataset['rowIdx'] = i + 1
            cell.dataset['colIdx'] = parseInt(c % 2) + 1
            c++
          } else {
            for (let j = 0; j < radius; j++) {
              const cell = document.createElement('td')
              row.appendChild(cell)
              cell.innerHTML = ``
            }
            const cell = document.createElement('td')
            row.appendChild(cell)
            cell.innerHTML = ``
            cell.dataset['rowIdx'] = i + 1
            cell.dataset['colIdx'] = parseInt(c % 2) + 1
          }
        }
      }
      tbody.dataset['rowNxtIdx'] = parseInt(data.length / 2) + 1
      tbody.dataset['colNxtIdx'] = parseInt(data.length % 2) + 1
    } else if (role === "addRow") {
      const tbody = target.closest('tfoot').previousElementSibling
      const rnxidx = tbody.dataset['rowNxtIdx']
      const cnxidx = tbody.dataset['colNxtIdx']

      const domain = target.dataset['domain']
      let modal = null
      if (domain === "spec") {
        modal = document.querySelector('#spec-modal')
        modal.querySelector('.modal-title').textContent = '添加规格'
        modal.querySelector('#spec-name').value = ''
        modal.querySelector('#spec-price').value = ''
        modal.querySelector('#spec-stock').value = ''
      } else if (domain === "param") {
        modal = document.querySelector('#param-modal')
        modal.querySelector('.modal-title').textContent = '添加参数'
        modal.querySelector('#param-name').value = ''
        modal.querySelector('#param-value').value = ''
      }
      modal.dataset['targetRow'] = rnxidx
      modal.dataset['targetCol'] = cnxidx
      modal.dataset['action'] = 'add'
      modal.dataset['target'] = target.dataset['target']
      modal.style.display = 'block'
    } else if (role === "productSave") {
      const product = serializeAccordionItem(target.dataset['accordionId'])
      if (product.product_name === "" || product.product_desc === "" ||
        product.product_express === "" || product.product_insurance === "" ||
        product.product_categories.length === 0 || product.product_specs.length === 0 || product.product_params.length === 0) {
        common.alertByModal('错误', '请填写完整信息。')
        return
      }
      common.readDB("imagebuf", "image", product.product_id).then(res => {
        const formData = new FormData()
        if (document.querySelector(target.dataset['accordionId']).querySelector('img').dataset['status'] === 'changed') {
          product.product_img = ''
          formData.append("image", res.data)
        }
        formData.append("product", JSON.stringify(product))
        return axios({
          "method": "post",
          "url": router.OperationRouters["updateProduct"],
          "data": formData,
          'headers': {
            'Content-Type': 'multipart/form-data'
          }
        })
      }).then(res => {
        const resp = res.data.resp
        const snapshot = resp['product']
        deserializeAccordionItem(target.dataset['accordionId'], snapshot)
        document.querySelector(target.dataset['accordionId']).classList.remove('edit')
      }).catch(err => {
        recoverSnapshot(target.dataset['accordionId'])
        console.log(err)
        common.alertByModal("错误", "一个错误发生了")
      })
    } else if (role === 'public-confirm') {
      const product = serializeAccordionItem(target.dataset['accordionId'])
      if (product.product_name === "" || product.product_desc === "" ||
        product.product_express === "" || product.product_insurance === "" ||
        product.product_categories.length === 0 || product.product_specs.length === 0 || product.product_params.length === 0) {
        common.alertByModal('错误', '请填写完整信息。')
        return
      }
      product.product_img = ''
      if (document.querySelector('#public-product img').dataset['status'] === 'no-change') {
        common.alertByModal('错误', '请上传图片。')
        return
      }
      common.readDB("imagebuf", "image", product.product_id).then(res => {
        const formData = new FormData()
        formData.append("image", res.data)
        formData.append("product", JSON.stringify(product))
        return axios({
          "method": "post",
          "url": router.OperationRouters["updateProduct"],
          "data": formData,
          'headers': {
            'Content-Type': 'multipart/form-data'
          }
        })
      }).then(res => {
        const resp = res.data.resp
        const snapshot = resp['product']
        const newAccordionItem = NewAccordionItem(snapshot)
        document.querySelector('#product-accordion').insertBefore(newAccordionItem, document.querySelector('#public-product'))
        accordionItemInputBind("#" + newAccordionItem.id)
        recoverSnapshot(target.dataset['accordionId'])
        document.querySelector('.detail-product .count span:nth-child(2)').textContent = parseInt(document.querySelector('#product-accordion').children.length) - 1
      }).catch(err => {
        recoverSnapshot(target.dataset['accordionId'])
        console.log(err)
        common.alertByModal("错误", "一个错误发生了")
      })
    } else if (role === 'public-discard') {
      recoverSnapshot(target.dataset['accordionId'])
      updateIndicators(target.dataset['accordionId'])
    } else if (role === 'productDel') {
      if (!confirm('确认删除？')) {
        return
      }
      const id = target.dataset['targetId']
      console.log(id)
      axios({
        "method": "post",
        "url": router.OperationRouters["deleteProduct"],
        "data": {
          product_id: id
        }
      }).then(res => {
        console.log(res)
        document.querySelector('#product-accordion').removeChild(target.closest('.accordion-item'))
        document.querySelector('.detail-product .count span:nth-child(2)').textContent = parseInt(document.querySelector('#product-accordion').children.length) - 1
      }).catch(err => {
        console.log(err)
      })

    } else if (role === 'updateImg') {
      document.querySelector(target.dataset['target']).click()
    }
  })

}()

!function () {
  const specModal = document.querySelector('#spec-modal')

  specModal.querySelector('#spec-price').addEventListener('blur', (e) => {
    const target = e.target
    if (target.value < 0) {
      target.value = 0
    }
    target.value = parseFloat(target.value).toFixed(2)
    if (isNaN(target.value)) {
      target.value = 0
      target.value = parseFloat(target.value).toFixed(2)
    }
  })

  specModal.querySelector('#spec-stock').addEventListener('blur', (e) => {
    const target = e.target
    if (target.value < 0) {
      target.value = 0
    }
    target.value = parseInt(target.value)
    if (isNaN(target.value)) {
      target.value = 0
      target.value = parseInt(target.value)
    }

  })

  specModal.querySelector('.btn-confirm').addEventListener('click', () => {
    const rowIdx = parseInt(specModal.dataset['targetRow'])
    const colIdx = parseInt(specModal.dataset['targetCol'])
    const action = specModal.dataset['action']
    const targetTbodyId = specModal.dataset['target']
    const name = specModal.querySelector('#spec-name').value.trim()
    const price = specModal.querySelector('#spec-price').value.trim()
    const stock = specModal.querySelector('#spec-stock').value.trim()

    if (name === '' || price === '' || stock === '') {
      specModal.style.display = "none"
      common.alertByModal('错误', '请填写完整信息。')
      return
    }

    const tbody = document.querySelector(targetTbodyId)

    if (rowIdx <= tbody.children.length) {
      // 修改
      const tr = tbody.children[rowIdx - 1]
      tr.children[4 * colIdx - 4].textContent = name
      tr.children[4 * colIdx - 3].textContent = price
      tr.children[4 * colIdx - 2].textContent = stock
      if (action === "add") {
        tr.children[4 * colIdx - 1].innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
                                                <span data-role="modifyRow" data-domain="spec" data-target="#${tbody.id}">修改</span>
                                              </a>
                                              <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
                                                <span data-role="deleteRow">删除</span>
                                              </a>`
        tr.children[4 * colIdx - 1].dataset['rowIdx'] = rowIdx
        tr.children[4 * colIdx - 1].dataset['colIdx'] = colIdx
      }
    } else {
      // 添加
      const tr = document.createElement('tr')
      tbody.appendChild(tr)
      for (let i = 0; i < 8; i++) {
        const td = document.createElement('td')
        tr.appendChild(td)
      }
      tr.children[0].textContent = name
      tr.children[1].textContent = price
      tr.children[2].textContent = stock
      tr.children[3].innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
                                                <span data-role="modifyRow" data-domain="spec" data-target="#${tbody.id}">修改</span>
                                              </a>
                                              <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
                                                <span data-role="deleteRow">删除</span>
                                              </a>`
      tr.children[3].dataset['rowIdx'] = rowIdx
      tr.children[3].dataset['colIdx'] = colIdx
    }

    if (action === "add") {
      tbody.dataset['rowNxtIdx'] = rowIdx + colIdx - 1
      tbody.dataset['colNxtIdx'] = colIdx % 2 + 1
    }
    specModal.style.display = "none"
  })
}()

!function () {
  const param_modal = document.querySelector('#param-modal')
  param_modal.querySelector('.btn-confirm').addEventListener('click', () => {
    const rowIdx = parseInt(param_modal.dataset['targetRow'])
    const colIdx = parseInt(param_modal.dataset['targetCol'])
    const action = param_modal.dataset['action']
    const targetTbodyId = param_modal.dataset['target']
    const name = param_modal.querySelector('#param-name').value.trim()
    const value = param_modal.querySelector('#param-value').value.trim()

    if (name === '' || value === '') {
      param_modal.style.display = "none"
      common.alertByModal('错误', '请填写完整信息。')
      return
    }

    const tbody = document.querySelector(targetTbodyId)
    if (rowIdx <= tbody.children.length) {
      const tr = tbody.children[rowIdx - 1]
      tr.children[3 * colIdx - 3].textContent = name
      tr.children[3 * colIdx - 2].textContent = value
      if (action === "add") {
        tr.children[3 * colIdx - 1].innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
                                                <span data-role="modifyRow" data-domain="param" data-target="#${tbody.id}">修改</span>
                                              </a>
                                              <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
                                                <span data-role="deleteRow">删除</span>
                                              </a>`
        tr.children[3 * colIdx - 1].dataset['rowIdx'] = rowIdx
        tr.children[3 * colIdx - 1].dataset['colIdx'] = colIdx
      }

    } else {
      const tr = document.createElement('tr')
      tbody.appendChild(tr)
      for (let i = 0; i < 6; i++) {
        const td = document.createElement('td')
        tr.appendChild(td)
      }
      tr.children[0].textContent = name
      tr.children[1].textContent = value
      tr.children[2].innerHTML = `<a href="javascript:void(0)" class="implicit-link change-success disabled-white">
                                                <span data-role="modifyRow" data-domain="param" data-target="#${tbody.id}">修改</span>
                                              </a>
                                              <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
                                                <span data-role="deleteRow">删除</span>
                                              </a>`
      tr.children[2].dataset['rowIdx'] = rowIdx
      tr.children[2].dataset['colIdx'] = colIdx
    }

    if (action === "add") {
      tbody.dataset['rowNxtIdx'] = rowIdx + colIdx - 1
      tbody.dataset['colNxtIdx'] = colIdx % 2 + 1
    }
    param_modal.style.display = "none"
  })
}()

async function ReadImg(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = (e) => {
      target.previousElementSibling.src = e.target.result
      resolve(e.target.result)
    }
    reader.readAsDataURL(file)
  })
}

function accordionItemInputBind(accordionId) {
  const element = document.querySelector(accordionId)
  element.querySelectorAll('input[data-role="img"]').forEach(ele => {
    ele.addEventListener('change', (e) => {
      const target = e.target
      console.log(target.dataset["productId"])

      const file = target.files[0]
      if (!file) {
        common.alertByModal('错误', '请选择一个图片文件。')
        return
      }
      if (file.size > 1024 * 1024) {
        common.alertByModal('错误', '图片文件大小不能超过1MB。')
        return
      }

      common.writeDB("imagebuf", "image", target.dataset["productId"], file, 1).then(res => {
        console.log(res)
        document.querySelector(ele.dataset["target"]).dataset["status"] = "changed"
      }).catch(err => {
        console.log(err)
      })
      const reader = new FileReader()
      reader.onload = (e) => {
        target.previousElementSibling.src = e.target.result
      }
      reader.readAsDataURL(file)

      // console.log(ele.dataset["productId"])
    })
  })

  element.querySelectorAll('input[data-role="name"]').forEach(ele => {
    ele.addEventListener('input', (e) => {
      const target = e.target
      if (target.value.length > 30) {
        target.value = target.value.slice(0, 30)
      }
      target.nextElementSibling.textContent = `(${target.value.length}/30)`
    })
  })
  element.querySelectorAll('textarea[data-role="desc"]').forEach(ele => {
    ele.addEventListener('input', (e) => {
      console.log(e.target)
      const target = e.target
      if (target.value.length > 200) {
        target.value = target.value.slice(0, 200)
      }
      target.nextElementSibling.textContent = `(${target.value.length}/200)`
    })
  })
}

function NewAccordionItem(product) {
  const productId = product.product_id;
  const productName = product.product_name;
  const productImg = product.product_img;
  const productDesc = product.product_desc;
  const productCategories = product.product_categories;
  const productInsurance = product.product_insurance;
  const productExpress = product.product_express;
  const productSpecs = product.product_specs;
  const productParams = product.product_params;

  const productCategoriesHtml = productCategories.map(cat => `
    <span class="category">
      <span class="category-name">${cat}</span>
      <a href="javascript:void(0)" class="category-del d-none" data-executor="#categories${productId}" data-role="categoryDel">&otimes;</a>
    </span>
  `).join('');

  const productSpecsHtml = productSpecs.map((spec, idx) => `
    ${idx % 2 === 0 ? '<tr>' : ''}
      <td>${spec.spec_name}</td>
      <td>${spec.spec_price}</td>
      <td>${spec.spec_stock}</td>
      <td data-row-idx="${Math.floor(idx / 2) + 1}" data-col-idx="${(idx % 2) + 1}">
        <a href="javascript:void(0)" class="implicit-link change-success disabled-white">
          <span data-role="modifyRow" data-domain="spec" data-target="#spec-tbody${productId}">修改</span>
        </a>
        <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
          <span data-role="deleteRow">删除</span>
        </a>
      </td>
    ${idx % 2 === 1 ? '</tr>' : ''}
  `).join('');

  const productParamsHtml = productParams.map((param, idx) => `
    ${idx % 2 === 0 ? '<tr>' : ''}
      <td>${param.param_name}</td>
      <td>${param.param_value}</td>
      <td data-row-idx="${Math.floor(idx / 2) + 1}" data-col-idx="${(idx % 2) + 1}">
        <a href="javascript:void(0)" class="implicit-link change-success disabled-white">
          <span data-role="modifyRow" data-domain="param" data-target="#params-tbody${productId}">修改</span>
        </a>
        <a href="javascript:void(0)" class="implicit-link change-danger disabled-white">
          <span data-role="deleteRow">删除</span>
        </a>
      </td>
    ${idx % 2 === 1 ? '</tr>' : ''}
  `).join('');

  const accordionItem = document.createElement('div');
  accordionItem.classList.add('accordion-item');
  accordionItem.id = `accordion-item${productId}`;
  accordionItem.dataset.productId = productId;
  accordionItem.innerHTML = `
      <h2 class="accordion-header">
        <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#accordion-item${productId}-body">
          ${productName}
        </button>
      </h2>
      <div class="accordion-collapse collapse show" id="accordion-item${productId}-body" data-bs-parent="#product-accordion">
        <div class="accordion-body">
          <div class="row">
            <div class="col-3 text-center">
              <img src="${productImg}" class="img-fluid" id="img${productId}" data-status="no-change"/>
              <input type="file" id="img-input${productId}" accept="image/*" style="display: none;" data-role="img" data-product-id="${productId}" data-target="img${productId}"/>
              <button class="btn btn-success mt-2 d-none" data-accordion-id="#accordion-item${productId}"><span data-role="updateImg" data-target="#img-input${productId}">修改商品图片</span></button>
            </div>
            <div class="col-8 myproduct">
              <div class="d-flex align-items-center">
                <label for="name${productId}" class="form-label mb-0 js-name">商品名：</label>
                <input id="name${productId}" type="text" class="form-control bg-body-tertiary w-75 ms-2 disabled" value="${productName}" data-role="name" />
                <span class="ms-2">(${productName.length}/30)</span>
              </div>
              <span>描述信息:</span>
              <textarea class="form-control bg-body-tertiary mt-1 disabled js-desc" id="desc${productId}" rows="5" data-role="desc">${productDesc}</textarea>
              <div class="text-end">
                <span class="">(${productDesc.length}/200)</span>
              </div>
            </div>
            <div class="category-area">
              <div class="categories" id="categories${productId}">
                <span>商品类别 (${productCategories.length}/5): </span>
                ${productCategoriesHtml}
              </div>
              <div class="category-control">
                <select class="form-select d-none" id="select${productId}">
                  <option value="" disabled selected>选择商品类别</option>
                  <option value="服装">服装</option>
                  <option value="鞋子">鞋子</option>
                  <option value="儿童">儿童</option>
                  <option value="食品">食品</option>
                  <option value="饼干">饼干</option>
                  <option value="软糖">软糖</option>
                  <option value="其他">其他</option>
                </select>
                <a href="javascript:void(0)" class="disable-link d-none">
                  <span data-executor="#categories${productId}" data-role="categoryAdd" data-selector="#select${productId}">+添加类别</span>
                </a>
              </div>
            </div>
            <div class="misc row">
              <div class="insurance col-6">
                <label for="insurance-select${productId}">保险类型</label>
                <select class="form-select disabled" id="insurance-select${productId}">
                  <option value="无保险" ${productInsurance === "无保险" ? "selected" : ""}>无保险</option>
                  <option value="退货险" ${productInsurance === "退货险" ? "selected" : ""}>退货险</option>
                  <option value="运输险" ${productInsurance === "运输险" ? "selected" : ""}>运输险</option>
                </select>
              </div>
              <div class="express col-6">
                <label for="express-select${productId}">快递类型</label>
                <select class="form-select disabled" id="express-select${productId}">
                  <option value="到付" ${productExpress === "到付" ? "selected" : ""}>到付</option>
                  <option value="包邮" ${productExpress === "包邮" ? "selected" : ""}>包邮</option>
                </select>
              </div>
            </div>
            <hr>
            <div class="row">
              <table class="text-center spec-table">
                <caption>商品规格表</caption>
                <thead>
                  <tr>
                    <th>规格名</th>
                    <th>价格</th>
                    <th>库存</th>
                    <th>操作</th>
                    <th>规格名</th>
                    <th>价格</th>
                    <th>库存</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody class="spec-table-body" data-row-nxt-idx="${Math.floor(productSpecs.length / 2) + 1}" data-col-nxt-idx="${(productSpecs.length % 2) + 1}" id="spec-tbody${productId}">
                  ${productSpecsHtml}
                  ${productSpecs.length % 2 === 1 ? `
                    <td></td>
                    <td></td>
                    <td></td>
                    <td data-row-idx="${Math.floor(productSpecs.length / 2) + 1}" data-col-idx="2"></td>
                  </tr>` : ''}
                </tbody>
                <tfoot class="d-none">
                  <tr>
                    <td colspan="8">
                      <a class="implicit-link change-primary" href="javascript:void(0)" data-role="addRow" data-domain="spec" data-target="#spec-tbody${productId}">添加规格</a>
                    </td>
                  </tr>
                </tfoot>
              </table>
            </div>
            <hr>
            <div class="row">
              <table class="text-center params-table">
                <caption>商品参数表</caption>
                <thead>
                  <tr>
                    <th>参数名</th>
                    <th>参数值</th>
                    <th>操作</th>
                    <th>参数名</th>
                    <th>参数值</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody data-row-nxt-idx="${Math.floor(productParams.length / 2) + 1}" data-col-nxt-idx="${(productParams.length % 2) + 1}" id="params-tbody${productId}">
                  ${productParamsHtml}
                  ${productParams.length % 2 === 1 ? `
                    <td></td>
                    <td></td>
                    <td data-row-idx="${Math.floor(productParams.length / 2) + 1}" data-col-idx="2"></td>
                  </tr>` : ''}
                </tbody>
                <tfoot class="d-none">
                  <tr>
                    <td colspan="8">
                      <a class="implicit-link change-primary" href="javascript:void(0)" data-role="addRow" data-domain="param" data-target="#params-tbody${productId}">添加参数</a>
                    </td>
                  </tr>
                </tfoot>
              </table>
            </div>
            <hr>
            <div class="text-end mt-3 onpaper">
              <button class="btn btn-outline-primary btn-sm w-25" data-accordion-id="#accordion-item${productId}" data-role="productEdit">编辑</button>
              <button class="btn btn-outline-danger btn-sm w-25" data-accordion-id="#accordion-item${productId}" data-role="productDel" data-target-id="${productId}">删除</button>
            </div>
            <div class="text-end mt-3 d-none">
              <button class="btn btn-outline-success btn-sm w-25" data-accordion-id="#accordion-item${productId}" data-role="productSave">保存</button>
              <button class="btn btn-outline-secondary btn-sm w-25" data-accordion-id="#accordion-item${productId}" data-role="productDiscard">放弃</button>
            </div>
          </div>
          <hr />
        </div>
      </div>
  `
  return accordionItem
}

function updateIndicators(accordionId) {
  const element = document.querySelector(accordionId)
  const nameInput = element.querySelector('input[data-role="name"]')
  nameInput.nextElementSibling.textContent = `(${nameInput.value.length}/30)`

  const descInput = element.querySelector('textarea[data-role="desc"]')
  descInput.nextElementSibling.textContent = `(${descInput.value.length}/200)`
}

