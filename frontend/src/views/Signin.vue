<template>
    <div class="h-screen mx-auto w-full px-10 lg:w-2/5 md:w-1/2 flex justify-center items-center flex-col">
        <div id="logo" class="overflow-hidden sm:my-5 sm:w-1/2 sm:h-5 h-30 w-full py-10">
            <img src="//www.baidu.com/img/flexible/logo/pc/result@2.png" alt="" class="max-w-full h-full">
        </div>

        <div class="tabs w-full justify-center">
            <a class="grow tab tab-lg tab-lifted" @click="mode = SigninMethod.Password"
                :class="{ 'tab-active': mode == SigninMethod.Password }">密码登录</a>
            <a class="grow tab tab-lg tab-lifted" @click="mode = SigninMethod.Code"
                :class="{ 'tab-active': mode == SigninMethod.Code }">验证码登录</a>
        </div>

        <SignInByPassword v-if="mode == SigninMethod.Password" />
        <SignInByCode v-if="mode == SigninMethod.Code" />
    </div>
</template>

<script lang="ts" setup>
import { SigninMethod } from '@/common';
import { ref } from 'vue';
import SignInByCode from './SigninByCode.vue';
import SignInByPassword from './SigninByPassword.vue';

components: { SignInByPassword, SignInByCode }

const mode = ref<SigninMethod>(SigninMethod.Code)

</script>@/common