<!-- src/views/Register.vue -->
<template>
    <div>
      <h2>注册</h2>
      <form @submit.prevent="register">
        <input v-model="username" placeholder="用户名" />
        <input v-model="email" type="email" placeholder="邮箱" />
        <input v-model="password" type="password" placeholder="密码" />
        <button type="submit">注册</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        username: '',
        email: '',
        password: '',
      };
    },
    methods: {
      async register() {
        try {
          const response = await axios.post('http://localhost:8080/register', {
            username: this.username,
            email: this.email,
            password: this.password,
          });
          alert('注册成功！');
          // 重定向到登录页或首页
          this.$router.push('/login'); // 假设你有一个登录页面
        } catch (error) {
          alert('注册失败：' + error.response.data.error);
        }
      },
    },
  };
  </script>
  