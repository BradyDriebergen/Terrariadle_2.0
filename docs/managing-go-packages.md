# Learning How to Design Systems

One of the biggest learning curves I faced was how to build and structure my packages in Go. In school, I never dealt with structuring my school projects. Most of the time, I would make my Java projects all within one folder. This never became a problem until I started dealing with a project this big.

Go has a very interesting way of allowing developers to organize their programs. There is a lot of freedom when it comes to how a Go project can be structured. It doesn't follow a strict file structure format similar to SvelteKit or React. Instead, Go projects consist of packages, often organized under conventions like an `/internal` folder for code that shouldn't be imported outside the module.

I wanted to write this document to show my learning process, going from someone who didn't know anything about structuring a project, to a developer that understands the importance of dependency direction and project structure.

## The Starting Point

Early on in developing this second iteration of my project, the backend was organized very poorly. I based logic loosely around grouped functionality. I didn't have any idea how packages should depend on one another. This caused my packages to be overly coupled without clear boundaries of functionality. I actually mapped out my old dependency diagram:

![Old dependency diagram](assets/Old-Dependency-Diagram.png)

You can also see how I originally had my project structured in [v2 File Structure](./legacy-notes/v2-file-structure.txt).

As you can see, my old structure was a web of imports between all of my packages. After finishing this draft, I started noticing a lot of issues with this approach.

- **No clear ownership of methods**: This structure had no clear ownership of functions or types. There were some packages like `jsonreader` and `store` that did very similar functionalities. I also had a trickle-down effect where I would declare something in one package, and use a package to pass it down to another package. Just from the diagram, you can see how messy it got.
- **Packages were dumping grounds**: the `server` package held handlers, services, and routing at the same time. `models` also defined the game data types and stored the game data from `db`. I basically had 2 stores, the `store` package and the `models` package. Finding where methods were sucked.
- **Nearly impossible to debug**: This was an issue with a lack of meaningful error messages, but trying to find where the code broke was impossible. I would try to follow paths and reach dead ends everywhere.
- **Also impossible to add features**: Like the issue with debugging, adding features sucked. Looking at the packages, I'd ask myself "where would I add this extra service function". Starting to add more to my backend quickly showed the flaws of this structure.
- **No obvious flow**: If you were to look at this and tell me how to start from the database and serve a client an HTTP response, you wouldn't be able to. This complex structure made it nearly impossible to understand the end-to-end flow of the packages.

## Working Through It

Refactoring this mess didn't happen immediately. I actually went through a few iterations before I really improved my design. Here are the main points that helped me when researching into how to design systems.

### Planning out the flow of the project

Before writing the code, the first step was planning out how the end-to-end flow should work. I started with the basics, "what do I want to get from my database and give to a user". I then broke it down into smaller parts. Such as "I'll need an HTTP server to serve my backend route," and "I need a cache, so I don't hammer my database so often." During planning, I kept the following things in mind:

- **Picture your program as a river**: This might not apply to massive projects, but I learned to picture my program as a flowing river. Each layer of the backend has its own responsibilities and should flow data from one package to another. Sometimes rivers converge, sometimes they break away. Your projects should have a clear path with a start and an end.
- **No circular imports**: Packages shouldn't pass data back and forth. This creates issues with packages depending on each other, so if either one changes, then you would need to change the other.
- **A package should have one clear responsibility**: As I stated above, there are multiple layers to programs. It's important to keep these layers separate because if you ever need to change anything, you know exactly where to look. Sometimes, these layers are ambiguous, but it's up to your best judgement to decide how to separate out your project.

The main point of this thinking is to maximize the simplicity of a project. I wanted to be able to have clear separations of functionality. I could then easily come back to the project and quickly understand what I need to change and where.

### Splitting the layers apart

Keeping the previous things in mind, I had to work on how to actually separate my layers. After trying out many iterations and doing more research into how my backend should work, I came up with the following:

- **Database Layer** - Responsible for getting and updating data in the database.
- **Store Layer** - Responsible for storing data from the database in memory
- **Jobs Layer** - Responsible for running background jobs like resetting the puzzle
- **Services Layer** - Responsible for using the data from the store and the client
- **Server Layer** - Responsible for serving data to the client

Of course, this wouldn't be my actual packages, just a list of the main tasks my backend had to accomplish. These were ended up being broken down more into smaller packages. Detail info on these packages can be found in the [packages](architecture/backend/packages.md) document.

### 3. Handling shared functionality

Unfortunately, it's common to have shared functionality throughout the project. Whether it's a `time` util method or a common type used throughout the app. While researching Go projects, I found that the `domain` package was built for this. This package holds common types and methods that can be used throughout your code. However, this should be used sparingly and not as a dumping ground for types and methods.

## The Result

![New dependency diagram](assets/Dependency-Diagram.png)

- **`db`** — generic database read/write functionality.
- **`repo`** — app-specific DB methods built on top of `db`.
- **`store`** — cache with read/write functionality built on `repo`, serves all game data.
- **`services`** — game functionality built on `store`.
- **`api`** — HTTP layer, consumes `services`, embeds the `web` frontend.
- **`domain`** — shared types and interfaces, no dependencies of its own.

Compared to the old diagram, every arrow now points in one direction. There is a clear flow and boundaries of responsibilities. Each package's role can be described in a single sentence.

Not only does simplify the flow, but I've found that adding new features is wicked fast now! I was able to build out the endpoints for the new game TerraTriva in a day! Having a structure like this makes developing, debugging, and maintaining SO much easier.

## Takeaway

This might've seemed obvious to other developers, but this was a revolutionary mindset shift for me. Like I said before, I've never built an application this big before. Comparing my first iteration to my new one really opened my eyes to what system design can do for a project. This dependency structure probably isn't perfect, but it's something I'm extremely proud of. Go structure was hard for me at first, and if you're dealing with the same issues, I hope I was able to help.
