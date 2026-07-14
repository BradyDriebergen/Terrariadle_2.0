import { sveltekit } from '@sveltejs/kit/vite';
import { execSync } from 'child_process';
import { defineConfig } from 'vite';

const appVersion = execSync('git describe --tags --always --dirty').toString().trim();

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/api': {
				target: 'http://localhost:8080',
				changeOrigin: false
			}
		}
	},
	define: {
		__APP_VERSION__: JSON.stringify(appVersion)
	}
});
