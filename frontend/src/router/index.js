import { createRouter, createWebHashHistory } from "vue-router";
import Home from '../pages/Home.vue'
import Setup from '../pages/Setup.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/Home',  name: "Home", component: Home },
    { path: '/Setup', component: Setup }
  ],
});


export default router;