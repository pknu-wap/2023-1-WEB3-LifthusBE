// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lifthus-auth/ent/predicate"
	"lifthus-auth/ent/session"
	"lifthus-auth/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetRegistered sets the "registered" field.
func (uu *UserUpdate) SetRegistered(b bool) *UserUpdate {
	uu.mutation.SetRegistered(b)
	return uu
}

// SetNillableRegistered sets the "registered" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRegistered(b *bool) *UserUpdate {
	if b != nil {
		uu.SetRegistered(*b)
	}
	return uu
}

// SetRegisteredAt sets the "registered_at" field.
func (uu *UserUpdate) SetRegisteredAt(t time.Time) *UserUpdate {
	uu.mutation.SetRegisteredAt(t)
	return uu
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRegisteredAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetRegisteredAt(*t)
	}
	return uu
}

// ClearRegisteredAt clears the value of the "registered_at" field.
func (uu *UserUpdate) ClearRegisteredAt() *UserUpdate {
	uu.mutation.ClearRegisteredAt()
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// ClearUsername clears the value of the "username" field.
func (uu *UserUpdate) ClearUsername() *UserUpdate {
	uu.mutation.ClearUsername()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetEmailVerified sets the "email_verified" field.
func (uu *UserUpdate) SetEmailVerified(b bool) *UserUpdate {
	uu.mutation.SetEmailVerified(b)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetGivenName sets the "given_name" field.
func (uu *UserUpdate) SetGivenName(s string) *UserUpdate {
	uu.mutation.SetGivenName(s)
	return uu
}

// SetFamilyName sets the "family_name" field.
func (uu *UserUpdate) SetFamilyName(s string) *UserUpdate {
	uu.mutation.SetFamilyName(s)
	return uu
}

// SetBirthdate sets the "birthdate" field.
func (uu *UserUpdate) SetBirthdate(t time.Time) *UserUpdate {
	uu.mutation.SetBirthdate(t)
	return uu
}

// SetNillableBirthdate sets the "birthdate" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBirthdate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetBirthdate(*t)
	}
	return uu
}

// ClearBirthdate clears the value of the "birthdate" field.
func (uu *UserUpdate) ClearBirthdate() *UserUpdate {
	uu.mutation.ClearBirthdate()
	return uu
}

// SetProfileImageURL sets the "profile_image_url" field.
func (uu *UserUpdate) SetProfileImageURL(s string) *UserUpdate {
	uu.mutation.SetProfileImageURL(s)
	return uu
}

// SetNillableProfileImageURL sets the "profile_image_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableProfileImageURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetProfileImageURL(*s)
	}
	return uu
}

// ClearProfileImageURL clears the value of the "profile_image_url" field.
func (uu *UserUpdate) ClearProfileImageURL() *UserUpdate {
	uu.mutation.ClearProfileImageURL()
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetUsercode sets the "usercode" field.
func (uu *UserUpdate) SetUsercode(s string) *UserUpdate {
	uu.mutation.SetUsercode(s)
	return uu
}

// SetNillableUsercode sets the "usercode" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsercode(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsercode(*s)
	}
	return uu
}

