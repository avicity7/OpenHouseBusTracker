package utils

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

func GenerateRandomToken(length int) string {
	arr := make([]byte, length)
	_, err := rand.Read(arr)
	if err != nil {
		fmt.Println("Error")
	}
	token := base64.URLEncoding.EncodeToString(arr)
	return token
}

func SendEmail(token string, user_email string) {
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
	email := hermes.Email{
		Body: hermes.Body{
			Greeting: "Hey there",
			Name:     user_email,
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
	emailBody, _ := h.GenerateHTML(email)
	from := "hellospopenhouse@gmail.com"
	to := user_email
	password := os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()

	m.SetHeader("From", "hellospopenhouse@gmail.com")

	m.SetHeader("To", to)

	m.SetHeader("Subject", "Verify your email")

	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
