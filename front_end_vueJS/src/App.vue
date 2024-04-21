<script setup>
import HelloWorld from './components/HelloWorld.vue'
import axios from 'axios'
import {reactive} from 'vue'

let datas = reactive({
  Name: "",
  Age: null,
  Address: "",
  Phone_Number: "",
})




 const test = () => {
  axios.post("http://localhost:3000/postData", {Name: datas.Name, Age: String(datas.Age), Address: datas.Address, Phone_Number: datas.Phone_Number}).then(res => console.log(res)).catch(error => console.log(error))
 }
</script>

<template>
  <div>
    <form @submit.prevent="test()">
      <input v-model="datas.Name" type="text" placeholder="Name">
      <input v-model="datas.Age" type="number" placeholder="Age">
      <input v-model="datas.Address" type="text" placeholder="Address">
      <input v-model="datas.Phone_Number" type="text" placeholder="Phone Number">
      <input  type="submit">
    </form>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>
  <HelloWorld msg="Vite + Vue" />
</template>

