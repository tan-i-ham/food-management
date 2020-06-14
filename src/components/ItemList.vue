<template>
  <div>
    <add-item></add-item>
      <b-row>
        <b-col >Food List</b-col>
      </b-row>
      <b-table striped hover :items="items" :fields="fields"></b-table>
  </div>
</template>

<script>
import AddItem from './AddItem'
import axios from 'axios'
export default {
  name: 'item',
  components: {
    AddItem
  },
  created: function created () {
    this.getFood()
  },
  data () {
    return {
      fields: ['name', 'purchase_date', 'expired_date'],
      items: []
    }
  },
  methods: {
    getFood () {
      axios.get('http://35b7a0e708b8.ngrok.io/foods')
        .then((resp) => {
          const foodData = resp.data
          foodData.forEach(element => {
            this.items.push(element)
          })
        })
        .catch(() => {})
    }
  }
}
</script>

<style>
.col {
  font-size: 25px;
  font-weight: bold;
  margin: 10px 10px 10px 10px;
}
</style>
