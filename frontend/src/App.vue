<script setup>
import { defineAsyncComponent } from 'vue';
import { ModalsContainer, useModal } from 'vue-final-modal'
import ComposeModal from './components/compose-modal.vue'
const { open, close } = useModal({
  component: ComposeModal,
  attrs: {
    title: 'Hello World!',
    onConfirm() {
      close()
    },
  },
})

const composeEmail = () => {
  console.log('here')
  open();
}
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
