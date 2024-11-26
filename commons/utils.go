package commons

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func ToByteArray(unMarshalledData interface{}) []byte {
	marshalledData, err := json.Marshal(unMarshalledData)
	if err != nil {
		fmt.Errorf("error in marshalling %v", err)
		return nil
	}
	return marshalledData
}

func WriteResponse(w http.ResponseWriter, code int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(data)
	if err != nil {
		fmt.Errorf("error while writing response %v", err)
		return
	}
}

func GetTemplate(htmlFile string) *template.Template {

	return template.Must(template.New(htmlFile).Funcs(template.FuncMap{
		"getFirstName": getFirstName, // Register the contains function
		"isPremium":    isPremium,
	}).ParseFiles(filepath.Join("templates", htmlFile)))
}

// Function to extract the first part of the username before any $
func getFirstName(username string) string {
	parts := strings.Split(username, "$")
	if len(parts) > 0 {
		return parts[0] // Return the first part (before $)
	}
	return username // Fallback to the original if no $ found
}

// Function to check if the username contains any premium segment
func isPremium(username string) bool {
	return strings.Contains(username, "$PREMIUM") // Adjust this check based on your logic
}
