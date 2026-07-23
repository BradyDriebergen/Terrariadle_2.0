# Backend Process Flow

![New dependency diagram](../../assets/Dependency-Diagram.png)

- **`db`** - generic database read/write functionality.
- **`repo`** - app-specific DB methods built on top of `db`.
- **`store`** - cache with read/write functionality built on `repo`, serves all game data.
- **`services`** - game functionality built on `store`.
- **`api`** - HTTP layer, consumes `services`, embeds the `web` frontend.
- **`web`** - embeds static files into the binary.
- **`domain`** - shared types and interfaces, no dependencies of its own.
- **`jobs`** - routine methods called at specific time intervals.

---

_For more details on the creation of this process flow, see [managing go packages](../../managing-go-packages.md)._

_For more info on specific package functionality, see [packages](packages.md)._
