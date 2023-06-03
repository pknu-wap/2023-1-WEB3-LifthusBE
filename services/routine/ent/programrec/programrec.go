// Code generated by ent, DO NOT EDIT.

package programrec

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the programrec type in the database.
	Label = "program_rec"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldProgramID holds the string denoting the program_id field in the database.
	FieldProgramID = "program_id"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProgram holds the string denoting the program edge name in mutations.
	EdgeProgram = "program"
	// EdgeWeeklyRoutineRecs holds the string denoting the weekly_routine_recs edge name in mutations.
	EdgeWeeklyRoutineRecs = "weekly_routine_recs"
	// EdgeDailyRoutineRecs holds the string denoting the daily_routine_recs edge name in mutations.
	EdgeDailyRoutineRecs = "daily_routine_recs"
	// EdgeBodyInfo holds the string denoting the body_info edge name in mutations.
	EdgeBodyInfo = "body_info"
	// EdgeOneRepMax holds the string denoting the one_rep_max edge name in mutations.
	EdgeOneRepMax = "one_rep_max"
	// Table holds the table name of the programrec in the database.
	Table = "program_recs"
	// ProgramTable is the table that holds the program relation/edge.
	ProgramTable = "program_recs"
	// ProgramInverseTable is the table name for the Program entity.
	// It exists in this package in order to avoid circular dependency with the "program" package.
	ProgramInverseTable = "programs"
	// ProgramColumn is the table column denoting the program relation/edge.
	ProgramColumn = "program_id"
	// WeeklyRoutineRecsTable is the table that holds the weekly_routine_recs relation/edge.
	WeeklyRoutineRecsTable = "weekly_routine_recs"
	// WeeklyRoutineRecsInverseTable is the table name for the WeeklyRoutineRec entity.
	// It exists in this package in order to avoid circular dependency with the "weeklyroutinerec" package.
	WeeklyRoutineRecsInverseTable = "weekly_routine_recs"
	// WeeklyRoutineRecsColumn is the table column denoting the weekly_routine_recs relation/edge.
	WeeklyRoutineRecsColumn = "program_rec_id"
	// DailyRoutineRecsTable is the table that holds the daily_routine_recs relation/edge.
	DailyRoutineRecsTable = "daily_routine_recs"
	// DailyRoutineRecsInverseTable is the table name for the DailyRoutineRec entity.
	// It exists in this package in order to avoid circular dependency with the "dailyroutinerec" package.
	DailyRoutineRecsInverseTable = "daily_routine_recs"
	// DailyRoutineRecsColumn is the table column denoting the daily_routine_recs relation/edge.
	DailyRoutineRecsColumn = "program_rec_id"
	// BodyInfoTable is the table that holds the body_info relation/edge.
	BodyInfoTable = "body_infos"
	// BodyInfoInverseTable is the table name for the BodyInfo entity.
	// It exists in this package in order to avoid circular dependency with the "bodyinfo" package.
	BodyInfoInverseTable = "body_infos"
	// BodyInfoColumn is the table column denoting the body_info relation/edge.
	BodyInfoColumn = "program_rec_id"
	// OneRepMaxTable is the table that holds the one_rep_max relation/edge.
	OneRepMaxTable = "one_rep_maxes"
	// OneRepMaxInverseTable is the table name for the OneRepMax entity.
	// It exists in this package in order to avoid circular dependency with the "onerepmax" package.
	OneRepMaxInverseTable = "one_rep_maxes"
	// OneRepMaxColumn is the table column denoting the one_rep_max relation/edge.
	OneRepMaxColumn = "program_rec_id"
)

// Columns holds all SQL columns for programrec fields.
var Columns = []string{
	FieldID,
	FieldAuthor,
	FieldProgramID,
	FieldStartDate,
	FieldEndDate,
	FieldStatus,
	FieldComment,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusHistory    Status = "history"
	StatusWaiting    Status = "waiting"
	StatusProceeding Status = "proceeding"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
	StatusCanceled   Status = "canceled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusHistory, StatusWaiting, StatusProceeding, StatusCompleted, StatusFailed, StatusCanceled:
		return nil
	default:
		return fmt.Errorf("programrec: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the ProgramRec queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByProgramID orders the results by the program_id field.
func ByProgramID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgramID, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProgramField orders the results by program field.
func ByProgramField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramStep(), sql.OrderByField(field, opts...))
	}
}

// ByWeeklyRoutineRecsCount orders the results by weekly_routine_recs count.
func ByWeeklyRoutineRecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWeeklyRoutineRecsStep(), opts...)
	}
}

// ByWeeklyRoutineRecs orders the results by weekly_routine_recs terms.
func ByWeeklyRoutineRecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWeeklyRoutineRecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDailyRoutineRecsCount orders the results by daily_routine_recs count.
func ByDailyRoutineRecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDailyRoutineRecsStep(), opts...)
	}
}

// ByDailyRoutineRecs orders the results by daily_routine_recs terms.
func ByDailyRoutineRecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDailyRoutineRecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBodyInfoCount orders the results by body_info count.
func ByBodyInfoCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBodyInfoStep(), opts...)
	}
}

// ByBodyInfo orders the results by body_info terms.
func ByBodyInfo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBodyInfoStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOneRepMaxCount orders the results by one_rep_max count.
func ByOneRepMaxCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOneRepMaxStep(), opts...)
	}
}

// ByOneRepMax orders the results by one_rep_max terms.
func ByOneRepMax(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOneRepMaxStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProgramStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProgramTable, ProgramColumn),
	)
}
func newWeeklyRoutineRecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeeklyRoutineRecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WeeklyRoutineRecsTable, WeeklyRoutineRecsColumn),
	)
}
func newDailyRoutineRecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DailyRoutineRecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DailyRoutineRecsTable, DailyRoutineRecsColumn),
	)
}
func newBodyInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BodyInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BodyInfoTable, BodyInfoColumn),
	)
}
func newOneRepMaxStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OneRepMaxInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OneRepMaxTable, OneRepMaxColumn),
	)
}