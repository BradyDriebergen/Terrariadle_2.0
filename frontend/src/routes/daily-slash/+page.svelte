<script lang="ts">
  import GuessPanel from './components/GuessPanel.svelte';
  import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';
  export let data: { userData?: any }; // type this to your shape if you know it
  const { userData } = data;
</script>

{#if userData}
  <div>
    <GuessPanel />
    <PrevWeaponGroup previousWeapon={null} />
  </div>
{:else}
  <!-- Optional: fallback UI if redirected logic ever returns a page -->
  <p>Loading…</p>
{/if}