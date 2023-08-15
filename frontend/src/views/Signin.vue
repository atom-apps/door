<template>
    <div class="h-screen mx-auto w-full px-10 lg:w-2/5 md:w-1/2 flex justify-start items-center flex-col">
        <div id="logo" class="overflow-hidden sm:my-5 sm:w-1/2 sm:h-5 h-30 w-full py-10">
            <img src="//www.baidu.com/img/flexible/logo/pc/result@2.png" alt="" class="max-w-full h-full">
        </div>

        <div class="tabs w-full justify-center">
            <a class="grow tab tab-lg tab-lifted" @click="authSignin.setPasswordLoginMode"
                :class="{ 'tab-active': authSignin.isPasswordLoginMode() }">密码登录</a>
            <a class="grow tab tab-lg tab-lifted" @click="authSignin.setCodeLoginMode"
                :class="{ 'tab-active': authSignin.isCodeLoginMode() }">验证码登录</a>
        </div>

        <form class="flex flex-col w-full my-5">

            <template v-if="authSignin.hasError()">
                <Alert type="error" :msg="authSignin.error" class="my-5" />
            </template>

            <div class="form-control w-full">
                <label class="label">用户名/手机/邮箱</label>
                <input type="text" name="username" v-model="authSignin.body.username"
                    class="input lg:input-lg input-bordered w-full" />
            </div>

            <template v-if="authSignin.isPasswordLoginMode()">
                <div class="form-control w-full my-5">
                    <label class="label">密码</label>
                    <input type="password" name="password" v-model="authSignin.body.password"
                        class="input lg:input-lg input-bordered w-full" />
                </div>
            </template>

            <div class="form-control w-full my-5" v-if="authSignin.isCodeLoginMode()">
                <label class="label">验证码</label>
                <div class="flex">
                    <input type="text" minlength="6" maxlength="6" name="password" v-model="authSignin.body.code"
                        class="input lg:input-lg input-bordered w-1/3" ref="codeInput" autocomplete="off" />
                    <div class="w-1"></div>
                    <a class="btn lg:btn-lg grow" @click.prevent="sendCode">发送验证码</a>
                </div>
            </div>

            <div class="flex justify-between items-center">
                <router-link :to="{ name: 'reset-password' }">忘记密码</router-link>
                <router-link :to="{ name: 'signup', params: { app: authSignin.body.app_name } }">注册新账号</router-link>
            </div>

            <button @click.prevent="authSignin.send" class="my-5 btn btn-lg lg:btn-xl btn-primary my-10"
                :disabled="authSignin.loading">
                <span class="loading loading-spinner" v-if="authSignin.loading"></span>
                <span v-else>登录</span>
            </button>


            <div class="divider mb-5">快速登录</div>

            <div class="flex flex-wrap gap-5">
                <button class="btn btn-circle btn-outline" v-for="n in 10" :key="n">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>
            </div>
        </form>
    </div>
</template>

<script lang="ts" setup>
import Alert from '@components/Alert.vue';
import { PostAuthSignin } from '@sdk/door/auth/signin';
import { reactive, ref, watch } from 'vue';


const codeInput = ref();

const authSignin = reactive<PostAuthSignin>(new PostAuthSignin())

// sendCode
const sendCode = () => {
    console.log('send code')
    codeInput.value.focus()
}

watch(() => authSignin.body.method, () => {
    authSignin.clearError()
})
</script>