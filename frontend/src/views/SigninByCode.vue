<template>
  <form class="flex flex-col w-full my-5">
    <Alert type="error" :msg="errors" class="my-5" v-if="errors.length > 0" />

    <div class="form-control w-full">
      <label class="label">手机/邮箱</label>
      <input
        type="text"
        name="username"
        v-model="form.username"
        @keypress.enter="usernamePressEnter"
        class="input lg:input-lg input-bordered w-full"
      />
    </div>

    <div class="form-control w-full my-5">
      <label class="label">验证码</label>
      <div class="flex">
        <input
          type="text"
          minlength="6"
          maxlength="6"
          name="password"
          v-model="form.code"
          class="input lg:input-lg input-bordered w-1/3"
          ref="codeInput"
          autocomplete="off"
        />
        <div class="w-1"></div>
        <SendVerifyCode :channel="channel" :duration="10" :to="form.username" />
      </div>
    </div>

    <div class="flex justify-between items-center">
      <router-link :to="{ name: 'reset-password' }" >忘记密码</router-link>
      <router-link :to="{ name: 'signup' }" >注册新账号</router-link>
    </div>

    <button
      @click.prevent="submit"
      class="my-10 btn btn-lg lg:btn-xl btn-primary"
      :disabled="loading"
    >
      <span class="loading loading-spinner" v-if="loading"></span>
      <span v-else>登录</span>
    </button>
  </form>
</template>

<script lang="ts" setup>
const channel = "signin";

import http from "@/axios";
import { SigninMethod } from "@/common";
import Alert from "@components/Alert.vue";
import SendVerifyCode from "@components/SendVerifyCode.vue";
import { UrlBuilder } from "@innova2/url-builder";
import { AxiosError } from "axios";
import { onMounted, reactive, ref } from "vue";

interface Form {
  method: SigninMethod;
  username: string;
  code: string;
}

interface ScopeResponse {
  scope: string;
  code: string;
  redirect: string;
}

const inputUsername = ref<HTMLInputElement>();
const inputPassword = ref<HTMLInputElement>();

const errors = ref<string[]>([]);
const loading = ref(false);
const form = reactive<Form>({
  method: SigninMethod.Code,
  username: "",
  code: "",
});

onMounted(() => {
  inputUsername.value?.focus();
});

const usernamePressEnter = () => {
  if (form.username == "") {
    errors.value.push("请输入用户名");
    return;
  }

  if (form.code == "") {
    inputPassword.value?.focus();
  }
};

const submit = () => {
  errors.value = [];

  if (form.username == "") {
    errors.value.push("请输入用户名");
    return;
  }

  if (form.code == "") {
    errors.value.push("请输入验证码");
    return;
  }

  loading.value = true;

  const url = UrlBuilder.createFromUrl(location.href);
  const query = url.getQueryParams();
  const redirect = query.get("redirect")?.toString();

  let formAction = `/v1/auth/signin`;
  if (redirect && redirect.length > 0) {
    formAction += `?redirect=${redirect}`;
  }

  http
    .post(formAction, form)
    .then((res) => {
      const resp: ScopeResponse = res.data;
      console.log("RESP: ", resp);
      window.location.href = resp.redirect;
    })
    .catch((err) => {
      console.log("ERR: ", err);

      loading.value = false;

      if (err.response) {
        const msg = err.response?.data?.message;
        switch (msg) {
          case "record not found":
            errors.value.push("用户不存在");
            break;
        }
        return;
      }

      if (err instanceof AxiosError) {
        errors.value.push(err.message);
      }
    })
    .finally(() => {
      loading.value = false;
    });
};
</script>
