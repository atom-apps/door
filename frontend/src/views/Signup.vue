<template>
    <div class="h-screen mx-auto w-full px-10 lg:w-2/5 md:w-1/2 flex justify-start items-center flex-col">
        <div id="logo"
            class="overflow-hidden sm:my-5 sm:w-1/2 sm:h-5 h-30 w-full py-10 bg-cover bg-[url('//www.baidu.com/img/flexible/logo/pc/result@2.png')]">
        </div>

        <form class="flex flex-col w-full my-5" method="post" action="http://localhost:9800/auth/signup">
            <Alert type="error" :msg="errors" v-if="errors.length > 0" />

            <div class="form-control w-full">
                <label class="label">用户名</label>
                <input type="text" name="username" v-model="form.username" ref="inputUsername"
                    class="input lg:input-lg input-bordered w-full" />
            </div>


            <div class="form-control w-full my-5">
                <label class="label">密码</label>
                <input type="password" name="password" v-model="form.password"
                    class="input lg:input-lg input-bordered w-full" />
            </div>


            <div class="form-control w-full">
                <label class="label">手机号</label>
                <input type="text" name="phone" v-model="form.phone" class="input lg:input-lg input-bordered w-full" />
            </div>

            <div class="form-control w-full my-5">
                <label class="label">验证码</label>
                <div class="flex">
                    <input type="text" minlength="6" maxlength="6" name="password" v-model="form.phone_code"
                        class="input lg:input-lg input-bordered w-1/3" ref="codeInput" autocomplete="off" />
                    <div class="w-1"></div>
                    <SendVerifyCode :channel="channel" :duration=120 :to="form.phone" />
                </div>
            </div>

            <div class="flex justify-between items-center mt-10">
                <router-link :to="{ name: 'reset-password' }">忘记密码</router-link>
                <router-link :to="{ name: 'signin', params: { app: form.app_name } }">已经账号？前往登录</router-link>
            </div>

            <button @click.prevent="submit" class="my-5 btn btn-lg lg:btn-xl btn-primary my-10" :disabled="loading">
                <span class="loading loading-spinner" v-if="loading"></span>
                <span v-else>注册新用户</span>
            </button>
        </form>
    </div>
</template>

<script lang="ts" setup>
import http from '@/axios';
import { checkError } from '@/common';
import Alert from '@components/Alert.vue';
import SendVerifyCode from '@components/SendVerifyCode.vue';
import { UrlBuilder } from '@innova2/url-builder';
import { onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

const channel = 'signup'

const router = useRoute();
const codeInput = ref();
const loading = ref(false)
const errors = ref<string[]>([])

const inputUsername = ref<HTMLInputElement>()

onMounted(() => {
    inputUsername.value?.focus()
})

interface Form {
    app_name: string,
    username: string,

    phone: string,
    phone_code: string,

    email: string,
    email_code: string,

    password: string,

    captcha: string,
    captcha_id: string,

}
const form = reactive<Form>({
    app_name: router.params['app'].toString(),
    username: '',
    password: '',
    phone: '',
    phone_code: '',
    email: '',
    email_code: '',
    captcha: '',
    captcha_id: '',
})

interface ScopeResponse {
    scope: string;
    code: string;
    redirect: string;
}

const submit = () => {
    errors.value = []

    if (form.username == '') {
        errors.value.push("请输入用户名")
        return
    }

    if (form.phone == '') {
        errors.value.push("请输入手机号")
        return
    }

    if (form.phone.length > 0 && form.phone_code == '') {
        errors.value.push("验证码错误")
        return
    }


    // if (form.email == '') {
    //     errors.value.push("请输入邮箱地址")
    //     return
    // }

    // if (form.email.length > 0 && form.email_code == '') {
    //     errors.value.push("验证码错误")
    //     return
    // }

    loading.value = true

    const url = UrlBuilder.createFromUrl(location.href)
    const query = url.getQueryParams()
    const redirect = query.get('redirect')?.toString()

    let formAction = `/v1/auth/signup`
    if (redirect && redirect.length > 0) {
        formAction += `?redirect=${redirect}`
    }

    http.post(formAction, form)
        .then(res => {
            const resp: ScopeResponse = res.data
            console.log("RESP: ", resp)
            window.location.href = resp.redirect
        })
        .catch(err => {
            loading.value = false
            errors.value.push(...checkError(err))
        })
        .finally(() => {
            loading.value = false
        })
}



</script>