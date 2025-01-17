// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"routine/ent/dailyroutine"
	"routine/ent/dailyroutinerec"
	"routine/ent/programrec"
	"routine/ent/weeklyroutinerec"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// DailyRoutineRec is the model entity for the DailyRoutineRec schema.
type DailyRoutineRec struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Author holds the value of the "author" field.
	Author uint64 `json:"author,omitempty"`
	// ProgramRecID holds the value of the "program_rec_id" field.
	ProgramRecID *uint64 `json:"program_rec_id,omitempty"`
	// WeeklyRoutineRecID holds the value of the "weekly_routine_rec_id" field.
	WeeklyRoutineRecID *uint64 `json:"weekly_routine_rec_id,omitempty"`
	// DailyRoutineID holds the value of the "daily_routine_id" field.
	DailyRoutineID *uint64 `json:"daily_routine_id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Status holds the value of the "status" field.
	Status dailyroutinerec.Status `json:"status,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment *string `json:"comment,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DailyRoutineRecQuery when eager-loading is set.
	Edges        DailyRoutineRecEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DailyRoutineRecEdges holds the relations/edges for other nodes in the graph.
type DailyRoutineRecEdges struct {
	// DailyRoutine holds the value of the daily_routine edge.
	DailyRoutine *DailyRoutine `json:"daily_routine,omitempty"`
	// ProgramRec holds the value of the program_rec edge.
	ProgramRec *ProgramRec `json:"program_rec,omitempty"`
	// WeeklyRoutineRec holds the value of the weekly_routine_rec edge.
	WeeklyRoutineRec *WeeklyRoutineRec `json:"weekly_routine_rec,omitempty"`
	// RoutineActRecs holds the value of the routine_act_recs edge.
	RoutineActRecs []*RoutineActRec `json:"routine_act_recs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// DailyRoutineOrErr returns the DailyRoutine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DailyRoutineRecEdges) DailyRoutineOrErr() (*DailyRoutine, error) {
	if e.loadedTypes[0] {
		if e.DailyRoutine == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: dailyroutine.Label}
		}
		return e.DailyRoutine, nil
	}
	return nil, &NotLoadedError{edge: "daily_routine"}
}

// ProgramRecOrErr returns the ProgramRec value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DailyRoutineRecEdges) ProgramRecOrErr() (*ProgramRec, error) {
	if e.loadedTypes[1] {
		if e.ProgramRec == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: programrec.Label}
		}
		return e.ProgramRec, nil
	}
	return nil, &NotLoadedError{edge: "program_rec"}
}

// WeeklyRoutineRecOrErr returns the WeeklyRoutineRec value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DailyRoutineRecEdges) WeeklyRoutineRecOrErr() (*WeeklyRoutineRec, error) {
	if e.loadedTypes[2] {
		if e.WeeklyRoutineRec == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: weeklyroutinerec.Label}
		}
		return e.WeeklyRoutineRec, nil
	}
	return nil, &NotLoadedError{edge: "weekly_routine_rec"}
}

// RoutineActRecsOrErr returns the RoutineActRecs value or an error if the edge
// was not loaded in eager-loading.
func (e DailyRoutineRecEdges) RoutineActRecsOrErr() ([]*RoutineActRec, error) {
	if e.loadedTypes[3] {
		return e.RoutineActRecs, nil
	}
	return nil, &NotLoadedError{edge: "routine_act_recs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DailyRoutineRec) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dailyroutinerec.FieldID, dailyroutinerec.FieldAuthor, dailyroutinerec.FieldProgramRecID, dailyroutinerec.FieldWeeklyRoutineRecID, dailyroutinerec.FieldDailyRoutineID:
			values[i] = new(sql.NullInt64)
		case dailyroutinerec.FieldStatus, dailyroutinerec.FieldComment:
			values[i] = new(sql.NullString)
		case dailyroutinerec.FieldDate, dailyroutinerec.FieldCreatedAt, dailyroutinerec.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DailyRoutineRec fields.
func (drr *DailyRoutineRec) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dailyroutinerec.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			drr.ID = uint64(value.Int64)
		case dailyroutinerec.FieldAuthor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				drr.Author = uint64(value.Int64)
			}
		case dailyroutinerec.FieldProgramRecID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field program_rec_id", values[i])
			} else if value.Valid {
				drr.ProgramRecID = new(uint64)
				*drr.ProgramRecID = uint64(value.Int64)
			}
		case dailyroutinerec.FieldWeeklyRoutineRecID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weekly_routine_rec_id", values[i])
			} else if value.Valid {
				drr.WeeklyRoutineRecID = new(uint64)
				*drr.WeeklyRoutineRecID = uint64(value.Int64)
			}
		case dailyroutinerec.FieldDailyRoutineID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field daily_routine_id", values[i])
			} else if value.Valid {
				drr.DailyRoutineID = new(uint64)
				*drr.DailyRoutineID = uint64(value.Int64)
			}
		case dailyroutinerec.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				drr.Date = value.Time
			}
		case dailyroutinerec.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				drr.Status = dailyroutinerec.Status(value.String)
			}
		case dailyroutinerec.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				drr.Comment = new(string)
				*drr.Comment = value.String
			}
		case dailyroutinerec.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				drr.CreatedAt = value.Time
			}
		case dailyroutinerec.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				drr.UpdatedAt = value.Time
			}
		default:
			drr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DailyRoutineRec.
// This includes values selected through modifiers, order, etc.
func (drr *DailyRoutineRec) Value(name string) (ent.Value, error) {
	return drr.selectValues.Get(name)
}

// QueryDailyRoutine queries the "daily_routine" edge of the DailyRoutineRec entity.
func (drr *DailyRoutineRec) QueryDailyRoutine() *DailyRoutineQuery {
	return NewDailyRoutineRecClient(drr.config).QueryDailyRoutine(drr)
}

// QueryProgramRec queries the "program_rec" edge of the DailyRoutineRec entity.
func (drr *DailyRoutineRec) QueryProgramRec() *ProgramRecQuery {
	return NewDailyRoutineRecClient(drr.config).QueryProgramRec(drr)
}

// QueryWeeklyRoutineRec queries the "weekly_routine_rec" edge of the DailyRoutineRec entity.
func (drr *DailyRoutineRec) QueryWeeklyRoutineRec() *WeeklyRoutineRecQuery {
	return NewDailyRoutineRecClient(drr.config).QueryWeeklyRoutineRec(drr)
}

// QueryRoutineActRecs queries the "routine_act_recs" edge of the DailyRoutineRec entity.
func (drr *DailyRoutineRec) QueryRoutineActRecs() *RoutineActRecQuery {
	return NewDailyRoutineRecClient(drr.config).QueryRoutineActRecs(drr)
}

// Update returns a builder for updating this DailyRoutineRec.
// Note that you need to call DailyRoutineRec.Unwrap() before calling this method if this DailyRoutineRec
// was returned from a transaction, and the transaction was committed or rolled back.
func (drr *DailyRoutineRec) Update() *DailyRoutineRecUpdateOne {
	return NewDailyRoutineRecClient(drr.config).UpdateOne(drr)
}

// Unwrap unwraps the DailyRoutineRec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (drr *DailyRoutineRec) Unwrap() *DailyRoutineRec {
	_tx, ok := drr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DailyRoutineRec is not a transactional entity")
	}
	drr.config.driver = _tx.drv
	return drr
}

// String implements the fmt.Stringer.
func (drr *DailyRoutineRec) String() string {
	var builder strings.Builder
	builder.WriteString("DailyRoutineRec(")
	builder.WriteString(fmt.Sprintf("id=%v, ", drr.ID))
	builder.WriteString("author=")
	builder.WriteString(fmt.Sprintf("%v", drr.Author))
	builder.WriteString(", ")
	if v := drr.ProgramRecID; v != nil {
		builder.WriteString("program_rec_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := drr.WeeklyRoutineRecID; v != nil {
		builder.WriteString("weekly_routine_rec_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := drr.DailyRoutineID; v != nil {
		builder.WriteString("daily_routine_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(drr.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", drr.Status))
	builder.WriteString(", ")
	if v := drr.Comment; v != nil {
		builder.WriteString("comment=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(drr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(drr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DailyRoutineRecs is a parsable slice of DailyRoutineRec.
type DailyRoutineRecs []*DailyRoutineRec
