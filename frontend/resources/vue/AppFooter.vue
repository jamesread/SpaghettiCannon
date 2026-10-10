<template>
	<footer title="footer">
		<p>
			<img
				v-if="logoUrl"
				class="logo"
				:src="logoUrl"
				:alt="appName + ' logo'"
				title="application icon"
				style="height: 1em;"
			>
			{{ appName }}
			<span v-if="showVersionNumber && currentVersion">{{ currentVersion }}</span>
		</p>

		<p>
			<span v-if="docsUrl">
				<a :href="docsUrl" target="_blank" rel="noopener noreferrer">Docs</a>
			</span>
			<span v-if="issuesUrl">
				<a :href="issuesUrl" target="_blank" rel="noopener noreferrer">Raise an issue</a>
			</span>
		</p>

		<p v-if="showUpdateLink">
			<a
				id="available-version"
				:href="updateUrl"
				target="_blank"
				rel="noopener noreferrer"
			>
				{{ availableVersion }}
			</a>
		</p>
	</footer>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
	appName: { type: String, required: true },
	logoUrl: { type: String, default: '' },
	currentVersion: { type: String, default: '' },
	showVersionNumber: { type: Boolean, default: true },
	showNewVersions: { type: Boolean, default: true },
	availableVersion: { type: String, default: '' },
	docsUrl: { type: String, default: '' },
	issuesUrl: { type: String, default: '' },
	updateUrl: { type: String, default: '' },
})

const showUpdateLink = computed(() => {
	if (!props.showVersionNumber || !props.showNewVersions) return false
	const v = (props.availableVersion || '').trim()
	if (!v || v === 'none' || v === '?' || v.startsWith('you-are-using')) return false
	return true
})
</script>
