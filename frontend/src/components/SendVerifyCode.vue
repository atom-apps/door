<script setup lang="ts">
import http from '@/axios';
import { onMounted, reactive, ref } from 'vue';
import Alert from './Alert.vue';


const props = defineProps({
  to: String,
  duration: { type: Number, required: false, default: 120 },
  theme: { type: String, required: false, default: "btn lg:btn-lg grow" },
})

interface CaptchaResponse {
  id: string
  image: string
}

interface Form {
  to: string
  code: string
  captcha_id: string
}


const form = reactive<Form>({
  to: '',
  code: '',
  captcha_id: '',
})

const errors = ref<string[]>([])
const countDown = ref(0)

const modal = ref(false)
const loading = ref(false)
const sendSuccess = ref(false)
const inputCode = ref<HTMLInputElement>();
const captcha = ref<string>('')

onMounted(() => {
  refreshCaptcha()
})

const refreshCaptcha = () => {
  errors.value = []
  sendSuccess.value = false
  form.code = ''

  http.get('/services/captcha/generate')
    .then(res => {
      const resp: CaptchaResponse = res.data
      captcha.value = resp.image
      form.captcha_id = resp.id
    }).catch(err => {
      console.log(err)
    }).finally(() => {
      console.log('finally')
    })
}
const rules = {
  email: /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/,
  phone: /^(?:(?:\+|00)86)?1[3-9]\d{9}$/,
}

const startCountDown = () => {
  countDown.value = props.duration
  if (countDown.value == 0) {
    countDown.value = 120
  }

  setInterval(() => {
    countDown.value--
  }, 1000)
}

// send verify code
const send = () => {
  form.to = props.to ?? ''

  let actionPath = ''
  if (rules.phone.test(form.to)) {
    actionPath = '/services/send/sms'
  } else if (rules.email.test(form.to)) {
    actionPath = '/services/send/email'
  } else {
    console.log(form.to)
    errors.value.push('请输入正确的手机号或邮箱')
    setTimeout(() => {
      modal.value = false
    }, 1000)
    return
  }

  if (form.code == '') {
    inputCode.value?.focus()
    errors.value.push('请输入验证码')
    return
  }

  sendSuccess.value = false
  loading.value = true


  http.post(actionPath, form)
    .then(() => {
      startCountDown()

      sendSuccess.value = true
      setTimeout(() => {
        modal.value = false
      }, 1000)

    }).catch(err => {
      loading.value = false
      if (err.response) {
        errors.value.push(err.response.data.message)
      } else {
        errors.value.push(err.message)
      }
      console.log(err)
    }).finally(() => {
      loading.value = false
      console.log('finally')
    })
}

</script>

<template>
  <label v-if="countDown > 0" :class="props.theme">{{ countDown }}秒后重新发送</label>
  <label v-else :to="props.to" for="captcha_modal" @click="refreshCaptcha" :class="props.theme">发送验证码</label>

  <input type="checkbox" id="captcha_modal" class="modal-toggle" v-model="modal" />
  <div class="modal">
    <div class="modal-box">
      <template v-if="sendSuccess">
        <div class="flex flex-col items-center justify-center py-10 prose">
          <svg t="1692097646348" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
            p-id="3978" width="128" height="128" xmlns:xlink="http://www.w3.org/1999/xlink">
            <path d="M512 512m-512 0a512 512 0 1 0 1024 0 512 512 0 1 0-1024 0Z" fill="#77C225" p-id="3979">
            </path>
            <path d="M433.2032 702.8224L256 525.6704l45.2608-45.312L433.152 612.352l292.352-292.352 45.2608 45.2608z"
              fill="#FFFFFF" p-id="3980"></path>
          </svg>
          <h2 class="my-10 text-green-600">发送成功</h2>
        </div>
      </template>

      <template v-else>
        <h3 class="text-lg font-bold">请输入验证码</h3>

        <img @click="refreshCaptcha" :src="captcha" :data-id="form.captcha_id" alt="验证码图片" class="my-2 w-full">

        <Alert class="my-5" type="error" :msg="errors" v-if="errors.length > 0" />

        <div class="flex">
          <input type="text" minlength="6" maxlength="6" name="code" class="input lg:input-lg input-bordered w-1/3"
            ref="inputCode" v-model="form.code" autocomplete="off" />
          <div class="w-1"></div>

          <button @click.prevent="send" class="btn btn-primary lg:btn-lg grow" :disabled="loading">
            <span class="loading loading-spinner" v-if="loading"></span>
            <span v-else>确认</span>
          </button>
        </div>
      </template>
    </div>
    <label class="modal-backdrop" for="captcha_modal">Close</label>
  </div>
</template>