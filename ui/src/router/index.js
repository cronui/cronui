import {createRouter, createWebHashHistory} from 'vue-router'
import {installService} from '../service';
import DashboardPage from '../page/DashboardPage.vue';
import LoginPage from '../page/LoginPage.vue';
import InstallPage from '../page/InstallPage.vue';
import ErrorPage from '../page/ErrorPage.vue';

const routes = [
    {
        path: '/',
        name: 'home',
        component: DashboardPage
    },
    {
        path: '/login',
        name: 'login',
        component: LoginPage
    },
    {
        path: '/install',
        name: 'install',
        component: InstallPage
    },
    {
        path: '/error',
        name: 'error',
        component: ErrorPage
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
});

router.beforeEach(async(to, from, next) => {
    const installed = await installService.isInstalled();
    if (installed === null && to.path !== '/error') {
        return next('/error');
    }
    if (installed !== null && to.path === '/error') {
        return installed ? next('/login') : next('/install');
    }
    if (installed === false && to.path !== '/install') {
        return next('/install');
    }
    if (installed === true && to.path === '/install') {
        return next('/login');
    }

    next();
})

router.beforeEach((to, from, next) => {
    const publicPages = ['/login', '/install', '/error'];
    const authRequired = !publicPages.includes(to.path);
    const loggedIn = localStorage.getItem('user');

    if (loggedIn && to.path === '/login') {
        return next('/');
    }
    if (authRequired && !loggedIn) {
        return next('/login');
    }

    next();
})

export default router
