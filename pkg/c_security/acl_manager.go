package security

import "github.com/polarsignals/frostdb/query/logicalplan"

type AclManager interface {
	AddRole(roleDef *RoleDef) error // RoleDef related
	GetRole(roleName string) *RoleDef

	AddUser(userDef *UserDef) error // UserDef related
	GetUser(userName string) *UserDef

	CheckQueryIntegrity(query *logicalplan.LogicalPlan, user *UserDef) error
}

var _ AclManager = &MockAclManager{}

type RoleDef struct {
	Id         int
	RoleName   string
	Privileges []Privilege
}

type Privilege int

const (
	Select Privilege = iota
	Insert
	Update
	Delete
)

type UserDef struct {
	UserName string
	Roles    []RoleDef
}
