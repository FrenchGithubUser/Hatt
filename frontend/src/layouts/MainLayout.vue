<template>
  <q-layout view="hHh lpR fFf">
    <div class="header">
      <div class="left"></div>
      <div class="right">
        <img
          src="images/github-mark.svg"
          class="github-logo svg icon"
          @click="openLink('https://github.com/FrenchGithubUser/Hatt')"
        />
        <q-icon
          name="home"
          class="home cursor-pointer icon"
          size="2.5em"
          v-if="$route.name !== 'home'"
          @click="$router.push({ name: 'home' })"
        />
        <q-icon
          name="settings"
          class="settings icon"
          size="2.5em"
          v-if="$route.name !== 'settings'"
          @click="$router.push('settings')"
        />
      </div>
    </div>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'MainLayout',
  components: {},
  setup() {},
  methods: {
    openLink(link) {
      window['runtime']['BrowserOpenURL'](link)
    },
  },
})
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
  .right {
    margin-right: 10px;
    display: flex;
    align-items: center;
    // justify-content: center;
    .svg {
      width: 32px;
    }
    .icon {
      margin-left: 10px;
      cursor: pointer;
    }
  }
  .left {
    margin-left: 10px;
  }
}
</style>
