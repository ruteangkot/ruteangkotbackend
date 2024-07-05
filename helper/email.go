package helper

import (
	"fmt"
	"net/smtp"
)

func SendResetPasswordEmail(email string, token string) error {
	from := "your-email@example.com"
	password := "your-email-password"
	to := email
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	message := fmt.Sprintf("Subject: Reset Password\n\nClick on the link to reset your password: https://yourdomain.com/reset-password?token=%s", token)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
}