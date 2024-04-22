package tests

import (
	"DebTour/controllers"
	"DebTour/database"
	"DebTour/models"
	"bytes"
	"encoding/json"
	"fmt"
	"gorm.io/gorm/utils/tests"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestReportProblem(t *testing.T) {
	type inputReportProblem struct {
		Message string `json:"message"`
		Image  string `json:"image"`
		Validation string `json:"validation"`
	}
	var testcases []inputReportProblem

// open file reportProblem.json in input folder
	jsonFile, err := os.Open("input/reportProblem.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened reportProblem.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &testcases)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := SetupRouter()

	// login
	router.POST("/api/v1/auth/login", controllers.Login)
	router.POST("/api/v1/reportProblem", controllers.CreateIssueReport)

	loginReqBody := models.FirstContactModel{
		Id: "sudo",
		Token: "123",
	}
	loginReqBodyJson, _ := json.Marshal(loginReqBody)
	loginReq, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(loginReqBodyJson))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRes := httptest.NewRecorder()
	router.ServeHTTP(loginRes, loginReq)
	var loginResponse map[string]interface{}
	err = json.Unmarshal([]byte(loginRes.Body.String()), &loginResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	token := loginResponse["token"]
	fmt.Println("Token:", token)

	for i, testcase := range testcases {
		fmt.Println("Test case number", i)
		t.Run("reportProblem" + fmt.Sprint(i), func(t *testing.T) {
			reportProblemReqBody := models.Issue{
				Message: testcase.Message,
				Image: testcase.Image,
				IssueType: "Problem",
				Status: "Open",
			}
			reportProblemReqBodyJson, _ := json.Marshal(reportProblemReqBody)
			reportProblemReq, _ := http.NewRequest("POST", "/api/v1/reportProblem", bytes.NewBuffer(reportProblemReqBodyJson))
			reportProblemReq.Header.Set("Content-Type", "application/json")
			reportProblemReq.Header.Set("Authorization", "Bearer " + token.(string))
			reportProblemRes := httptest.NewRecorder()
			router.ServeHTTP(reportProblemRes, reportProblemReq)

			var reportProblemResponse map[string]interface{}
			err = json.Unmarshal([]byte(reportProblemRes.Body.String()), &reportProblemResponse)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(reportProblemResponse)
			if testcase.Validation == "valid" {
				tests.AssertEqual(t, reportProblemRes.Code, http.StatusOK)
			} else {
				// assert response success is false
				tests.AssertEqual(t, reportProblemResponse["success"], false)
			}
		})
	}

	// clean up
	database.MainDB.Exec("DELETE FROM issues")
}