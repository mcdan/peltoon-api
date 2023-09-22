package payloads

type LoginRequestBody struct {
	Username string `json:"username_or_email"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
}

const AddClassToSchedule = `mutation AddClassToSchedule($id: ID!, $startTime: DateTime!) {\n  addClassToSchedule(scheduledClassInput: {id: $id, scheduledStartTime: $startTime}) {\n    ... on ScheduledClass {\n      ...ScheduledClassDetails\n      ...ScheduledUsers\n      __typename\n    }\n    ... on ScheduledClassDoesNotExistError {\n      code\n      message\n      __typename\n    }\n    ... on CannotScheduleClassError {\n      code\n      message\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ScheduledClassDetails on ScheduledClass {\n  id\n  isScheduled\n  joinToken\n  pelotonId\n  pelotonClass {\n    classId\n    description\n    fitnessDiscipline {\n      ...FitnessDiscipline\n      __typename\n    }\n    id\n    originLocale {\n      code\n      __typename\n    }\n    title\n    ... on LiveClass {\n      airTime\n      actualStartTime\n      difficultyLevel {\n        ...DifficultyLevel\n        __typename\n      }\n      instructor {\n        ...Instructor\n        __typename\n      }\n      liveClassCategory\n      explicitRating\n      __typename\n    }\n    ... on OnDemandInstructorClass {\n      airTime\n      difficultyLevel {\n        ...DifficultyLevel\n        __typename\n      }\n      instructor {\n        ...Instructor\n        __typename\n      }\n      explicitRating\n      __typename\n    }\n    contentAvailability\n    contentAvailabilityLevel\n    isLimitedRide\n    freeForLimitedTime\n    __typename\n  }\n  scheduleSource\n  scheduledStartTime\n  scheduledEndTime\n  totalUsersCountedIn\n  __typename\n}\n\nfragment FitnessDiscipline on FitnessDiscipline {\n  displayName\n  slug\n  __typename\n}\n\nfragment DifficultyLevel on DifficultyLevel {\n  displayName\n  slug\n  __typename\n}\n\nfragment Instructor on Instructor {\n  assets {\n    profileImage {\n      location\n      __typename\n    }\n    __typename\n  }\n  name\n  __typename\n}\n\nfragment ScheduledUsers on ScheduledClass {\n  scheduledUsers {\n    totalCount\n    edges {\n      ...ScheduledUser\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment ScheduledUser on UserEdge {\n  node {\n    id\n    assets {\n      image {\n        location\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n`

type GraphOperation struct {
	OperationName string            `json:"operationName"`
	Variables     map[string]string `json:"variables"`
	Query         string            `json:"query"`
}
