import { boot } from 'quasar/wrappers'
import axios from 'axios'
// import { displayErrors } from 'src/helpers/helpers'

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
axios.defaults.withCredentials = true
const api = axios.create({
  baseURL: 'http://localhost:8081/',
  headers: {
    'Access-Control-Allow-Origin': '*',
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
})

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
})

export { api }

// api.interceptors.request.use(
//   function (config) {
//     if (config.url !== "/user/register") {
//       const accessToken = localStorage.getItem('access_token');
//       if (accessToken) {
//         config.headers.Authorization="access_token_cookie=" + accessToken;
//       }
//       console.log(config.headers)
//     }
//     return config;
//   },
//   function (error) {
//     // Do something with request error
//     return Promise.reject(error);
//   }
// );

api.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    // displayErrors([error.response.data.error])
  }
)
