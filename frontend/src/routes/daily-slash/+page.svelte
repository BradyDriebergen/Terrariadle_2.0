<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';
  import GuessPanel from './components/GuessPanel.svelte';
  import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';

  interface SearchResult {
    weaponId: number;
    name: string;
    path: string;
  }

  const items: string[] = ['option a', 'option b', 'option c', 'option d', 'option e', 'option f', 'option g'];
</script>

<div>
  <GuessPanel />
  <PrevWeaponGroup previousWeapon={null} />

  <Dropdown 
    selectItem={(item: string) => {console.log(item)}}
    items={items}
  />
</div>