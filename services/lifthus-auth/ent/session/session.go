// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldConnectedAt holds the string denoting the connected_at field in the database.
	FieldConnectedAt = "connected_at"
	// FieldSignedAt holds the string denoting the signed_at field in the database.
	FieldSignedAt = "signed_at"
	// FieldUsed holds the string denoting the used field in the database.
	FieldUsed = "used"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the session in the database.
	Table = "sessions"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "sessions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "uid"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldConnectedAt,
	FieldSignedAt,
	FieldUsed,
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
	// DefaultConnectedAt holds the default value on creation for the "connected_at" field.
	DefaultConnectedAt func() time.Time
	// UpdateDefaultSignedAt holds the default value on update for the "signed_at" field.
	UpdateDefaultSignedAt func() time.Time
	// DefaultUsed holds the default value on creation for the "used" field.
	DefaultUsed bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Session queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUID orders the results by the uid field.
func ByUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUID, opts...).ToFunc()
}

// ByConnectedAt orders the results by the connected_at field.
func ByConnectedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConnectedAt, opts...).ToFunc()
}

// BySignedAt orders the results by the signed_at field.
func BySignedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignedAt, opts...).ToFunc()
}

// ByUsed orders the results by the used field.
func ByUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsed, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
