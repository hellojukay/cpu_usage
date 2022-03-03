<template>
  <div>
    <ol>
      <li v-for="(core, i) in cores" v-bind:key="i">
        <el-progress
          :percentage="cpus[i].toFixed(2)"
          :color="customColorMethod"
        ></el-progress>
      </li>
    </ol>
  </div>
</template>

<script>
export default {
  name: "CPU",
  data() {
    return {
      cpus: [],
      cores: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11],
    };
  },
  methods: {
    format(n) {
      return n.toFixed(2);
    },
    usage() {
      window.go.main.CPU.Usage().then((data) => {
        let r = JSON.parse(data);
        console.log(r);
        if (r) {
          this.cpus = r;
        }
      });
    },
  },
  created() {
    setInterval(this.usage, 1000);
  },
};
</script>

<style>
ol > li {
  list-style: none;
}
</style>