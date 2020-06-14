<template>
  <b-container fluid>
    <b-row class="my-1">
      <b-col sm="2">
        <label for="input-small">Food:</label>
      </b-col>
      <b-col sm="10">
        <b-form-input v-model="foodObject.name" id="input-small" placeholder="Enter the food name"></b-form-input>
      </b-col>
    </b-row>

    <b-row class="my-1">
      <b-col sm="2">
        <label for="input-default">Purchase Date:</label>
      </b-col>
      <b-col sm="10">
        <b-form-input v-model="foodObject.purchase_date" id="input-default" type="date"></b-form-input>
      </b-col>
    </b-row>

    <b-row class="my-1">
      <b-col sm="2">
        <label for="input-default">Expired Date:</label>
      </b-col>
      <b-col sm="10">
        <b-form-input v-model="foodObject.expired_date" id="input-default" type="date"></b-form-input>
      </b-col>
    </b-row>
    <b-button block variant="outline-primary" @click="submitFood">Subbmit</b-button>
  </b-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'addItem',
  data () {
    return {
      foodObject: {
        name: '',
        purchase_date: null,
        expired_date: null
      }
    }
  },
  methods: {
    submitFood () {
      console.log(this.foodObject)
      axios.post('http://35b7a0e708b8.ngrok.io/food', this.foodObject)
        .then(() => {})
        .catch(() => {})
    }
  }
}
</script>
<style>

</style>
