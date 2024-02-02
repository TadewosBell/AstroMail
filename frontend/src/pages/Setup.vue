<script setup>
import { ref } from 'vue'
import 'vue3-carousel/dist/carousel.css';
import { Carousel, Slide, Pagination, Navigation } from 'vue3-carousel';
import SetupSlide from '../components/setup-slide.vue';
import FirstSlide from '../components/first-slide.vue';
import SecondSlide from '../components/second-slide.vue';
import ThirdSlide from '../components/third-slide.vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter()
const Instructions = ref(null)
let end = false;
const Slides = [FirstSlide, SecondSlide, ThirdSlide]

function Next_Slide() {
    if(end)router.push('/Inbox');
    // Methods are available in this reference
    Instructions.value.next()
    
    console.log(Instructions.value.data.currentSlide.value, Instructions.value.data.maxSlide.value)
    if (Instructions.value.data.currentSlide.value === Instructions.value.data.maxSlide.value) {
        end = true

}
}


</script>

<template>
    <carousel ref="Instructions" :mouseDrag="false" :touchDrag="false">
        <slide v-for="(slideComponent, index) in Slides" :key="index">
            <SetupSlide>
                <component @next-slide="Next_Slide" :is="slideComponent" />
            </SetupSlide>
        </slide>
    </carousel>
</template>

<style>
.carousel {
    background-color: #f2f6f6;
}
</style>
  