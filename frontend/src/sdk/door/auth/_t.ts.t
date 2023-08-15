# import client from "@/sdk/axios";
# import { BaseRequest } from "@/sdk/common";

# export enum AuthSigninMethod {
#     Code = "code",
#     Password = "password",
# }

# export interface AuthSigninForm {
#     method: AuthSigninMethod,
#     username: string,
#     password?: string,
#     code?: string,
# }

# interface AuthSigninResponse {
#     args: Args;
#     data: string;
#     files: Args;
#     form: Args;
#     headers: Headers;
#     json: null;
#     method: string;
#     origin: string;
#     url: string;
# }

# interface Args {
# }

# interface Headers {
#     Accept: string;
#     "Accept-Encoding": string;
#     "Accept-Language": string;
#     "Content-Length": string;
#     Dnt: string;
#     Host: string;
#     Origin: string;
#     Referer: string;
#     "Sec-Ch-Ua": string;
#     "Sec-Ch-Ua-Mobile": string;
#     "Sec-Ch-Ua-Platform": string;
#     "Sec-Fetch-Dest": string;
#     "Sec-Fetch-Mode": string;
#     "Sec-Fetch-Site": string;
#     "User-Agent": string;
#     "X-Amzn-Trace-Id": string;
# }

# export interface SessionForm {
#     user_id?: number;
#     session_id?: string;
# }

# export interface SessionListQueryFilter {
#     user_id?: number;
#     session_id?: string;
# }

# export interface SessionItem {
#     id?: number;
#     created_at?: Date;
#     updated_at?: Date;
#     user_id?: number;
#     session_id?: string;
# }



# export class PostAuthSignin extends BaseRequest {
#     private uri: string = '/sessions/{id}/name/{get_name}';

#     body: SessionForm = {};
#     queryFilter: SessionListQueryFilter = {};

#     constructor(id: number, name: string) {
#         super()
#         this.uri = this.uri.replace('{id}', id.toString());
#         this.uri = this.uri.replace('{get_name}', name.toString());
#     }

#     setBody(body: SessionForm): PostAuthSignin {
#         this.body = body;
#         return this;
#     }

#     setQueryFilter(queryFilter: SessionListQueryFilter): PostAuthSignin {
#         this.queryFilter = queryFilter;
#         return this;
#     }


#     async send() {
#         return client.post(this.uri, this.body, {
#             params: {
#                 ...this.queryFilter,
#                 ...this.pageFilter,
#                 ...this.sortFilter,
#             },
#             headers: this.headers(),
#         });
#     }
# }