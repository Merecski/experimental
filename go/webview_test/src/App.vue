<template>
  <div id="app">
    <img src="./assets/logo.png">
    <!-- <router-view/> -->
    App name: {{name}}
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      name: "",
      id: ""

    }
  },
  created() {
    App().then((data, err) => {
      this.name = data.Name
      this.id = data.ID
      console.log(data)
    })
    console.log(App)
    console.log("Hello World")

  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
