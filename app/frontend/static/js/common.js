!function () {
  try {
    document.querySelectorAll('.modal').forEach((modal) => {
      const modalClose = modal.querySelector('.modal .btn-close')
      const modalConfirm = modal.querySelector('.modal-footer button')
      modalClose.addEventListener('click', () => {
        modal.style.display = 'none'
      })
      modalConfirm.addEventListener('click', () => {
        modal.style.display = 'none'
      })
    })
  } catch (e) {
    console.error(e)
  }
}();

function alertByModal(title, message) {
  const modal = document.querySelector('#message-modal')
  const modalTitle = modal.querySelector('.modal-title')
  const modalBody = modal.querySelector('.modal-body')

  modalTitle.innerText = title
  modalBody.innerText = message
  modal.style.display = 'block'
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
    return null;
}

function deleteAllCookies() {
  // 获取所有 Cookie
  const cookies = document.cookie.split(";");

  // 遍历每个 Cookie
  for (let i = 0; i < cookies.length; i++) {
    const cookie = cookies[i];
    const eqPos = cookie.indexOf("=");
    const name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;

    // 设置过期时间为过去的日期
    document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT;path=/";
  }
}

function openDB(dbName, osName, dbVersion = 1) {
  return new Promise((resolve, reject) => {

    const request = window.indexedDB.open(dbName, dbVersion);
    request.onupgradeneeded = (event) => {
      const db = event.target.result;
      if (!db.objectStoreNames.contains(osName)) {
        const objectStore = db.createObjectStore(osName, { keyPath: 'id' });
        objectStore.createIndex('id', 'id', { unique: false });
      }
    };

    request.onsuccess = (event) => {
      resolve(event.target.result);
    };

    request.onerror = (event) => {
      reject(event);
    };
  });
}

async function writeDB(dbName, osName, id, data, dbVersion = 1) {
  const db = await openDB(dbName, osName, dbVersion);
  return await new Promise((resolve, reject) => {
    const transaction = db.transaction(osName, 'readwrite');
    const objectStore = transaction.objectStore(osName);
    const req = objectStore.put({ id: id, data: data });

    req.onsuccess = (event) => {
      resolve(event.target.result);
    };

    req.onerror = (event_1) => {
      console.log('writeDB error', event_1);
      reject(event_1);
    };

    transaction.oncomplete = () => {
      console.log('Transaction completed');
    };

    transaction.onerror = (event_2) => {
      console.log('Transaction error', event_2);
      reject(event_2);
    };
  });
}

async function readDB(dbName, osName, id, dbVersion = 1) {
  const db = await openDB(dbName, osName, dbVersion);
  console.log(db)
  return await new Promise((resolve, reject) => {
    const transaction = db.transaction(osName, 'readonly');
    const objectStore = transaction.objectStore(osName);
    const req = objectStore.get(id);

    req.onsuccess = (event) => {
      resolve(event.target.result);
    };

    req.onerror = (event_1) => {
      console.log('readDB error', event_1);
      reject(event_1);
    };

    transaction.oncomplete = () => {
      console.log('Transaction completed');
    };

    transaction.onerror = (event_2) => {
      console.log('Transaction error', event_2);
      reject(event_2);
    };
  });
}

!function logout() {
  document.querySelectorAll('[data-role="logout"]').forEach(ele => {
    ele.addEventListener('click', () => {
      localStorage.clear();
      deleteAllCookies();
      window.location.href = '/';
    });
  });
}()

export { alertByModal };
export { setCookie, delCookie, getCookie, deleteAllCookies };
export { writeDB, readDB };