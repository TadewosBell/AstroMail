<script setup>
const emit = defineEmits(['inFocus', 'FolderSelected'])
function FolderSelected(folder) {
  emit('FolderSelected', folder)
}
</script>
<template>
    <div class="sidenav">
        <ul class="items">
            <li class="folders" v-on:click="FolderSelected('inbox')" >Inbox</li>
        </ul>
    </div>
</template>
<style>
.sidenav {
    grid-area: 2 / 1 / 13 / 3;
    overflow: hidden;
    background-color: #dfe3e3;
}

.folders {
    cursor: pointer;
    position: relative;
    padding: 05px 0px 5px 0px;
    list-style-type: none;
}
</style>