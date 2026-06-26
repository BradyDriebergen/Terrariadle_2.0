import { writable } from 'svelte/store';

export const userIdStore = writable<string>('');

export function initSession() {
  if (typeof window !== 'undefined') {
    userIdStore.set(localStorage.getItem('user_id') ?? '');
  }
}