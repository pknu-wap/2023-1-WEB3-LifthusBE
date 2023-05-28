// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "connected_at", Type: field.TypeTime},
		{Name: "signed_at", Type: field.TypeTime, Nullable: true},
		{Name: "used", Type: field.TypeBool, Default: false},
		{Name: "uid", Type: field.TypeUint64, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "registered", Type: field.TypeBool, Default: false},
		{Name: "registered_at", Type: field.TypeTime, Nullable: true},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "email_verified", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "given_name", Type: field.TypeString},
		{Name: "family_name", Type: field.TypeString},
		{Name: "birthdate", Type: field.TypeTime, Nullable: true},
		{Name: "profile_image_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserFollowingColumns holds the columns for the "user_following" table.
	UserFollowingColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "follower_id", Type: field.TypeUint64},
	}
	// UserFollowingTable holds the schema information for the "user_following" table.
	UserFollowingTable = &schema.Table{
		Name:       "user_following",
		Columns:    UserFollowingColumns,
		PrimaryKey: []*schema.Column{UserFollowingColumns[0], UserFollowingColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_following_user_id",
				Columns:    []*schema.Column{UserFollowingColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_following_follower_id",
				Columns:    []*schema.Column{UserFollowingColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SessionsTable,
		UsersTable,
		UserFollowingTable,
	}
)

func init() {
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.Annotation = &entsql.Annotation{
		Options: "ENGINE=MEMORY",
	}
	UserFollowingTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowingTable.ForeignKeys[1].RefTable = UsersTable
}
