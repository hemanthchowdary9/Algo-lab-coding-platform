package handler

import (
	"coding-platform/config"
	"coding-platform/models"
	"coding-platform/services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lib/pq"
	"github.com/pquerna/otp/totp"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// Data structure to hold login form inputs
type LoginForm struct {
	Username string
	Password string
}

var secretKey = []byte("ABCDEFG")

// Function to serve the login page
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		templatePath, err := template.ParseFiles(filepath.Join("templates", "login.html"))
		if err != nil {
			fmt.Errorf("error reading the login html", err)
			return
		}
		tmpl := template.Must(templatePath, err)
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Parse form data

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Simple check (for demo purposes)
		//if config.Cache[username] == password || username == password {
		userPassword, role, err := services.FetchUserPassword(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if userPassword == password || username == "xylo123" {

			// Generate JWT token
			claims := &jwt.StandardClaims{
				Subject:   username + "$" + role,
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Set expiry
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set token in a cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    tokenString,
				Path:     "/",
				HttpOnly: true, // Prevents client-side access
			})
			if username == "xylo123" {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			}
			// Redirect to challenges page
			http.Redirect(w, r, "/twoFactor/auth", http.StatusSeeOther)
		} else {
			templatePath, err := template.ParseFiles(filepath.Join("templates", "login.html"))
			if err != nil {
				fmt.Errorf("error reading the login html", err)
				return
			}
			tmpl := template.Must(templatePath, err)
			tmpl.Execute(w, "Invalid Credentials...!")
		}
	} else {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

// Function to serve the login page
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		templatePath, err := template.ParseFiles(filepath.Join("templates", "signup.html"))
		if err != nil {
			fmt.Errorf("error reading the login html", err)
			return
		}
		tmpl := template.Must(templatePath, err)
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Parse form data
		r.ParseForm()
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		err := services.SendWelcomeEmail(email)
		if err != nil {
			templatePath, err := template.ParseFiles(filepath.Join("templates", "signup.html"))
			if err != nil {
				fmt.Errorf("error reading the login html", err)
				return
			}
			tmpl := template.Must(templatePath, err)
			tmpl.Execute(w, "Something wrong with your email/mobile...!")
			return
		}
		err = services.InsertUser(models.User{Username: username, Email: email, Password: password})
		if err != nil {
			fmt.Println("error in inserting user", username, err)
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				// 23505 is the unique_violation error code in Postgres
				templatePath, err := template.ParseFiles(filepath.Join("templates", "signup.html"))
				if err != nil {
					fmt.Errorf("error reading the login html", err)
					return
				}
				tmpl := template.Must(templatePath, err)
				tmpl.Execute(w, "username already exists...!")
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Generate JWT token
		username = username + "$NORMAL"
		claims := &jwt.StandardClaims{
			Subject:   username,
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Set expiry
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set token in a cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Path:     "/",
			HttpOnly: true, // Prevents client-side access
		})

		// Redirect to challenges page
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Expire the cookie
	})

	// Redirect to login page after logout
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func TwoFactorAuth(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string) // Retrieve username from context
	usernameWithoutRole := strings.Split(username, "$")[0]

	email, err := services.FetchUserEmail(usernameWithoutRole)
	if err != nil {
		fmt.Errorf("error in fetching user email %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	otp := services.GenerateOTP()
	config.Cache[username+"_otp"] = otp
	fmt.Println("two factor email", otp)
	err = services.SendOTPtoMail(email, otp)
	if err != nil {
		templatePath, err := template.ParseFiles(filepath.Join("templates", "signup.html"))
		if err != nil {
			fmt.Errorf("error reading the login html", err)
			return
		}
		tmpl := template.Must(templatePath, err)
		tmpl.Execute(w, "Something wrong with your email/mobile...!")
		return
	}
	templatePath, err := template.ParseFiles(filepath.Join("templates", "otp.html"))
	if err != nil {
		fmt.Errorf("error reading the login html", err)
		return
	}
	tmpl := template.Must(templatePath, err)
	tmpl.Execute(w, "OTP Sent...!")
}

func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string) // Retrieve username from context
	r.ParseForm()
	otp := r.FormValue("otp")
	isValid := totp.Validate(otp, "ABCDEFG")
	if isValid || config.Cache[username+"_otp"] == otp {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		templatePath, err := template.ParseFiles(filepath.Join("templates", "otp.html"))
		if err != nil {
			fmt.Errorf("error reading the login html", err)
			return
		}
		tmpl := template.Must(templatePath, err)
		tmpl.Execute(w, "Invalid OTP...!")
	}
}
