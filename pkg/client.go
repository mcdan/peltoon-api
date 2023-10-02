package pkg

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/mcdan/peltoon-api/pkg/dto"
	"github.com/mcdan/peltoon-api/pkg/internal/payloads"
	"golang.org/x/net/publicsuffix"
)

type Client struct {
	username   *string
	password   *string
	httpClient *http.Client
	userID     *string
	sessionID  *string
}

// https://github.com/justmedude/pylotoncycle/blob/main/pylotoncycle/pylotoncycle.py
const apiURL = "https://api.onepeloton.com"
const graphURL = "https://gql-graphql-gateway.prod.k8s.onepeloton.com/graphql"

// First token is class id
// Second token is the invite's peloton home id
// third is the join token.
const inviteURLTemplate = "https://members.onepeloton.com/scheduled/class/%s/%s?join_token=%s&start=0&type=0&locale=en-US"

func InitializeClient(c *Client) {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c.httpClient = &http.Client{
		Timeout: 30 * time.Second,
		Jar:     jar,
	}
}

func (c Client) Login(username string, password string) error {
	loginPayload := payloads.LoginRequestBody{
		Username: username,
		Password: password,
	}
	loginBody, err := json.Marshal(loginPayload)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(loginBody)
	loginRequest, err := createRequest(http.MethodPost, fmt.Sprintf("%s/auth/login", apiURL), bodyReader)
	if err != nil {
		return err
	}

	loginResponse, err := c.httpClient.Do(loginRequest)
	if err != nil {
		return err
	}
	responseBytes, err := io.ReadAll(loginResponse.Body)
	if err != nil {
		return err
	}
	if loginResponse.StatusCode != 200 {
		return fmt.Errorf("unable to login: %s", responseBytes)
	}
	structResponse := payloads.LoginResponseBody{}
	err = json.Unmarshal(responseBytes, &structResponse)
	if err != nil {
		return err
	}
	c.userID = &structResponse.UserID
	c.sessionID = &structResponse.SessionID
	return nil
}

type scheduleJWT struct {
	HomePelotonID *string `json:"home_peloton_id"`
	RideID        string  `json:"ride_id"`
	// StartEpoch      int64   `json:"scheduled_start_time"`
	StudioPelotonID *string `json:"studio_peloton_id"`
	Type            string  `json:"type"`
}

func (c Client) GenerateInviteLink(classID string, startTime time.Time) (string, error) {
	scheduleToken := scheduleJWT{
		HomePelotonID:   nil,
		RideID:          classID,
		StudioPelotonID: nil,
		Type:            "on_demand",
	}

	scheduleTokenBytes, err := json.Marshal(scheduleToken)
	if err != nil {
		return "", err
	}
	easternLocation, err := time.LoadLocation("America/New_York")
	if err != nil {
		return "", err
	}
	easternStartTime := startTime.In(easternLocation)
	encodedToken := base64.StdEncoding.EncodeToString(scheduleTokenBytes)

	invitePayload, err := payloads.GetInvite(encodedToken, easternStartTime)
	if err != nil {
		return "", err
	}
	bodyReader := bytes.NewReader([]byte(invitePayload))

	createInviteReq, err := createRequest(http.MethodPost, graphURL, bodyReader)
	if err != nil {
		return "", err
	}
	inviteResponse, err := c.httpClient.Do(createInviteReq)
	if err != nil {
		return "", err
	}
	if inviteResponse.StatusCode != http.StatusOK {
		errorBody, err := io.ReadAll(inviteResponse.Body)
		if err == nil {
			return "", fmt.Errorf("http status code %d:%s\n%s", inviteResponse.StatusCode, inviteResponse.Status, errorBody)
		}
		return "", fmt.Errorf("http status code %d:%s", inviteResponse.StatusCode, inviteResponse.Status)

	}
	responseBytes, err := io.ReadAll(inviteResponse.Body)
	if err != nil {
		return "", err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return "", err
	}
	dataPayload := result["data"].(map[string]interface{})
	methodResponse := dataPayload["addClassToSchedule"].(map[string]interface{})
	joinToken := methodResponse["id"].(string)
	pelotonID := methodResponse["pelotonId"].(string)
	return renderJoinUrl(classID, pelotonID, joinToken), nil
}

