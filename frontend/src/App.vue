<script setup>
import { onBeforeMount } from 'vue';
import { ModalsContainer, useModal } from 'vue-final-modal'
import ComposeModal from './components/compose-modal.vue'
import { Send_Email, Is_Setup } from '../wailsjs/go/main/App'
import { useRouter } from 'vue-router';

const router = useRouter()

const { open, close } = useModal({
  component: ComposeModal,
  attrs: {
    title: 'Hello World!',
    onConfirm() {
      close()
    },
    onSend(email) {

      console.log(email)
      Send_Email(email.to, email.subject,email.body).then(result => {
    }).catch(error => {
        console.error('Send Email failed:', error);
        // Handle the error appropriately
    });
    },
  },
})

const composeEmail = () => {
  console.log('here')
  open();
}

onBeforeMount(() => {
  Is_Setup().then(result => {
    const status = result;
    console.log("Status: ", status)
    if(status){
      router.push('/Home')
    } else{
      router.push('/Setup')
    }
  })
})


</script>

<template>
  <div>
    <ModalsContainer />
    <router-view @compose-email="composeEmail" />
  </div>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}

body {
  color: black;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
    "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
    sans-serif;
}
</style>
