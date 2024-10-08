// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"demo/ent/migrate"

	"demo/ent/course"
	"demo/ent/department"
	"demo/ent/enrollment"
	"demo/ent/instructor"
	"demo/ent/student"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Course is the client for interacting with the Course builders.
	Course *CourseClient
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// Enrollment is the client for interacting with the Enrollment builders.
	Enrollment *EnrollmentClient
	// Instructor is the client for interacting with the Instructor builders.
	Instructor *InstructorClient
	// Student is the client for interacting with the Student builders.
	Student *StudentClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Course = NewCourseClient(c.config)
	c.Department = NewDepartmentClient(c.config)
	c.Enrollment = NewEnrollmentClient(c.config)
	c.Instructor = NewInstructorClient(c.config)
	c.Student = NewStudentClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Course:     NewCourseClient(cfg),
		Department: NewDepartmentClient(cfg),
		Enrollment: NewEnrollmentClient(cfg),
		Instructor: NewInstructorClient(cfg),
		Student:    NewStudentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:        ctx,
		config:     cfg,
		Course:     NewCourseClient(cfg),
		Department: NewDepartmentClient(cfg),
		Enrollment: NewEnrollmentClient(cfg),
		Instructor: NewInstructorClient(cfg),
		Student:    NewStudentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Course.
//		Query().
//		Count(ctx)
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
	c.Course.Use(hooks...)
	c.Department.Use(hooks...)
	c.Enrollment.Use(hooks...)
	c.Instructor.Use(hooks...)
	c.Student.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Course.Intercept(interceptors...)
	c.Department.Intercept(interceptors...)
	c.Enrollment.Intercept(interceptors...)
	c.Instructor.Intercept(interceptors...)
	c.Student.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CourseMutation:
		return c.Course.mutate(ctx, m)
	case *DepartmentMutation:
		return c.Department.mutate(ctx, m)
	case *EnrollmentMutation:
		return c.Enrollment.mutate(ctx, m)
	case *InstructorMutation:
		return c.Instructor.mutate(ctx, m)
	case *StudentMutation:
		return c.Student.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CourseClient is a client for the Course schema.
type CourseClient struct {
	config
}

