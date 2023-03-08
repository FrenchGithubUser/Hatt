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
          :saved="saving"
          ref="generalSettings"
        />
        <WebsiteLogins v-if="selectedSection === 'website_logins'" />
        <CustomCategories v-if="selectedSection === 'custom_categories'" />
      </div>
      <q-btn
        color="primary"
        :label="$t('settings.save')"
        icon="save"
        no-caps
        @click="updateSettings"
        :loading="saving"
      />
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import SidebarItem from 'components/settings/SidebarItem.vue'
import GeneralSettings from 'src/components/settings/GeneralSettings.vue'
import WebsiteLogins from 'src/components/settings/WebsiteLogins.vue'
import CustomCategories from 'src/components/settings/CustomCategories.vue'

export default defineComponent({
  name: 'SettingsPage',
  components: { SidebarItem, GeneralSettings, WebsiteLogins, CustomCategories },
  data() {
    return {
      settingSections: ['general', 'custom_categories', 'website_logins'],
      selectedSection: 'general',
      saving: false,
      settingsValues: null,
    }
  },
  methods: {
    updateSettings() {
      this.saving = true
      let updatedSettings = {}
      updatedSettings.general = this.$refs.generalSettings.values
      window['go']['main']['App']
        ['UpdateUserSettings'](updatedSettings)
        .then(() => {
          this.$q.notify(this.$t('settings.saved'))
          this.saving = false
        })
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
