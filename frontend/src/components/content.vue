<script setup>
import { parseEmailAndName } from '../../util/address';
const props = defineProps({
  email: Object
})
</script>
<template>
    <div class="content_div">
        <div class="content" v-if="email" >
          <div class="content_subject" > {{ email?.subject }}</div>
          <div class="body" v-if="email?.html.trim() !== ''"  v-html="email?.html"></div>
          <div v-else>{{ email?.text }}</div>
          <div class="to" > to: {{email?.to}}</div>
        </div>
    </div>
</template>
<style>
.content_div {
    grid-area: 2 / 6 / 13 / 13;
  background-color: #fff;
}

.content {
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 0.1fr 1fr 0.1fr;
  grid-column-gap: 0px;
  grid-row-gap: 0px;
  min-height: 95vh;
  max-height: 95vh;
}

.content_subject{
  max-height: 5vh;
  border-bottom: 1px solid black;
}

.body{
  overflow: scroll;
}

.to{
  background-color: #dfe3e3;
}
</style>