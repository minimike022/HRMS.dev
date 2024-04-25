<script setup>
import { defineAsyncComponent } from 'vue'
import LoginForms from "../components/login_components/LoginForms.vue"
</script>

<template>
    <div class="bg-[#008DEA] w-screen h-[70dvh]
    absolute bottom-0
    "></div>
    
    <div class="w-screen flex justify-center absolute top-[22dvh] font-Poppins">
        <LoginForms />
    </div>
    <div class="w-screen flex justify-center absolute top-14">
        <img src="/src/assets/motivIT_logo.png" alt="" class="h-[15dvh]">
    </div>
    <h1 class="absolute bottom-2 text-sm left-2 text-white">Copyright © 2024 MotivIT Group. All Rights Reserved. Privacy Policy </h1>
    <h1 class="absolute bottom-2 text-sm right-2 text-white">MotivIT RMS Version 1.0</h1>
</template>