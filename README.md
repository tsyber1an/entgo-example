EntGo example
-------------

We will design ER models and generate Ent entities by example. Let's start.
Given we have a Player model.

```go
type Player struct {
}
```

to turn `Player` mode into an `Ent`'s entity, it is just enought embed `ent.Schema`, so it becomes:

```go
type Player struct {
    ent.Schema
}
```

After we have scheme embed, we will need to define our schema parts such as `Fields`, `Relationships` and, if neccessary, some database specifics like `indexes`.
Let's add some fields to our `Player` entity, e.g. `nickname`, `email` and `scores`:

```go
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname"),
		field.String("email"),
		field.Int("scores"),
	}
}
```

Let alone to have `Fields` is good, so next step of using `Ent` lib, would be to generate `Ent` files called `Ent`'s templates. they will be based on our defined schemas. To do so, we need first to install codegen tool (if you have already have it):
```bash
go get entgo.io/ent/cmd/ent
```

after you got a tool, time to generate you first template:
```bash
go run entgo.io/ent/cmd/ent init Player
```

the `init` command will create a folder `ent` which will contain:
```
./ent
    schema/
        player.go <---- your template for Player struct
    generate.go
```

those files should be kept without manual editing and every time you change something in your entities and you got an uprade of Ent, you will invoke `init` to overwrite those files ones more.

Tsyren Ochirov (c) 2021