// SetCompany sets the "company" field.
func (uu *UserUpdate) SetCompany(s string) *UserUpdate {
	uu.mutation.SetCompany(s)
	return uu
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCompany(s *string) *UserUpdate {
	if s != nil {
		uu.SetCompany(*s)
	}
	return uu
}

// ClearCompany clears the value of the "company" field.
func (uu *UserUpdate) ClearCompany() *UserUpdate {
	uu.mutation.ClearCompany()
	return uu
}

// SetLocation sets the "location" field.
func (uu *UserUpdate) SetLocation(s string) *UserUpdate {
	uu.mutation.SetLocation(s)
	return uu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLocation(s *string) *UserUpdate {
	if s != nil {
		uu.SetLocation(*s)
	}
	return uu
}

// ClearLocation clears the value of the "location" field.
func (uu *UserUpdate) ClearLocation() *UserUpdate {
	uu.mutation.ClearLocation()
	return uu
}

// SetContact sets the "contact" field.
func (uu *UserUpdate) SetContact(s string) *UserUpdate {
	uu.mutation.SetContact(s)
	return uu
}

// SetNillableContact sets the "contact" field if the given value is not nil.
func (uu *UserUpdate) SetNillableContact(s *string) *UserUpdate {
	if s != nil {
		uu.SetContact(*s)
	}
	return uu
}

// ClearContact clears the value of the "contact" field.
func (uu *UserUpdate) ClearContact() *UserUpdate {
	uu.mutation.ClearContact()
	return uu
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (uu *UserUpdate) AddSessionIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddSessionIDs(ids...)
	return uu
}

// AddSessions adds the "sessions" edges to the Session entity.
func (uu *UserUpdate) AddSessions(s ...*Session) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSessionIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowingIDs(ids ...uint64) *UserUpdate {
	uu.mutation.AddFollowingIDs(ids...)
	return uu
}

// AddFollowing adds the "following" edges to the User entity.
func (uu *UserUpdate) AddFollowing(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowingIDs(ids...)
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowerIDs(ids ...uint64) *UserUpdate {
	uu.mutation.AddFollowerIDs(ids...)
	return uu
}

// AddFollowers adds the "followers" edges to the User entity.
func (uu *UserUpdate) AddFollowers(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowerIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (uu *UserUpdate) ClearSessions() *UserUpdate {
	uu.mutation.ClearSessions()
	return uu
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (uu *UserUpdate) RemoveSessionIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveSessionIDs(ids...)
	return uu
}

// RemoveSessions removes "sessions" edges to Session entities.
func (uu *UserUpdate) RemoveSessions(s ...*Session) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSessionIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uu *UserUpdate) ClearFollowing() *UserUpdate {
	uu.mutation.ClearFollowing()
	return uu
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowingIDs(ids ...uint64) *UserUpdate {
	uu.mutation.RemoveFollowingIDs(ids...)
	return uu
}

// RemoveFollowing removes "following" edges to User entities.
func (uu *UserUpdate) RemoveFollowing(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowingIDs(ids...)
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uu *UserUpdate) ClearFollowers() *UserUpdate {
	uu.mutation.ClearFollowers()
	return uu
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowerIDs(ids ...uint64) *UserUpdate {
	uu.mutation.RemoveFollowerIDs(ids...)
	return uu
}

// RemoveFollowers removes "followers" edges to User entities.
func (uu *UserUpdate) RemoveFollowers(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Registered(); ok {
		_spec.SetField(user.FieldRegistered, field.TypeBool, value)
	}
	if value, ok := uu.mutation.RegisteredAt(); ok {
		_spec.SetField(user.FieldRegisteredAt, field.TypeTime, value)
	}
	if uu.mutation.RegisteredAtCleared() {
		_spec.ClearField(user.FieldRegisteredAt, field.TypeTime)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if uu.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.EmailVerified(); ok {
		_spec.SetField(user.FieldEmailVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.GivenName(); ok {
		_spec.SetField(user.FieldGivenName, field.TypeString, value)
	}
	if value, ok := uu.mutation.FamilyName(); ok {
		_spec.SetField(user.FieldFamilyName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Birthdate(); ok {
		_spec.SetField(user.FieldBirthdate, field.TypeTime, value)
	}
	if uu.mutation.BirthdateCleared() {
		_spec.ClearField(user.FieldBirthdate, field.TypeTime)
	}
	if value, ok := uu.mutation.ProfileImageURL(); ok {
		_spec.SetField(user.FieldProfileImageURL, field.TypeString, value)
	}
	if uu.mutation.ProfileImageURLCleared() {
		_spec.ClearField(user.FieldProfileImageURL, field.TypeString)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Usercode(); ok {
		_spec.SetField(user.FieldUsercode, field.TypeString, value)
	}
	if value, ok := uu.mutation.Company(); ok {
		_spec.SetField(user.FieldCompany, field.TypeString, value)
	}
	if uu.mutation.CompanyCleared() {
		_spec.ClearField(user.FieldCompany, field.TypeString)
	}
	if value, ok := uu.mutation.Location(); ok {
		_spec.SetField(user.FieldLocation, field.TypeString, value)
	}
	if uu.mutation.LocationCleared() {
		_spec.ClearField(user.FieldLocation, field.TypeString)
	}
	if value, ok := uu.mutation.Contact(); ok {
		_spec.SetField(user.FieldContact, field.TypeString, value)
	}
	if uu.mutation.ContactCleared() {
		_spec.ClearField(user.FieldContact, field.TypeString)
	}
	if uu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !uu.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetRegistered sets the "registered" field.
func (uuo *UserUpdateOne) SetRegistered(b bool) *UserUpdateOne {
	uuo.mutation.SetRegistered(b)
	return uuo
}

// SetNillableRegistered sets the "registered" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRegistered(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetRegistered(*b)
	}
	return uuo
}

// SetRegisteredAt sets the "registered_at" field.
func (uuo *UserUpdateOne) SetRegisteredAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetRegisteredAt(t)
	return uuo
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRegisteredAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetRegisteredAt(*t)
	}
	return uuo
}

// ClearRegisteredAt clears the value of the "registered_at" field.
func (uuo *UserUpdateOne) ClearRegisteredAt() *UserUpdateOne {
	uuo.mutation.ClearRegisteredAt()
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// ClearUsername clears the value of the "username" field.
func (uuo *UserUpdateOne) ClearUsername() *UserUpdateOne {
	uuo.mutation.ClearUsername()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetEmailVerified sets the "email_verified" field.
func (uuo *UserUpdateOne) SetEmailVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetEmailVerified(b)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetGivenName sets the "given_name" field.
func (uuo *UserUpdateOne) SetGivenName(s string) *UserUpdateOne {
	uuo.mutation.SetGivenName(s)
	return uuo
}

// SetFamilyName sets the "family_name" field.
func (uuo *UserUpdateOne) SetFamilyName(s string) *UserUpdateOne {
	uuo.mutation.SetFamilyName(s)
	return uuo
}

// SetBirthdate sets the "birthdate" field.
func (uuo *UserUpdateOne) SetBirthdate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetBirthdate(t)
	return uuo
}

// SetNillableBirthdate sets the "birthdate" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBirthdate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetBirthdate(*t)
	}
	return uuo
}

// ClearBirthdate clears the value of the "birthdate" field.
func (uuo *UserUpdateOne) ClearBirthdate() *UserUpdateOne {
	uuo.mutation.ClearBirthdate()
	return uuo
}

// SetProfileImageURL sets the "profile_image_url" field.
func (uuo *UserUpdateOne) SetProfileImageURL(s string) *UserUpdateOne {
	uuo.mutation.SetProfileImageURL(s)
	return uuo
}

// SetNillableProfileImageURL sets the "profile_image_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProfileImageURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetProfileImageURL(*s)
	}
	return uuo
}

// ClearProfileImageURL clears the value of the "profile_image_url" field.
func (uuo *UserUpdateOne) ClearProfileImageURL() *UserUpdateOne {
	uuo.mutation.ClearProfileImageURL()
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetUsercode sets the "usercode" field.
func (uuo *UserUpdateOne) SetUsercode(s string) *UserUpdateOne {
	uuo.mutation.SetUsercode(s)
	return uuo
}

// SetNillableUsercode sets the "usercode" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsercode(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsercode(*s)
	}
	return uuo
}

// SetCompany sets the "company" field.
func (uuo *UserUpdateOne) SetCompany(s string) *UserUpdateOne {
	uuo.mutation.SetCompany(s)
	return uuo
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCompany(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCompany(*s)
	}
	return uuo
}

// ClearCompany clears the value of the "company" field.
func (uuo *UserUpdateOne) ClearCompany() *UserUpdateOne {
	uuo.mutation.ClearCompany()
	return uuo
}

// SetLocation sets the "location" field.
func (uuo *UserUpdateOne) SetLocation(s string) *UserUpdateOne {
	uuo.mutation.SetLocation(s)
	return uuo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLocation(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLocation(*s)
	}
	return uuo
}

// ClearLocation clears the value of the "location" field.
func (uuo *UserUpdateOne) ClearLocation() *UserUpdateOne {
	uuo.mutation.ClearLocation()
	return uuo
}

// SetContact sets the "contact" field.
func (uuo *UserUpdateOne) SetContact(s string) *UserUpdateOne {
	uuo.mutation.SetContact(s)
	return uuo
}

// SetNillableContact sets the "contact" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableContact(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetContact(*s)
	}
	return uuo
}

// ClearContact clears the value of the "contact" field.
func (uuo *UserUpdateOne) ClearContact() *UserUpdateOne {
	uuo.mutation.ClearContact()
	return uuo
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (uuo *UserUpdateOne) AddSessionIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddSessionIDs(ids...)
	return uuo
}

// AddSessions adds the "sessions" edges to the Session entity.
func (uuo *UserUpdateOne) AddSessions(s ...*Session) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSessionIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowingIDs(ids ...uint64) *UserUpdateOne {
	uuo.mutation.AddFollowingIDs(ids...)
	return uuo
}

// AddFollowing adds the "following" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowingIDs(ids...)
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowerIDs(ids ...uint64) *UserUpdateOne {
	uuo.mutation.AddFollowerIDs(ids...)
	return uuo
}

// AddFollowers adds the "followers" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowerIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (uuo *UserUpdateOne) ClearSessions() *UserUpdateOne {
	uuo.mutation.ClearSessions()
	return uuo
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (uuo *UserUpdateOne) RemoveSessionIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveSessionIDs(ids...)
	return uuo
}

// RemoveSessions removes "sessions" edges to Session entities.
func (uuo *UserUpdateOne) RemoveSessions(s ...*Session) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSessionIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowing() *UserUpdateOne {
	uuo.mutation.ClearFollowing()
	return uuo
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowingIDs(ids ...uint64) *UserUpdateOne {
	uuo.mutation.RemoveFollowingIDs(ids...)
	return uuo
}

// RemoveFollowing removes "following" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowingIDs(ids...)
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowers() *UserUpdateOne {
	uuo.mutation.ClearFollowers()
	return uuo
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowerIDs(ids ...uint64) *UserUpdateOne {
	uuo.mutation.RemoveFollowerIDs(ids...)
	return uuo
}

// RemoveFollowers removes "followers" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowerIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Registered(); ok {
		_spec.SetField(user.FieldRegistered, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.RegisteredAt(); ok {
		_spec.SetField(user.FieldRegisteredAt, field.TypeTime, value)
	}
	if uuo.mutation.RegisteredAtCleared() {
		_spec.ClearField(user.FieldRegisteredAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if uuo.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EmailVerified(); ok {
		_spec.SetField(user.FieldEmailVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.GivenName(); ok {
		_spec.SetField(user.FieldGivenName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.FamilyName(); ok {
		_spec.SetField(user.FieldFamilyName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Birthdate(); ok {
		_spec.SetField(user.FieldBirthdate, field.TypeTime, value)
	}
	if uuo.mutation.BirthdateCleared() {
		_spec.ClearField(user.FieldBirthdate, field.TypeTime)
	}
	if value, ok := uuo.mutation.ProfileImageURL(); ok {
		_spec.SetField(user.FieldProfileImageURL, field.TypeString, value)
	}
	if uuo.mutation.ProfileImageURLCleared() {
		_spec.ClearField(user.FieldProfileImageURL, field.TypeString)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Usercode(); ok {
		_spec.SetField(user.FieldUsercode, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Company(); ok {
		_spec.SetField(user.FieldCompany, field.TypeString, value)
	}
	if uuo.mutation.CompanyCleared() {
		_spec.ClearField(user.FieldCompany, field.TypeString)
	}
	if value, ok := uuo.mutation.Location(); ok {
		_spec.SetField(user.FieldLocation, field.TypeString, value)
	}
	if uuo.mutation.LocationCleared() {
		_spec.ClearField(user.FieldLocation, field.TypeString)
	}
	if value, ok := uuo.mutation.Contact(); ok {
		_spec.SetField(user.FieldContact, field.TypeString, value)
	}
	if uuo.mutation.ContactCleared() {
		_spec.ClearField(user.FieldContact, field.TypeString)
	}
	if uuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !uuo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
