<script setup>
</script>
<template>
    <div class="parent">
        <slot></slot>
        <div class="div3"></div>
        <div class="div4"></div>
        <div class="div5"></div>
        <div class="div6"></div>
    </div>
</template>
<style>
.parent {
    display: grid;
    grid-template-columns: 0.1fr repeat(8, 1fr) 0.1fr;
    grid-template-rows: repeat(10, 1fr);
    grid-column-gap: 0px;
    grid-row-gap: 0px;
    width: 100vw;
    min-height: 100vh;
}

.div1 {
    grid-area: 2 / 6 / 10 / 10;
    background-color: #dbe0e0;
}

.div2 {
    grid-area: 2 / 2 / 10 / 6;
    background-color: #dbe0e0;
}

.div3 {
    grid-area: 1 / 2 / 2 / 10;
    background-color: #f2f6f6;
    --wails-draggable:drag
}

.div4 {
    grid-area: 1 / 10 / 11 / 11;
    background-color: #f2f6f6;
}

.div5 {
    grid-area: 10 / 2 / 11 / 10;
    background-color: #f2f6f6;
}

.div6 {
    grid-area: 1 / 1 / 11 / 2;
    background-color: #f2f6f6;
}
</style>