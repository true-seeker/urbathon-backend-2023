//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Users struct {
	ID             int32 `sql:"primary_key"`
	Email          *string
	Password       *[]byte
	Salt           *[]byte
	Role           *int32
	OrganizationID *int32
	FirstName      *string
	LastName       *string
	Patronymic     *string
}