<template>
  <div class="index-page" id="index-page">
    <q-input
      bg-color="blue-grey-3"
      borderless
      v-model="input"
      class="search-bar"
      @keyup.enter="search"
      filled
    >
      <template v-slot:append>
        <q-icon name="search" @click="search" />
      </template>
    </q-input>

    <CategorySelector ref="categories" />

    <q-table
      title="Results"
      class="items-table"
      :rows="rows"
      :columns="columns"
      :rows-per-page-options="[0, 25, 50]"
      row-key="name"
      @row-click="rowClicked"
    >
      <template v-slot:body-cell-Thumbnail="props">
        <q-td :props="props"
          ><img :src="props.row.Thumbnail" alt="thumbnail" class="thumbnail"
        /></q-td>
      </template>
    </q-table>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { searchItems } from 'src/helpers/apiCalls.js'
import CategorySelector from 'src/components/categories/CategorySelector.vue'
import { copyToClipboard, Notify } from 'quasar'

export default defineComponent({
  name: 'IndexPage',
  components: { CategorySelector },
  data() {
    return {
      input: '',
      columns: [
        {
          name: 'Thumbnail',
          label: '',
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
        {
          name: 'Website',
          label: 'Website',
          field: 'Website',
          align: 'left',
          sortable: true,
        },
      ],
      rows: [],
    }
  },
  methods: {
    rowClicked(e, row) {
      copyToClipboard(row.Link).then(() => {
        Notify.create('Link copied to clipboard')
      })
    },
    search() {
      let categories = this.$refs.categories.getSelectedCategories
      searchItems({
        input: this.input,
        categories: categories.categories,
      }).then((data) => {
        if (data) {
          this.rows = data
        }
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
    margin-bottom: 20px;
  }
  .items-table {
    margin: 20px 0px;
    .thumbnail {
      max-width: 150px;
      max-height: 150px;
    }
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
