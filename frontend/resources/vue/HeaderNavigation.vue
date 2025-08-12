<template>
	<nav class="header-navigation" ref="nav">
		<ul role = "menubar">
			<li v-for="el in navLinks">
				<a href = "#" :class = "el.className" @click = "showSection(el.className)" role = "menuitem">
					<Icon :icon = "el.icon" width = "1.6em" height = "1.6em" />
				</a>
			</li>
		</ul>
	</nav>
</template>

<script setup>
	import { ref, onMounted } from 'vue';
	import { Icon } from '@iconify/vue';
	import { useRouter } from 'vue-router';

	const router = useRouter();

	const navLinks = [
		{ icon: "mdi-light:home", className: "home" },
		{ icon: "healthicons:weight-outline", className: "weight" },
		{ icon: "healthicons:heart-outline", className: "heart" }
	];

	const nav = ref(null);

	function showSection(section) {

		router.push({ name: section });
	}

	onMounted(() => {
	});
</script>
