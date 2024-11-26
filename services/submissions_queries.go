package services

import (
	"coding-platform/commons"
	"coding-platform/database"
	"coding-platform/models"
	"encoding/json"
	"fmt"
)

func UpdateSubmissions(username string, newSubmissions []models.Submission) error {
	fmt.Println("updating submissions ", username)
	db := database.GetPsqlConnection()
	submissions, err := FetchSubmissions(username)
	if err != nil {
		return err
	}
	for _, newSubmission := range newSubmissions {
		foundSub := false
		for ind, sub := range submissions {
			if sub.Language == newSubmission.Language && sub.ProblemTitle == newSubmission.ProblemTitle {
				newSubmission.Attempts += sub.Attempts
				submissions[ind] = newSubmission
				foundSub = true
			}
		}
		if !foundSub {
			submissions = append(submissions, newSubmission)
		}
	}
	submissionsJson := commons.ToByteArray(submissions)

	fmt.Println("Inserting into submissions with latest data")
	_, err = db.Exec(`
    INSERT INTO user_submission (username, submissions) 
    VALUES ($1, $2)
    ON CONFLICT (username) DO UPDATE 
    SET submissions = EXCLUDED.submissions`,
		username, submissionsJson)
	if err != nil {
		fmt.Errorf("error in creating challenge %v", err)
		return err
	}
	return nil
}

func FetchSubmissions(username string) ([]models.Submission, error) {
	db := database.GetPsqlConnection()
	rows, err := db.Query(`select submissions from user_submission where username=$1`, username)
	if err != nil {
		fmt.Errorf("error in creating challenge %v", err)
		return []models.Submission{}, err
	}

	var (
		submissions []models.Submission
	)
	for rows.Next() {

		var submissionData []byte // Read JSONB as byte slice

		err = rows.Scan(&submissionData)
		if err != nil {
			fmt.Errorf("error scanning Submission: %v", err)
			return nil, fmt.Errorf("error scanning Submission: %v", err)
		}

		// Unmarshal JSONB data into the Challenge struct
		err = json.Unmarshal(submissionData, &submissions)
		if err != nil {
			fmt.Errorf("error unmarshaling Submission JSON: %v", err)
			return nil, fmt.Errorf("error unmarshaling submission JSON: %v", err)
		}
	}
	return submissions, nil
}
