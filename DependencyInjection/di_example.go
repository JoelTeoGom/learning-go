package main

import "fmt"

// Service interface that defines the behavior we depend on
type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

// Concrete implementation 1 of EmailService
type ExpressEmailService struct{}

func (e *ExpressEmailService) SendEmail(to string, subject string, body string) error {
	fmt.Printf("ExpressEmailService sending email to: %s, Subject: %s\n", to, subject)
	return nil
}

// Concrete implementation 2 of EmailService
type PremiumEmailService struct{}

func (p *PremiumEmailService) SendEmail(to string, subject string, body string) error {
	fmt.Printf("PremiumEmailService sending email to: %s, Subject: %s\n", to, subject)
	return nil
}

// User structure that holds email and associated email service
type User struct {
	Email        string
	EmailService EmailService
}

// Function that takes an EmailService (dependency injection)
func NotifyUsers(users []User) {
	for _, user := range users {
		user.EmailService.SendEmail(user.Email, "Hello from Go!", "This is an injected email service!")
	}
}

func main() {
	// List of users with different email services
	users := []User{
		{Email: "user1@example.com", EmailService: &ExpressEmailService{}},
		{Email: "user2@example.com", EmailService: &PremiumEmailService{}},
		{Email: "user3@example.com", EmailService: &ExpressEmailService{}},
	}

	// Notify each user using their specific email service
	NotifyUsers(users)
}
