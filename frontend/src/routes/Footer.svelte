<script>
	import { resolve } from '$app/paths';
</script>

<footer>
	<a href="https://x.com/terrariadle" target="_blank" rel="noopener noreferrer">
		<img class="icon" src="/logos/x_logo.png" alt="X link" title="x.com" />
	</a>
	<a href={resolve('/about')}>
		<img class="sign" src="/logos/terraria_sign.png" alt="About/help" title="How to play" />
	</a>
	<a href="https://ko-fi.com/bytebeard" target="_blank" rel="noopener noreferrer">
		<img class="icon" src="/logos/kofi_logo.png" alt="Kofi link" title="ko-fi.com" />
	</a>
</footer>

<style>
	footer {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.sign {
		width: auto;
		height: 40px;
		padding: 0 10px;
	}

	.icon {
		width: auto;
		height: 24px;
		transition: height 0.2s ease;
	}
	
	.icon:hover,
	.sign:hover {
		transform: scale(1.05);
	}
</style>
