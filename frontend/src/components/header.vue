<script setup>
</script>
<template>
    <div class="header">
        <h2>Inbox</h2>  
      <input class="search" placeholder="search" />
    </div>
</template>
<style>
.header {
    grid-area: 1 / 3 / 2 / 13;
  width: 100%;
  display: flex;
  align-items: center;
  padding: 0 20px;
  background-color: #f2f6f6;
  --wails-draggable: drag;
  border-bottom: 1px solid rgba(5, 5, 5, 0.06);
}

.search {
  padding: 12px 16px 12px 16px;
  margin: 0 0 0 20px;
}
</style>