<template>
  <router-view />
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',
  created() {
    window['go']['main']['App']['ReadUserSettings']().then((data) => {
      this.root = document.documentElement
      this.root.style.setProperty(
        '--thumbnails-size',
        data.general.thumbnailsSize + 'px'
      )
    })
  },
})
</script>
