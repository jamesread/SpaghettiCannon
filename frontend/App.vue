<template>
	<header>
		<img src = "./SpaghettiCannonLogo.png" class = "logo" />
		<h1>SpaghettiCannon</h1>
	</header>

	<section v-if = "configError" class = "with-header-and-content config-error-banner">
		<div class = "section-header">
			<h2>Configuration error</h2>
		</div>
		<div class = "section-content padding">
			<p>{{ configError }}</p>
			<p v-if = "configFile">
				Expected file: <code>{{ configFile }}</code>
			</p>
		</div>
	</section>

	<Navigation ref = "navigation">
		<div id = "layout">

			<Sidebar ref = "sidebar" />

			<div id = "content">
				<main>
					<router-view />
				</main>

				<footer>
					<span><a href = "https://github.com/jamesread/SpaghettiCannon">SpaghettiCannon</a></span>
					<span id = "version">?</span>
				</footer>
			</div>
		</div>
	</Navigation>
</template>

<script setup>
	import Navigation from 'picocrank/vue/components/Navigation.vue';
	import Sidebar from 'picocrank/vue/components/Sidebar.vue';
	import { onMounted, ref } from 'vue';

	const navigation = ref(null);
	const sidebar = ref(null);
	const configError = ref('');
	const configFile = ref('');

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

	  navigation.value.addRouterLink('home');
	  navigation.value.addRouterLink('weight');
	  navigation.value.addRouterLink('heart');
	  sidebar.value.toggle();
	  sidebar.value.stick();
	});


</script>
