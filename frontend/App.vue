<template>
	<header>
		<img src = "./SpaghettiCannonLogo.png" class = "logo" />
		<h1>SpaghettiCannon</h1>
	</header>

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

	onMounted(() => {
	  navigation.value.addRouterLink('home');
	  navigation.value.addRouterLink('weight');
	  navigation.value.addRouterLink('heart');
	  sidebar.value.toggle();
	  sidebar.value.stick();
	});


</script>
