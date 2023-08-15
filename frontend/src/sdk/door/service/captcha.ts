import Http from "@/sdk/axios";
import { BaseRequest } from "@/sdk/common";
import { AxiosError } from "axios";

export interface ServiceCaptchaGenerate {
    image: string,
    id: string,
}



export class GetServiceCaptchaGenerate extends BaseRequest {
    private uri: string = '/services/captcha/generate';

    data: ServiceCaptchaGenerate = {
        image: "",
        id: "",
    }

    async send() {
        this.loader.start()
        return Http.get<ServiceCaptchaGenerate>(this.uri)
            .then(res => {
                this.data.image = res.data.image
                this.data.id = res.data.id

                console.log(this.data)
            })
            .catch(err => {
                if (err instanceof AxiosError) {
                    this.error = err.message
                }
                this.loader.stop()
            })
            .finally(() => {
                this.loader.stop()
            })
    }


}