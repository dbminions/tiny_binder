package security

import "github.com/polarsignals/frostdb/query/logicalplan"

type MockAclManager struct {
}

func NewMockAclManager() *MockAclManager {
	return &MockAclManager{}
}

func (m *MockAclManager) AddRole(roleDef *RoleDef) error {
	panic("implement me")
}

func (m *MockAclManager) GetRole(roleName string) *RoleDef {
	panic("implement me")
}

func (m *MockAclManager) AddUser(userDef *UserDef) error {
	panic("implement me")
}

func (m *MockAclManager) GetUser(userName string) *UserDef {
	panic("implement me")
}

func (m *MockAclManager) CheckQueryIntegrity(query *logicalplan.LogicalPlan, user *UserDef) error {
	return nil
}
