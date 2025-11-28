<script lang="ts">
	let { previousWeapon } = $props();

	type Rarity = 'White' | 'Blue' | 'Green' | 'Orange' | 'Light Red' | 'Pink' | 'Light Purple' | 'Lime' | 'Yellow' | 'Cyan' | 'Red';

	const colors: Record<Rarity, string> = {
		White: '#ffffff', 
		Blue: '#177cff', 
		Green: '#35d400', 
		Orange: '#ffa600', 
		'Light Red': '#ff8080', 
		Pink: '#ffbdf0', 
		'Light Purple': '#be9bdc', 
		Lime: '#69f75c',
		Yellow: '#fbff00', 
		Cyan: '#00fff7', 
		Red: '#ff0000'
	};
</script>

<div>
	<p>5 people have guessed today's weapon</p>
	<p>Yesterday's weapon was:</p>

	<img 
		style="border-color: {colors[previousWeapon.rarity as Rarity]}" 
		src={`/weapons/${previousWeapon.path}`} 
		alt='Previous weapon'
	/>
	<p style="color: {colors[previousWeapon.rarity as Rarity]}">{previousWeapon.name}</p>

	<p>Guess any weapon to begin.</p>
</div>

<style>
	img {
		background-color: var(--color-button);
		padding: 18px;
		width: 45px;
		height: 45px;
		object-fit: contain;

		border-radius: 15px;
		border: 2px solid;
	}
</style>
