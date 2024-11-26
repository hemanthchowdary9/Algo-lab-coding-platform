package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

type ChallengeDescription struct {
	Title       string
	Description string
	Difficulty  string
	Examples    []SampleIO
}

type SampleIO struct {
	Input       string
	Output      string
	Explanation string
}

type Challenge struct {
	ID         int
	Title      string
	Difficulty string
}

var (
	IdToChallengeMap map[string]ChallengeDescription
	Challenges       []Challenge
	Configurations   map[string]map[string]string
)

func init() {
	IdToChallengeMap = make(map[string]ChallengeDescription)
	IdToChallengeMap = map[string]ChallengeDescription{
		"0": {
			Title:       "Two Sum",
			Difficulty:  "Easy",
			Description: "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.",
			Examples: []SampleIO{
				{
					Input:       "nums = [2,7,11,15], target = 9",
					Output:      "[0,1]",
					Explanation: "Because nums[0] + nums[1] == 9, we return [0, 1].",
				}, {
					Input:  "nums = [3,2,4], target = 6",
					Output: "[1,2]",
				}, {
					Input:  "nums = [3,3], target = 6",
					Output: "[0,1]",
				},
			},
		},
		"1": {
			Title:       "Longest Substring Without Repeating Characters",
			Difficulty:  "Medium",
			Description: "Given a string s, find the length of the longest substring without repeating characters.",
			Examples: []SampleIO{
				{
					Input:       "s = \"abcabcbb\"",
					Output:      "3",
					Explanation: "The answer is \"abc\", with the length of 3.",
				}, {
					Input:       "s = \"bbbbb\"",
					Output:      "1",
					Explanation: "The answer is \"b\", with the length of 1.",
				}, {
					Input:       "s = \"pwwkew\"",
					Output:      "3",
					Explanation: "The answer is \"wke\", with the length of 3.\nNotice that the answer must be a substring, \"pwke\" is a subsequence and not a substring.",
				},
			},
		},
		"2": {
			Title:       "Valid Parentheses",
			Difficulty:  "Easy",
			Description: "Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.\n\nAn input string is valid if:\n\nOpen brackets must be closed by the same type of brackets.\nOpen brackets must be closed in the correct order.\nEvery close bracket has a corresponding open bracket of the same type.",
			Examples: []SampleIO{
				{
					Input:  "s = \"()\"",
					Output: "true",
				}, {
					Input:  "s = \"(){}[]\"",
					Output: "true",
				}, {
					Input:  "s = \"(){[}]\"",
					Output: "false",
				},
			},
		},
		"3": {
			Title:       "Median of Two Sorted Arrays",
			Difficulty:  "Hard",
			Description: "Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.\n\nThe overall run time complexity should be O(log (m+n)).",
			Examples: []SampleIO{
				{
					Input:       "nums1 = [1,3], nums2 = [2]",
					Output:      "2.000",
					Explanation: "merged array = [1,2,3] and median is 2.",
				}, {
					Input:       "nums1 = [1,2], nums2 = [3,4]",
					Output:      "2.500",
					Explanation: "merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.",
				},
			},
		},
	}

	var keys []string
	for key := range IdToChallengeMap {
		keys = append(keys, key)
	}

	// Sort the keys alphabetically
	sort.Strings(keys)
	for _, key := range keys {
		id, err := strconv.Atoi(key)
		if err != nil {
			log.Fatal("Error in converting key to integer index", err)
			return
		}
		value := IdToChallengeMap[key]
		Challenges = append(Challenges, Challenge{
			ID:         id,
			Title:      strconv.Itoa(id+1) + ". " + value.Title,
			Difficulty: value.Difficulty,
		})
	}

	Configurations = make(map[string]map[string]string)
}

func LoadYamlConfigurations() {
	yamlFile, err := ioutil.ReadFile("config/properties.yaml")
	if err != nil {
		log.Fatal("unable to read properties yaml config file", err)
	}
	yamlConfigs := make(map[string][]map[string]string)
	err = yaml.Unmarshal(yamlFile, &yamlConfigs)
	if err != nil {
		log.Fatal("unable to parse properties yaml config file", err)
	}
	parseYamlConfigs(yamlConfigs)
}

func parseYamlConfigs(yamlConfigs map[string][]map[string]string) {
	for configKey, configValuesList := range yamlConfigs {
		configValueMap := make(map[string]string)
		for _, configMap := range configValuesList {
			for key, value := range configMap {
				configValueMap[key] = value
			}
		}
		Configurations[configKey] = configValueMap
	}
	fmt.Println("Successfully loaded yaml configurations", Configurations)
}

func FetchDatabaseConfigs() map[string]string {
	if os.Getenv("DB_USER") != "" {
		dbDetails := map[string]string{
			"host":     os.Getenv("DB_HOST"),
			"port":     os.Getenv("DB_PORT"),
			"user":     os.Getenv("DB_USER"),
			"password": os.Getenv("DB_PASSWORD"),
			"name":     os.Getenv("DB_NAME"),
		}
		fmt.Println("Using environment db details", dbDetails)
		return dbDetails
	}
	return Configurations["database"]
}

// algolab logo design ui link : https://www.design.com/ai-logo-generator/results?businessDescription=AlgoLab&FilterByTags=&Colors=&text=AlgoLab&searchText=Artificial+Intelligence%2C+Machine+Learning%2C+Data+Analytics&customPrompt=&isFromAILogoGenerator=true
