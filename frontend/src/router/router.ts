import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Home from '@/views/Home.vue'
import ExportField from '@/views/ExportField.vue'
import SecretConfig from '@/views/SecretConfig.vue'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        meta: {
            title: "开始执行",
            icon: "zhihang",
            isMenu: true
        }
    },
    {
        path: '/sk',
        name: 'secretConfig',
        component: SecretConfig,
        meta: {
            title: "密钥配置",
            icon: "peizhi-xitongpeizhi",
            isMenu: true
        }
    },
    {
        path: '/field',
        name: 'field',
        component: ExportField,
        meta: {
            title: "导出字段",
            icon: "lishixiao",
            isMenu: true,
        }
    },

];


const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})


export default router;
