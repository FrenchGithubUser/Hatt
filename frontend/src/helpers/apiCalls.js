import { api } from 'boot/axios'

export function searchItems(payload) {
  return new Promise((resolve, reject) => {
    return api
      .get(
        'itemsList?input=' + payload.input + '&categories=' + payload.categories
      )
      .then(({ data }) => {
        resolve(data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}
