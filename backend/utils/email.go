// Package utils defines functions that do not necessarily update entities, but are re-used.
package utils

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

// This function is used throughout different Controllers for different purposes, hence the logic.
func SendEmail(token string, user_email string, pwd string, name string) error {
	godotenv.Load()
	env := os.Getenv("ENV")
	var link string
	if env == "PROD" {
		link = "https://open-house-bus-tracker.vercel.app"
	} else {
		link = "http://localhost:5173"
	}

	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      "SP OH Bus Team",
			Link:      "https://www.sp.edu.sg/",
			Logo:      "https://www.sp.edu.sg/images/default-source/default-album/sp-70-logo.png?sfvrsn=7dfc222d_0",
			Copyright: "Copyright © 2024 Singapore Polytechnic. All rights reserved.",
		},
	}
	var email hermes.Email
	if pwd == "" {
		email = hermes.Email{
			Body: hermes.Body{
				Greeting: "Hey",
				Name:     name,
				Intros: []string{
					"Welcome to OH Bus Tracker!",
				},
				Actions: []hermes.Action{
					{
						Instructions: "To get started, click here:",
						Button: hermes.Button{
							Color: "#b91c1c",
							Text:  "Verify your email",
							Link:  link + "/verify/" + token,
						},
					},
				},
				Outros: []string{
					"Need help, or have questions? Just reply to this email, we'd love to help.",
				},
			},
		}
	} else if pwd == "reset" {
		email = hermes.Email{
			Body: hermes.Body{
				Greeting: "Hey",
				Name:     name,
				Intros: []string{
					"We've received a request to reset your password.",
				},
				Actions: []hermes.Action{
					{
						Instructions: "To reset your password, click here:",
						Button: hermes.Button{
							Color: "#b91c1c",
							Text:  "Reset password",
							Link:  link + "/profile/reset-password/" + token,
						},
					},
				},
				Outros: []string{
					"Need help, or have questions? Just reply to this email, we'd love to help.",
				},
			},
		}
	} else {
		email = hermes.Email{
			Body: hermes.Body{
				Greeting: "Hey",
				Name:     name,
				Intros: []string{
					"Welcome to OH Bus Tracker!",
				},
				Actions: []hermes.Action{
					{
						Instructions: "Here's your password!",
						InviteCode:   pwd,
					},
				},
				Outros: []string{
					"Need help, or have questions? Just reply to this email, we'd love to help.",
				},
			},
		}
	}

	emailBody, _ := h.GenerateHTML(email)
	from := "hellospopenhouse@gmail.com"
	to := user_email
	password := os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()

	m.SetHeader("From", "hellospopenhouse@gmail.com")

	m.SetHeader("To", to)

	if pwd == "" {
		m.SetHeader("Subject", "Verify your email")
	} else if pwd == "reset" {
		m.SetHeader("Subject", "Reset your password")
	} else {
		m.SetHeader("Subject", "You're signed up!")
	}

	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}
