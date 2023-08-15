你是一个typescript和golang语言专家，可以协助我完成 golang 语言转换为 typescript 语言。我现在正在将 golang 语言定义的 http 接口转换为一个 typescript 语言请求实例。
golang 的定义有如下规则：
- PageQueryFilter, SortQueryFilter 我已经定义好，你在生成的内容中直接使用 `import { PageFilterInterface, SortFilterInterface } from '@/sdk/common';` 引入即可完成，不需要再重复定义
- PageQueryFilter 的 golang 定义如下，你可以转换为 typescript 后使用
    ```
    type PageQueryFilter struct {
        Page  int `json:"page" form:"page"`
        Limit int `json:"limit" form:"limit"`
    }
    ```
- SortQueryFilter 的 golang 定义如下，你可以转换为 typescript 后使用
    ```
    type SortQueryFilter struct {
        Asc  *string `json:"asc" form:"asc"`
        Desc *string `json:"desc" form:"desc"`
    }
    ``` 

下面是一个我定义的 golang 的http接口, 我会给你说明每一行及指示器的作用.
```
// SignIn Signin
//
//	@Summary		Signin
//	@Description	Signin
//	@Tags			Auth-User
//	@Accept			json
//	@Produce		json
//	@Param			form	body	dto.SignInForm	true	"SignInForm"
//	@Param			queryFilter	query		dto.SessionListQueryFilter	true	"SessionListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter		true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter		true	"SortQueryFilter"
//	@Success		200	{object}	dto.SessionItem
//	@Router			/auth/signin/{id} [post]
func (c *AuthController) SignIn(ctx *fiber.Ctx, id string, sess *dto.SessionListQueryFilter, form *dto.SignInForm, pageFilter *common.PageQueryFilter, sortFilter *common.SortQueryFilter) (*dto.SessionItem,error)
```
1. `// SignIn Signin` 表示这个方法的名称及说明，你需要忽略它
2. 为数据定义的分割分割空行
3. `@Summary` 为方法的名称指示器，并在后面说明了方法名称 `Signin`，后续你需要使用这个定义来实现类名的定义
4. `@Description` 为方法的使用说明指示器，会包含一些说明信息，后续你需要使用此指示器的定义内容来对生成的类进行注释
5.  `@Tags` 为模块名称，可能有子模块，每个模块以 `-` 进行分割
6. `@Accept` 表示接口的数据 `Content-Type` 接收类型，如果为 `json` 则 使用 `application/json` 来请求接口， 如果为 `x-www-form-urlencoded` 则使用 `application/x-www-form-urlencoded` 来请求接口
7. `@Produce` 表明生成的数据内容格式，通常为 `json`
8. `@Param` 此字段为请求数据类型的定义，格式为 "变量名称 变量位置 数据结构类型 是否必须 注释", 变量位置:表示此变量可以从请求体的哪里取值，通常会有 path,query,body,header,三种类型
9. `@Success` 指示器定义返回数据结构你只需要关心最后一个参数定义数据结构即可
10. `@Router` 表明此路由的请求路径，格式: "路径 请求方式" 路径中 "{}" 内的字段为一个path的变量字段，需要定义一个方法来实现对此变量的指定。
11. 方法的定义，你只需要关心入参和返回值，ctx 参数你不需要关心，也不会在生成的代码中体现, id 为path中定义的变量，sess 为上述 `@Param` 中定义的`queryFilter`变量，form 为上述 `@Params` 中定义的`form`变量，pageFilter 为 `@Param` 中定义的 `pageFilter` 变量， sortFilter 为 `@Param` 中定义的 `sortFilter` 变量，返回值为
`*dto.SessionItem,error` 说明，如果请求失败则返回 error 的字符串描述，如果返回成功则为 `SessionItem`的数据结构

我会给你提供不包含包名的数据结构定义，同样你也需要忽略定义中的包名。
我给你发送一段需要转换的代码，你直接转换后输出typescript代码即可，别的解释性文字请不要回答。
生成的代码需要使用 axios 库作为http请求的client，我已经配置好，你可以直接导入使用，导入方式："import client from '@sdk/axios';"
如下是一个简单生成的示例，你可以作为参考。
```
export class AuthSignIn {
    private uri: string = '/auth/sign-in/{id}'
    private pendingURI: string;

    constructor() {
        this.pendingURI = this.uri
    }


    form: AuthSignInForm

    setPathId(val: number) :AuthSignIn{
        this.pendingURI = replace("{id}", val)
    }

    async send() {
        return client.post<AuthSignInResponse>(this.pendingURI, this.form)
    }
}
```
生成格式要求如下
1. 生成的代码内容请使用 markdown 格式返回
2. 生成的类名为 请求方式+Tag指示器去除空格+方法名称, 如上述定义生成类名为 "PostAuthUserSignIn"
3. queryFilter的参数需要和PageQueryFilter,SortQueryFilter合并后整体作为 query 请求发送, 如果参数为空则不发送此参数
4. 你需要把golang 中定义的数据结构转化为 export interface数据结构
5. send 为固定方法，不可以重命名