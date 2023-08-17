import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { name:"reset-password" ,path: '/auth/reset-password/:app', component: () => import('@views/ResetPassword.vue') },
    { name: "signin", path: '/auth/signin/:app', component: () => import('@views/Signin.vue') },
    { name: "signup", path: '/auth/signup/:app', component: () => import('@views/Signup.vue') },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export { router };
