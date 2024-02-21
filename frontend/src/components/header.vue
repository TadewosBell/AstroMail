<script setup>
// MdEmailRound
import {OhVueIcon}  from "oh-vue-icons";
const props = defineProps({
  title: String,
  current_page: Number,
})

const emit = defineEmits(['composeEmail', 'refreshEmail',  'previousPage', 'nextPage'])

</script>
<template>
    <div class="header">
        <h2>{{ title }}</h2>  
        <button class="" v-on:click="emit('refreshEmail')"><OhVueIcon name="md-refresh" ></OhVueIcon> </button>
        <button class="" v-on:click="emit('previousPage')"><OhVueIcon name="md-navigatebefore" ></OhVueIcon> </button>
        <div>{{ current_page }}</div>
        <button class="" v-on:click="emit('nextPage')"><OhVueIcon name="md-navigatenext" ></OhVueIcon> </button>
      <button class="view_code" v-on:click="emit('composeEmail')"><OhVueIcon name="md-email-round" ></OhVueIcon> </button>
    </div>
</template>
<style>
.header {
    grid-area: 1 / 2 / 2 / 13;
  width: 100%;
  display: flex;
  align-items: center;
  padding: 0 20px;
  background-color: #f2f6f6;
  --wails-draggable: drag;
  border-bottom: 1px solid rgba(5, 5, 5, 0.06);
}

.search {
  padding: 12px 16px 12px 16px;
  margin: 0 0 0 20px;
}

.view_code {
  margin-left: 65vw; /* Push the button to the right */
  padding: 8px 16px; /* Add some padding to the button */
  /* Additional styling for the button if needed */
}

</style>