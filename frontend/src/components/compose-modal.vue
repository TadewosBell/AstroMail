<script setup lang="ts">
import { ref } from 'vue';
import { VueFinalModal } from 'vue-final-modal';
import {OhVueIcon}  from "oh-vue-icons";
import { EventsOn } from '../../wailsjs/runtime/runtime';

EventsOn('Sent', () => {
  emit('confirm')
});
// No need for title as a prop since "New Message" will be static
defineProps<{
  // You can extend this if you need to accept more props
}>()

const emit = defineEmits<{
  (e: 'send', message: { to: string; subject: string; body: string }): void,
  (e: 'confirm'): void
}>()

// Reactive states for email fields
const to = ref('');
const subject = ref('');
const body = ref('');

const close_modal = ref(true);

const exitCompose = () => {
    emit('confirm')
}

// Function to emit send event with email data
const sendEmail = () => {
  emit('send', { to: to.value, subject: subject.value, body: body.value });
}


</script>

<template >
  <VueFinalModal
    v-model="close_modal"
    :click-to-close="true"

    class="email-compose-modal"
    content-class="email-compose-modal-content"
    overlay-transition="vfm-fade"
    content-transition="vfm-fade"
    background="interactive"
  >
    <h1>New Message</h1>
    <button class="close" v-on:click="exitCompose" >close</button>
    <input v-model="to" placeholder="To" type="email" class="email-input"/>
    <input v-model="subject" placeholder="Subject" class="subject-input"/>
    <textarea  rows="50" cols="10" v-model="body" placeholder="Your message here..." class="body-textarea"></textarea>
    <button class="send" @click="sendEmail">
        <OhVueIcon name="io-send" ></OhVueIcon>
    </button>
  </VueFinalModal>
</template>

<style>
.email-compose-modal {
  display: flex;
  justify-content: center;
  align-items: center;
}
.email-compose-modal-content {
  display: flex;
  flex-direction: column;
  padding: 1rem;
  background: #fff;
  border-radius: 0.5rem;
  bottom: 0;
  right: 0;
  position: absolute;
  width: 50vw;
  height: 50vh;
}
.email-compose-modal-content > * + * {
  margin: 0.5rem 0;
}
.email-compose-modal-content h1 {
  font-size: 1rem;
  margin-right: 540px;
}
.close {
  /* font-size: 1rem; */
  margin-left: 600px;
    margin-top: 0;
    position: absolute;
    background-color: red;
    color: white;
    border-radius: 10px;
    border-style: solid red;
}

.email-compose-modal-content .email-input, 
.email-compose-modal-content .subject-input,
.email-compose-modal-content .body-textarea {
  width: 100%;
  padding: 0.5rem;
  margin: 0.25rem 0;
  border: 1px solid #ccc;
  border-radius: 0.25rem;
}
.send {
  margin: 0.25rem 0 0 auto;
  padding: 0.5rem 8px;
  border: 1px solid;
  border-radius: 0.5rem;
  cursor: pointer;
  color: blue
}
.dark .email-compose-modal-content {
  background: #333;
  color: #fff;
}
.dark .email-compose-modal-content input, 
.dark .email-compose-modal-content textarea {
  background: #222;
  color: #fff;
  border-color: #555;
}

</style>
