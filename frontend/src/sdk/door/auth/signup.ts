import { Http } from "@/sdk/axios";
import { AuthScopeResponse, BaseRequest, GlobalVar } from "@/sdk/common";
import { v4 as uuidv4 } from 'uuid';
import { useRoute } from "vue-router";

declare const globalVar: GlobalVar;

export enum AuthSignupMethod {
    Code = "code",
    Password = "password",
}

export interface AuthSignupForm {
    sid?: string,
    app_name?: string,

    username: string,

    phone?: string,
    phone_code?: string,

    email?: string,
    email_code?: string,

    password?: string,
}



export class PostAuthSignup extends BaseRequest {
    private uri: string = '/auth/signup';
    constructor() {
        super()
        const router = useRoute();
        this.body.app_name = router.params['app'].toString()

        let storageSID = localStorage.getItem('sid') || '';
        if (storageSID == '') {
            storageSID = uuidv4();
            localStorage.setItem('sid', storageSID);
        }
        this.body.sid = storageSID;
    }

    body: AuthSignupForm = {
        app_name: "",
        username: "",
    };

    async send() {
        return Http.post<AuthScopeResponse>(this.uri, this.body, {
            params: {
                ...this.pageFilter,
                ...this.sortFilter,
            },
        });
    }
}