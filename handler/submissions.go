package handler

import (
	"coding-platform/commons"
	"coding-platform/models"
	"coding-platform/services"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func SaveSubmission(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string) // Retrieve username from context
	username = strings.Split(username, "$")[0]
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "dashboard.html")))
	data := struct {
		Username string
	}{
		Username: username,
	}

	tmpl.Execute(w, data)
}

func FetchSubmissions(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string) // Retrieve username from context
	usernameWithoutRole := strings.Split(username, "$")[0]
	tmpl := commons.GetTemplate("submission.html")

	//tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "submission.html")))
	submissions := []models.Submission{
		{
			ProblemTitle: "ABC oialkval oikaf fa oqifomexqoiefiuq3oq2p[4jroq24iojq oqmq2[pcqlkejcnoq204tkpoq24krmocqj4orqj2o oitcmoqojrcehrigj[oqrj[gwor ow3jm0wrotcwrcl;tjmopw3i",
			RunTime:      "1s",
			Time:         time.Now().UTC(),
			Language:     "Java",
			Status:       "Passed",
			Code:         "import java.utils.* \n public class Main{ \n public static void main(args []String){\n\n}\n}",
		}, {
			ProblemTitle: "ABC",
			RunTime:      "1s",
			Time:         time.Now().UTC(),
			Language:     "Java",
			Status:       "Failed",
			Code:         "import java.utils.* \n public class Main{ \n public static void main(args []String){\n\n}\n}",
		}, {
			ProblemTitle: "ABC",
			RunTime:      "1s",
			Time:         time.Now().UTC(),
			Language:     "Java",
			Status:       "Partial",
			Code:         "import java.utils.* \n public class Main{ \n public static void main(args []String){\n\n}\n}",
		}, {
			ProblemTitle: "ABC",
			RunTime:      "1s",
			Time:         time.Now().UTC(),
			Language:     "Java",
			Status:       "Passed",
			Code:         "import java.utils.* \n public class Main{ \n public static void main(args []String){\n\n}\n}",
		},
	}
	var err error = nil
	submissions, err = services.FetchSubmissions(usernameWithoutRole)
	fmt.Println("Submissions fetched", submissions)
	if err == nil {

		data := struct {
			Username    string
			Submissions []models.Submission
		}{
			Username:    username,
			Submissions: submissions,
		}

		fmt.Println("data for submissions \n", data)
		tmpl.Execute(w, data)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
