package handler

import (
	"coding-platform/models"
	"coding-platform/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CompileRequest struct {
	Language    string `json:"language"`
	Code        string `json:"code"`
	ChallengeId string `json:"challengeId"`
}

func Compile(w http.ResponseWriter, r *http.Request) {

	var compileReq CompileRequest

	// Decode the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&compileReq)
	if err != nil {
		fmt.Errorf("error ", err)
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	cID, _ := strconv.Atoi(compileReq.ChallengeId)
	challenge, _ := services.FetchChallenges(cID, "")
	// Now you can use compileReq.Language and compileReq.Code
	fmt.Printf("Received body: %d\n", compileReq.ChallengeId)

	var compileRes []models.CompileResponse
	for _, input := range challenge[0].Examples {
		response, err := services.Execute(compileReq.Code, input.Input, compileReq.Language)
		if err == nil {
			response.ExpectedOutput = input.Output
			response.Input = input.Input
			if response.IsExecutionSuccess && strings.EqualFold(response.OutPut, input.Output) {
				response.IsTestCasePassed = true
			}
			compileRes = append(compileRes, response)
		}
	}
	// You can send a response back
	username := r.Context().Value("username").(string) // Retrieve username from context
	username = strings.Split(username, "$")[0]
	go TransformAndSaveSubmissions(username, challenge[0].Title, compileReq, compileRes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compileRes)
}

func CompileTest(w http.ResponseWriter, r *http.Request) {

	time.Sleep(5 * time.Second)
	response := []models.CompileResponse{
		{
			Input:              "2\n1 2",
			OutPut:             "\n# command-line-arguments\n./jdoodle.go:2: non-declaration statement outside function body\n./jdoodle.go:6: undefined: fmt in fmt.Scan\n./jdoodle.go:9: undefined: fmt in fmt.Scan\n./jdoodle.go:15: undefined: fmt in fmt.Println\nCommand exited with non-zero status 2",
			ExpectedOutput:     "5 5 5",
			CpuTime:            "0.1",
			IsExecutionSuccess: false,
			IsCompiled:         true,
			IsTestCasePassed:   false,
		}, {
			Input:              "2\n1 2",
			OutPut:             "5 5 5",
			ExpectedOutput:     "5 5 5",
			CpuTime:            "0.75",
			IsExecutionSuccess: true,
			IsCompiled:         true,
			IsTestCasePassed:   true,
		}, {
			Input:              "2\n1 2",
			OutPut:             "5 5 5",
			ExpectedOutput:     "5 5 5",
			CpuTime:            "0.75",
			IsExecutionSuccess: true,
			IsCompiled:         true,
			IsTestCasePassed:   true,
		}, {
			Input:              "2\n1 2",
			OutPut:             "5 5 5",
			ExpectedOutput:     "5 5 5",
			CpuTime:            "0.75",
			IsExecutionSuccess: true,
			IsCompiled:         true,
			IsTestCasePassed:   true,
		},
		{
			Input:              "2\n1 2",
			OutPut:             "5 5 5",
			ExpectedOutput:     "5 5 5",
			CpuTime:            "0.75",
			IsExecutionSuccess: true,
			IsCompiled:         true,
			IsTestCasePassed:   true,
		},
	}
	username := r.Context().Value("username").(string) // Retrieve username from context
	username = strings.Split(username, "$")[0]
	go TransformAndSaveSubmissions(username, "Sample_Title", CompileRequest{Code: "Sample_Code", Language: "Go"}, response)
	json.NewEncoder(w).Encode(response)
}

func TransformAndSaveSubmissions(username, title string, compileReq CompileRequest, response []models.CompileResponse) {
	var submissions []models.Submission
	var sub models.Submission
	sub.ProblemTitle = title
	sub.Code = compileReq.Code
	sub.Language = compileReq.Language
	sub.Error = response[0].Error
	sub.Time = time.Now()
	sub.RunTime = response[0].CpuTime
	sub.Attempts = 1
	allCasesPassed := true
	allCaseFailed := true
	for _, outCome := range response {
		if !outCome.IsTestCasePassed {
			allCasesPassed = false
		}
		if outCome.IsTestCasePassed {
			allCaseFailed = false
		}

	}
	if allCasesPassed {
		sub.Status = "Passed"
	} else if allCaseFailed {
		sub.Status = "Failed"
	} else {
		sub.Status = "Partial"
	}
	submissions = append(submissions, sub)
	fmt.Println("submissions ", submissions)
	err := services.UpdateSubmissions(username, submissions)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
}
