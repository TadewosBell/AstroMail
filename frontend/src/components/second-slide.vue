<script setup>
const emit = defineEmits(['NextSlide'])
function NextSlide(folder) {
  emit('NextSlide', folder)
}
</script>
<template>
        <div class="div1">
            <img class="slide_image" alt="Wails logo" src="../assets/images/setup2.png" />
        </div>
        <div class="div2">
            <h1>
                Second Slide
            </h1>
            <p>We are going to go through some instructions now to setup your email server</p>
            <button v-on:click="NextSlide" >Next</button>
        </div>
</template>
<style>
.slide_image {
    width: 100%;
}

.div2 {
    background-color: #f2f6f6;
    width: 100%;
}
</style>