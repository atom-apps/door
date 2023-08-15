import axios, { AxiosRequestConfig } from 'axios'

const defaults: AxiosRequestConfig = {
    method: 'GET',
    baseURL: 'http://localhost:9800/',
    headers: {
    },
    timeout: 10 * 1000,
}

class LoadingInspector {
    private _loading: boolean = false

    public get loading(): boolean {
        return this._loading
    }

    public set loading(value: boolean) {
        this._loading = value
    }

    start() {
        this.loading = true
    }

    stop() {
        this.loading = false
    }

}
const Loader = new LoadingInspector()
const Http = axios.create(defaults)

Http.interceptors.request.use((config) => {
    Loader.start()
    return config
})
Http.interceptors.response.use(
    response => {
        Loader.stop()
        return response
    },
    error => {
        console.log("ERR: ", error)
        Loader.stop()
        return Promise.reject(error)
    }
)
Http.defaults.headers.common['Cookie'] = "sessionid=d88a1db96a3470b1945fbd527d244479";

export { Http, Loader }
