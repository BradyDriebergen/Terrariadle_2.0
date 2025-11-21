<script lang="ts">
  import GuessPanel from './components/GuessPanel.svelte';
  import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';
  import { setContext } from 'svelte';

  let { data } = $props()

  setContext('weapons', data.weapons);
</script>

<div>
  <GuessPanel />
  <PrevWeaponGroup previousWeapon={null} />
</div>