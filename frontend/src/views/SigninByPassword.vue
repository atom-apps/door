<template>
        <form class="flex flex-col w-full my-5">

            <template v-if="authSignin.hasError()">
                <Alert type="error" :msg="authSignin.error" class="my-5" />
            </template>

            <div class="form-control w-full">
                <label class="label">用户名/手机/邮箱</label>
                <input type="text" name="username" v-model="authSignin.body.username"
                    class="input lg:input-lg input-bordered w-full" />
            </div>

                <div class="form-control w-full my-5">
                    <label class="label">密码</label>
                    <input type="password" name="password" v-model="authSignin.body.password"
                        class="input lg:input-lg input-bordered w-full" />
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
</template>

<script lang="ts" setup>
import { PostAuthSignin } from '@sdk/door/auth/signin';
import { reactive } from 'vue';

const authSignin = reactive<PostAuthSignin>(new PostAuthSignin())

</script>