package handlers

import (
	"database/sql"
//	"fmt"
	"go-chat-server/models"
	"go-chat-server/utils"
	"log"
	"math/rand"
	"strconv"
//	"time"

	"gopkg.in/telebot.v3"
)

func RegisterAuthHandlers(bot *telebot.Bot, db *sql.DB) {
	bot.Handle("/signup", func(c telebot.Context) error {
		email := c.Message().Payload
		if email == "" {
			return c.Send("Usage: /signup your_email@example.com")
		}

		otp := strconv.Itoa(rand.Intn(899999) + 100000)
		query := "INSERT INTO users (email, otp, role) VALUES (?, ?, 'client')"
		_, err := db.Exec(query, email, otp)
		if err != nil {
			return c.Send("Error: Could not create user.")
		}

		// Send OTP via Email
		go utils.SendOTP(email, otp)

		return c.Send("An OTP has been sent to your email. Use /verify_otp <OTP> to verify.")
	})

	bot.Handle("/verify_otp", func(c telebot.Context) error {
		args := c.Message().Payload
		if args == "" {
			return c.Send("Usage: /verify_otp <OTP>")
		}

		var user models.User
		err := db.QueryRow("SELECT id, email FROM users WHERE otp = ?", args).Scan(&user.ID, &user.Email)
		if err != nil {
			return c.Send("Invalid OTP.")
		}

		// Clear OTP after successful verification
		_, err = db.Exec("UPDATE users SET otp = NULL WHERE id = ?", user.ID)
		if err != nil {
			log.Println(err)
		}

		return c.Send("You are now verified! Use /login to access your account.")
	})
}
