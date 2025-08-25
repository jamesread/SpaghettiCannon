<template>
	<Section title = 'Update Box'>

		<p>Type an update and press Enter to log it.</p>
		<form>
			<input type = "text" placeholder = "Enter update here" />

			<fieldset>
				<button type = "submit">Log</button>
			</fieldset>
		</form>
	</Section>
</template>

<script>
	import Section from 'picocrank/vue/components/Section.vue'
	/**
  export default {
    setup() {
      const onKeyDown = (e) => {
        console.log(e.target.value)

        if (e.keyCode === 13) {
        }
      }

      return {
        onKeyDown
      }
    }
  }
  */
</script>
