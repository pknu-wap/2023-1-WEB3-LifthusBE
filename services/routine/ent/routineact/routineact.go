// Code generated by ent, DO NOT EDIT.

package routineact

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the routineact type in the database.
	Label = "routine_act"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActID holds the string denoting the act_id field in the database.
	FieldActID = "act_id"
	// FieldDailyRoutineID holds the string denoting the daily_routine_id field in the database.
	FieldDailyRoutineID = "daily_routine_id"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldWRatio holds the string denoting the w_ratio field in the database.
	FieldWRatio = "w_ratio"
	// FieldReps holds the string denoting the reps field in the database.
	FieldReps = "reps"
	// FieldLap holds the string denoting the lap field in the database.
	FieldLap = "lap"
	// FieldWarmup holds the string denoting the warmup field in the database.
	FieldWarmup = "warmup"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAct holds the string denoting the act edge name in mutations.
	EdgeAct = "act"
	// EdgeDailyRoutine holds the string denoting the daily_routine edge name in mutations.
	EdgeDailyRoutine = "daily_routine"
	// EdgeRoutineActRecs holds the string denoting the routine_act_recs edge name in mutations.
	EdgeRoutineActRecs = "routine_act_recs"
	// Table holds the table name of the routineact in the database.
	Table = "routine_acts"
	// ActTable is the table that holds the act relation/edge.
	ActTable = "routine_acts"
	// ActInverseTable is the table name for the Act entity.
	// It exists in this package in order to avoid circular dependency with the "act" package.
	ActInverseTable = "acts"
	// ActColumn is the table column denoting the act relation/edge.
	ActColumn = "act_id"
	// DailyRoutineTable is the table that holds the daily_routine relation/edge.
	DailyRoutineTable = "routine_acts"
	// DailyRoutineInverseTable is the table name for the DailyRoutine entity.
	// It exists in this package in order to avoid circular dependency with the "dailyroutine" package.
	DailyRoutineInverseTable = "daily_routines"
	// DailyRoutineColumn is the table column denoting the daily_routine relation/edge.
	DailyRoutineColumn = "daily_routine_id"
	// RoutineActRecsTable is the table that holds the routine_act_recs relation/edge.
	RoutineActRecsTable = "routine_act_recs"
	// RoutineActRecsInverseTable is the table name for the RoutineActRec entity.
	// It exists in this package in order to avoid circular dependency with the "routineactrec" package.
	RoutineActRecsInverseTable = "routine_act_recs"
	// RoutineActRecsColumn is the table column denoting the routine_act_recs relation/edge.
	RoutineActRecsColumn = "routine_act_id"
)

// Columns holds all SQL columns for routineact fields.
var Columns = []string{
	FieldID,
	FieldActID,
	FieldDailyRoutineID,
	FieldOrder,
	FieldWRatio,
	FieldReps,
	FieldLap,
	FieldWarmup,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// OrderValidator is a validator for the "order" field. It is called by the builders before save.
	OrderValidator func(int) error
	// WRatioValidator is a validator for the "w_ratio" field. It is called by the builders before save.
	WRatioValidator func(float64) error
	// RepsValidator is a validator for the "reps" field. It is called by the builders before save.
	RepsValidator func(int) error
	// LapValidator is a validator for the "lap" field. It is called by the builders before save.
	LapValidator func(int) error
	// DefaultWarmup holds the default value on creation for the "warmup" field.
	DefaultWarmup bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the RoutineAct queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByActID orders the results by the act_id field.
func ByActID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActID, opts...).ToFunc()
}

// ByDailyRoutineID orders the results by the daily_routine_id field.
func ByDailyRoutineID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDailyRoutineID, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByWRatio orders the results by the w_ratio field.
func ByWRatio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWRatio, opts...).ToFunc()
}

// ByReps orders the results by the reps field.
func ByReps(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReps, opts...).ToFunc()
}

// ByLap orders the results by the lap field.
func ByLap(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLap, opts...).ToFunc()
}

// ByWarmup orders the results by the warmup field.
func ByWarmup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWarmup, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByActField orders the results by act field.
func ByActField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActStep(), sql.OrderByField(field, opts...))
	}
}

// ByDailyRoutineField orders the results by daily_routine field.
func ByDailyRoutineField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDailyRoutineStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoutineActRecsCount orders the results by routine_act_recs count.
func ByRoutineActRecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoutineActRecsStep(), opts...)
	}
}

// ByRoutineActRecs orders the results by routine_act_recs terms.
func ByRoutineActRecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoutineActRecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newActStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ActTable, ActColumn),
	)
}
func newDailyRoutineStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DailyRoutineInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DailyRoutineTable, DailyRoutineColumn),
	)
}
func newRoutineActRecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineActRecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutineActRecsTable, RoutineActRecsColumn),
	)
}
