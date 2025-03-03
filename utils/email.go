package utils

import (
	"log"
	"net/smtp"
	"os"
)

func SendOTP(email, otp string) {
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{email}

	msg := "Subject: Your OTP Code\n\nYour OTP is: " + otp

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", from, password, "smtp.example.com"),
		from, to, []byte(msg))

	if err != nil {
		log.Println("Failed to send email:", err)
	}
}