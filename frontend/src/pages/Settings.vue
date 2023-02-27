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
    <div class="main-content">main content</div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import SidebarItem from 'components/settings/SidebarItem.vue'

export default defineComponent({
  name: 'SettingsPage',
  components: { SidebarItem },
  data() {
    return {
      settingSections: ['appearance', 'custom_categories', 'website_logins'],
      selectedSection: '',
    }
  },
  methods: {},
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
