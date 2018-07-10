<template>
   <div id="app">
    <b-container>
      <div class="text-center">
        <router-link class="link-home"  to="/">Главная</router-link>
        <router-link class="link-home"  to="/about">О программе</router-link>
      </div>
      <img src="./assets/logo.png" class="logo">
      <router-view/>
    </b-container>
  </div>
</template>

<style>
*{
  box-sizing: border-box;
}
.logo{
  display: block;
  margin: 10px auto;
  height: 75px;
}
.link-home{
  margin-left: 15px;
}
</style>
