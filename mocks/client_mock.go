package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockClient is the mock implementation of ClientInterface
type MockClient struct {
	mock.Mock
}

// Mock SendCommand method
func (m *MockClient) SendCommand(cmd, hostname, port string) (string, error) {
	args := m.Called(cmd, hostname, port)
	return args.String(0), args.Error(1)
}

// Mock Authenticate method
func (m *MockClient) Authenticate(username, password, database, hostname, port string) error {
	args := m.Called(username, password, database, hostname, port)
	return args.Error(0)
}
