package tests

import (
	"DebTour/controllers"
	"DebTour/database"
	"DebTour/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils/tests"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// read input from file and create testcases for joinTour

func SetupRouter() *gin.Engine {
	database.InitDB()
	router := gin.Default()

	return router
}

func TestJoinTour(t *testing.T) {
	// input for joinTour
	type inputjoinTour struct {
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Age      int    `json:"age"`
		Expected string `json:"expected"`
	}
	var testcases []inputjoinTour

	// open file joinTour.json in input folder
	jsonFile, err := os.Open("input/joinTour.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened joinTour.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &testcases)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := SetupRouter()

	// login
	router.POST("/api/v1/auth/login", controllers.Login)
	router.POST("/api/v1/joinings", controllers.JoinTour)

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

	// create testcases
	for i, testcase := range testcases {
		fmt.Println("Test case number", i)
		fmt.Println(testcase)
		t.Run("joinTour" + fmt.Sprint(i), func(t *testing.T) {
			member := models.JoinedMembers{
				FirstName: testcase.Firstname,
				LastName:  testcase.Lastname,
				Age: uint(testcase.Age),
			}
			// create request body
			requestBody := models.JoinTourRequest{
				TourId:        1,
				JoinedMembers: []models.JoinedMembers{member},
			}
			reqBodyJson, _ := json.Marshal(requestBody)

			// create request
			req, _ := http.NewRequest("POST", "/api/v1/joinings", bytes.NewBuffer(reqBodyJson))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer " + token.(string))

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// log response
			fmt.Println(w.Body.String())

			// check response
			fmt.Println("Expected: " + testcase.Expected)
			if testcase.Expected == "valid" {
				tests.AssertEqual(t, http.StatusOK, w.Code)
			} else {
				// read success status from response
				var response map[string]interface{}
				err := json.Unmarshal([]byte(w.Body.String()), &response)
				if err != nil {
					return 
				}

				// check if success status is false
				tests.AssertEqual(t, false, response["success"])
				// show error message
				fmt.Println(response["error"])
			}
		})
	}
	// clean up
	database.MainDB.Exec("DELETE FROM joinings")
}