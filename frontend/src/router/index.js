import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import Layout from "@/layout";

export const constantRoutes = [
  {
    path: "/redirect",
    component: Layout,
    hidden: true,
    children: [
      {
        path: "/redirect/:path(.*)",
        component: () => import("@v/redirect/index"),
      },
    ],
  },
  {
    path: "/login",
    component: () => import("@v/login/index"),
    hidden: true,
  },
  {
    path: "/404",
    component: () => import("@v/error-page/404"),
    hidden: true,
  },
  {
    path: "/401",
    component: () => import("@v/error-page/401"),
    hidden: true,
  },
  {
    path: "/",
    component: Layout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: () => import("@v/dashboard/index"),
        meta: {
          title: "仪表盘",
          icon: "el-icon-location-information",
          affix: true,
        },
      },
    ],
  },
  {
    path: "/other",
    component: Layout,
    redirect: "/other/account",
    meta: { title: "商机", icon: "el-icon-s-promotion" },
    children: [
      {
        path: "account",
        name: "AccountList",
        component: () => import("@v/account/list"),
        meta: { title: "商机客户" },
      },
    ],
  },
  {
    path: "/listen",
    component: Layout,
    redirect: "/listen/list",
    meta: { title: "课程设置", icon: "el-icon-collection" },
    children: [
      {
        path: "images",
        name: "ListenImg",
        component: () => import("@v/listen/images"),
        meta: { title: "课程图上传"},
      },
      {
        path: "list",
        name: "ListenList",
        component: () => import("@v/listen/list"),
        meta: { title: "课程列表"},
      },
    ],
  },
  {
    path: "/rbac",
    component: Layout,
    redirect: "/rbac/user",
    meta: { title: "用户权限", icon: "el-icon-user" },
    children: [
      {
        path: "user",
        name: "UserList",
        component: () => import("@v/rbac/user"),
        meta: { title: "后台用户" },
      },
    ],
  },
  {
    path: "/conf",
    component: Layout,
    redirect: "/conf/list",
    meta: { title: "配置", icon: "el-icon-setting" },
    children: [
      {
        path: "list",
        name: "ConfigList",
        component: () => import("@v/config/list"),
        meta: { title: "配置列表" },
      },
    ],
  },
  {
    path: "/profile",
    component: Layout,
    hidden: true,
    children: [
      {
        path: "",
        component: () => import("@v/profile/index"),
        name: "Profile",
        meta: { title: "个人资料", icon: "user", noCache: true },
      },
    ],
  },
  { path: "*", redirect: "/404", hidden: true },
];

// 路由规则：
// 权限控制为 meta 中的 auth 属性，填写规则：后端路由去掉前缀 [/api/]
// 如果未设置 auth 属性，表示无需权限都可以访问
export const asyncRoutes = [
  { path: "*", redirect: "/404", hidden: true },
];

const createRouter = () =>
  new Router({
    // mode: 'history',
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes,
  });

const router = createRouter();

export function resetRouter() {
  const newRouter = createRouter();
  router.matcher = newRouter.matcher; // reset router
}

export default router;
