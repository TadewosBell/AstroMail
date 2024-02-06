<script setup>
// MdEmailRound
import {OhVueIcon}  from "oh-vue-icons";
const props = defineProps({
  title: String,
})

const emit = defineEmits(['composeEmail'])

</script>
<template>
    <div class="header">
        <h2>{{ title }}</h2>  
      <button class="view_code" v-on:click="emit('composeEmail')"><OhVueIcon name="md-email-round" ></OhVueIcon> </button>
    </div>
</template>
<style>
.header {
    grid-area: 1 / 3 / 2 / 13;
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
  margin-left: 800px; /* Push the button to the right */
  padding: 8px 16px; /* Add some padding to the button */
  /* Additional styling for the button if needed */
}

</style>