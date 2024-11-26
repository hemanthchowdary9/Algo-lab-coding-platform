package main

import (
	"coding-platform/config"
	"coding-platform/database"
	"coding-platform/handler"
	"coding-platform/middlewares"
	"fmt"
	"github.com/stripe/stripe-go/v81"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Main Started - coding platform service")
	stripe.Key = "sk_test_51QFrciHUG3HSOntwDif36ftSsnlg4OSnhdIHZ5jPHSuBOypwejJYUUQQnxolzkL0l2su65jfZVbmTzKRYOCVDE8600JCQzpsID"

	config.LoadYamlConfigurations()
	database.InitializeDB()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//http.HandleFunc("/login", handler.LoginPage)
	http.Handle("/", http.HandlerFunc(handler.Login))
	http.Handle("/signup", http.HandlerFunc(handler.SignUp))
	http.Handle("/logout", http.HandlerFunc(handler.Logout))
	http.Handle("/challenge/create", http.HandlerFunc(handler.CreateChallenge))
	http.Handle("/challenge/fetch", http.HandlerFunc(handler.FetchChallengeJSON))
	http.Handle("/challenge/update", http.HandlerFunc(handler.ChallengeInfo))
	http.Handle("/challenges", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.ChallengesPage)))
	http.Handle("/customer-care", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.CustomerCare)))
	http.Handle("/dashboard", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.DashboardPage)))
	http.Handle("/challenge/description", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.ChallengeInfo)))
	http.Handle("/twoFactor/auth", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.TwoFactorAuth)))
	http.Handle("/verify-otp", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.VerifyOTPHandler)))
	http.Handle("/compile", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.Compile)))
	http.Handle("/compile-test", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.CompileTest)))
	http.Handle("/save-submission", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.SaveSubmission)))
	http.Handle("/submissions", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.FetchSubmissions)))
	http.Handle("/create-checkout-session", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.CreateCheckoutSession)))
	http.Handle("/premium", middlewares.JWTAuthMiddleware(http.HandlerFunc(handler.PremiumPage)))

	// Start server on port 8080
	fmt.Println("Server starting at :7074")
	log.Fatal(http.ListenAndServe(":7074", nil))
}
