<template>
  <q-layout view="hHh lpR fFf">
    <div class="header">
      <div class="left">
        <q-icon
          name="home"
          class="home cursor-pointer"
          size="2.5em"
          v-if="$route.name !== 'home'"
          @click="$router.push({ name: 'home' })"
        />
      </div>
      <div class="right">
        <q-icon
          name="settings"
          class="settings cursor-pointer"
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
})
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
  .right {
    margin-right: 10px;
  }
  .left {
    margin-left: 10px;
  }
}
</style>
