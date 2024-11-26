package handler

import (
	"coding-platform/commons"
	"coding-platform/models"
	"coding-platform/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ChallengesPage(w http.ResponseWriter, r *http.Request) {

	username := r.Context().Value("username").(string) // Retrieve username from context
	//tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "challengesV3.html")))
	tmpl := commons.GetTemplate("challengesV3.html")

	fmt.Println("calling challenges", username)
	category := r.URL.Query().Get("id")
	fmt.Println("calling challenges", username, category)
	// Placeholder data: coding challenges
	var header string
	today := time.Now()
	formattedDate := today.Format("2006-01-02")

	if category == formattedDate {
		header = "Today's Featured Challenges"
	} else if category == "java" {
		header = "OOP Excellence"
	} else if category == "go" {
		header = "Concurrency and Simplicity"
	} else if category == "js" {
		header = "Dynamic Web Interactions"
	} else if category == "top_picks" {
		header = "Curated Coding Excellence"
	} else if category == "interview_ready" {
		header = "Essential Problem-Solving skills"
	} else if category == "advanced" {
		header = "Complex Algorithms Mastery"
	} else if category == "master_syntax" {
		header = "Excel with Syntaxing"
	} else if category == "unt_special" {
		header = "Curated List of Challenges by UNT"
	} else if category == "dsa" {
		header = "Master Data Structures and Algorithms"
	} else {
		header = "Most Liked Challenges"
		category = "public"
	}

	challenges, err := services.FetchChallenges(-1, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var previewChallenges []models.ChallengePreview
	for _, challenge := range challenges {
		premium := false
		if challenge.Premium == "PREMIUM" {
			premium = true
			if strings.Split(username, "$")[1] == "PREMIUM" {
				premium = false
			}

		}
		fmt.Println(" challenge premium ", challenge.Id, challenge.Premium)
		previewChallenges = append(previewChallenges, models.ChallengePreview{
			Title:      challenge.Title,
			Difficulty: challenge.Difficulty,
			ID:         challenge.Id,
			Premium:    premium,
		})
	}

	data := struct {
		HeaderTitle string
		Challenges  []models.ChallengePreview
		Username    string
	}{
		HeaderTitle: header,
		Challenges:  previewChallenges,
		Username:    username,
	}

	tmpl.Execute(w, data)
}

func ChallengeInfo(w http.ResponseWriter, r *http.Request) {

	username := r.Context().Value("username").(string) // Retrieve username from context
	//tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "challenge_detailsV2.html")))
	tmpl := commons.GetTemplate("challenge_detailsV2.html")

	challengeId := r.URL.Query().Get("id")
	fmt.Println("calling challenge info", username, challengeId)
	// Placeholder data: coding challenges
	//challenge := config.IdToChallengeMap[challengeId]
	id, err := strconv.Atoi(challengeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	challenges, err := services.FetchChallenges(id, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var challenge models.Challenge
	if len(challenges) > 0 && (challenges[0].Premium != "PREMIUM" || challenges[0].Premium == "PREMIUM" && strings.Split(username, "$")[1] == "PREMIUM") {
		challenge = challenges[0]
	}
	tmpl.Execute(w, challenge)
}

func FetchChallengeJSON(w http.ResponseWriter, r *http.Request) {

	challengeId := r.URL.Query().Get("id")
	fmt.Println("calling challenge info", challengeId)
	// Placeholder data: coding challenges
	//challenge := config.IdToChallengeMap[challengeId]
	id, err := strconv.Atoi(challengeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	challenges, err := services.FetchChallenges(id, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var challenge models.Challenge
	if len(challenges) > 0 {
		challenge = challenges[0]
	}
	commons.WriteResponse(w, http.StatusOK, commons.ToByteArray(challenge))
}

func CreateChallenge(w http.ResponseWriter, r *http.Request) {
	var requestBody models.Challenge
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdChallenge, err := services.CreateChallenge(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	commons.WriteResponse(w, http.StatusOK, commons.ToByteArray(createdChallenge))
}
