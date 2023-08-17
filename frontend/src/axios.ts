import axios, { AxiosRequestConfig } from 'axios'

const defaults: AxiosRequestConfig = {
    method: 'GET',
    baseURL: 'http://localhost:9800/',
    headers: {
    },
    timeout: 10 * 1000,
}



const http = axios.create(defaults)

// Http.interceptors.request.use((config) => {
//     Loader.start()
//     return config
// })
// Http.interceptors.response.use(
//     response => {
//         Loader.stop()
//         return response
//     },
//     error => {
//         console.log("ERR: ", error)
//         Loader.stop()
//         return Promise.reject(error)
//     }
// )
// Http.defaults.headers.common['Cookie'] = "sessionid=d88a1db96a3470b1945fbd527d244479";

export default http

