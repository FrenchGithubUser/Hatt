<template>
  <div class="selected-websites shadow-3">
    <div class="websites">
      <div
        class="website"
        v-for="(selected, website) in selectedWebsites"
        :key="website"
      >
        <q-checkbox v-model="selectedWebsites[website]" />
        {{ website }}
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SelectedWebsites',
  components: {},
  props: { websites: { type: Array } },
  data() {
    return {
      selectedWebsites: {},
    }
  },
  created() {
    this.updateSelectedWebsites()
  },
  methods: {
    updateSelectedWebsites() {
      let newSelectedWebsites = {}
      this.websites.forEach((website) => {
        if (!(website in this.selectedWebsites)) {
          newSelectedWebsites[website] = true
        } else {
          newSelectedWebsites[website] = this.selectedWebsites[website]
        }
      })
      this.selectedWebsites = newSelectedWebsites
    },
  },
  watch: {
    websites() {
      this.updateSelectedWebsites()
    },
  },
  computed: {
    getSelectedWebsites() {
      let selectedWebsites = []
      Object.keys(this.selectedWebsites).forEach((website) => {
        if (this.selectedWebsites[website]) {
          selectedWebsites.push(website)
        }
      })
      return selectedWebsites
    },
  },
})
</script>
<style lang="scss" scoped>
.selected-websites {
  padding: 10px;
  border-radius: 15px;
  max-width: 80%;
  margin-top: 10px;
  .websites {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }
}
</style>
