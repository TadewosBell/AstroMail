import { createRouter, createWebHashHistory } from "vue-router";
import Home from '../pages/Home.vue'
import Setup from '../pages/Setup.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/Inbox',  name: "Home", component: Home },
    { path: '/', component: Setup }
  ],
});


export default router;