<template>
	<header>
		<img src = "./SpaghettiCannonLogo.png" class = "logo" />
		<h1>SpaghettiCannon</h1>

        <HeaderNavigation />
	</header>

	<main>
		<router-view />

		<footer>
			<span><a href = "https://github.com/jamesread/SpaghettiCannon">SpaghettiCannon</a></span>
			<span id = "version">?</span>
		</footer>
	</main>
</template>
