<template>
  <div class="website-logins">
    <div class="websites">
      <div
        class="website shadow-3"
        v-for="website in websites"
        :key="website"
        @click="editCredentials(website)"
      >
        <div class="name">
          {{ website.Name }}
          <q-icon name="edit" />
          <!-- <q-icon :name="areCredentialsSaved(website) ? 'done' : 'cross'" /> -->
        </div>
      </div>
    </div>
    <q-dialog v-model="editingCredentials">
      <EditCredentials :website="currentWebsite" />
    </q-dialog>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import EditCredentials from './EditCredentials.vue'

export default defineComponent({
  name: 'WebsiteLogins',
  components: { EditCredentials },
  data() {
    return {
      websites: [],
      currentWebsite: '',
      editingCredentials: false,
    }
  },
  props: {},
  methods: {
    editCredentials(website) {
      this.currentWebsite = website
      this.editingCredentials = true
    },
    // areCredentialsSaved(website) {
    //   let credsSaved = false
    //   this.websites.forEach((w) => {
    //     if (w.Name === website.Name) {
    //       credsSaved = true
    //       console.log(this.savedCredentials)
    //       let websiteIndex = this.savedCredentials.map((c) => c.Name).indexOf(w)
    //       w.Fields.forEach((field) => {
    //         if (this.savedCredentials[websiteIndex].LoginInfo[field] === '') {
    //           credsSaved = false
    //         }
    //       })
    //     }
    //   })
    //   console.log(credsSaved)
    //   return credsSaved
    // },
  },
  computed: {},
  created() {
    window['go']['main']['App']['GetWebsitesWithLogin']().then((websites) => {
      this.websites = websites
    })
  },
})
</script>
<style lang="scss" scoped>
.website-logins {
  margin-bottom: 20px;
  .websites {
    display: flex;
    .website {
      padding: 10px;
      border-radius: 10px;
    }
  }
}
</style>
