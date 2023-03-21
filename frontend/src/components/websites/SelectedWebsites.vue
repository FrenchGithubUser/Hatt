<template>
  <div class="selected-websites shadow-3">
    <div class="websites">
      <div
        :class="{
          website: true,
          searching: searching,
          done: doneWebsites.indexOf(website) >= 0 && searching,
        }"
        v-for="(selected, website) in selectedWebsites"
        :key="website"
      >
        <q-checkbox v-model="selectedWebsites[website]" class="checkbox" />
        {{ website }}
      </div>
    </div>
    <q-linear-progress
      :value="getSearchProgress"
      size="5px"
      :animation-speed="500"
      v-if="searching"
      key="something"
    />
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SelectedWebsites',
  components: {},
  props: {
    websites: { type: Array },
    searching: { type: Boolean },
    doneWebsites: { type: Array },
  },
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
    getSearchProgress() {
      let selectedWebsitesAmount = 0
      Object.values(this.selectedWebsites).forEach((selected) => {
        if (selected) {
          selectedWebsitesAmount++
        }
      })
      return this.doneWebsites.length / selectedWebsitesAmount
    },
  },
})
</script>
<style lang="scss" scoped>
.selected-websites {
  border-radius: 15px;
  max-width: 80%;
  margin-top: 10px;
  overflow: hidden;
  .websites {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    padding: 10px;
    .website {
      margin-right: 10px;
      padding-right: 7px;
      margin-bottom: 2px;
      margin-top: 2px;
      border-radius: 14px;
      &.searching {
        border: solid 2px $primary;
      }
      &.done {
        border: solid 2px $positive;
      }
      .checkbox {
        margin-right: -4px;
      }
    }
  }
}
</style>
