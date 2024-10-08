package main

import (
	"testing"
)

// MockEmailService is a mock implementation of EmailService for testing
type MockEmailService struct {
	EmailsSent []string
}

func (m *MockEmailService) SendEmail(to string, subject string, body string) error {
	m.EmailsSent = append(m.EmailsSent, to)
	return nil
}

func TestNotifyUsers(t *testing.T) {
	mockService := &MockEmailService{}
	users := []User{
		{Email: "test1@example.com", EmailService: mockService},
		{Email: "test2@example.com", EmailService: mockService},
	}

	NotifyUsers(users)

	if len(mockService.EmailsSent) != len(users) {
		t.Errorf("Expected %d emails to be sent, but got %d", len(users), len(mockService.EmailsSent))
	}

	for i, user := range users {
		if mockService.EmailsSent[i] != user.Email {
			t.Errorf("Expected to send email to %s, but sent to %s", user.Email, mockService.EmailsSent[i])
		}
	}
}
