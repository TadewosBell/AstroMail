<script setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'

const data = reactive({
  email: null,
})

function GetEmail() {
  Greet(data.name).then(result => {
    console.log(result)
    data.email = JSON.parse(result)
  })
}

</script>

<template>
  <main class="parent">
      <div class="div1"> </div>
      <div class="div2">
      <h2>Inbox</h2>  
      <input class="search" placeholder="search" />
      </div>
      <div class="div3">
        <ul class="items">
          <li class="email" v-on:click="GetEmail">Email one</li>
          <li class="email">Email two</li>
        </ul>
      </div>
      <div class="div4">
        <ul class="items">
          <li class="folders">Inbox</li>
          <li class="folders">compose</li>
        </ul>  
      </div>
      <div class="div5">{{ data.email?.text }}</div>
  </main>
</template>

<style scoped>

.search {
  padding: 12px 16px 12px 16px;
  margin: 0 0 0 20px;
}


.parent {
  display: grid;
  grid-template-columns: repeat(2, 1.5fr) 0.1fr repeat(4, 1fr) 0.1fr 0fr repeat(3, 2.25fr);
  grid-template-rows: 1fr 0.1fr repeat(10, 1fr);
  grid-column-gap: 0px;
  grid-row-gap: 0px;
  height: 100vh;
  widows: 100%;
}

.div1 {
  grid-area: 1 / 1 / 2 / 3;
  overflow: hidden;
  background-color: #dfe3e3;
}

.div2 {
  grid-area: 1 / 3 / 2 / 13;
  width: 100%;
  display: flex;
  align-items: center;
  padding: 0 20px;
  background-color: #f2f6f6;
  --wails-draggable: drag;
  border-bottom: 1px solid rgba(5, 5, 5, 0.06);
}

.div3 {
  grid-area: 2 / 3 / 13 / 9;
  /* background-color: blue; */
  border-right: 1px solid rgba(5, 5, 5, 0.06);
  background-color: #fff;
}

.div4 {
  grid-area: 2 / 1 / 13 / 3;
  overflow: hidden;
  background-color: #dfe3e3;
}

.div5 {
  grid-area: 2 / 10 / 13 / 13;
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

.folders {
  cursor: pointer;
  position: relative;
  padding: 05px 0px 5px 0px;
  list-style-type: none;
}
</style>
