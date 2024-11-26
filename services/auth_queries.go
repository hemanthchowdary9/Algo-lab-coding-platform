package services

import (
	"coding-platform/commons"
	"coding-platform/database"
	"coding-platform/models"
	"fmt"
)

func InsertUser(user models.User) error {
	db := database.GetPsqlConnection()
	_, err := db.Exec("INSERT INTO coder(username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		fmt.Errorf("error in creating user %v", err)
		return err
	}
	var submissions []models.Submission
	submissionsJson := commons.ToByteArray(submissions)

	_, err = db.Exec("INSERT INTO user_submission(username, submissions) VALUES ($1, $2)", user.Username, submissionsJson)
	if err != nil {
		fmt.Errorf("error in creating user %v", err)
		return err
	}
	return nil
}

func FetchUserPassword(username string) (string, string, error) {
	db := database.GetPsqlConnection()
	rows, err := db.Query("SELECT password, role from coder where username=$1", username)
	if err != nil {
		fmt.Errorf("error in fetching user %v", err)
		return "", "", err
	}
	password, role := "", "NORMAL"
	for rows.Next() {
		err = rows.Scan(&password, &role)
		if err != nil {
			fmt.Errorf("error in scanning user %v", err)
			return "", "", err
		}
	}
	return password, role, nil
}

func FetchUserEmail(username string) (string, error) {
	db := database.GetPsqlConnection()
	rows, err := db.Query("SELECT email from coder where username=$1", username)
	if err != nil {
		fmt.Errorf("error in fetching user %v", err)
		return "", err
	}
	var email string
	for rows.Next() {
		err = rows.Scan(&email)
		if err != nil {
			fmt.Errorf("error in scanning user %v", err)
			return "", err
		}
	}
	return email, nil
}

func UpdateUserRole(username, role string) error {
	db := database.GetPsqlConnection()
	_, err := db.Exec("UPDATE coder SET role = $1 where username=$2", role, username)
	if err != nil {
		fmt.Errorf("error in fetching user %v", err)
		return err
	}
	return nil
}
