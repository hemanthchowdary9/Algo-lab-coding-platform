package handler

import (
	"coding-platform/commons"
	"coding-platform/services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

func DashboardPage(w http.ResponseWriter, r *http.Request) {

	username := r.Context().Value("username").(string) // Retrieve username from context
	//tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "dashboard.html")))
	tmpl := commons.GetTemplate("dashboard.html")

	data := struct {
		Username string
	}{
		Username: username,
	}

	tmpl.Execute(w, data)
}

func PremiumPage(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string) // Retrieve username from context

	fmt.Println("new premium user", username)
	usernameWithoutRole := strings.Split(username, "$")[0]
	newRole := "PREMIUM"
	err := services.UpdateUserRole(usernameWithoutRole, newRole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userWithNewRole := usernameWithoutRole + "$" + newRole
	claims := &jwt.StandardClaims{
		Subject:   userWithNewRole,
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
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func CustomerCare(w http.ResponseWriter, r *http.Request) {

	username := r.Context().Value("username").(string) // Retrieve username from context
	//tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "customer-care.html")))
	tmpl := commons.GetTemplate("customer-care.html")

	data := struct {
		Username string
	}{
		Username: username,
	}

	tmpl.Execute(w, data)
}
