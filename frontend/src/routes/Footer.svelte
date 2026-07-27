<script>
	import { resolve } from '$app/paths';
</script>

<footer>
	<div class="links">
		<a
			href="https://github.com/BradyDriebergen/Terrariadle_2.0"
			target="_blank"
			rel="noopener noreferrer"
		>
			<img class="github-icon" src="/logos/github-logo.png" alt="github link" title="github.com" />
		</a>
		<a href={resolve('/about')}>
			<img class="sign" src="/logos/terraria_sign.png" alt="About/help" title="How to play" />
		</a>
		<a href="https://ko-fi.com/bytebeard" target="_blank" rel="noopener noreferrer">
			<img class="kofi-icon" src="/logos/kofi_logo.png" alt="Kofi link" title="ko-fi.com" />
		</a>
	</div>

	<p class="version">{__APP_VERSION__}</p>
</footer>

<style>
	.links {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 0;
	}

	.sign {
		width: auto;
		height: 40px;
		padding: 0 10px;
	}

	.github-icon {
		width: auto;
		height: 26px;
		transition: height 0.2s ease;
	}

	.kofi-icon {
		width: auto;
		height: 24px;
		transition: height 0.2s ease;
	}

	img:hover {
		transform: scale(1.05);
	}

	.version {
		font-size: 14px;
	}

	@media (max-width: 700px) {
		.sign {
			height: 50px;
			padding: 0 15px;
		}

		.kofi-icon,
		.github-icon {
			height: 30px;
		}
	}
</style>
