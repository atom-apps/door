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

        <SignInByCode v-if="authSignin.isCodeLoginMode()" />
        <SignInByPassword v-if="authSignin.isPasswordLoginMode()" />
    </div>


</template>

<script lang="ts" setup>
import { PostAuthSignin } from '@sdk/door/auth/signin';
import { reactive } from 'vue';
import SignInByCode from './SigninByCode.vue';
import SignInByPassword from './SigninByPassword.vue';

const authSignin = reactive<PostAuthSignin>(new PostAuthSignin())
</script>