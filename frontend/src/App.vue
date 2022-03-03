<template>
  <div id="app">
    <h1>CPU 占用情况</h1>
    <CPU />
  </div>
</template>

<script>
import CPU from "./components/CPU.vue";

export default {
  name: "App",
  components: {
    CPU,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
