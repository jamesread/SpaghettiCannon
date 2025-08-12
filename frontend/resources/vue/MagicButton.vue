<template>
	<section>
		<h2>Magic Button</h2>
		<p>Click the button to increase the count.</p>
		<button @click="onClick">{{ state.count }}</button>
	</section>
</template>

<script>
  import { reactive, computed } from 'vue'

  export default {
    setup() {
      const state = reactive({
        count: 0
      })

      const onClick = () => {
        state.count++
      }

      return {
        state,
        onClick
      }
    }
  }
</script>

<style lang = "css" scoped>
  button {
    background-color: #dee3e7;
    color: black;
  }
</style>
