package services

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"gopkg.in/gomail.v2"

	"log"
	"time"
)

func GenerateOTP() string {
	otp, err := totp.GenerateCode("ABCDEFG", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return otp
}

func SendOTPtoMail(to, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "algolab.coding@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "One Time Password - AlgoLab")
	m.SetBody("text/plain", "OTP to Login AlgoLab : "+otp)

	// mail password: !Nothwsinkr8cvcy
	d := gomail.NewDialer("smtp.gmail.com", 587, "algolab.coding@gmail.com", "zdhv jjyz uomq gdul")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error while sending mail", err)
		return err
	}
	return nil
}

func SendWelcomeEmail(to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "algolab.coding@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Welcome to AlgoLab!")

	// HTML content for a beautiful email body
	htmlBody := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Welcome Email</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					color: #333;
					background-color: #f4f4f9;
					padding: 20px;
				}
				.container {
					max-width: 600px;
					margin: auto;
					background-color: #ffffff;
					border-radius: 8px;
					padding: 20px;
					box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.2);
				}
				.header {
					text-align: center;
					color: #4CAF50;
					margin-bottom: 20px;
				}
				.header h1 {
					margin: 0;
				}
				.content p {
					line-height: 1.6;
				}
				.button {
					display: inline-block;
					margin: 20px 0;
					padding: 10px 20px;
					color: #ffffff;
					background-color: #4CAF50;
					border-radius: 5px;
					text-decoration: none;
					font-weight: bold;
				}
				.footer {
					text-align: center;
					font-size: 12px;
					color: #888;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Welcome to AlgoLab!</h1>
				</div>
				<div class="content">
					<p>Hello,</p>
					<p>Thank you for signing up for AlgoLab! We’re thrilled to have you on board and excited to help you with your coding journey.</p>
					<p>To get started, explore our challenges and start solving problems to improve your skills.</p>
					<p>If you have any questions or need assistance, feel free to reach out to us with this email.</p>
				</div>
				<div class="footer">
					<p>Happy Coding!<br>– The AlgoLab Team</p>
				</div>
			</div>
		</body>
		</html>
	`

	// Set HTML body
	m.SetBody("text/html", htmlBody)

	// Send email
	d := gomail.NewDialer("smtp.gmail.com", 587, "algolab.coding@gmail.com", "zdhv jjyz uomq gdul")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error while sending mail", err)
		return err
	}
	return nil
}
