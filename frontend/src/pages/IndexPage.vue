<template>
  <div class="index-page" id="index-page">
    <q-input
      bg-color="blue-grey-3"
      borderless
      v-model="input"
      class="search-bar"
      filled
    >
      <template v-slot:append>
        <q-icon name="search" @click="search" />
      </template>
    </q-input>

    <q-table title="Results" :rows="rows" :columns="columns" row-key="name">
      <template v-slot:body-cell-Thumbnail="props">
        <q-td :props="props"
          ><img :src="props.row.Thumbnail" alt="thumbnail"
        /></q-td>
      </template>
    </q-table>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { searchItems } from 'src/helpers/apiCalls.js'

export default defineComponent({
  name: 'IndexPage',
  data() {
    return {
      input: '',
      columns: [
        {
          name: 'Thumbnail',
          label: 'Thumbnail',
          field: 'Thumbnail',
          align: 'left',
        },
        {
          name: 'Name',
          label: 'Name',
          field: 'Name',
          align: 'left',
          sortable: true,
        },
      ],
      rows: [],
    }
  },
  methods: {
    search() {
      searchItems({ input: this.input }).then((data) => {
        this.rows = data
      })
    },
  },
})
</script>

<style lang="scss" scoped>
.index-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  .search-bar {
    width: 40%;
    margin-top: 50px;
  }
}
</style>
<style lang="scss">
#index-page {
  .search-bar {
    .q-field__control {
      border-radius: 15px !important;
    }
  }
}
</style>
