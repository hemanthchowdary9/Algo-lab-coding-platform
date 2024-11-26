package models

type Challenge struct {
	Id          int        `json:"id,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Difficulty  string     `json:"difficulty"`
	Category    string     `json:"category"`
	Premium     string     `json:"premium"`
	Examples    []SampleIO `json:"examples"`
}

type SampleIO struct {
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation,omitempty"`
}

type ChallengePreview struct {
	ID         int
	Title      string
	Difficulty string
	Premium    bool
}

type User struct {
	Username string
	Email    string
	Password string
}

type CompileResponse struct {
	Input              string `json:"input,omitempty"`
	Error              string `json:"error,omitempty"`
	ExpectedOutput     string `json:"expectedOutput,omitempty"`
	OutPut             string `json:"output,omitempty"`
	CpuTime            string `json:"cpuTime,omitempty"`
	IsExecutionSuccess bool   `json:"isExecutionSuccess,omitempty"`
	IsCompiled         bool   `json:"isCompiled,omitempty"`
	IsTestCasePassed   bool   `json:"isTestCasePassed,omitempty"`
}
