<script lang="ts">
	import { fly, scale } from 'svelte/transition';

	function handleScroll(event: MouseEvent, id: string) {
		event.preventDefault();
		const element = document.getElementById(id);
		if (element) element.scrollIntoView({ behavior: 'smooth' });
	}
</script>

<div class="how-to-play" in:fly={{ x: 80, duration: 500 }}>
	<div class="directory">
		<h3>How to Play</h3>
		<nav>
			<a class="link" href="#dw" on:click={(e) => handleScroll(e, 'dw')}>Daily Weapon</a>
			<a class="link" href="#tc" on:click={(e) => handleScroll(e, 'tc')}>Connections</a>
			<a class="link" href="#gtn" on:click={(e) => handleScroll(e, 'gtn')}>Guess The NPC</a>
			<a class="link" href="#hm" on:click={(e) => handleScroll(e, 'hm')}>Hangman</a>
			<a class="link" href="#tt" on:click={(e) => handleScroll(e, 'tt')}>TerraTrivia</a>
		</nav>
	</div>

	<h1>About</h1>
	<hr />

	<p>
		Welcome to Terrariadle! This is a passion project based on the best game ever,
		<a class="link" href="https://terraria.org/" target="_blank" rel="noopener noreferrer"
			>Terraria</a
		>. I appreciate you stopping by and taking the time to read this.
	</p>

	<p>
		I'm a solo developer that loved playing daily games such as
		<a
			class="link"
			href="https://www.nytimes.com/games/wordle/index.html"
			target="_blank"
			rel="noopener noreferrer">Wordle</a
		>,
		<a class="link" href="https://loldle.net/" target="_blank" rel="noopener noreferrer">Loldle</a>,
		<a
			class="link"
			href="https://www.nytimes.com/games/connections"
			target="_blank"
			rel="noopener noreferrer">Connections</a
		>, and
		<a class="link" href="https://puzzles.usatoday.com/" target="_blank" rel="noopener noreferrer"
			>Crossword</a
		>. I wanted to make a version of these daily guessing games based on my all-time favorite game.
	</p>

	<p>
		This project aims to remake another developer's project:
		<a class="link" href="https://www.terradle.com/" target="_blank" rel="noopener noreferrer"
			>Terradle</a
		>. They did an amazing job on his site but I wanted to expand and improve on the idea.
	</p>

	<p>
		If you're a developer or just curious about how this was built, the full source code is
		available on GitHub. Feel free to poke around, and if you build something inspired by it, I'd
		love to hear about it.
	</p>

	<p>
		Found a bug <img src="/emojis/Bug.png" alt="" class="emoji-icons" />, have a suggestion
		<img src="/emojis/Suggestion.png" alt="" class="emoji-icons" />, or you want to say hi
		<img src="/emojis/Heart.png" alt="" class="emoji-icons" /> ?
	</p>

	<p>
		If you'd like to support the game, you can
		<a class="link" href="https://ko-fi.com/bytebeard" target="_blank" rel="noopener noreferrer">
			buy me a cup of coffee!
		</a>
		<img src="/emojis/Coffee.png" alt="" id="coffee" />
	</p>

	<h2 id="t">How to play: Terrariadle</h2>
	<hr />
	<p>
		Terrariadle is made up of 5 different mini-games that test you on your knowledge of the game.
		Each puzzle involves different aspects of Terraria, such as weapons, enemies, npcs, and more.
		Below are the five different puzzles you can solve:
	</p>
	<ul style="font-size: 18px;">
		<li>Daily Slash</li>
		<li>Connections</li>
		<li>Guess the NPC</li>
		<li>Hangman</li>
		<li>TerraTrivia</li>
	</ul>
	<p>
		Each of these games have their own unique styles. Some allow you to guess until you solve it,
		others give you a certain amount of chances. Players can solve the daily puzzles every 24 hours.
		Puzzles refresh at mightnight every night. To see specific game functionality, the sections
		below discribe these games in detail.
	</p>

	<h2 id="dw">How to play: Daily Slash</h2>
	<hr />
	<p>
		The Daily Slash plays a lot like Wordle, where you try to guess the daily weapon using hints
		from previous guesses. To start guessing, you will see a module on the top of the page
		containing a search bar. Typing in this search bar lists all of the weapons in the game. You can
		select any of these weapons by clicking/tapping, or you can also navigate using arrow keys and
		guess using the enter key.
	</p>
	<p>
		Throughout guessing, you'll unlock hints that reveal traits about the daily weapon. Each hint is
		more revealing than the last. You can reveal the hint once you've guessed enough times, shown by
		the progress bar.
	</p>
	<p>
		Guessed weapons are displayed below the input prompt. Each guess contains traits specific to the
		guessed weapon. Each trait can have one of three backgrounds:
	</p>
	<ul>
		<li>
			<span style="color: var(--color-red);">Red</span> - Trait is not the same as the daily weapon.
		</li>
		<li>
			<span style="color: var(--color-yellow);">Yellow</span> - Trait is partialy similar to the daily
			weapon.
		</li>
		<li>
			<span style="color: var(--color-green);">Green</span> - Trait is the same as the daily weapon.
		</li>
	</ul>

	<h5>
		<em>
			*Note: Some weapons traits match exactly with others. If you get all the traits right, it
			doesn't mean you won. It means that there is another weapon with the exact same traits.
		</em>
	</h5>

	<p>
		You know you've won when a prompt appears above the guesses displaying "You Got It!" The prompt
		will also display the remaining time until the daily weapon refreshes.
	</p>

	<h4>Extra notes about the game:</h4>
	<ul style="font-size: 16px;">
		<li>- Melee weapons that are thrown count at throwables.</li>
		<li>- Lances count as spears</li>
		<li>- Repeaters and crossbows count as bows</li>
		<li>- Specialty guns count as guns</li>
		<li>- Staffs and wands are grouped together to avoid confusion</li>
		<li>- Wooden swords and phaseblades just count as phaseblade and wooden sword/bow</li>
	</ul>

	<h2 id="tc">How to play: Terraria Connections</h2>
	<hr />
	<p>
		This game is based heavily on the NYT game:
		<a
			class="link"
			href="https://www.nytimes.com/games/connections"
			target="_blank"
			rel="noopener noreferrer">Connections</a
		>. The goal of this game is to find groups of four objects that share something in common. Once
		you've selected four objects, click submit to check if your guess is correct. You have to find
		all four groups without making four mistakes.
	</p>
	<p>
		Because this is based on Terraria, each catagory will be based on different parts of the game.
		This can range from accessories, potions, fish, weapons, biomes, etc. However, catagories are
		not as straightforward as they seem. Use your best judgement to find the similarities between
		the objects.
	</p>

	<h3>Examples of catagories:</h3>
	<ul style="margin-top: -10px;">
		<li>- Slime Bonus Drops</li>
		<li>- Types of Sharks</li>
		<li>- Snow Critters</li>
		<li>- Underground Plants</li>
		<li>- Rarity Colors</li>
	</ul>

	<p>
		Unlike Connections, this game doesn't have difficulty tiers. Each catagory ranges in difficulty,
		so you can have days where the solution is obvious, and others where its more challenging.
	</p>

	<h2 id="gtn">How to play: Guess The NPC</h2>
	<hr />
	<p>
		Terraria's NPCs say all kinds of things. Everyday, a random quote said by a NPC is picked and
		you are prompted to guess who said it. The quote can be from any NPC found in the game excluding
		Santa Claus. You can guess through entering the trade of the NPC (Guide, Dryad, Arms Dealer,
		etc.). You will know when your right when the winning screen appears.
	</p>

	<p>
		When you guess correctly, you will be prompted to play a mini-game. Like NPC quotes, many NPCs
		have unique names. This mini-game will task you with guessing a name the Daily NPC could have.
		This game can only be played once the daily NPC has been guessed.
	</p>

	<p>
		To play this game, you are given four name options. Only one of the names is what the daily NPC
		could be named in the game. The rest of the three are names that other NPCs could be named. You
		get one chance to make your guess. Once the guess is made, the answer will be revealed alongside
		your guess.
	</p>

	<h2 id="hm">How to play: Hangman</h2>
	<hr />
	<p>
		This gamemode is a twist on the classic Hangman game. The goal of this game is to guess the name
		of a random enemy before you hang the Guide. Enemies can consist of any enemy or boss found in
		the game.
	</p>
	<p>
		You can guess by selecting letters on the on-screen keyboard. If the letter is within the
		enemy's name, it will show up on the position of the letter in respect to the enemy's name. If
		the letter is not in the name, it will add a limb to the Guide. If you add all 6 body parts to
		the Guide, you lose and the correct name is revealed.
	</p>

	<p>
		If you struggle at first, don't worry. The more you play, the more you'll develop patterns that
		can help you. Good luck guessing Terrarian!
	</p>

	<h2 id="tt">How to play: TerraTrivia</h2>
	<hr />
	<p>
		TerraTrivia is a new gamemode based on the niche game,
		<a class="link" href="https://7littlewords.com/" target="_blank" rel="noopener noreferrer">
			Seven Little Words
		</a>. The goal of this game is to build all of the Terraria-based terms using the letter bundles
		given to you.
	</p>
	<p>
		You can start building out the terms by selecting letter 'chunks' via the buttons. There is no
		check guess button, rather, you can keep guessing by trying out different combinations. Once
		you've solved a word, it will automatically update the row with the associated clue. Once you've
		solved all 7 terms, you've won the game.
	</p>
	<p>
		Don't be fooled by all the words you can create. While you can create all sorts of whacky terms,
		you can only solve the puzzle with specific Terraria terms. Pay attention to the clues and try
		your best to guess.
	</p>
