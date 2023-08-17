import { AxiosError } from "axios"

export enum SigninMethod {
    Code = "code",
    Password = "password",
}


export const checkError = function (err: AxiosError): string[] {
    let errors: string[] = []
    if (err.response) {
        const msg = err.response?.data?.message
        switch (msg) {
            case 'record not found':
                errors.push("找不到记录")
                break
            default:
                errors.push(msg)
        }
        return errors
    }

    errors.push(err.message)
    return errors
}
