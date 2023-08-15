import Http from "@/sdk/axios";
import { BaseRequest } from "@/sdk/common";

export interface SendSmsForm {
    captcha_id: string,
    phone: string,
    code: string,
}

export class PostServiceSendSms extends BaseRequest {
    private uri: string = '/services/send/sms';

    body: SendSmsForm = {
        captcha_id: "",
        code: "",
        phone: "",
    }

    async send() {
        this.loader.start()
        return Http.post(this.uri, this.body)
    }

    reset() {
        this.body = {
            captcha_id: "",
            code: "",
            phone: "",
        }
    }
}