<script setup>
import { OhVueIcon } from "oh-vue-icons";
import { reactive } from "vue";

const emit = defineEmits(['inFocus', 'FolderSelected'])
const data = reactive({
    selectedFolder: 'inbox'
})

function FolderSelected(folder) {
  data.selectedFolder = folder;
  emit('FolderSelected', folder);
}

const folders = {
  'inbox': {
    text: 'Inbox',
    icon: 'md-inbox'
  },
  'sent': {
    text: 'Sent',
    icon: 'io-send'
  }
}
</script>

<template>
  <div class="sidenav">
    <ul class="items">
      <li class="folders" v-for="(folder, key) in folders" :key="key" @click="FolderSelected(key)">
        <OhVueIcon name="md-navigatenext" v-if="data.selectedFolder === key"></OhVueIcon>
        <OhVueIcon :name="folder.icon"></OhVueIcon>
        {{ folder.text }}
      </li>
    </ul>
  </div>
</template>

<style>
.sidenav {
  grid-area: 2 / 1 / 13 / 2;
  overflow: hidden;
  background-color: #dfe3e3;
}

.folders {
  cursor: pointer;
  position: relative;
  padding: 8px 0px;
  list-style-type: none;
}
</style>
