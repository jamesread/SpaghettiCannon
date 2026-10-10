<template>
	<Section title="Update Box">
		<p>Type an update and press Enter to log it.</p>
		<form @submit.prevent="onSubmit">
			<input
				v-model="content"
				type="text"
				placeholder="Enter update here"
			>

			<fieldset>
				<button type="submit" :disabled="submitting || !content.trim()">Log</button>
			</fieldset>
		</form>
	</Section>
</template>

<script setup>
	import Section from 'picocrank/vue/components/Section.vue'
	import { ref } from 'vue'

	const content = ref('')
	const submitting = ref(false)

	async function onSubmit () {
		const value = content.value.trim()
		if (!value || submitting.value) {
			return
		}

		submitting.value = true
		try {
			await window.client.addUpdate({ content: value })
			content.value = ''
		} catch (error) {
			console.error('Failed to add update:', error)
		} finally {
			submitting.value = false
		}
	}
</script>
