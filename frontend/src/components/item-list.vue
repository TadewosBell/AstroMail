<script setup>
import { createDisplayDate } from "../../util/time"
import { parseEmailAndName, previewString } from '../../util/address';
import { reactive } from "vue";

const data = reactive({
  selected_item: null,
})

defineProps({
  folder: Array,
})
const emit = defineEmits(['inFocus', 'itemSelected'])

const selectItem = (id) => {
  emit('itemSelected', id)
  data.selected_item = id;
}


</script>
<template>
  <div class="item-list">
    <ul class="items">
      <li :class="item.id == data.selected_item? 'selected email': 'email'" v-for="item, in folder" v-on:click="selectItem(item.id)">
        <h3 class="name">
          {{
            previewString(parseEmailAndName(item?.from).name? parseEmailAndName(item?.from).name: parseEmailAndName(item?.from).email, 10)
          }}
        </h3>
        <h3 class="subject">
          {{ previewString(item.subject, 20) }}
        </h3>
        <span class="time">
          {{ createDisplayDate(item.date) }}
        </span>
      </li>
    </ul>
  </div>
</template>
<style>
.item-list {
  grid-area: 2 / 2 / 13 / 6;
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
  display: flex;
  align-items: center;
  height: 50px;
  border-bottom: 1px solid whitesmoke;
  cursor: pointer;
  z-index: 999;
  min-width: 100%;
}

.email:hover {
  background: #dfe3e3;
  border-top: 1px solid whitesmoke;
  box-shadow: 0 3px 3px -2px rgba(0, 0, 0, 0.24);
}

.selected {
  background: #dfe3e3;
  border-top: 1px solid whitesmoke;
  border-bottom: 1px solid whitesmoke;
}

.name{
  display: flex;
  flex: 0.5;
  align-items: center;
  font-size: 12px;
  padding-left: 10px;
}

.subject{
  display: flex;
  flex: 1;
  text-align: left;
  font-size: 15px;
  /* padding-left: 10px; */
}


.time {
  padding-right: 15px;
  font-size: 12px;
  font-weight: bold;
}

</style>