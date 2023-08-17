import { AxiosError } from "axios"

export enum SigninMethod {
    Code = "code",
    Password = "password",
}


interface errMsg {
    code: number
    message: string
    data?: errMsgData
}

interface errMsgData {
    status?: string
    details?: errMsgDataDetails[]
}
interface errMsgDataDetails {
    description: string
    field?: string
}

export const checkError = function (err: AxiosError): string[] {
    let errors: string[] = []
    if (err.response) {
        const msg: errMsg = err.response?.data as errMsg
        switch (msg.message) {
            case 'record not found':
                errors.push("找不到记录")
                break
            default:
                errors.push(msg.message)
        }
        return errors
    }

    errors.push(err.message)
    return errors
}
