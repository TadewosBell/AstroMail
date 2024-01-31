<script setup>
defineProps({
    folder: Array,
})
const emit = defineEmits(['inFocus', 'itemSelected'])

</script>
<template>
    <div class="item-list">
        <ul class="items">
            <li class="email" v-for="item, in folder" v-on:click="emit('itemSelected', item.id)">
                {{ item.headers }}
            </li>
        </ul>
    </div>
</template>
<style>
.item-list {
    grid-area: 2 / 3 / 13 / 9;
  /* background-color: blue; */
  border-right: 1px solid rgba(5, 5, 5, 0.06);
  background-color: #fff;
}

.items {
  margin: 0;
  padding: 0;
  color: black
}

/* Style the list items */
.email {
  cursor: pointer;
  position: relative;
  padding: 12px 8px 12px 40px;
  list-style-type: none;
  background: rgb(255, 255, 255);
  font-size: 18px;
  transition: 0.2s;

  /* make the list items unselectable */
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  border-bottom: 2px solid rgba(5, 5, 5, 0.06);
}

.email:hover{
  background: #dfe3e3;
}
</style>