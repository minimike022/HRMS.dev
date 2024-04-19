<script setup>
import HelloWorld from './components/HelloWorld.vue'
import axios from 'axios'


const test = () => {
  axios.get("http://localhost:3000/").then(res => console.log(res.data))
}

test()
</script>

<template>
  <div>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>
  <HelloWorld msg="Vite + Vue" />
</template>

