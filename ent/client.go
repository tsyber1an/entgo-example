// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Funfun/entgo-example/ent/migrate"

	"github.com/Funfun/entgo-example/ent/book"
	"github.com/Funfun/entgo-example/ent/person"
	"github.com/Funfun/entgo-example/ent/player"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Book is the client for interacting with the Book builders.
	Book *BookClient
	// Person is the client for interacting with the Person builders.
	Person *PersonClient
	// Player is the client for interacting with the Player builders.
	Player *PlayerClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Book = NewBookClient(c.config)
	c.Person = NewPersonClient(c.config)
	c.Player = NewPlayerClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Book:   NewBookClient(cfg),
		Person: NewPersonClient(cfg),
		Player: NewPlayerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Book:   NewBookClient(cfg),
		Person: NewPersonClient(cfg),
		Player: NewPlayerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Book.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Book.Use(hooks...)
	c.Person.Use(hooks...)
	c.Player.Use(hooks...)
}

// BookClient is a client for the Book schema.
type BookClient struct {
	config
}

// NewBookClient returns a client for the Book from the given config.
func NewBookClient(c config) *BookClient {
	return &BookClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `book.Hooks(f(g(h())))`.
func (c *BookClient) Use(hooks ...Hook) {
	c.hooks.Book = append(c.hooks.Book, hooks...)
}

// Create returns a create builder for Book.
func (c *BookClient) Create() *BookCreate {
	mutation := newBookMutation(c.config, OpCreate)
	return &BookCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Book entities.
func (c *BookClient) CreateBulk(builders ...*BookCreate) *BookCreateBulk {
	return &BookCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Book.
func (c *BookClient) Update() *BookUpdate {
	mutation := newBookMutation(c.config, OpUpdate)
	return &BookUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookClient) UpdateOne(b *Book) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBook(b))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookClient) UpdateOneID(id int) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBookID(id))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Book.
func (c *BookClient) Delete() *BookDelete {
	mutation := newBookMutation(c.config, OpDelete)
	return &BookDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BookClient) DeleteOne(b *Book) *BookDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BookClient) DeleteOneID(id int) *BookDeleteOne {
	builder := c.Delete().Where(book.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookDeleteOne{builder}
}

// Query returns a query builder for Book.
func (c *BookClient) Query() *BookQuery {
	return &BookQuery{
		config: c.config,
	}
}

// Get returns a Book entity by its id.
func (c *BookClient) Get(ctx context.Context, id int) (*Book, error) {
	return c.Query().Where(book.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookClient) GetX(ctx context.Context, id int) *Book {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Book.
func (c *BookClient) QueryAuthor(b *Book) *PersonQuery {
	query := &PersonQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.AuthorTable, book.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookClient) Hooks() []Hook {
	return c.hooks.Book
}

// PersonClient is a client for the Person schema.
type PersonClient struct {
	config
}

// NewPersonClient returns a client for the Person from the given config.
func NewPersonClient(c config) *PersonClient {
	return &PersonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `person.Hooks(f(g(h())))`.
func (c *PersonClient) Use(hooks ...Hook) {
	c.hooks.Person = append(c.hooks.Person, hooks...)
}

// Create returns a create builder for Person.
func (c *PersonClient) Create() *PersonCreate {
	mutation := newPersonMutation(c.config, OpCreate)
	return &PersonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Person entities.
func (c *PersonClient) CreateBulk(builders ...*PersonCreate) *PersonCreateBulk {
	return &PersonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Person.
func (c *PersonClient) Update() *PersonUpdate {
	mutation := newPersonMutation(c.config, OpUpdate)
	return &PersonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PersonClient) UpdateOne(pe *Person) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPerson(pe))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PersonClient) UpdateOneID(id int) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPersonID(id))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Person.
func (c *PersonClient) Delete() *PersonDelete {
	mutation := newPersonMutation(c.config, OpDelete)
	return &PersonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PersonClient) DeleteOne(pe *Person) *PersonDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PersonClient) DeleteOneID(id int) *PersonDeleteOne {
	builder := c.Delete().Where(person.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PersonDeleteOne{builder}
}

// Query returns a query builder for Person.
func (c *PersonClient) Query() *PersonQuery {
	return &PersonQuery{
		config: c.config,
	}
}

// Get returns a Person entity by its id.
func (c *PersonClient) Get(ctx context.Context, id int) (*Person, error) {
	return c.Query().Where(person.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PersonClient) GetX(ctx context.Context, id int) *Person {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBooks queries the books edge of a Person.
func (c *PersonClient) QueryBooks(pe *Person) *BookQuery {
	query := &BookQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, person.BooksTable, person.BooksColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PersonClient) Hooks() []Hook {
	return c.hooks.Person
}

// PlayerClient is a client for the Player schema.
type PlayerClient struct {
	config
}

// NewPlayerClient returns a client for the Player from the given config.
func NewPlayerClient(c config) *PlayerClient {
	return &PlayerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `player.Hooks(f(g(h())))`.
func (c *PlayerClient) Use(hooks ...Hook) {
	c.hooks.Player = append(c.hooks.Player, hooks...)
}

// Create returns a create builder for Player.
func (c *PlayerClient) Create() *PlayerCreate {
	mutation := newPlayerMutation(c.config, OpCreate)
	return &PlayerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Player entities.
func (c *PlayerClient) CreateBulk(builders ...*PlayerCreate) *PlayerCreateBulk {
	return &PlayerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Player.
func (c *PlayerClient) Update() *PlayerUpdate {
	mutation := newPlayerMutation(c.config, OpUpdate)
	return &PlayerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlayerClient) UpdateOne(pl *Player) *PlayerUpdateOne {
	mutation := newPlayerMutation(c.config, OpUpdateOne, withPlayer(pl))
	return &PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlayerClient) UpdateOneID(id int) *PlayerUpdateOne {
	mutation := newPlayerMutation(c.config, OpUpdateOne, withPlayerID(id))
	return &PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Player.
func (c *PlayerClient) Delete() *PlayerDelete {
	mutation := newPlayerMutation(c.config, OpDelete)
	return &PlayerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlayerClient) DeleteOne(pl *Player) *PlayerDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlayerClient) DeleteOneID(id int) *PlayerDeleteOne {
	builder := c.Delete().Where(player.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlayerDeleteOne{builder}
}

// Query returns a query builder for Player.
func (c *PlayerClient) Query() *PlayerQuery {
	return &PlayerQuery{
		config: c.config,
	}
}

// Get returns a Player entity by its id.
func (c *PlayerClient) Get(ctx context.Context, id int) (*Player, error) {
	return c.Query().Where(player.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlayerClient) GetX(ctx context.Context, id int) *Player {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PlayerClient) Hooks() []Hook {
	return c.hooks.Player
}
