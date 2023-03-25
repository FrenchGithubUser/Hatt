<template>
  <div class="cryptocurrency-address">
    <img :src="'images/cryptocurrencies/' + crypto.name + '.png'" />
    <div class="name">
      {{ crypto.name }}
      <q-icon
        name="content_copy"
        class="cursor-pointer"
        @click="copyAddress"
        size="1.2em"
      />
    </div>
  </div>
</template>

<script>
import { copyLink } from 'src/helpers/helpers'
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SupportButton',
  components: {},
  props: {
    crypto: { type: Object },
  },
  data() {
    return {}
  },
  created() {},
  methods: {
    copyAddress() {
      copyLink(this.crypto.address)
    },
  },
  watch: {},
  computed: {},
})
</script>
<style lang="scss" scoped>
.cryptocurrency-address {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 150px;
  img {
    max-width: 100%;
  }
  .name {
    color: var(--q-primary);
    font-weight: bold;
    margin-top: 5px;
  }
}
</style>
