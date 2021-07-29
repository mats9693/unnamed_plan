import Vue from "vue"
import VueRouter, { RouteConfig } from "vue-router"

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "",
    component: () => import("@/views/home/home.vue"),
    children: [
      {
        path: "",
        name: "home",
        component: () => import("@/views/home/content.vue")
      },
      {
        path: "games",
        name: "games",
        component: () => import("@/views/games/games.vue")
      }
    ]
  },
  {
    path: "/404",
    name: "notFound",
    redirect: { name: "home" }
  },
  {
    path: "*",
    redirect: { name: "notFound" }
  }
]

const router = new VueRouter({
  routes
})

export default router
