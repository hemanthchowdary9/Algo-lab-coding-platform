package handler

import (
	"fmt"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"log"
	"net/http"
	"os"
)

func CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:7074"
	if os.Getenv("DOMAIN_URL") != "" {
		domain = os.Getenv("DOMAIN_URL")
		fmt.Println("fetched domain url from env variables: ", domain)
	}
	price_id := "price_1QFsr2HUG3HSOntwPhCxZDuX"
	if os.Getenv("PRICE_ID") != "" {
		price_id = os.Getenv("PRICE_ID")
		fmt.Println("fetched price id from env variables: ", price_id)
	}
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String(price_id),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/premium"),
		CancelURL:  stripe.String(domain + "/dashboard"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	fmt.Println("end of create session")
	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
