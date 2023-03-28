<template>
  <q-layout view="hHh lpR fFf">
    <div class="header">
      <div class="left"><div class="title">Hatt</div></div>
      <div class="right">
        <img
          src="images/github-mark.svg"
          class="github-logo svg icon custom-icon"
          @click="openLink('https://github.com/FrenchGithubUser/Hatt')"
        />
        <SupportButton class="icon" />
        <LanguagePicker class="icon" />
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
    <q-dialog v-model="updatePopup">
      <UpdatePopup :details="updateDetails" @close="updatePopup = false" />
    </q-dialog>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import SupportButton from 'src/components/header/SupportButton.vue'
import UpdatePopup from 'src/components/header/UpdatePopup.vue'
import LanguagePicker from 'src/components/settings/LanguagePicker.vue'
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'MainLayout',
  components: { LanguagePicker, SupportButton, UpdatePopup },
  data() {
    return {
      updatePopup: true,
      updateDetails: {},
    }
  },
  methods: {
    openLink(link) {
      window['runtime']['BrowserOpenURL'](link)
    },
  },
  created() {
    window['go']['main']['App']['CheckForUpdate']().then((data) => {
      if (data.Name !== '') {
        this.updateDetails = data
        this.updatePopup = true
      }
    })
  },
})
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  justify-content: space-between;
  margin-top: 5px;
  .right {
    margin-right: 10px;
    display: flex;
    align-items: center;
    // justify-content: center;
    .svg {
      width: 32px;
    }
    .icon {
      margin-left: 12px;
      cursor: pointer;
    }
  }
  .left {
    margin-left: 10px;
    .title {
      font-weight: bold;
      font-size: 1.9em;
    }
  }
}
</style>
