- Mutate $state variables through the declared variable, not through a reference extracted from it.
For example, this is the wrong way to do it:

const hint = hints[num - 1]; // pulls a reference out
hint.visible = true;          // mutates through that reference

And this is the right way of doing it:

hints[num - 1] = { ...hints[num - 1], visible: true }


- This taught me about reading up on developer notes and discussion posts. There is a new warning
circulating that throws when trying to assign a prop before it's initialized. This is the specific
error:

This reference only captures the initial value of `data`. Did you mean to reference it inside a closure instead?
https://svelte.dev/e/state_referenced_locally

I searched through the following resources to find out more about this issue because I thought
it was weird because of the { data } prop. Here are some of the links I looked at:

https://github.com/sveltejs/svelte/issues/16343

https://github.com/sveltejs/svelte/pull/17266

https://svelte.dev/docs/svelte/compiler-warnings