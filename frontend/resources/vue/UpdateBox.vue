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

<script setup>
	import Section from 'picocrank/vue/components/Section.vue'
</script>
