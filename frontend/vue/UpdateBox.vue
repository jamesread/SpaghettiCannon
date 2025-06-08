<template>
	<section class = "weight">
		<h2>Update Box</h2>
		<p>Type an update and press Enter to log it.</p>
		<form>
			<input type = "text" placeholder = "Enter update here" />

			<fieldset>
				<button type = "submit">Log</button>
			</fieldset>
		</form>
	</section>
</template>

<script>
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
