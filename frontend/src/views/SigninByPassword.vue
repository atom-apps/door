<template>
    <form class="flex flex-col w-full my-5">


        <Alert type="error" :msg="errors" class="my-5" v-if="errors.length > 0" />

        <div class="form-control w-full">
            <label class="label">用户名/手机/邮箱</label>
            <input type="text" name="username" @keypress.enter="usernamePressEnter" ref="inputUsername"
                v-model="form.username" class="input lg:input-lg input-bordered w-full" />
        </div>

        <div class="form-control w-full my-5">
            <label class="label">密码</label>
            <input type="password" name="password" v-model="form.password" ref="inputPassword"
                class="input lg:input-lg input-bordered w-full" />
        </div>


        <div class="flex justify-between items-center">
            <router-link :to="{ name: 'reset-password' }">忘记密码</router-link>
            <router-link :to="{ name: 'signup', params: { app: form.app_name } }">注册新账号</router-link>
        </div>

        <button @click.prevent="submit" class="my-5 btn btn-lg lg:btn-xl btn-primary my-10" :disabled="loading">
            <span class="loading loading-spinner" v-if="loading"></span>
            <span v-else>登录</span>
        </button>
    </form>
</template>

<script lang="ts" setup>
import http from "@/axios";
import { SigninMethod } from '@/common';
import Alert from '@components/Alert.vue';
import { UrlBuilder } from '@innova2/url-builder';
import { AxiosError } from "axios";
import { v4 as uuidv4 } from 'uuid';
import { onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

components: {
    Alert
}

interface Form {
    method: SigninMethod,
    app_name: string,
    sid?: string,
    username: string,
    password?: string,
}

interface ScopeResponse {
    scope: string;
    code: string;
    redirect: string;
}

const inputUsername = ref<HTMLInputElement>()
const inputPassword = ref<HTMLInputElement>()

onMounted(() => {
    inputUsername.value?.focus()
})

const usernamePressEnter = () => {
    if (form.username == '') {
        errors.value.push('请输入用户名')
        return
    }

    if (form.password == '') {
        inputPassword.value?.focus()
    }
}


const router = useRoute();
const errors = ref<string[]>([])
const loading = ref<boolean>(false)

const form = reactive<Form>({
    method: SigninMethod.Password,
    app_name: router.params['app'].toString(),
    sid: localStorage.getItem('sid') || '',
    username: '',
    password: '',
})

// 表单提交
const submit = () => {
    errors.value = []

    if (form.username == '') {
        errors.value.push('请输入用户名')
        return
    }

    if (form.password == '') {
        errors.value.push('请输入密码')
        return
    }

    if (form.password && form.password?.length < 6) {
        errors.value.push('密码长度不能小于6位')
        return
    }

    if (form.method != SigninMethod.Password) {
        errors.value.push('登录方式错误')
        return
    }

    if (form.app_name == '') {
        errors.value.push('应用名称不能为空')
        return
    }

    if (form.sid == '') {
        let storageSID = uuidv4();
        localStorage.setItem('sid', storageSID);
        form.sid = storageSID;
        return
    }


    const url = UrlBuilder.createFromUrl(location.href)
    const query = url.getQueryParams()
    const redirect = query.get('redirect')?.toString()

    let formAction = `/auth/signin`
    if (redirect && redirect.length > 0) {
        formAction += `?redirect=${redirect}`
    }

    loading.value = true
    http.post(formAction, form)
        .then(res => {
            const resp: ScopeResponse = res.data
            console.log("RESP: ", resp)
            window.location.href = resp.redirect
        }).catch(err => {
            console.log("ERR: ", err)

            loading.value = false

            if (err.response) {
                const msg = err.response?.data?.message
                switch (msg) {
                    case 'record not found':
                        errors.value.push('用户不存在')
                        break
                }
                return
            }

            if (err instanceof AxiosError) {
                errors.value.push(err.message)
            }

        }).finally(() => {
            loading.value = false
        })
}

</script>