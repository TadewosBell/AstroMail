<script setup>
const props = defineProps({
  email: Object
})
</script>
<template>
    <div class="inbox">
        <div class="div5">{{ email?.text }}</div>
    </div>
</template>
<style>
.inbox {
    grid-area: 2 / 10 / 13 / 13;
  background-color: #fff;
}
</style>