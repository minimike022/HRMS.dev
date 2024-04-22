<script setup>
import HelloWorld from './components/HelloWorld.vue'
import axios from 'axios'
import {ref, onMounted} from 'vue'
const testData = ref(null)
const testPrint = async() => {
  const response = await axios.get("http://127.0.0.1:3000/")
  testData.value = response.data
}
onMounted(async() => {
  testPrint()
})

</script>

<template>
  
    {{ testData }}

    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  <HelloWorld msg="Vite + Vue" />
</template>

