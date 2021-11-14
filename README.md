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

In official introduction of Entgo, it says to init model by running codegeltool. Go ahead and install it first:

```bash
go get entgo.io/ent/cmd/ent
```

After you got a tool, time to generate you first template:
```bash
go run entgo.io/ent/cmd/ent init Player
```

The `init` command will create a folder `ent` which will contain:
```
./ent
    schema/
        player.go <---- your template for Player struct
    generate.go
```

You will find it has the same structure as we describe in the beginning, plus an aadditional method:
```
func (Player) Edges() []ent.Edge {
	return nil
}
```

Now, you can copy your `Field()` method into `./ent/schema/player.go`. To verify that you did all right, run `describe` command, you should get something like bellow:
```
➜  entgo-example git:(main) go run entgo.io/ent/cmd/ent describe ./ent/schema
Player:
        +----------+--------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+
        |  Field   |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |         StructTag         | Validators |
        +----------+--------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+
        | id       | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"       |          0 |
        | nickname | string | false  | false    | false    | false   | false         | false     | json:"nickname,omitempty" |          0 |
        | email    | string | false  | false    | false    | false   | false         | false     | json:"email,omitempty"    |          0 |
        | scores   | int    | false  | false    | false    | false   | false         | false     | json:"scores,omitempty"   |          0 |
```

We can remove `./player.go` file, as it is not gonna be needed anymore. From now on, you will edit your models in `./ent/schema` folder.

Now, you migth wonder how do we call SQL queries. To do that, you would need to generate your SQL interfaces for your schema. Invoke the following:

Full version:
```bash
go run entgo.io/ent/cmd/ent generate ./ent/schema
```

Short version:
```
go generate ./ent
```

Inspect results in `./ent`. Previously, we had there our schame file. Now, there are bunch of files, do not edit them yet. Let's understand their purpose by using it.
let's pick our `Postgres` as our RDMS engine. Steps are:

Spin up `Postgres` server in docker:
```bash
docker run -it --rm --name dev-postgres -e POSTGRES_DB=entgo_example -e POSTGRES_PASSWORD=topsecret -p 5432:5432 postgres
```

our connection setttings:
```
Host: localhost
Port: 5432
User: postgres
Password: topsecret
Database: entgo_example
```

Code to connect to our db:
```go
sql.Open("postgres", "user=postgres password=topsecret dbname=entgo_example sslmode=disable")
```


Tsyren Ochirov (c) 2021
