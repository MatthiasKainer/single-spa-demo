import Vue from "vue";
import VueRouter from "vue-router";
import Dashboard from "../components/dashboard";

Vue.use(VueRouter);

const routes = [
  {
    path: "/vue",
    name: "dashboard",
    component: Dashboard
  },
  {
    path: "/both",
    name: "dashboard",
    component: Dashboard
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
