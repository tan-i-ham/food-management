<template>
  <div id="app">
      <navi-bar></navi-bar>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld'
import NaviBar from './components/NaviBar'

export default {
  name: 'App',
  components: {
    HelloWorld,
    naviBar: NaviBar
  }
}
</script>

<style>
</style>
