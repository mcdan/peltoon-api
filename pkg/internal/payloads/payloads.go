package payloads

import (
	"encoding/json"
	"strings"
	"time"
)

type LoginRequestBody struct {
	Username string `json:"username_or_email"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
}

const addClassToScheduleQuery = `mutation AddClassToSchedule($id: ID!, $startTime: DateTime!) {\n  addClassToSchedule(scheduledClassInput: {id: $id, scheduledStartTime: $startTime}) {\n    ... on ScheduledClass {\n      ...ScheduledClassDetails\n      ...ScheduledUsers\n      __typename\n    }\n    ... on ScheduledClassDoesNotExistError {\n      code\n      message\n      __typename\n    }\n    ... on CannotScheduleClassError {\n      code\n      message\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ScheduledClassDetails on ScheduledClass {\n  id\n  isScheduled\n  joinToken\n  pelotonId\n  pelotonClass {\n    classId\n    description\n    fitnessDiscipline {\n      ...FitnessDiscipline\n      __typename\n    }\n    id\n    originLocale {\n      code\n      __typename\n    }\n    title\n    ... on LiveClass {\n      airTime\n      actualStartTime\n      difficultyLevel {\n        ...DifficultyLevel\n        __typename\n      }\n      instructor {\n        ...Instructor\n        __typename\n      }\n      liveClassCategory\n      explicitRating\n      __typename\n    }\n    ... on OnDemandInstructorClass {\n      airTime\n      difficultyLevel {\n        ...DifficultyLevel\n        __typename\n      }\n      instructor {\n        ...Instructor\n        __typename\n      }\n      explicitRating\n      __typename\n    }\n    contentAvailability\n    contentAvailabilityLevel\n    isLimitedRide\n    freeForLimitedTime\n    __typename\n  }\n  scheduleSource\n  scheduledStartTime\n  scheduledEndTime\n  totalUsersCountedIn\n  __typename\n}\n\nfragment FitnessDiscipline on FitnessDiscipline {\n  displayName\n  slug\n  __typename\n}\n\nfragment DifficultyLevel on DifficultyLevel {\n  displayName\n  slug\n  __typename\n}\n\nfragment Instructor on Instructor {\n  assets {\n    profileImage {\n      location\n      __typename\n    }\n    __typename\n  }\n  name\n  __typename\n}\n\nfragment ScheduledUsers on ScheduledClass {\n  scheduledUsers {\n    totalCount\n    edges {\n      ...ScheduledUser\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment ScheduledUser on UserEdge {\n  node {\n    id\n    assets {\n      image {\n        location\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n`
const yourScheduleItemsQuery = `query YourScheduleItems($startTime: DateTime!, $endTime: DateTime!, $limit: Int = 7) {  userScheduledItemsList(startTime: $startTime, endTime: $endTime) {    ...ScheduleItemsList    __typename  }}fragment FitnessDiscipline on FitnessDiscipline {  displayName  slug  __typename}fragment DifficultyLevel on DifficultyLevel {  displayName  slug  __typename}fragment Instructor on Instructor {  assets {    profileImage {      location      __typename    }    __typename  }  name  __typename}fragment ScheduledClassDetails on ScheduledClass {  id  isScheduled  joinToken  pelotonId  pelotonClass {    classId    description    fitnessDiscipline {      ...FitnessDiscipline      __typename    }    id    originLocale {      code      __typename    }    title    ... on LiveClass {      airTime      actualStartTime      difficultyLevel {        ...DifficultyLevel        __typename      }      instructor {        ...Instructor        __typename      }      liveClassCategory      explicitRating      __typename    }    ... on OnDemandInstructorClass {      airTime      difficultyLevel {        ...DifficultyLevel        __typename      }      instructor {        ...Instructor        __typename      }      explicitRating      __typename    }    contentAvailability    contentAvailabilityLevel    isLimitedRide    freeForLimitedTime    __typename  }  scheduleSource  scheduledStartTime  scheduledEndTime  totalUsersCountedIn  __typename}fragment InviteUser on InviteUser {  invitedUserStatus  invitedUser {    id    username    getInvitedUserStatus    assets {      image {        location        __typename      }      __typename    }    __typename  }  __typename}fragment ScheduledInviteDetails on InviteDetails {  authedInvitedUser {    invitedUserStatus    invitorUser {      username      assets {        image {          location          __typename        }        __typename      }      __typename    }    __typename  }  getInvitedUsers(limit: $limit, cursor: "") {    ... on InviteFriendsError {      message      __typename    }    ... on InvitedUsersPaginatedList {      totalCount      edges {        inviteUserItem {          ...InviteUser          __typename        }        __typename      }      __typename    }    __typename  }  createdAt  hostUser {    id    username    assets {      image {        location        __typename      }      __typename    }    __typename  }  inviteId  inviteStatus  occasion {    fitnessDiscipline    occasionSlug    assets {      smallImage {        location        __typename      }      __typename    }    translatedOccasionName    __typename  }  visibility  scheduledClassId  __typename}fragment InviteDetailsWithClassDetails on InviteDetails {  ...ScheduledInviteDetails  scheduledClass {    ...ScheduledClassDetails    __typename  }  __typename}fragment PendingInvites on YourScheduleItemList {  pendingInvites {    ... on PendingInvites {      countOfPendingInvites      __typename    }    ... on InviteFriendsError {      message      __typename    }    __typename  }  __typename}fragment ScheduleItemsList on YourScheduleItemList {  yourScheduleItems {    ...ScheduledClassDetails    ...InviteDetailsWithClassDetails    ... on InviteFriendsError {      message      __typename    }    ... on InviteFriendsHttpException {      message      __typename    }    ... on InviteFriendsNotAuthorized {      message      __typename    }    ... on InviteFriendsInvalidDataFormat {      message      __typename    }    ... on InviteFriendsItemNotFoundError {      message      __typename    }    __typename  }  ...PendingInvites  __typename}`

const removeClassFromScheduleQuery = `mutation RemoveClassFromSchedule($id: String!) {\n  removeClassFromSchedule(scheduledClassId: $id) {\n    ... on ScheduledClass {\n      id\n      __typename\n    }\n    ... on ScheduledClassDoesNotExistError {\n      code\n      message\n      __typename\n    }\n    ... on CannotScheduleClassError {\n      code\n      message\n      __typename\n    }\n    __typename\n  }\n}`

type yourScheduleItemsVariables struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Limit     int    `json:"limit"`
}

type removeClassVariables struct {
	Id string `json:"id"`
}
type GraphOperation struct {
	OperationName string      `json:"operationName"`
	Variables     interface{} `json:"variables"`
	Query         string      `json:"query"`
}

func GetInvite(token string, startTime time.Time) (string, error) {
	queryVariables := map[string]string{
		"id":        token,
		"startTime": startTime.Format("Mon Jan 02 2006 15:04:05 GMT-0700"),
	}
	return toJsonString(GraphOperation{
		OperationName: "AddClassToSchedule",
		Variables:     queryVariables,
		Query:         addClassToScheduleQuery,
	})
}

// GetSchedule Times will be converted to UTC
func GetSchedule(startTime time.Time, endTime time.Time, limit int) (string, error) {
	// 2023-10-16T03:59:59.000+00:00
	formatString := "2006-01-02T15:04:05.999+00:00"
	utcStart := startTime.UTC()
	utcEnd := endTime.UTC()

	return toJsonString(GraphOperation{
		OperationName: "YourScheduleItems",
		Variables: yourScheduleItemsVariables{
			StartTime: utcStart.Format(formatString),
			EndTime:   utcEnd.Format(formatString),
			Limit:     limit,
		},
		Query: yourScheduleItemsQuery,
	})
}

func RemoveScheduledClass(joinToken string) (string, error) {
	return toJsonString(GraphOperation{
		OperationName: "RemoveClassFromSchedule",
		Variables: removeClassVariables{
			Id: joinToken,
		},
		Query: removeClassFromScheduleQuery,
	})

}

func toJsonString(operation GraphOperation) (string, error) {
	inviteBody, err := json.Marshal(operation)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(inviteBody), `\\n`, `\n`), nil
}
