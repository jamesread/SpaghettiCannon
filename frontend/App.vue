<template>
	<header>
		<img src="./SpaghettiCannonLogo.png" class="logo" />
		<h1>{{ pageTitle }}</h1>
	</header>

	<section v-if="configError" class="with-header-and-content config-error-banner">
		<div class="section-header">
			<h2>Configuration error</h2>
		</div>
		<div class="section-content padding">
			<p>{{ configError }}</p>
			<p v-if="configFile">
				Expected file: <code>{{ configFile }}</code>
			</p>
		</div>
	</section>

	<Navigation ref="navigation">
		<div id="layout">

			<Sidebar ref="sidebar" />

			<div id="content">
				<main>
					<router-view />
				</main>

				<AppFooter
					v-if="showFooter"
					:app-name="pageTitle"
					:logo-url="logoUrl"
					:current-version="currentVersion"
					:show-version-number="showVersionNumber"
					:show-new-versions="showNewVersions"
					:available-version="availableVersion"
					:docs-url="docsUrl"
					:issues-url="issuesUrl"
					:update-url="updateUrl"
				/>
			</div>
		</div>
	</Navigation>
</template>

<script setup>
	import Navigation from 'picocrank/vue/components/Navigation.vue';
	import Sidebar from 'picocrank/vue/components/Sidebar.vue';
	import AppFooter from './resources/vue/AppFooter.vue';
	import logoUrl from './SpaghettiCannonLogo.png';
	import { onMounted, onUnmounted, ref } from 'vue';

	const navigation = ref(null);
	const sidebar = ref(null);
	const configError = ref('');
	const configFile = ref('');

	const docsUrl = 'https://github.com/jamesread/SpaghettiCannon';
	const issuesUrl = 'https://github.com/jamesread/SpaghettiCannon/issues/new/choose';
	const updateUrl = 'https://github.com/jamesread/SpaghettiCannon/releases';

	const pageTitle = ref('SpaghettiCannon');
	const showFooter = ref(true);
	const currentVersion = ref('');
	const showVersionNumber = ref(true);
	const showNewVersions = ref(true);
	const availableVersion = ref('');

	function applyInitResponse () {
		const init = window.initResponse;
		if (!init) {
			return;
		}

		showFooter.value = init.showFooter ?? true;
		currentVersion.value = init.currentVersion || '';
		showVersionNumber.value = init.showVersionNumber ?? true;
		showNewVersions.value = init.showNewVersions ?? true;
		availableVersion.value = init.availableVersion || '';
		pageTitle.value = init.pageTitle || 'SpaghettiCannon';
		document.title = pageTitle.value;
	}

	onMounted(async () => {
		try {
			const response = await fetch('/webUiSettings.json');
			if (response.ok) {
				const settings = await response.json();
				configError.value = settings.configError ?? '';
				configFile.value = settings.configFile ?? '';
			}
		} catch (error) {
			console.error('Failed to load web UI settings:', error);
		}

		applyInitResponse();
		window.addEventListener('init-reloaded', applyInitResponse);

		navigation.value.addRouterLink('home');
		navigation.value.addRouterLink('weight');
		navigation.value.addRouterLink('heart');
		navigation.value.addRouterLink('settings');
		sidebar.value.toggle();
		sidebar.value.stick();
	});

	onUnmounted(() => {
		window.removeEventListener('init-reloaded', applyInitResponse);
	});
</script>
