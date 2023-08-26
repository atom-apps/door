import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { name:"reset-password" ,path: '/auth/reset-password', component: () => import('@views/ResetPassword.vue') },
    { name: "signin", path: '/auth/signin', component: () => import('@views/Signin.vue') },
    { name: "signup", path: '/auth/signup', component: () => import('@views/Signup.vue') },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export { router };
