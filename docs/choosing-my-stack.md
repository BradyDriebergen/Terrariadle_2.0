# Why I Chose This Stack

### Table of Contents

- [SvelteKit](#sveltekit)
- [Go](#go)
- [MongoDB](#mongodb-atlass-free-tier)
- [Caddy](#caddy)
- [Oracle Cloud](#oracle-cloud)

## SvelteKit

Svelte is an open-source frontend framework that helps you build UI components using HTML, CSS, and JavaScript/Typescript. SvelteKit is the official full-stack companion framework built on top of Svelte to handle routing, data loading, and server-side operations.

As stated in other parts of this documentation, this project's frontend was originally programmed in React. When looking at frameworks to use for this refactor, Svelte + SvelteKit stood out to me because of its simplicity and efficiency. I've only ever used React and Angular in the past, so seeing Svelte was a breath of fresh air. Alongside performance, I chose SvelteKit for the following reasons:

- The tutorial was amazing! It included interactive coding modules that touched on all of what both Svelte and SvelteKit had to offer. Super easy to learn from.
- Everything is in HTML/CSS/JS, which meant it had a small learning curve.
- It's a modern framework with awesome tools built in (Vite, Linter, prettier, etc.).
- It compiles down directly to HTML/JS. No virtual DOM, and allows for embedding in go binary.
- Super simple reactivity with states and runes.
- Built in scoped styles.
- File structure and page layout is very simple.

While Svelte came with some amazing features and easy-to-understand tools, it also had some drawbacks. While developing, I noticed the following:

- SvelteKit has a tiny ecosystem. It's hard to find abundant projects, dependencies, stack-overflow posts, and documentation.
- Server side rendering vs static site generating. SvelteKit relies pretty heavily on a Node process to pre-render pages.
- Updates are frequent and tend to change a lot of things (svelte 4 -> svelte 5).
- No clear line between legacy features and current features, (svelte/store vs svelte/state).

I want to highlight a major pain point I had when developing. If you've been working with recent versions of Svelte 5 (I'm on version v5.56.3), you might've noticed a new warning that pops up when using props:

```
This reference only captures the initial value of `data`. Did you mean to reference it inside a closure instead?
https://svelte.dev/e/state_referenced_locally
```

This new warning refers to trying to use a `$prop` before it is initialized. For the most part, this is a good thing to check. Sometimes, when you try and use a `$prop` and it isn't initialized, it could cause problems with rendering/logic without any explicit errors. However, this warning currently gets thrown anytime you want to use a prop at all. It makes it super frustrating because props should be initialized before the component renders. This happens even for `data` props coming from a `+page.ts` load function

To compare these updated changes, here is how `$props` were called in the past:

```ts
let { data }: { data: PageData } = $props();

let guesses: WeaponGuess[] = $state(data.gameContext.guesses);
let prevWeapon: WeaponPreview | null = $state(data.gameContext.previous_weapon);
let finished: boolean = $state(data.gameContext.finished);

let correctWeapon: Weapon | null = $derived(
	finished ? guesses[0].weapon : null,
);
```

However, if you tried implementing props this way now, it would throw you a warning on each line you called `data.gamecontext...`. As a result of this persistent warning, this is the fix I ended up finding:

```ts
let { data }: { data: PageData } = $props();

let guesses: WeaponGuess[] = $state([]);
let prevWeapon: WeaponPreview | null = $state(null);
let finished: boolean = $state(false);

let correctWeapon: Weapon | null = $derived(
	finished ? guesses[0].weapon : null,
);

$effect(() => {
	// Initialize data once pre-fetch is finished
	if (data.gameContext) {
		guesses = data.gameContext.guesses;
		prevWeapon = data.gameContext.previous_weapon;
		finished = data.gameContext.finished;
	}
});
```

As you can see, it adds extra logic there to a framework that's supposed to make reactivity easy. It also doesn't help that there isn't a lot of documentation on this specific issue.

This is a common pain point I've found with working with newer languages. Tooling and features tend to change constantly and cause you to relearn how to do something that was intuitive before. Despite my criticisms, I still really love Svelte. I would choose it over Angular and React in a heartbeat. It's just good to know the potential pain points you might hit when working with newer languages/frameworks.

## Go

Go (also referred to as Golang) is an open-source, statically typed, and compiled programming language. It aims to combine high runtime performance with simplicity and readability.

One of the things I desperately wanted to fix was my backend implementation. Originally, I was using `Node.js` with 2 JavaScript files to host my entire backend. When planning this project, I didn't want to deal with constant crashes and a lack of any structure.

Go was a language I touched on in school and found myself really drawn to it. I always liked the simplicity of Java, being able to focus on the logic rather than the syntax, and I also liked C, having access to references and pointers allowed for some really efficient programming. I feel like Go combines what's best about both of these style of languages.

Alongside its obvious benefits, I chose Go because:

- Go is extremely lightweight, combining the speed of languages like C while also having an efficient garbage collector.
- Concurrency is extremely easy. Mutexes and goroutines have a bit of a learning curve, but are super easy to use.
- Type assignment, using `:=` without having to declare every type.
- Go's standard library is AMAZING! I was able to make a full backend with SSE events with only 2 outside dependencies.
- Go `embed` allows me to run the backend and frontend from the same binary, only needing one process to run for the whole project.
- Error handling is very easy to read and follow.
- Comes with `gofmt`, which most editors can wire up to run on save

Challenges:

- Slightly higher learning curve. I started out programming with OOP languages. Learning Go's package structure/procedural language practices was a bit of a challenge at first.
- Very easy to muddy up a project without proper system design fundamentals.
- Error handling can get a little messy at times. Busies up the code and makes certain parts hard to read.
- Easy to create race conditions with goroutines.

I don't have much to complain about with this language. Everything seems to be extremely easy to work with, like goroutines, mutexes, and HTTP servers. However, the one nuance with this language I've noticed is that it gives the developer little to no guidelines to work with. In something like a frontend framework, you're usually supposed to follow strict rules on file structures, logic hierarchies, and process flows. Go doesn't restrict developers to any of these rules. It's completely up to the user on how they want to design their Go Projects. It's great for control, but it also requires developers to think hard about how they should design their projects. I recommend reading [Managing Go Packages](./managing-go-packages.md) if you want to hear about one of my biggest, self-caused pain points I dealt with when designing with Go.

Overall, I think this control over your Go projects is a great feature, but it's something to consider if you're starting a project and want a little structure to work with. Go is an amazing language to use with so many helpful built in tools.

## MongoDB (Atlas's free tier)

MongoDB is a highly scalable, open-source NoSQL document database. Instead of the rigid tables and rows used in traditional relational (SQL) databases, it stores data in flexible, JSON-like documents called `BSON`.

MongoDB was originally in my previous iteration of this project. I decided to use it again because of its ease of use, how adaptable its schema is, and the web tooling Atlas provides. It was a no-brainer compared to a relational database for a project this size. Treating users and game types as JSON objects compared to table records also let me iterate quickly between schema designs during implementation.

I chose to continue with MongoDB because:

- Records maps easily onto Go structs. A puzzle, a user, or a game mode's state can be marshaled straight to/from `BSON` without the impedance mismatch you get mapping rows to objects in a relational DB.
- Schema flexibility. I tested schemas for multiple game modes that each have different shapes of data. I didn't want to be writing migrations every time I changed a field.
- MongoDB Atlas's free tier is free, and meant zero database ops for me. No patching, no backups to manage, no server to secure.
- MongoDB Atlas provides a super useful web-app for managing data on the database.

Challenges:

- Records don't keep order in the database. I have to explicitly sort records when initially pulling from the database.
- No joins. Anything that would be a simple SQL join becomes either an aggregation pipeline or multiple round trips.
- Atlas's free tier has a low connection limit, which is part of why I ended up building a write-behind caching layer instead of hitting the database directly on every request.
- It's a little difficult to set up locally.

I wanted to compare MongoDB with SQL to show how different my site would've been if I didn't have access to Mongo's generic record finders. Here is a generic function in my program that I use to pull a single item from a collection in my Mongo database:

```go
func FindOne[T any](ctx context.Context, m *MongoDB, collectionName string, filter Filter) (*T, error) {
    collection := m.client.Database(m.dbName).Collection(collectionName)

    var result T

    err := collection.FindOne(ctx, bson.M(filter)).Decode(&result)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, ErrNotFound
        }
        return nil, fmt.Errorf("findone %s: %w", collectionName, err)
    }

    return &result, nil
}
```

With SQL (using something like Go's standard database/sql package), the equivalent lookup looks more like this:

```go
func FindOneUser(ctx context.Context, db *sql.DB, id string) (*User, error) {
    query := `SELECT * FROM users WHERE id = $1`

    var u User
    err := db.QueryRowContext(ctx, query, id).Scan(
        &u.ID,
        &u.Username,
        &u.Email,
        &u.CreatedAt,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotFound
        }
        return nil, fmt.Errorf("findoneuser: %w", err)
    }

    return &u, nil
}
```

With SQL, you have to be deliberate on what you're reading/writing into the database. In my implementation, I decided to make generic helper functions for all my database calls. This allowed my `repo` package with tagged structs to take care of the specific calls for user, game data, catalog, etc. If I were to use a SQL implementation, I would've been forced to make a helper method for every record type in my database, increasing the code and decreasing readability. There's nothing wrong on the surface with this approach, but it didn't fit how I wanted to implement my project.

I've had nothing but pleasant experiences with MongoDB. Setting up in Go is super easy, lookup times are great, and I've never had lost data. Even the Mongo CLI tool is super nice to use. My project isn't super heavy in the data side of things, so I haven't really stressed tested Mongo. However, for a small project like this, it is a great, low-effort database to integrate.

## Caddy

I chose Caddy as the reverse proxy in front of the Go binary because:

- Automatic HTTPS. Caddy provisions and renews Let's Encrypt certificates on its own.
- The config is way simpler than an nginx config for the same job.
- It's super lightweight.

Challenges:

- Like Svelte, Caddy has a much smaller ecosystem and plugin base than nginx.
- Caddy needs to reach the origin server directly, so if you use something like Cloudflare's proxy, it can't access the server so it fails the cert.
- Ran into a problem getting the first certificate issued: Let's Encrypt's HTTP-01 challenge needs to reach the origin server directly, and Cloudflare's proxied ("orange-cloud") DNS mode got in the way of that. I had to set the DNS record to DNS-only ("grey-cloud") for the initial issuance before switching back.

## Oracle Cloud

I chose Oracle Cloud Infrastructure (OCI) for hosting because:

- Oracle's free tier is a steal compared to other providers. I run a server with 1 OCPU and 1 gig of ram unlimited for free.
- Two-layer firewall, one at the instance level and one at the shape level. Helps provide extra levels of security.

Challenges:

- I intended on using their ARM instance option (1 OCPU and 6 gigs of ram), but the availability is hard to come by.
- Steeper learning curve around OCI-specific networking concepts (VCNs, subnets, CIDR blocks, Shielded Instance options) compared to simpler providers like DigitalOcean or Heroku.
- A bit trickier to set up. You have to know about networks, subnets, shapes, and securities to properly set up your instance.
- UI is slow and not that user-friendly

I hoped to have a better reason to use this cloud provider, but it really boiled down to their free tier compute. I'm trying to keep this project monetized-free, and that means to find deals where ever I can.

Oracle Cloud is a pretty decent experience to use as long as you have networking fundamentals down. I recommend for anyone looking to save a few bucks a month if they're running a pretty lightweight app.