func (c Client) GetRide(id string, ride *dto.RideDetails) error {
	getClassRequest, err := createRequest(http.MethodGet, fmt.Sprintf("%s/api/ride/%s/details", apiURL, id), nil)
	if err != nil {
		return err
	}
	getClassResponse, err := c.httpClient.Do(getClassRequest)
	if err != nil {
		return err
	}
	if getClassResponse.StatusCode != 200 {
		return fmt.Errorf("could not get class %d", getClassResponse.StatusCode)
	}

	responseBytes, err := io.ReadAll(getClassResponse.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseBytes, &ride)
	if err != nil {
		return err
	}
	return nil
}
func (c Client) GetScheduled(startTime time.Time, endTime time.Time, limit int) ([]*dto.YourScheduleItem, error) {
	scheduleListPayload, err := payloads.GetSchedule(startTime, endTime, limit)
	if err != nil {
		return []*dto.YourScheduleItem{}, err
	}
	bodyReader := bytes.NewReader([]byte(scheduleListPayload))
	getScheduleRequest, err := createRequest(http.MethodPost, fmt.Sprintf("%s/graphql?query=YourScheduleItem", graphURL), bodyReader)
	if err != nil {
		return []*dto.YourScheduleItem{}, err
	}
	getClassResponse, err := c.httpClient.Do(getScheduleRequest)
	if err != nil {
		return []*dto.YourScheduleItem{}, err
	}
	responseBytes, err := io.ReadAll(getClassResponse.Body)
	if err != nil {
		return []*dto.YourScheduleItem{}, err
	}
	if getClassResponse.StatusCode != 200 {
		return []*dto.YourScheduleItem{}, fmt.Errorf("could not get scheduled classes for user: %s", responseBytes)
	}
	structuredResponse := dto.YourScheduledItemsResponse{}
	err = json.Unmarshal(responseBytes, &structuredResponse)
	if err != nil {
		return []*dto.YourScheduleItem{}, err
	}

	if len(structuredResponse.Data.UserScheduledItemsList.YourScheduleItems) <= 0 {
		return []*dto.YourScheduleItem{}, err
	}
	scheduledClasses := make([]*dto.YourScheduleItem, len(structuredResponse.Data.UserScheduledItemsList.YourScheduleItems))
	for i := range structuredResponse.Data.UserScheduledItemsList.YourScheduleItems {
		item := structuredResponse.Data.UserScheduledItemsList.YourScheduleItems[i]
		item.JoinURL = renderJoinUrl(item.PelotonClass.ClassId, item.PelotonId, item.JoinToken)
		scheduledClasses[i] = &item
	}
	return scheduledClasses, nil
}

func (c Client) DeleteScheduledClass(joinToken string) error {
	scheduleListPayload, err := payloads.RemoveScheduledClass(joinToken)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader([]byte(scheduleListPayload))
	removeScheduledRequest, err := createRequest(http.MethodPost, fmt.Sprintf("%s/graphql", graphURL), bodyReader)
	getClassResponse, err := c.httpClient.Do(removeScheduledRequest)
	if err != nil {
		return err
	}
	responseBytes, err := io.ReadAll(getClassResponse.Body)
	if err != nil {
		return err
	}
	if getClassResponse.StatusCode != 200 {
		return fmt.Errorf("could not get scheduled classes for user: %s", responseBytes)
	}

	return nil
}
func createRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "golang-peltoon-api")
	return req, err
}

func renderJoinUrl(classID string, pelotonID string, joinToken string) string {
	return fmt.Sprintf(inviteURLTemplate, classID, pelotonID, joinToken)
}
