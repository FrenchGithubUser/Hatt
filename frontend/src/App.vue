<template>
  <router-view v-if="ready" />
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',
  data() {
    return {
      ready: false,
    }
  },
  created() {
    window['go']['main']['App']['ReadUserSettings']().then((data) => {
      window.settings = data
      this.ready = true

      this.root = document.documentElement
      this.root.style.setProperty(
        '--thumbnails-size',
        data.general.thumbnailsSize + 'px'
      )
    })
  },
})
</script>
