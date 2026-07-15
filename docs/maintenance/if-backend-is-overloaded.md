## If the backend is getting hammered:

If you noticed that your backend is getting hammered with requests, there are simple/temporary fixes that might help it for a bit:

### Removing preloading on hover

Currently, the site preloads `+page.ts` data when a user hovers over a link in the site. While this keeps the site quick and responsive, it can sometimes cause users to call the `init-game` quite frequently when they hover over links. If this ever becomes a problem for the backend, you can disable this for the time being to reduce the load by a small amount. To do this, change the following line:

`frontend/src/routes/app.html` - line 9:

```
<body data-sveltekit-preload-data="hover">
```

into:

```
<body data-sveltekit-preload-data="tap">
```

This changes this process so rather than preloading the data when a user hovers over a link, they click on the link for the data to load. This also creates a bit of a delay between clicking on a link and loading it, but it's not much.
