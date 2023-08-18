<template>
    <div class="h-screen mx-auto w-full px-10 lg:w-2/5 md:w-1/2 flex justify-center items-center flex-col">
        <div id="logo" class="overflow-hidden sm:my-5 sm:w-1/2 sm:h-5 h-30 w-full py-10">
            <img src="//www.baidu.com/img/flexible/logo/pc/result@2.png" alt="" class="max-w-full h-full">
        </div>

        <template v-if="success">
            <div class="flex flex-col items-center justify-center py-10 prose">
                <svg t="1692097646348" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
                    p-id="3978" width="128" height="128" xmlns:xlink="http://www.w3.org/1999/xlink">
                    <path d="M512 512m-512 0a512 512 0 1 0 1024 0 512 512 0 1 0-1024 0Z" fill="#77C225" p-id="3979">
                    </path>
                    <path
                        d="M433.2032 702.8224L256 525.6704l45.2608-45.312L433.152 612.352l292.352-292.352 45.2608 45.2608z"
                        fill="#FFFFFF" p-id="3980"></path>
                </svg>
                <h2 class="my-5 text-green-600">密码重置成功</h2>
                <router-link class="btn" :to="{ name: 'signin', params: { app: router.params.app } }">前往登录</router-link>
            </div>
        </template>

        <template v-else>
            <div class="tabs w-full justify-center mb-10">
                <span class="grow tab tab-lg tab-lifted" :class="{ 'tab-active': step == 1 }">第一步</span>
                <span class="grow tab tab-lg tab-lifted" :class="{ 'tab-active': step == 2 }">第二步</span>
            </div>

            <Alert :msg="errors" type="error" />

            <template v-if="step == 1">
                <div class="form-control w-full">
                    <label class="label">手机/邮箱</label>
                    <input type="text" v-model="formStep1.username" name="username"
                        class="input lg:input-lg input-bordered w-full" />
                </div>

                <div class="form-control w-full">
                    <label class="label">验证码</label>
                    <div class="flex">
                        <input type="text" minlength="6" maxlength="6" name="password" v-model="formStep1.code"
                            class="input lg:input-lg input-bordered w-1/3" ref="codeInput" autocomplete="off" />
                        <div class="w-1"></div>
                        <SendVerifyCode :channel="channel" :duration=120 :to="formStep1.username" />
                    </div>
                </div>

                <button @click.prevent="checkStep1" class="my-5 btn btn-lg lg:btn-xl btn-primary w-full"
                    :disabled="loading">
                    <span class="loading loading-spinner" v-if="loading"></span>
                    <span v-else>下一步</span>
                </button>
            </template>

            <template v-if="step == 2">
                <div class="form-control w-full my-5">
                    <label class="label">新密码</label>
                    <input type="text" name="password" class="input lg:input-lg input-bordered w-full"
                        v-model="formStep2.password" />
                </div>

                <button @click.prevent="checkStep2" class="my-5 btn btn-lg lg:btn-xl btn-primary w-full"
                    :disabled="loading">
                    <span class="loading loading-spinner" v-if="loading"></span>
                    <span v-else>确认修改</span>
                </button>
            </template>
        </template>

    </div>
</template>

<script lang="ts" setup>
import http from '@/axios';
import { checkError } from '@/common';
import Alert from '@components/Alert.vue';
import SendVerifyCode from '@components/SendVerifyCode.vue';
import { reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

const router = useRoute();
const channel = 'reset-password'

interface step1Form {
    channel: string,
    username: string,
    code: string
}
interface step1Response {
    token: string
}

interface step2Form {
    app_name: string
    token: string
    password: string
}

const formStep1 = reactive<step1Form>({
    channel: channel,
    username: "",
    code: "",
})

const formStep2 = reactive<step2Form>({
    app_name: router.params['app'].toString(),
    token: "",
    password: "",
})

const step = ref<number>(1)
const loading = ref<boolean>(false)
const success = ref<boolean>(false)
const errors = ref<string[]>([])

const checkStep1 = () => {
    errors.value = []

    loading.value = true
    const action = "/v1/auth/check-reset-password-code"
    http.post(action, formStep1).then(res => {
        const resp: step1Response = res.data
        console.log(resp)

        formStep2.token = resp.token
        step.value = 2

        errors.value = []
    }).catch(err => {
        loading.value = false
        errors.value.push(...checkError(err))
    }).finally(() => {
        loading.value = false
    })
}

const checkStep2 = () => {
    errors.value = []

    loading.value = true
    const action = "/v1/auth/reset-password-by-token"
    http.post(action, formStep2).then(() => {
        success.value = true
    }).catch(err => {
        loading.value = false
        errors.value.push(...checkError(err))
    }).finally(() => {
        loading.value = false
    })
}

</script>@/common