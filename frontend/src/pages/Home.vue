<script setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <main class="parent">
      <div class="div1">div1 </div>
      <div class="div2">
      <h2>Inbox</h2>  
      </div>
      <div class="div3">
        <ul class="items">
          <li>Hit the gym</li>
          <li class="checked">Pay bills</li>
        </ul>
      </div>
      <div class="div4">div4 </div>
      <div class="div5">div5</div>
  </main>
</template>

<style scoped>
.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
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
  background-color: red;
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
  background-color: green;
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
.items li {
  cursor: pointer;
  position: relative;
  padding: 12px 8px 12px 40px;
  list-style-type: none;
  background: #eee;
  font-size: 18px;
  transition: 0.2s;

  /* make the list items unselectable */
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