</div>

<style>
	.directory {
		padding: 10px 20px;
		margin: auto;
		margin-top: -45px;
		width: fit-content;

		background-color: var(--color-lightblue);
		border: 2px solid black;
		border-radius: 25px;
	}
	.directory h3 {
		margin: 5px;
	}
	.directory nav {
		display: flex;
		gap: 20px;
	}
	.directory a {
		font-size: 18px;
	}
	@media (max-width: 600px) {
		.directory a {
			font-size: 14px;
		}
	}

	.link {
		color: rgb(111, 255, 255);
	}
	.link:hover {
		color: yellow;
	}

	.how-to-play {
		background-color: var(--color-button);
		text-align: left;
		margin: 20px auto;
		margin-top: 60px;
		max-width: 600px;
		padding: 1px 15px 1px 15px;
		font-size: 20px;

		border: 2px solid #000;
		border-radius: 15px;
		background-color: var(--color-button);
	}
	.how-to-play h1 {
		margin-bottom: -5px;
	}
	.how-to-play h2 {
		padding-top: 20px;
		margin-bottom: -5px;
	}
	.how-to-play ul {
		list-style-type: none;
		margin: 0;
		padding: 0;
	}

	.emoji-icons {
		max-width: 15px;
		max-height: 15px;
		width: 15px;
		height: 15px;
		object-fit: contain;
	}
</style>
