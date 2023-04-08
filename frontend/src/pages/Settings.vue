<template>
  <div class="settings">
    <div class="sidebar shadow-1">
      <SidebarItem
        v-for="section in settingSections"
        :key="section"
        :name="section"
        :selected="section === selectedSection"
        @click="selectedSection = section"
      />
    </div>
    <div class="main-content">
      <div class="setting-components" v-if="settingsValues !== null">
        <GeneralSettings
          v-if="selectedSection === 'general'"
          :originalValues="settingsValues.general"
          ref="generalSettings"
          @updated="settingsUpdated"
        />
        <WebsiteLogins v-if="selectedSection === 'website_logins'" />
        <CustomLists v-if="selectedSection === 'custom_lists'" />
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import SidebarItem from 'components/settings/SidebarItem.vue'
import GeneralSettings from 'src/components/settings/GeneralSettings.vue'
import WebsiteLogins from 'src/components/settings/WebsiteLogins.vue'
import CustomLists from 'src/components/settings/CustomLists.vue'

export default defineComponent({
  name: 'SettingsPage',
  components: { SidebarItem, GeneralSettings, WebsiteLogins, CustomLists },
  data() {
    return {
      settingSections: ['general', 'custom_lists', 'website_logins'],
      selectedSection: 'general',
      settingsValues: null,
    }
  },
  methods: {
    settingsUpdated() {
      this.settingsValues = window.settings
    },
  },
  created() {
    window['go']['main']['App']['ReadUserSettings']().then((data) => {
      this.settingsValues = data
    })
  },
})
</script>

<style lang="scss" scoped>
.settings {
  display: flex;
  justify-content: space-around;
  .sidebar {
    width: 15%;
    border-radius: 15px;
    padding: 10px;
  }
  .main-content {
    width: 80%;
  }
}
</style>
