<template>
    <form class="flex flex-col w-full my-5">

        <template v-if="authSignin.hasError()">
            <Alert type="error" :msg="authSignin.error" class="my-5" />
        </template>

        <div class="form-control w-full">
            <label class="label">手机/邮箱</label>
            <input type="text" name="username" v-model="authSignin.body.username"
                class="input lg:input-lg input-bordered w-full" />
        </div>


        <div class="form-control w-full my-5">
            <label class="label">验证码</label>
            <div class="flex">
                <input type="text" minlength="6" maxlength="6" name="password" v-model="authSignin.body.code"
                    class="input lg:input-lg input-bordered w-1/3" ref="codeInput" autocomplete="off" />
                <div class="w-1"></div>
                <label for="captcha_modal" class="btn lg:btn-lg grow">发送验证码</label>
            </div>
        </div>

        <div class="flex justify-between items-center">
            <router-link :to="{ name: 'reset-password' }">忘记密码</router-link>
            <router-link :to="{ name: 'signup', params: { app: authSignin.body.app_name } }">注册新账号</router-link>
        </div>

        <button @click.prevent="authSignin.send" class="my-5 btn btn-lg lg:btn-xl btn-primary my-10"
            :disabled="authSignin.loader.loading">
            <span class="loading loading-spinner" v-if="authSignin.loader.loading"></span>
            <span v-else>登录</span>
        </button>


    </form>


    <input type="checkbox" id="captcha_modal" class="modal-toggle" v-model="modal" />
    <div class="modal">
        <div class="modal-box">
            <template v-if="sendSmsSuccess">
                <div class="flex items-center justify-center py-10">
                    <svg t="1692097646348" class="icon" viewBox="0 0 1024 1024" version="1.1"
                        xmlns="http://www.w3.org/2000/svg" p-id="3978" width="128" height="128"
                        xmlns:xlink="http://www.w3.org/1999/xlink">
                        <path d="M512 512m-512 0a512 512 0 1 0 1024 0 512 512 0 1 0-1024 0Z" fill="#77C225" p-id="3979">
                        </path>
                        <path
                            d="M433.2032 702.8224L256 525.6704l45.2608-45.312L433.152 612.352l292.352-292.352 45.2608 45.2608z"
                            fill="#FFFFFF" p-id="3980"></path>
                    </svg>
                </div>
            </template>

            <template v-else>
                <h3 class="text-lg font-bold">请输入验证码</h3>

                <img @click="captcha.send" :src="captcha.data.image" :data-id="captcha.data.id" alt="验证码图片"
                    class="my-2 w-full">
                <div class="flex">
                    <input type="text" minlength="6" maxlength="6" name="password"
                        class="input lg:input-lg input-bordered w-1/3" v-model="sendSms.body.code" autocomplete="off" />
                    <div class="w-1"></div>

                    <button @click.prevent="userSendSms" class="btn btn-primary lg:btn-lg grow"
                        :disabled="sendSms.loader.loading">
                        <span class="loading loading-spinner" v-if="sendSms.loader.loading"></span>
                        <span v-else>确认</span>
                    </button>
                </div>
            </template>
        </div>
        <label class="modal-backdrop" for="captcha_modal">Close</label>
    </div>
</template>

<script lang="ts" setup>
import { PostAuthSignin } from '@sdk/door/auth/signin';
import { GetServiceCaptchaGenerate } from '@sdk/door/service/captcha';
import { PostServiceSendSms } from '@sdk/door/service/send-sms';
import { onMounted, reactive, ref, watch } from 'vue';


const sendSmsSuccess = ref<boolean>(false)
const modal = ref<boolean>(false)
const codeInput = ref();

const authSignin = reactive<PostAuthSignin>(new PostAuthSignin())
const captcha = reactive(new GetServiceCaptchaGenerate());
const sendSms = reactive(new PostServiceSendSms());

onMounted(() => {
    captcha.send()
})


watch(() => authSignin.body.method, () => {
    authSignin.clearError()
})

watch(() => captcha.data.id, () => {
    sendSms.body.captcha_id = captcha.data.id
})

const userSendSms = () => {
    sendSms.body.phone = authSignin.body.username
    sendSms.send().then(() => {
        sendSmsSuccess.value = true
        setTimeout(() => {
            modal.value = false
            sendSms.reset()
            captcha.send()
            setTimeout(() => {
                sendSmsSuccess.value = false
            }, 1000)
        }, 1000)

        codeInput.value.focus()
    }).catch(() => {
        sendSms.loader.loading = false
    }).finally(() => {
        sendSms.loader.loading = false
    })
}

</script>