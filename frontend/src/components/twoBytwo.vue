<script setup>
// add children with div 1, 2, 3 and 4 classnames
/*
    <div class="div1"> 1</div>
    <div class="div2"> 2</div>
    <div class="div3"> 3</div>
    <div class="div4"> 4</div>
*/
</script>
<template>
    <div class="parent">
        <slot></slot>
    </div>
</template>
<style>
.parent {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(2, 1fr);
    grid-column-gap: 0px;
    grid-row-gap: 0px;
    width: 100vw;
    min-height: 100vh;
}

.div1 {
    grid-area: 1 / 1 / 2 / 2;
}

.div2 {
    grid-area: 1 / 2 / 2 / 3;
}

.div3 {
    grid-area: 2 / 1 / 3 / 2;
}

.div4 {
    grid-area: 2 / 2 / 3 / 3;
}</style>
