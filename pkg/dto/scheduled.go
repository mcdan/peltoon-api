package dto

import "time"

type YourScheduleItem struct {
	Id           string `json:"id"`
	IsScheduled  bool   `json:"isScheduled"`
	JoinToken    string `json:"joinToken"`
	PelotonId    string `json:"pelotonId"`
	JoinURL      string
	PelotonClass struct {
		ClassId           string `json:"classId"`
		Description       string `json:"description"`
		FitnessDiscipline struct {
			DisplayName string `json:"displayName"`
			Slug        string `json:"slug"`
			Typename    string `json:"__typename"`
		} `json:"fitnessDiscipline"`
		Id           string `json:"id"`
		OriginLocale struct {
			Code     string `json:"code"`
			Typename string `json:"__typename"`
		} `json:"originLocale"`
		Title           string      `json:"title"`
		AirTime         time.Time   `json:"airTime"`
		DifficultyLevel interface{} `json:"difficultyLevel"`
		Instructor      struct {
			Assets struct {
				ProfileImage struct {
					Location string `json:"location"`
					Typename string `json:"__typename"`
				} `json:"profileImage"`
				Typename string `json:"__typename"`
			} `json:"assets"`
			Name     string `json:"name"`
			Typename string `json:"__typename"`
		} `json:"instructor"`
		ExplicitRating           int    `json:"explicitRating"`
		Typename                 string `json:"__typename"`
		ContentAvailability      string `json:"contentAvailability"`
		ContentAvailabilityLevel string `json:"contentAvailabilityLevel"`
		IsLimitedRide            bool   `json:"isLimitedRide"`
		FreeForLimitedTime       bool   `json:"freeForLimitedTime"`
	} `json:"pelotonClass"`
	ScheduleSource      string    `json:"scheduleSource"`
	ScheduledStartTime  time.Time `json:"scheduledStartTime"`
	ScheduledEndTime    time.Time `json:"scheduledEndTime"`
	TotalUsersCountedIn int       `json:"totalUsersCountedIn"`
	Typename            string    `json:"__typename"`
}

type YourScheduledItemsResponse struct {
	Data struct {
		UserScheduledItemsList struct {
			YourScheduleItems []YourScheduleItem
		} `json:"userScheduledItemsList"`
	} `json:"data"`
}
