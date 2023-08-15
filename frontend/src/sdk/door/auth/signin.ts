import Http from "@/sdk/axios";
import { AuthScopeResponse, BaseRequest, GlobalVar } from "@/sdk/common";
import { AxiosError } from "axios";
import { v4 as uuidv4 } from 'uuid';
import { useRoute } from 'vue-router';

declare const globalVar: GlobalVar;

export enum AuthSigninMethod {
    Code = "code",
    Password = "password",
}

export interface AuthSigninForm {
    method: AuthSigninMethod,
    app_name: string,
    sid?: string,
    username: string,
    password?: string,
    code?: string,
}

export class PostAuthSignin extends BaseRequest {
    private uri: string = '/auth/signin';

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

    body: AuthSigninForm = {
        method: AuthSigninMethod.Password,
        sid: "",
        username: "",
        app_name: "",
    };

    setBody(body: AuthSigninForm): PostAuthSignin {
        this.body = body;
        return this;
    }

    async send() {
        return Http.post<AuthScopeResponse>(this.uri, this.body, {
            params: {
                ...this.pageFilter,
                ...this.sortFilter,
            },
        }).catch(err => {
            this.loader.stop()
            if (err instanceof AxiosError) {
                this.error = err.message
            }
        }).finally(() => {
            this.loader.stop()
        })
    }

    isCodeLoginMode(): boolean {
        return this.body.method === AuthSigninMethod.Code;
    }

    isPasswordLoginMode(): boolean {
        return this.body.method === AuthSigninMethod.Password;
    }

    setCodeLoginMode(): PostAuthSignin {
        this.body.method = AuthSigninMethod.Code;
        return this;
    }

    setPasswordLoginMode(): PostAuthSignin {
        this.body.method = AuthSigninMethod.Password;
        return this;
    }
}