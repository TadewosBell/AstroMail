import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // Import the router
import { OhVueIcon, addIcons } from "oh-vue-icons";
import './style.css';
import 'vue-final-modal/style.css'
import { MdEmailRound, IoSend } from "oh-vue-icons/icons";


import { createVfm } from 'vue-final-modal'


const vfm = createVfm()

addIcons(MdEmailRound, IoSend);

const app = createApp(App)
app.use(vfm)
app.use(router)
app.component("v-icon", OhVueIcon);
app.mount('#app');
