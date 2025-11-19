<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';
  import GuessPanel from './components/GuessPanel.svelte';
  import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';
</script>

<div>
  <GuessPanel />
  <PrevWeaponGroup previousWeapon={null} />

  <Dropdown />
</div>