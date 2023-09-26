package dto

type Instructor struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Ride struct {
	ClassTypeIDs []string   `json:"class_type_ids"`
	Description  string     `json:"description"`
	HomeID       string     `json:"home_peloton_id"`
	ID           string     `json:"id"`
	StudioID     string     `json:"studio_peloton_id"`
	Title        string     `json:"title"`
	Instructor   Instructor `json:"instructor"`
	ImageUrl     string     `json:"image_url"`
}
type RideDetails struct {
	Ride Ride `json:"ride"`
}

// type T struct {
// 	Ride struct {
// 		FreeForLimitedTime       bool   `json:"free_for_limited_time"`
// 		ContentAvailability      string `json:"content_availability"`
// 		ContentAvailabilityLevel string `json:"content_availability_level"`
// 		IsLimitedRide            bool   `json:"is_limited_ride"`
// 		Availability             struct {
// 			IsAvailable bool        `json:"is_available"`
// 			Reason      interface{} `json:"reason"`
// 		} `json:"availability"`
// 		ClassTypeIds                 []string      `json:"class_type_ids"`
// 		ContentProvider              string        `json:"content_provider"`
// 		ContentFormat                string        `json:"content_format"`
// 		Description                  string        `json:"description"`
// 		DifficultyEstimate           float64       `json:"difficulty_estimate"`
// 		OverallEstimate              float64       `json:"overall_estimate"`
// 		DifficultyRatingAvg          float64       `json:"difficulty_rating_avg"`
// 		DifficultyRatingCount        int           `json:"difficulty_rating_count"`
// 		DifficultyLevel              interface{}   `json:"difficulty_level"`
// 		Duration                     int           `json:"duration"`
// 		EquipmentIds                 []interface{} `json:"equipment_ids"`
// 		EquipmentTags                []interface{} `json:"equipment_tags"`
// 		ExplicitRating               int           `json:"explicit_rating"`
// 		ExtraImages                  []interface{} `json:"extra_images"`
// 		FitnessDiscipline            string        `json:"fitness_discipline"`
// 		FitnessDisciplineDisplayName string        `json:"fitness_discipline_display_name"`
// 		HasClosedCaptions            bool          `json:"has_closed_captions"`
// 		HasPedalingMetrics           bool          `json:"has_pedaling_metrics"`
// 		HomePelotonId                string        `json:"home_peloton_id"`
// 		Id                           string        `json:"id"`
// 		ImageUrl                     string        `json:"image_url"`
// 		InstructorId                 string        `json:"instructor_id"`
// 		IndividualInstructorIds      []interface{} `json:"individual_instructor_ids"`
// 		IsArchived                   bool          `json:"is_archived"`
// 		IsClosedCaptionShown         bool          `json:"is_closed_caption_shown"`
// 		IsExplicit                   bool          `json:"is_explicit"`
// 		HasFreeMode                  bool          `json:"has_free_mode"`
// 		IsLiveInStudioOnly           bool          `json:"is_live_in_studio_only"`
// 		Language                     string        `json:"language"`
// 		OriginLocale                 string        `json:"origin_locale"`
// 		Length                       int           `json:"length"`
// 		LiveStreamId                 string        `json:"live_stream_id"`
// 		LiveStreamUrl                interface{}   `json:"live_stream_url"`
// 		Location                     string        `json:"location"`
// 		Metrics                      []string      `json:"metrics"`
// 		OriginalAirTime              int           `json:"original_air_time"`
// 		OverallRatingAvg             float64       `json:"overall_rating_avg"`
// 		OverallRatingCount           int           `json:"overall_rating_count"`
// 		PedalingStartOffset          int           `json:"pedaling_start_offset"`
// 		PedalingEndOffset            int           `json:"pedaling_end_offset"`
// 		PedalingDuration             int           `json:"pedaling_duration"`
// 		Rating                       int           `json:"rating"`
// 		RideTypeId                   string        `json:"ride_type_id"`
// 		RideTypeIds                  []string      `json:"ride_type_ids"`
// 		SampleVodStreamUrl           interface{}   `json:"sample_vod_stream_url"`
// 		SamplePreviewStreamUrl       interface{}   `json:"sample_preview_stream_url"`
// 		ScheduledStartTime           int           `json:"scheduled_start_time"`
// 		SeriesId                     string        `json:"series_id"`
// 		SoldOut                      bool          `json:"sold_out"`
// 		StudioPelotonId              string        `json:"studio_peloton_id"`
// 		Title                        string        `json:"title"`
// 		TotalRatings                 int           `json:"total_ratings"`
// 		TotalInProgressWorkouts      int           `json:"total_in_progress_workouts"`
// 		TotalWorkouts                int           `json:"total_workouts"`
// 		VodStreamUrl                 string        `json:"vod_stream_url"`
// 		VodStreamId                  string        `json:"vod_stream_id"`
// 		Captions                     []string      `json:"captions"`
// 		JoinTokens                   struct {
// 			OnDemand string `json:"on_demand"`
// 		} `json:"join_tokens"`
// 		Flags                          []interface{} `json:"flags"`
// 		IsDynamicVideoEligible         bool          `json:"is_dynamic_video_eligible"`
// 		IsFixedDistance                bool          `json:"is_fixed_distance"`
// 		DynamicVideoRecordedSpeedInMph int           `json:"dynamic_video_recorded_speed_in_mph"`
// 		ThumbnailTitle                 interface{}   `json:"thumbnail_title"`
// 		ThumbnailLocation              interface{}   `json:"thumbnail_location"`
// 		Distance                       interface{}   `json:"distance"`
// 		DistanceUnit                   interface{}   `json:"distance_unit"`
// 		DistanceDisplayValue           interface{}   `json:"distance_display_value"`
// 		IsOutdoor                      bool          `json:"is_outdoor"`
// 		Instructor                     struct {
// 			Id                                string        `json:"id"`
// 			Bio                               string        `json:"bio"`
// 			ShortBio                          string        `json:"short_bio"`
// 			CoachType                         string        `json:"coach_type"`
// 			IsFilterable                      bool          `json:"is_filterable"`
// 			IsInstructorGroup                 bool          `json:"is_instructor_group"`
// 			IndividualInstructorIds           []interface{} `json:"individual_instructor_ids"`
// 			IsVisible                         bool          `json:"is_visible"`
// 			IsAnnounced                       bool          `json:"is_announced"`
// 			ListOrder                         int           `json:"list_order"`
// 			FeaturedProfile                   bool          `json:"featured_profile"`
// 			FilmLink                          string        `json:"film_link"`
// 			FacebookFanPage                   string        `json:"facebook_fan_page"`
// 			MusicBio                          string        `json:"music_bio"`
// 			SpotifyPlaylistUri                string        `json:"spotify_playlist_uri"`
// 			Background                        string        `json:"background"`
// 			OrderedQAndAs                     [][]string    `json:"ordered_q_and_as"`
// 			InstagramProfile                  string        `json:"instagram_profile"`
// 			StravaProfile                     string        `json:"strava_profile"`
// 			TwitterProfile                    string        `json:"twitter_profile"`
// 			Quote                             string        `json:"quote"`
// 			Username                          string        `json:"username"`
// 			Name                              string        `json:"name"`
// 			FirstName                         string        `json:"first_name"`
// 			LastName                          string        `json:"last_name"`
// 			UserId                            string        `json:"user_id"`
// 			LifeStyleImageUrl                 string        `json:"life_style_image_url"`
// 			BikeInstructorListDisplayImageUrl interface{}   `json:"bike_instructor_list_display_image_url"`
// 			WebInstructorListDisplayImageUrl  string        `json:"web_instructor_list_display_image_url"`
// 			IosInstructorListDisplayImageUrl  string        `json:"ios_instructor_list_display_image_url"`
// 			AboutImageUrl                     string        `json:"about_image_url"`
// 			ImageUrl                          string        `json:"image_url"`
// 			JumbotronUrl                      interface{}   `json:"jumbotron_url"`
// 			JumbotronUrlDark                  string        `json:"jumbotron_url_dark"`
// 			JumbotronUrlIos                   string        `json:"jumbotron_url_ios"`
// 			WebInstructorListGifImageUrl      interface{}   `json:"web_instructor_list_gif_image_url"`
// 			InstructorHeroImageUrl            string        `json:"instructor_hero_image_url"`
// 			WorkoutShareImages                []struct {
// 				FitnessDiscipline string `json:"fitness_discipline"`
// 				ImageUrl          string `json:"image_url"`
// 			} `json:"workout_share_images"`
// 			FitnessDisciplines []string `json:"fitness_disciplines"`
// 		} `json:"instructor"`
// 		MuscleGroupScore []struct {
// 			MuscleGroup string `json:"muscle_group"`
// 			Score       int    `json:"score"`
// 			Percentage  int    `json:"percentage"`
// 			Bucket      int    `json:"bucket"`
// 			DisplayName string `json:"display_name"`
// 		} `json:"muscle_group_score"`
// 		IsFavorite             bool        `json:"is_favorite"`
// 		TotalUserWorkouts      int         `json:"total_user_workouts"`
// 		TotalFollowingWorkouts int         `json:"total_following_workouts"`
// 		LeaderboardFilterType  interface{} `json:"leaderboard_filter_type"`
// 		MembershipTieringInfo  struct {
// 			TierType                string      `json:"tier_type"`
// 			LimitedClassesTaken     interface{} `json:"limited_classes_taken"`
// 			LimitedClassesTotal     interface{} `json:"limited_classes_total"`
// 			LimitedClassesResetDate interface{} `json:"limited_classes_reset_date"`
// 		} `json:"membership_tiering_info"`
// 	} `json:"ride"`
// 	ClassTypes []struct {
// 		Id   string `json:"id"`
// 		Name string `json:"name"`
// 	} `json:"class_types"`
// 	Playlist struct {
// 		Id     string `json:"id"`
// 		RideId string `json:"ride_id"`
// 		Songs  []struct {
// 			Id      string `json:"id"`
// 			Title   string `json:"title"`
// 			Artists []struct {
// 				ArtistId   string `json:"artist_id"`
// 				ArtistName string `json:"artist_name"`
// 			} `json:"artists"`
// 			Album struct {
// 				Id       string `json:"id"`
// 				ImageUrl string `json:"image_url"`
// 				Name     string `json:"name"`
// 			} `json:"album"`
// 			ExplicitRating  int  `json:"explicit_rating"`
// 			CueTimeOffset   int  `json:"cue_time_offset"`
// 			StartTimeOffset int  `json:"start_time_offset"`
// 			Liked           bool `json:"liked"`
// 		} `json:"songs"`
// 		IsTopArtistsShown   bool `json:"is_top_artists_shown"`
// 		IsPlaylistShown     bool `json:"is_playlist_shown"`
// 		IsInClassMusicShown bool `json:"is_in_class_music_shown"`
// 		TopArtists          []struct {
// 			ArtistId   string `json:"artist_id"`
// 			ArtistName string `json:"artist_name"`
// 		} `json:"top_artists"`
// 		TopAlbums []struct {
// 			Id       string `json:"id"`
// 			ImageUrl string `json:"image_url"`
// 			Name     string `json:"name"`
// 		} `json:"top_albums"`
// 		StreamId  string      `json:"stream_id"`
// 		StreamUrl interface{} `json:"stream_url"`
// 	} `json:"playlist"`
// 	Averages struct {
// 		AverageTotalWork            int     `json:"average_total_work"`
// 		AverageDistance             float64 `json:"average_distance"`
// 		AverageCalories             int     `json:"average_calories"`
// 		AverageAvgPower             int     `json:"average_avg_power"`
// 		AverageAvgSpeed             float64 `json:"average_avg_speed"`
// 		AverageAvgCadence           int     `json:"average_avg_cadence"`
// 		AverageAvgResistance        int     `json:"average_avg_resistance"`
// 		AverageEffortScore          int     `json:"average_effort_score"`
// 		TotalHeartRateZoneDurations struct {
// 			HeartRateZ1Duration int `json:"heart_rate_z1_duration"`
// 			HeartRateZ2Duration int `json:"heart_rate_z2_duration"`
// 			HeartRateZ3Duration int `json:"heart_rate_z3_duration"`
// 			HeartRateZ4Duration int `json:"heart_rate_z4_duration"`
// 			HeartRateZ5Duration int `json:"heart_rate_z5_duration"`
// 		} `json:"total_heart_rate_zone_durations"`
// 	} `json:"averages"`
// 	Segments struct {
// 		SegmentList []struct {
// 			Id              string  `json:"id"`
// 			Length          int     `json:"length"`
// 			StartTimeOffset int     `json:"start_time_offset"`
// 			IconUrl         string  `json:"icon_url"`
// 			IntensityInMets float64 `json:"intensity_in_mets"`
// 			MetricsType     string  `json:"metrics_type"`
// 			IconName        string  `json:"icon_name"`
// 			IconSlug        string  `json:"icon_slug"`
// 			Name            string  `json:"name"`
// 			IsDrill         bool    `json:"is_drill"`
// 			SubsegmentsV2   []struct {
// 				Id                         string      `json:"id"`
// 				Type                       string      `json:"type"`
// 				DisplayName                string      `json:"display_name"`
// 				ScheduledOffset            int         `json:"scheduled_offset"`
// 				Offset                     int         `json:"offset"`
// 				Length                     int         `json:"length"`
// 				Rounds                     interface{} `json:"rounds"`
// 				TrackableMovementsDisabled bool        `json:"trackable_movements_disabled"`
// 				Movements                  []struct {
// 					Id           string      `json:"id"`
// 					Name         string      `json:"name"`
// 					Note         interface{} `json:"note"`
// 					Slug         interface{} `json:"slug"`
// 					SkillLevel   string      `json:"skill_level"`
// 					MuscleGroups []struct {
// 						MuscleGroup string `json:"muscle_group"`
// 						DisplayName string `json:"display_name"`
// 						Ranking     int    `json:"ranking"`
// 					} `json:"muscle_groups"`
// 					ShortVideoUrl       interface{}   `json:"short_video_url"`
// 					LongVideoUrl        interface{}   `json:"long_video_url"`
// 					MovementVideos      []interface{} `json:"movement_videos"`
// 					ImageUrl            interface{}   `json:"image_url"`
// 					TalkbackDescription string        `json:"talkback_description"`
// 				} `json:"movements"`
// 			} `json:"subsegments_v2"`
// 			IsTransition bool `json:"is_transition"`
// 		} `json:"segment_list"`
// 		SegmentCategoryDistribution struct {
// 			CyclingWarmup   string `json:"Cycling Warmup"`
// 			Cycling         string `json:"cycling"`
// 			CyclingCoolDown string `json:"Cycling Cool Down"`
// 		} `json:"segment_category_distribution"`
// 		SegmentBodyFocusDistribution struct {
// 			Cardio string `json:"cardio"`
// 		} `json:"segment_body_focus_distribution"`
// 		MovementsByWeightCategory struct {
// 		} `json:"movements_by_weight_category"`
// 	} `json:"segments"`
// 	DefaultAlbumImages struct {
// 		DefaultInClassImageUrl     string `json:"default_in_class_image_url"`
// 		DefaultClassDetailImageUrl string `json:"default_class_detail_image_url"`
// 	} `json:"default_album_images"`
// 	ExcludedPlatforms          []interface{} `json:"excluded_platforms"`
// 	IsFtpTest                  bool          `json:"is_ftp_test"`
// 	DisabledLeaderboardFilters struct {
// 		JustMe       bool `json:"just_me"`
// 		AgeAndGender bool `json:"age_and_gender"`
// 		Following    bool `json:"following"`
// 	} `json:"disabled_leaderboard_filters"`
// 	SampledTopTags []string `json:"sampled_top_tags"`
// 	InstructorCues []struct {
// 		Offsets struct {
// 			Start int `json:"start"`
// 			End   int `json:"end"`
// 		} `json:"offsets"`
// 		ResistanceRange struct {
// 			Upper int `json:"upper"`
// 			Lower int `json:"lower"`
// 		} `json:"resistance_range"`
// 		CadenceRange struct {
// 			Upper int `json:"upper"`
// 			Lower int `json:"lower"`
// 		} `json:"cadence_range"`
// 	} `json:"instructor_cues"`
// 	TargetClassMetrics struct {
// 		TargetGraphMetrics []struct {
// 			GraphData struct {
// 				Upper   []int `json:"upper"`
// 				Lower   []int `json:"lower"`
// 				Average []int `json:"average"`
// 			} `json:"graph_data"`
// 			Max     int    `json:"max"`
// 			Min     int    `json:"min"`
// 			Average int    `json:"average"`
// 			Type    string `json:"type"`
// 		} `json:"target_graph_metrics"`
// 		TotalExpectedOutput struct {
// 			ExpectedUpperOutput int `json:"expected_upper_output"`
// 			ExpectedLowerOutput int `json:"expected_lower_output"`
// 		} `json:"total_expected_output"`
// 	} `json:"target_class_metrics"`
// 	TargetMetricsData struct {
// 		TargetMetrics []struct {
// 			Offsets struct {
// 				Start int `json:"start"`
// 				End   int `json:"end"`
// 			} `json:"offsets"`
// 			SegmentType string `json:"segment_type"`
// 			Metrics     []struct {
// 				Name  string `json:"name"`
// 				Upper int    `json:"upper"`
// 				Lower int    `json:"lower"`
// 			} `json:"metrics"`
// 		} `json:"target_metrics"`
// 		TotalExpectedOutput struct {
// 			ExpectedUpperOutput int `json:"expected_upper_output"`
// 			ExpectedLowerOutput int `json:"expected_lower_output"`
// 		} `json:"total_expected_output"`
// 	} `json:"target_metrics_data"`
// 	Events struct {
// 		Data []interface{} `json:"data"`
// 	} `json:"events"`
// 	LanebreakInfo   interface{} `json:"lanebreak_info"`
// 	LastTimeMetrics interface{} `json:"last_time_metrics"`
// 	RelatedRides    struct {
// 		Id          string `json:"id"`
// 		Name        string `json:"name"`
// 		DisplayName string `json:"display_name"`
// 		Rides       []struct {
// 			FitnessDiscipline  string `json:"fitness_discipline"`
// 			ContentProvider    string `json:"content_provider"`
// 			Title              string `json:"title"`
// 			InstructorId       string `json:"instructor_id"`
// 			Id                 string `json:"id"`
// 			ImageUrl           string `json:"image_url"`
// 			Language           string `json:"language"`
// 			OriginLocale       string `json:"origin_locale"`
// 			OriginalAirTime    int    `json:"original_air_time"`
// 			ScheduledStartTime int    `json:"scheduled_start_time"`
// 		} `json:"rides"`
// 	} `json:"related_rides"`
// }