// NewCourseClient returns a client for the Course from the given config.
func NewCourseClient(c config) *CourseClient {
	return &CourseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `course.Hooks(f(g(h())))`.
func (c *CourseClient) Use(hooks ...Hook) {
	c.hooks.Course = append(c.hooks.Course, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `course.Intercept(f(g(h())))`.
func (c *CourseClient) Intercept(interceptors ...Interceptor) {
	c.inters.Course = append(c.inters.Course, interceptors...)
}

// Create returns a builder for creating a Course entity.
func (c *CourseClient) Create() *CourseCreate {
	mutation := newCourseMutation(c.config, OpCreate)
	return &CourseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Course entities.
func (c *CourseClient) CreateBulk(builders ...*CourseCreate) *CourseCreateBulk {
	return &CourseCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CourseClient) MapCreateBulk(slice any, setFunc func(*CourseCreate, int)) *CourseCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CourseCreateBulk{err: fmt.Errorf("calling to CourseClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CourseCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CourseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Course.
func (c *CourseClient) Update() *CourseUpdate {
	mutation := newCourseMutation(c.config, OpUpdate)
	return &CourseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CourseClient) UpdateOne(co *Course) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourse(co))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CourseClient) UpdateOneID(id int) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourseID(id))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Course.
func (c *CourseClient) Delete() *CourseDelete {
	mutation := newCourseMutation(c.config, OpDelete)
	return &CourseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CourseClient) DeleteOne(co *Course) *CourseDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CourseClient) DeleteOneID(id int) *CourseDeleteOne {
	builder := c.Delete().Where(course.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CourseDeleteOne{builder}
}

// Query returns a query builder for Course.
func (c *CourseClient) Query() *CourseQuery {
	return &CourseQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCourse},
		inters: c.Interceptors(),
	}
}

// Get returns a Course entity by its id.
func (c *CourseClient) Get(ctx context.Context, id int) (*Course, error) {
	return c.Query().Where(course.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CourseClient) GetX(ctx context.Context, id int) *Course {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDepartment queries the department edge of a Course.
func (c *CourseClient) QueryDepartment(co *Course) *DepartmentQuery {
	query := (&DepartmentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, course.DepartmentTable, course.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEnrollments queries the enrollments edge of a Course.
func (c *CourseClient) QueryEnrollments(co *Course) *EnrollmentQuery {
	query := (&EnrollmentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(course.Table, course.FieldID, id),
			sqlgraph.To(enrollment.Table, enrollment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, course.EnrollmentsTable, course.EnrollmentsColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CourseClient) Hooks() []Hook {
	return c.hooks.Course
}

// Interceptors returns the client interceptors.
func (c *CourseClient) Interceptors() []Interceptor {
	return c.inters.Course
}

func (c *CourseClient) mutate(ctx context.Context, m *CourseMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CourseCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CourseUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CourseDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Course mutation op: %q", m.Op())
	}
}

// DepartmentClient is a client for the Department schema.
type DepartmentClient struct {
	config
}

// NewDepartmentClient returns a client for the Department from the given config.
func NewDepartmentClient(c config) *DepartmentClient {
	return &DepartmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `department.Hooks(f(g(h())))`.
func (c *DepartmentClient) Use(hooks ...Hook) {
	c.hooks.Department = append(c.hooks.Department, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `department.Intercept(f(g(h())))`.
func (c *DepartmentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Department = append(c.inters.Department, interceptors...)
}

// Create returns a builder for creating a Department entity.
func (c *DepartmentClient) Create() *DepartmentCreate {
	mutation := newDepartmentMutation(c.config, OpCreate)
	return &DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Department entities.
func (c *DepartmentClient) CreateBulk(builders ...*DepartmentCreate) *DepartmentCreateBulk {
	return &DepartmentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DepartmentClient) MapCreateBulk(slice any, setFunc func(*DepartmentCreate, int)) *DepartmentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DepartmentCreateBulk{err: fmt.Errorf("calling to DepartmentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DepartmentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DepartmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Department.
func (c *DepartmentClient) Update() *DepartmentUpdate {
	mutation := newDepartmentMutation(c.config, OpUpdate)
	return &DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepartmentClient) UpdateOne(d *Department) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartment(d))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepartmentClient) UpdateOneID(id int) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartmentID(id))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Department.
func (c *DepartmentClient) Delete() *DepartmentDelete {
	mutation := newDepartmentMutation(c.config, OpDelete)
	return &DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DepartmentClient) DeleteOne(d *Department) *DepartmentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DepartmentClient) DeleteOneID(id int) *DepartmentDeleteOne {
	builder := c.Delete().Where(department.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepartmentDeleteOne{builder}
}

// Query returns a query builder for Department.
func (c *DepartmentClient) Query() *DepartmentQuery {
	return &DepartmentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDepartment},
		inters: c.Interceptors(),
	}
}

// Get returns a Department entity by its id.
func (c *DepartmentClient) Get(ctx context.Context, id int) (*Department, error) {
	return c.Query().Where(department.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepartmentClient) GetX(ctx context.Context, id int) *Department {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstructors queries the instructors edge of a Department.
func (c *DepartmentClient) QueryInstructors(d *Department) *InstructorQuery {
	query := (&InstructorClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(instructor.Table, instructor.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.InstructorsTable, department.InstructorsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCourses queries the courses edge of a Department.
func (c *DepartmentClient) QueryCourses(d *Department) *CourseQuery {
	query := (&CourseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.CoursesTable, department.CoursesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStudents queries the students edge of a Department.
func (c *DepartmentClient) QueryStudents(d *Department) *StudentQuery {
	query := (&StudentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.StudentsTable, department.StudentsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DepartmentClient) Hooks() []Hook {
	return c.hooks.Department
}

// Interceptors returns the client interceptors.
func (c *DepartmentClient) Interceptors() []Interceptor {
	return c.inters.Department
}

func (c *DepartmentClient) mutate(ctx context.Context, m *DepartmentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Department mutation op: %q", m.Op())
	}
}

// EnrollmentClient is a client for the Enrollment schema.
type EnrollmentClient struct {
	config
}

// NewEnrollmentClient returns a client for the Enrollment from the given config.
func NewEnrollmentClient(c config) *EnrollmentClient {
	return &EnrollmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `enrollment.Hooks(f(g(h())))`.
func (c *EnrollmentClient) Use(hooks ...Hook) {
	c.hooks.Enrollment = append(c.hooks.Enrollment, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `enrollment.Intercept(f(g(h())))`.
func (c *EnrollmentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Enrollment = append(c.inters.Enrollment, interceptors...)
}

// Create returns a builder for creating a Enrollment entity.
func (c *EnrollmentClient) Create() *EnrollmentCreate {
	mutation := newEnrollmentMutation(c.config, OpCreate)
	return &EnrollmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Enrollment entities.
func (c *EnrollmentClient) CreateBulk(builders ...*EnrollmentCreate) *EnrollmentCreateBulk {
	return &EnrollmentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EnrollmentClient) MapCreateBulk(slice any, setFunc func(*EnrollmentCreate, int)) *EnrollmentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EnrollmentCreateBulk{err: fmt.Errorf("calling to EnrollmentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EnrollmentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EnrollmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Enrollment.
func (c *EnrollmentClient) Update() *EnrollmentUpdate {
	mutation := newEnrollmentMutation(c.config, OpUpdate)
	return &EnrollmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EnrollmentClient) UpdateOne(e *Enrollment) *EnrollmentUpdateOne {
	mutation := newEnrollmentMutation(c.config, OpUpdateOne, withEnrollment(e))
	return &EnrollmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EnrollmentClient) UpdateOneID(id int) *EnrollmentUpdateOne {
	mutation := newEnrollmentMutation(c.config, OpUpdateOne, withEnrollmentID(id))
	return &EnrollmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Enrollment.
func (c *EnrollmentClient) Delete() *EnrollmentDelete {
	mutation := newEnrollmentMutation(c.config, OpDelete)
	return &EnrollmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EnrollmentClient) DeleteOne(e *Enrollment) *EnrollmentDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EnrollmentClient) DeleteOneID(id int) *EnrollmentDeleteOne {
	builder := c.Delete().Where(enrollment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EnrollmentDeleteOne{builder}
}

// Query returns a query builder for Enrollment.
func (c *EnrollmentClient) Query() *EnrollmentQuery {
	return &EnrollmentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEnrollment},
		inters: c.Interceptors(),
	}
}

// Get returns a Enrollment entity by its id.
func (c *EnrollmentClient) Get(ctx context.Context, id int) (*Enrollment, error) {
	return c.Query().Where(enrollment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EnrollmentClient) GetX(ctx context.Context, id int) *Enrollment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStudent queries the student edge of a Enrollment.
func (c *EnrollmentClient) QueryStudent(e *Enrollment) *StudentQuery {
	query := (&StudentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(enrollment.Table, enrollment.FieldID, id),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, enrollment.StudentTable, enrollment.StudentColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCourse queries the course edge of a Enrollment.
func (c *EnrollmentClient) QueryCourse(e *Enrollment) *CourseQuery {
	query := (&CourseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(enrollment.Table, enrollment.FieldID, id),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, enrollment.CourseTable, enrollment.CourseColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EnrollmentClient) Hooks() []Hook {
	return c.hooks.Enrollment
}

// Interceptors returns the client interceptors.
func (c *EnrollmentClient) Interceptors() []Interceptor {
	return c.inters.Enrollment
}

func (c *EnrollmentClient) mutate(ctx context.Context, m *EnrollmentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EnrollmentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EnrollmentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EnrollmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EnrollmentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Enrollment mutation op: %q", m.Op())
	}
}

// InstructorClient is a client for the Instructor schema.
type InstructorClient struct {
	config
}

// NewInstructorClient returns a client for the Instructor from the given config.
func NewInstructorClient(c config) *InstructorClient {
	return &InstructorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instructor.Hooks(f(g(h())))`.
func (c *InstructorClient) Use(hooks ...Hook) {
	c.hooks.Instructor = append(c.hooks.Instructor, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `instructor.Intercept(f(g(h())))`.
func (c *InstructorClient) Intercept(interceptors ...Interceptor) {
	c.inters.Instructor = append(c.inters.Instructor, interceptors...)
}

// Create returns a builder for creating a Instructor entity.
func (c *InstructorClient) Create() *InstructorCreate {
	mutation := newInstructorMutation(c.config, OpCreate)
	return &InstructorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Instructor entities.
func (c *InstructorClient) CreateBulk(builders ...*InstructorCreate) *InstructorCreateBulk {
	return &InstructorCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *InstructorClient) MapCreateBulk(slice any, setFunc func(*InstructorCreate, int)) *InstructorCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &InstructorCreateBulk{err: fmt.Errorf("calling to InstructorClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*InstructorCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &InstructorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Instructor.
func (c *InstructorClient) Update() *InstructorUpdate {
	mutation := newInstructorMutation(c.config, OpUpdate)
	return &InstructorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstructorClient) UpdateOne(i *Instructor) *InstructorUpdateOne {
	mutation := newInstructorMutation(c.config, OpUpdateOne, withInstructor(i))
	return &InstructorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstructorClient) UpdateOneID(id int) *InstructorUpdateOne {
	mutation := newInstructorMutation(c.config, OpUpdateOne, withInstructorID(id))
	return &InstructorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Instructor.
func (c *InstructorClient) Delete() *InstructorDelete {
	mutation := newInstructorMutation(c.config, OpDelete)
	return &InstructorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *InstructorClient) DeleteOne(i *Instructor) *InstructorDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *InstructorClient) DeleteOneID(id int) *InstructorDeleteOne {
	builder := c.Delete().Where(instructor.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstructorDeleteOne{builder}
}

// Query returns a query builder for Instructor.
func (c *InstructorClient) Query() *InstructorQuery {
	return &InstructorQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeInstructor},
		inters: c.Interceptors(),
	}
}

// Get returns a Instructor entity by its id.
func (c *InstructorClient) Get(ctx context.Context, id int) (*Instructor, error) {
	return c.Query().Where(instructor.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstructorClient) GetX(ctx context.Context, id int) *Instructor {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDepartment queries the department edge of a Instructor.
func (c *InstructorClient) QueryDepartment(i *Instructor) *DepartmentQuery {
	query := (&DepartmentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instructor.Table, instructor.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instructor.DepartmentTable, instructor.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstructorClient) Hooks() []Hook {
	return c.hooks.Instructor
}

// Interceptors returns the client interceptors.
func (c *InstructorClient) Interceptors() []Interceptor {
	return c.inters.Instructor
}

func (c *InstructorClient) mutate(ctx context.Context, m *InstructorMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&InstructorCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&InstructorUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&InstructorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&InstructorDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Instructor mutation op: %q", m.Op())
	}
}

// StudentClient is a client for the Student schema.
type StudentClient struct {
	config
}

// NewStudentClient returns a client for the Student from the given config.
func NewStudentClient(c config) *StudentClient {
	return &StudentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `student.Hooks(f(g(h())))`.
func (c *StudentClient) Use(hooks ...Hook) {
	c.hooks.Student = append(c.hooks.Student, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `student.Intercept(f(g(h())))`.
func (c *StudentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Student = append(c.inters.Student, interceptors...)
}

// Create returns a builder for creating a Student entity.
func (c *StudentClient) Create() *StudentCreate {
	mutation := newStudentMutation(c.config, OpCreate)
	return &StudentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Student entities.
func (c *StudentClient) CreateBulk(builders ...*StudentCreate) *StudentCreateBulk {
	return &StudentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StudentClient) MapCreateBulk(slice any, setFunc func(*StudentCreate, int)) *StudentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StudentCreateBulk{err: fmt.Errorf("calling to StudentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StudentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StudentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Student.
func (c *StudentClient) Update() *StudentUpdate {
	mutation := newStudentMutation(c.config, OpUpdate)
	return &StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudentClient) UpdateOne(s *Student) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudent(s))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudentClient) UpdateOneID(id int) *StudentUpdateOne {
	mutation := newStudentMutation(c.config, OpUpdateOne, withStudentID(id))
	return &StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Student.
func (c *StudentClient) Delete() *StudentDelete {
	mutation := newStudentMutation(c.config, OpDelete)
	return &StudentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StudentClient) DeleteOne(s *Student) *StudentDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StudentClient) DeleteOneID(id int) *StudentDeleteOne {
	builder := c.Delete().Where(student.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudentDeleteOne{builder}
}

// Query returns a query builder for Student.
func (c *StudentClient) Query() *StudentQuery {
	return &StudentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStudent},
		inters: c.Interceptors(),
	}
}

// Get returns a Student entity by its id.
func (c *StudentClient) Get(ctx context.Context, id int) (*Student, error) {
	return c.Query().Where(student.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudentClient) GetX(ctx context.Context, id int) *Student {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDepartment queries the department edge of a Student.
func (c *StudentClient) QueryDepartment(s *Student) *DepartmentQuery {
	query := (&DepartmentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(student.Table, student.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, student.DepartmentTable, student.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEnrollments queries the enrollments edge of a Student.
func (c *StudentClient) QueryEnrollments(s *Student) *EnrollmentQuery {
	query := (&EnrollmentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(student.Table, student.FieldID, id),
			sqlgraph.To(enrollment.Table, enrollment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, student.EnrollmentsTable, student.EnrollmentsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StudentClient) Hooks() []Hook {
	return c.hooks.Student
}

// Interceptors returns the client interceptors.
func (c *StudentClient) Interceptors() []Interceptor {
	return c.inters.Student
}

func (c *StudentClient) mutate(ctx context.Context, m *StudentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StudentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StudentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StudentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Student mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Course, Department, Enrollment, Instructor, Student []ent.Hook
	}
	inters struct {
		Course, Department, Enrollment, Instructor, Student []ent.Interceptor
	}
)
