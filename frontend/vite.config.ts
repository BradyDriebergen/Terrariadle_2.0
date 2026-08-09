import { sveltekit } from '@sveltejs/kit/vite';
import { execSync } from 'child_process';
import { defineConfig } from 'vite';

const appVersion = execSync('git describe --tags --abbrev=0').toString().trim();

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/api': {
				target: 'http://localhost:8080',
				changeOrigin: false
			}
		},
		allowedHosts: ['iodize-maritime-frosting.ngrok-free.dev']
	},
	define: {
		__APP_VERSION__: JSON.stringify(appVersion)
	}
});
