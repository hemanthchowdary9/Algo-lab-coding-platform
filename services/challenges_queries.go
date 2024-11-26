package services

import (
	"coding-platform/commons"
	"coding-platform/database"
	"coding-platform/models"
	"encoding/json"
	"errors"
	"fmt"
)

func CreateChallenge(challenge models.Challenge) (models.Challenge, error) {
	db := database.GetPsqlConnection()
	var lastInsertID int
	challengeJson := commons.ToByteArray(challenge)
	if challengeJson == nil {
		return models.Challenge{}, errors.New("marshalling error while inserting to db")
	}
	err := db.QueryRow(`INSERT INTO challenges(title, difficulty, challenge, category) VALUES ($1, $2, $3, $4) RETURNING id`,
		challenge.Title, challenge.Difficulty, challengeJson, challenge.Category).Scan(&lastInsertID)
	if err != nil {
		fmt.Errorf("error in creating challenge %v", err)
		return models.Challenge{}, err
	}
	challenge.Id = lastInsertID
	return challenge, nil
}

func FetchChallenges(id int, category string) ([]models.Challenge, error) {
	db := database.GetPsqlConnection()
	query := "SELECT id, challenge FROM challenges"
	var args []interface{}

	// Check for conditions and append to query and args
	if id >= 0 {
		query += " WHERE id = $1"
		args = append(args, id) // Add id to args slice
	}

	if category != "" {
		if len(args) > 0 {
			query += " AND category = $2" // Use AND if id was already added
			args = append(args, category) // Add category to args slice
		} else {
			query += " WHERE category = $1" // Use $1 for category if it's the first condition
			args = append(args, category)   // Add category to args slice
		}
	}
	fmt.Println("query ", query)
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Errorf("error in creating challenge %v", err)
		return []models.Challenge{}, err
	}

	var (
		challenges []models.Challenge
	)
	for rows.Next() {
		var id int
		var challengeData []byte // Read JSONB as byte slice

		err = rows.Scan(&id, &challengeData)
		if err != nil {
			fmt.Errorf("error scanning challenge: %v", err)
			return nil, fmt.Errorf("error scanning challenge: %v", err)
		}

		var challenge models.Challenge
		challenge.Id = id // Set the ID from the database

		// Unmarshal JSONB data into the Challenge struct
		err = json.Unmarshal(challengeData, &challenge)
		if err != nil {
			fmt.Errorf("error unmarshaling challenge JSON: %v", err)
			return nil, fmt.Errorf("error unmarshaling challenge JSON: %v", err)
		}

		// Append the challenge to the slice
		challenges = append(challenges, challenge)
	}
	return challenges, nil
}
