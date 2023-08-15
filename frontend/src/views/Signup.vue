<template>
    <div class="h-screen mx-auto w-full px-10 lg:w-2/5 md:w-1/2 flex justify-start items-center flex-col">
        <div id="logo"
            class="overflow-hidden sm:my-5 sm:w-1/2 sm:h-5 h-30 w-full py-10 bg-cover bg-[url('//www.baidu.com/img/flexible/logo/pc/result@2.png')]">
        </div>

        <form class="flex flex-col w-full my-5" method="post" action="http://localhost:9800/auth/signup">
            <div class="form-control w-full">
                <label class="label">用户名</label>
                <input type="text" name="username" v-model="authSignup.body.username"
                    class="input lg:input-lg input-bordered w-full" />
            </div>


            <div class="form-control w-full my-5">
                <label class="label">密码</label>
                <input type="password" name="password" v-model="authSignup.body.password"
                    class="input lg:input-lg input-bordered w-full" />
            </div>


            <div class="form-control w-full">
                <label class="label">手机号</label>
                <input type="text" name="phone" v-model="authSignup.body.phone"
                    class="input lg:input-lg input-bordered w-full" />
            </div>

            <div class="form-control w-full my-5">
                <label class="label">验证码</label>
                <div class="flex">
                    <input type="text" minlength="6" maxlength="6" name="password" v-model="authSignup.body.phone_code"
                        class="input lg:input-lg input-bordered w-1/3" ref="codeInput" autocomplete="off" />
                    <div class="w-1"></div>
                    <button class="btn lg:btn-lg grow" @click.prevent="sendCode">发送验证码</button>
                </div>
            </div>

            <div class="flex justify-between items-center mt-10">
                <router-link :to="{ name: 'reset-password'}">忘记密码</router-link>
                <router-link :to="{ name: 'signin', params: { app: authSignup.body.app_name}}">已经账号？前往登录</router-link>
            </div>

            <button @click.prevent="authSignup.send" class="my-5 btn btn-lg lg:btn-xl btn-primary my-10" :disabled="authSignup.loading">
                <span class="loading loading-spinner" v-if="authSignup.loading"></span>
                <span v-else>注册新用户</span>
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
import { PostAuthSignup } from '@sdk/door/auth/signup';
import { reactive, ref } from 'vue';


const codeInput = ref();

const authSignup = reactive<PostAuthSignup>(new PostAuthSignup())

// sendCode
const sendCode = () => {
    console.log('send code')
    codeInput.value.focus()
}

</script>