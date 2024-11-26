package services

import (
	"bytes"
	"coding-platform/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ClientIDs = []string{
	"d9ae345661f6659da0ec1f1e38abbd30",
	"2f6f63ea03c79a5d2d8c8c410d9ad256",
	"f28a454f42d209553cc256b72b02b87e",
	"6aa2858f151251a2333d85792417528f",
	"2cfbc30fc32deab1910f5c3c30fcb2c0",
	"26307392e64ef2176ccb5f6292c81f07",
	"a98400fa43eb63b92b8e3e81a23ba9a6",
	"561559d12a19e50d66553cc41effc306",
	"e7ae8708aa3bd1b92aca208b33af27c2",
	"dd772f9f30eef9d24145a142aac3ccda",
}

var ClientSecrets = []string{
	"466530a51c2036dbf8601b206f26d8409884d4ca9e04adb0025e8e8169513db0",
	"3cf9e1aace8b8d93c9fc73edb45cece287bec4954f8fd7f1f3ffde3bc636e682",
	"a39561fa9031652c9a511661d9366050ad580276abe67b926161e52ad862c1cc",
	"6967346dcab9b95671d17fc5a83debc7837110c899b192c6a4fe14b70d0c295c",
	"17ab2daa7715f1ba693d4abad02180c38c795c324f106ecec5accb61da294148",
	"4ac4e2e61b9c6c3d034a31d55711d82e3d9f126f0414805f57eaabb2304360d6",
	"ee31dc3faa7a4d618816779376b08f3bd4b9cb34f5a38ba76fd1e38eb7f2e2ed",
	"7aa80a09ca531c0447f2f0241270239d683550cd3cf7f5200d7b49d854b1c1eb",
	"8f1d6b8e00346db28b8ea32d32cd25fca94e87b1292c46b7baeec15e753ee0e",
	"42ce904fdb591255a1d64e38b3b94f1ca8cd70aec40ce8b2e462c5dda03aa3eb",
}

func Execute(code, stdin, language string) (models.CompileResponse, error) {
	// Replace with your actual JDoodle client ID and secret
	var (
		clientID, clientSecret string
		creditsAvailable       bool
		ind, end               = 0, len(ClientIDs)
	)

	for ind < end {
		clientID = ClientIDs[ind]
		clientSecret = ClientSecrets[ind]
		spent, err := CrediChecker(models.JDoodleRequest{ClientID: clientID, ClientSecret: clientSecret})
		if err != nil {
			return models.CompileResponse{}, err
		}
		if spent < 20 {
			fmt.Println("spent ", spent)
			creditsAvailable = true
			break
		} else {
			ind++
		}
	}
	var version string
	if language != "python3" {
		version = "0"
	}
	if creditsAvailable {
		script := code
		requestBody := models.JDoodleRequest{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Script:       script,
			Stdin:        stdin,
			Language:     language,
			VersionIndex: version,
		}
		fmt.Println("request body ", requestBody)

		// Marshal the request body to JSON
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return models.CompileResponse{}, err
		}

		// Send the request to JDoodle API
		response, err := http.Post("https://api.jdoodle.com/v1/execute", "application/json", bytes.NewBuffer(jsonBody))
		if err != nil {
			fmt.Println("Error sending request:", err)
			return models.CompileResponse{}, err
		}
		defer response.Body.Close()

		// Read and print the response
		var responseBody models.CompileResponse
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			fmt.Println("Error decoding response:", err)
			return models.CompileResponse{}, err
		}
		fmt.Println("response ", responseBody, "\n", response.Body)
		return responseBody, nil
	}
	return models.CompileResponse{}, errors.New("credits are not available to execute. Please try after sometime")
}

func CrediChecker(requestBody models.JDoodleRequest) (int, error) {

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return 0, err
	}

	// Send the request to JDoodle API
	response, err := http.Post("https://api.jdoodle.com/v1/credit-spent", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0, err
	}
	defer response.Body.Close()

	// Read and print the response
	var spent models.CreditSpent
	if err := json.NewDecoder(response.Body).Decode(&spent); err != nil {
		fmt.Println("Error decoding response:", err)
		return 0, err
	}
	return spent.Used, nil
}
