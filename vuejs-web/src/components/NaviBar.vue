<template>
<div>
  <b-navbar type="dark" variant="dark">
    <b-navbar-nav>
      <b-nav-item href="#">FoodManagement</b-nav-item>
      <b-nav-item href="#">Fridge</b-nav-item>
      <b-nav-item href="#">Cabinet</b-nav-item>

      <!-- Navbar dropdowns -->
      <b-nav-item-dropdown text="Food Category" right>
        <b-dropdown-item href="#">Fruits</b-dropdown-item>
        <b-dropdown-item href="#">Meat</b-dropdown-item>
        <b-dropdown-item href="#">Veg</b-dropdown-item>
        <b-dropdown-item href="#">Snacks</b-dropdown-item>
      </b-nav-item-dropdown>

    </b-navbar-nav>
  </b-navbar>
</div>
</template>
<script>
export default {
  name: 'NaviBar',
  data () {
    return {

    }
  }
}

</script>
<style>

</style>
