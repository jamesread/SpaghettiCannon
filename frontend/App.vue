<template>
	<header>
		<img src = "./SpaghettiCannonLogo.png" class = "logo" />
		<h1>SpaghettiCannon</h1>
	</header>

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
</template>

<script setup>
	import Sidebar from 'picocrank/vue/components/Sidebar.vue';
	import { onMounted, ref } from 'vue';

	const sidebar = ref(null);

	onMounted(() => {
	  sidebar.value.addNavigationLink({path: '/', id: 'home', icon: 'mdi-light:home', title: 'Home'});
	  sidebar.value.addNavigationLink({path: '/weight', id: 'weight', icon: 'mdi-light:home', title: 'Weight'});
	  sidebar.value.addNavigationLink({path: '/heart', id: 'heart', icon: 'mdi-light:home', title: 'Heart'});
	  sidebar.value.toggle();
	  sidebar.value.stick();
	});


</script>
