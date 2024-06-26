package model

// GameOverResp ...
type GameOverResp struct {
	ResponseData []any `json:"response_data"`
	ReleaseInfo  []any `json:"release_info"`
	StatusCode   int   `json:"status_code"`
}

// NormalLiveStatusList ...
type NormalLiveStatusList struct {
	LiveDifficultyID   int   `json:"live_difficulty_id"`
	Status             int   `json:"status"`
	HiScore            int   `json:"hi_score"`
	HiComboCount       int   `json:"hi_combo_count"`
	ClearCnt           int   `json:"clear_cnt"`
	AchievedGoalIDList []int `json:"achieved_goal_id_list"`
}

// SpecialLiveStatusList ...
type SpecialLiveStatusList struct {
	LiveDifficultyID   int   `json:"live_difficulty_id"`
	Status             int   `json:"status"`
	HiScore            int   `json:"hi_score"`
	HiComboCount       int   `json:"hi_combo_count"`
	ClearCnt           int   `json:"clear_cnt"`
	AchievedGoalIDList []int `json:"achieved_goal_id_list"`
}

// TrainingLiveStatusList ...
type TrainingLiveStatusList struct {
	LiveDifficultyID   int   `json:"live_difficulty_id"`
	Status             int   `json:"status"`
	HiScore            int   `json:"hi_score"`
	HiComboCount       int   `json:"hi_combo_count"`
	ClearCnt           int   `json:"clear_cnt"`
	AchievedGoalIDList []int `json:"achieved_goal_id_list"`
}

// LiveStatusRes ...
type LiveStatusRes struct {
	NormalLiveStatusList   []NormalLiveStatusList   `json:"normal_live_status_list"`
	SpecialLiveStatusList  []SpecialLiveStatusList  `json:"special_live_status_list"`
	TrainingLiveStatusList []TrainingLiveStatusList `json:"training_live_status_list"`
	MarathonLiveStatusList []any                    `json:"marathon_live_status_list"`
	FreeLiveStatusList     []any                    `json:"free_live_status_list"`
	CanResumeLive          bool                     `json:"can_resume_live"`
}

// LiveStatusResp ...
type LiveStatusResp struct {
	Result     LiveStatusRes `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

// LiveList ...
type LiveList struct {
	LiveDifficultyID int    `json:"live_difficulty_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	IsRandom         bool   `json:"is_random"`
}

// LimitedBonusCommonList ...
type LimitedBonusCommonList struct {
	LiveType          int    `json:"live_type"`
	LimitedBonusType  int    `json:"limited_bonus_type"`
	LimitedBonusValue int    `json:"limited_bonus_value"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
}

// RandomLiveList ...
type RandomLiveList struct {
	AttributeID int    `json:"attribute_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

// TrainingLiveList ...
type TrainingLiveList struct {
	LiveDifficultyID int    `json:"live_difficulty_id"`
	StartDate        string `json:"start_date"`
	IsRandom         bool   `json:"is_random"`
}

// LiveScheduleRes ...
type LiveScheduleRes struct {
	EventList              []any                    `json:"event_list"`
	LiveList               []LiveList               `json:"live_list"`
	LimitedBonusList       []any                    `json:"limited_bonus_list"`
	LimitedBonusCommonList []LimitedBonusCommonList `json:"limited_bonus_common_list"`
	RandomLiveList         []RandomLiveList         `json:"random_live_list"`
	FreeLiveList           []any                    `json:"free_live_list"`
	TrainingLiveList       []TrainingLiveList       `json:"training_live_list"`
}

// LiveScheduleResp ...
type LiveScheduleResp struct {
	Result     LiveScheduleRes `json:"result"`
	Status     int             `json:"status"`
	CommandNum bool            `json:"commandNum"`
	TimeStamp  int64           `json:"timeStamp"`
}

// PlayReq ...
type PlayReq struct {
	Module           string `json:"module"`
	PartyUserID      int64  `json:"party_user_id"`
	Action           string `json:"action"`
	Mgd              int    `json:"mgd"`
	IsTraining       bool   `json:"is_training"`
	UnitDeckID       int    `json:"unit_deck_id"`
	LiveDifficultyID string `json:"live_difficulty_id"`
	TimeStamp        int    `json:"timeStamp"`
	LpFactor         int    `json:"lp_factor"`
	CommandNum       string `json:"commandNum"`
}

// RankInfo ...
type RankInfo struct {
	Rank    int `json:"rank"`
	RankMin int `json:"rank_min"`
	RankMax int `json:"rank_max"`
}

// NotesList ...
type NotesList struct {
	TimingSec      float64 `json:"timing_sec"`
	NotesAttribute int     `json:"notes_attribute"`
	NotesLevel     int     `json:"notes_level"`
	Effect         int     `json:"effect"`
	EffectValue    float64 `json:"effect_value"`
	Position       int     `json:"position"`
}

// LiveInfo ...
type LiveInfo struct {
	LiveDifficultyID int         `json:"live_difficulty_id"`
	IsRandom         bool        `json:"is_random"`
	AcFlag           int         `json:"ac_flag"`
	SwingFlag        int         `json:"swing_flag"`
	NotesList        []NotesList `json:"notes_list"`
}

// PlayCostume ...
type PlayCostume struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

// UnitList ...
type UnitList struct {
	Smile   int         `json:"smile"`
	Cute    int         `json:"cute"`
	Cool    int         `json:"cool"`
	Costume PlayCostume `json:"costume,omitempty"`
}

// DeckInfo ...
type DeckInfo struct {
	UnitDeckID       int        `json:"unit_deck_id"`
	TotalSmile       int        `json:"total_smile"`
	TotalCute        int        `json:"total_cute"`
	TotalCool        int        `json:"total_cool"`
	TotalHp          int        `json:"total_hp"`
	PreparedHpDamage int        `json:"prepared_hp_damage"`
	UnitList         []UnitList `json:"unit_list"`
}

// PlayLiveList ...
type PlayLiveList struct {
	LiveInfo LiveInfo `json:"live_info"`
	DeckInfo DeckInfo `json:"deck_info"`
}

// PlayRes ...
type PlayRes struct {
	RankInfo            []RankInfo     `json:"rank_info"`
	EnergyFullTime      string         `json:"energy_full_time"`
	OverMaxEnergy       int            `json:"over_max_energy"`
	AvailableLiveResume bool           `json:"available_live_resume"`
	LiveList            []PlayLiveList `json:"live_list"`
	IsMarathonEvent     bool           `json:"is_marathon_event"`
	MarathonEventID     any            `json:"marathon_event_id"`
	NoSkill             bool           `json:"no_skill"`
	CanActivateEffect   bool           `json:"can_activate_effect"`
	ServerTimestamp     int64          `json:"server_timestamp"`
}

// PlayResp ...
type PlayResp struct {
	ResponseData PlayRes `json:"response_data"`
	ReleaseInfo  []any   `json:"release_info"`
	StatusCode   int     `json:"status_code"`
}

// PlayScoreReq ...
type PlayScoreReq struct {
	Module           string `json:"module"`
	Action           string `json:"action"`
	TimeStamp        int64  `json:"timeStamp"`
	Mgd              int    `json:"mgd"`
	LiveDifficultyID string `json:"live_difficulty_id"`
	CommandNum       string `json:"commandNum"`
}

// On ...
type On struct {
	HasRecord   bool     `json:"has_record"`
	LiveInfo    LiveInfo `json:"live_info"`
	RandomSeed  any      `json:"random_seed"`
	MaxCombo    any      `json:"max_combo"`
	UpdateDate  any      `json:"update_date"`
	PreciseList any      `json:"precise_list"`
	DeckInfo    any      `json:"deck_info"`
	TapAdjust   any      `json:"tap_adjust"`
	CanReplay   bool     `json:"can_replay"`
}

// Off ...
type Off struct {
	HasRecord   bool     `json:"has_record"`
	LiveInfo    LiveInfo `json:"live_info"`
	RandomSeed  any      `json:"random_seed"`
	MaxCombo    any      `json:"max_combo"`
	UpdateDate  any      `json:"update_date"`
	PreciseList any      `json:"precise_list"`
	DeckInfo    any      `json:"deck_info"`
	TapAdjust   any      `json:"tap_adjust"`
	CanReplay   bool     `json:"can_replay"`
}

// PlayScoreRes ...
type PlayScoreRes struct {
	On                On         `json:"on"`
	Off               Off        `json:"off"`
	RankInfo          []RankInfo `json:"rank_info"`
	CanActivateEffect bool       `json:"can_activate_effect"`
	ServerTimestamp   int64      `json:"server_timestamp"`
}

// PlayScoreResp ...
type PlayScoreResp struct {
	ResponseData PlayScoreRes `json:"response_data"`
	ReleaseInfo  []any        `json:"release_info"`
	StatusCode   int          `json:"status_code"`
}

// PlayRewardReq ...
type PlayRewardReq struct {
	Module           string          `json:"module"`
	Action           string          `json:"action"`
	GoodCnt          int             `json:"good_cnt"`
	MissCnt          int             `json:"miss_cnt"`
	IsTraining       bool            `json:"is_training"`
	GreatCnt         int             `json:"great_cnt"`
	CommandNum       string          `json:"commandNum"`
	LoveCnt          int             `json:"love_cnt"`
	RemainHp         int             `json:"remain_hp"`
	MaxCombo         int             `json:"max_combo"`
	ScoreSmile       int             `json:"score_smile"`
	PerfectCnt       int             `json:"perfect_cnt"`
	BadCnt           int             `json:"bad_cnt"`
	Mgd              int             `json:"mgd"`
	EventPoint       int             `json:"event_point"`
	LiveDifficultyID int             `json:"live_difficulty_id"`
	TimeStamp        int             `json:"timeStamp"`
	PreciseScoreLog  PreciseScoreLog `json:"precise_score_log"`
	ScoreCute        int             `json:"score_cute"`
	EventID          any             `json:"event_id"`
	ScoreCool        int             `json:"score_cool"`
}

// Icon ...
type Icon struct {
	SlideID  int `json:"slide_id"`
	JustID   int `json:"just_id"`
	NormalID int `json:"normal_id"`
}

// LiveSetting ...
type LiveSetting struct {
	StringSize                 int     `json:"string_size"`
	PreciseScoreAutoUpdateFlag bool    `json:"precise_score_auto_update_flag"`
	SeID                       int     `json:"se_id"`
	CutinBrightness            int     `json:"cutin_brightness"`
	RandomValue                int     `json:"random_value"`
	PreciseScoreUpdateType     int     `json:"precise_score_update_type"`
	EffectFlag                 bool    `json:"effect_flag"`
	NotesSpeed                 float64 `json:"notes_speed"`
	Icon                       Icon    `json:"icon"`
	CutinType                  int     `json:"cutin_type"`
}

// PreciseList ...
type PreciseList struct {
	Effect     int     `json:"effect"`
	Count      int     `json:"count"`
	Tap        float64 `json:"tap"`
	NoteNumber int     `json:"note_number"`
	Position   int     `json:"position"`
	Accuracy   int     `json:"accuracy"`
	IsSame     bool    `json:"is_same"`
}

// BackgroundScore ...
type BackgroundScore struct {
	Smile int `json:"smile"`
	Cute  int `json:"cute"`
	Cool  int `json:"cool"`
}

// TriggerLog ...
type TriggerLog struct {
	ActivationRate int `json:"activation_rate"`
	Position       int `json:"position"`
}

// PreciseScoreLog ...
type PreciseScoreLog struct {
	LiveSetting     LiveSetting     `json:"live_setting"`
	TapAdjust       int             `json:"tap_adjust"`
	PreciseList     []PreciseList   `json:"precise_list"`
	BackgroundScore BackgroundScore `json:"background_score"`
	IsLogOn         bool            `json:"is_log_on"`
	ScoreLog        []int           `json:"score_log"`
	IsSkillOn       bool            `json:"is_skill_on"`
	TriggerLog      []TriggerLog    `json:"trigger_log"`
	RandomSeed      int             `json:"random_seed"`
}

// RewardLiveInfo ...
type RewardLiveInfo struct {
	LiveDifficultyID int  `json:"live_difficulty_id"`
	IsRandom         bool `json:"is_random"`
	AcFlag           int  `json:"ac_flag"`
	SwingFlag        int  `json:"swing_flag"`
}

// PlayerExpUnitMax ...
type PlayerExpUnitMax struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

// PlayerExpFriendMax ...
type PlayerExpFriendMax struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

// PlayerExpLpMax ...
type PlayerExpLpMax struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

// BaseRewardInfo ...
type BaseRewardInfo struct {
	PlayerExp             int                `json:"player_exp"`
	PlayerExpUnitMax      PlayerExpUnitMax   `json:"player_exp_unit_max"`
	PlayerExpFriendMax    PlayerExpFriendMax `json:"player_exp_friend_max"`
	PlayerExpLpMax        PlayerExpLpMax     `json:"player_exp_lp_max"`
	GameCoin              int                `json:"game_coin"`
	GameCoinRewardBoxFlag bool               `json:"game_coin_reward_box_flag"`
	SocialPoint           int                `json:"social_point"`
}

// LiveClear ...
type LiveClear struct {
	AddType                    int   `json:"add_type"`
	Amount                     int   `json:"amount"`
	ItemCategoryID             int   `json:"item_category_id"`
	UnitID                     int   `json:"unit_id"`
	UnitOwningUserID           int64 `json:"unit_owning_user_id"`
	IsSupportMember            bool  `json:"is_support_member"`
	Exp                        int   `json:"exp"`
	NextExp                    int   `json:"next_exp"`
	MaxHp                      int   `json:"max_hp"`
	Level                      int   `json:"level"`
	MaxLevel                   int   `json:"max_level"`
	LevelLimitID               int   `json:"level_limit_id"`
	SkillLevel                 int   `json:"skill_level"`
	Rank                       int   `json:"rank"`
	Love                       int   `json:"love"`
	IsRankMax                  bool  `json:"is_rank_max"`
	IsLevelMax                 bool  `json:"is_level_max"`
	IsLoveMax                  bool  `json:"is_love_max"`
	IsSigned                   bool  `json:"is_signed"`
	NewUnitFlag                bool  `json:"new_unit_flag"`
	RewardBoxFlag              bool  `json:"reward_box_flag"`
	UnitSkillExp               int   `json:"unit_skill_exp"`
	DisplayRank                int   `json:"display_rank"`
	UnitRemovableSkillCapacity int   `json:"unit_removable_skill_capacity"`
	RemovableSkillIds          []any `json:"removable_skill_ids"`
}

// LiveRank ...
type LiveRank struct {
	AddType                    int   `json:"add_type"`
	Amount                     int   `json:"amount"`
	ItemCategoryID             int   `json:"item_category_id"`
	UnitID                     int   `json:"unit_id"`
	UnitOwningUserID           int64 `json:"unit_owning_user_id"`
	IsSupportMember            bool  `json:"is_support_member"`
	Exp                        int   `json:"exp"`
	NextExp                    int   `json:"next_exp"`
	MaxHp                      int   `json:"max_hp"`
	Level                      int   `json:"level"`
	MaxLevel                   int   `json:"max_level"`
	LevelLimitID               int   `json:"level_limit_id"`
	SkillLevel                 int   `json:"skill_level"`
	Rank                       int   `json:"rank"`
	Love                       int   `json:"love"`
	IsRankMax                  bool  `json:"is_rank_max"`
	IsLevelMax                 bool  `json:"is_level_max"`
	IsLoveMax                  bool  `json:"is_love_max"`
	IsSigned                   bool  `json:"is_signed"`
	NewUnitFlag                bool  `json:"new_unit_flag"`
	RewardBoxFlag              bool  `json:"reward_box_flag"`
	UnitSkillExp               int   `json:"unit_skill_exp"`
	DisplayRank                int   `json:"display_rank"`
	UnitRemovableSkillCapacity int   `json:"unit_removable_skill_capacity"`
	RemovableSkillIds          []any `json:"removable_skill_ids"`
}

// RewardUnitList ...
type RewardUnitList struct {
	LiveClear []LiveClear `json:"live_clear"`
	LiveRank  []LiveRank  `json:"live_rank"`
	LiveCombo []any       `json:"live_combo"`
}

// Rewards ...
type Rewards struct {
	Rarity         int    `json:"rarity"`
	ItemID         int    `json:"item_id"`
	AddType        int    `json:"add_type"`
	Amount         int    `json:"amount"`
	ItemCategoryID int    `json:"item_category_id"`
	RewardBoxFlag  bool   `json:"reward_box_flag"`
	InsertDate     string `json:"insert_date"`
}

// EffortPoint ...
type EffortPoint struct {
	LiveEffortPointBoxSpecID int       `json:"live_effort_point_box_spec_id"`
	Capacity                 int       `json:"capacity"`
	Before                   int       `json:"before"`
	After                    int       `json:"after"`
	Rewards                  []Rewards `json:"rewards"`
}

// PlayRewardUnitList ...
type PlayRewardUnitList struct {
	ID               int   `xorm:"id pk autoincr" json:"-"`
	UserDeckID       int   `xorm:"user_deck_id" json:"-"`
	UnitOwningUserID int   `xorm:"unit_owning_user_id" json:"unit_owning_user_id"`
	UnitID           int   `xorm:"unit_id" json:"unit_id"`
	Position         int   `xorm:"position" json:"position"`
	Level            int   `xorm:"level" json:"level"`
	LevelLimitID     int   `xorm:"level_limit_id" json:"level_limit_id"`
	DisplayRank      int   `xorm:"display_rank" json:"display_rank"`
	Love             int   `xorm:"love" json:"love"`
	UnitSkillLevel   int   `xorm:"unit_skill_level" json:"unit_skill_level"`
	IsRankMax        bool  `xorm:"is_rank_max" json:"is_rank_max"`
	IsLoveMax        bool  `xorm:"is_love_max" json:"is_love_max"`
	IsLevelMax       bool  `xorm:"is_level_max" json:"is_level_max"`
	IsSigned         bool  `xorm:"is_signed" json:"is_signed"`
	BeforeLove       int   `xorm:"before_love" json:"before_love"`
	MaxLove          int   `xorm:"max_love" json:"max_love"`
	InsertData       int64 `xorm:"insert_date" json:"-"`
}

// BeforeUserInfo ...
type BeforeUserInfo struct {
	Level                          int    `json:"level"`
	Exp                            int    `json:"exp"`
	PreviousExp                    int    `json:"previous_exp"`
	NextExp                        int    `json:"next_exp"`
	GameCoin                       int    `json:"game_coin"`
	SnsCoin                        int    `json:"sns_coin"`
	FreeSnsCoin                    int    `json:"free_sns_coin"`
	PaidSnsCoin                    int    `json:"paid_sns_coin"`
	SocialPoint                    int    `json:"social_point"`
	UnitMax                        int    `json:"unit_max"`
	WaitingUnitMax                 int    `json:"waiting_unit_max"`
	CurrentEnergy                  int    `json:"current_energy"`
	EnergyMax                      int    `json:"energy_max"`
	TrainingEnergy                 int    `json:"training_energy"`
	TrainingEnergyMax              int    `json:"training_energy_max"`
	EnergyFullTime                 string `json:"energy_full_time"`
	LicenseLiveEnergyRecoverlyTime int    `json:"license_live_energy_recoverly_time"`
	FriendMax                      int    `json:"friend_max"`
	TutorialState                  int    `json:"tutorial_state"`
	OverMaxEnergy                  int    `json:"over_max_energy"`
	UnlockRandomLiveMuse           int    `json:"unlock_random_live_muse"`
	UnlockRandomLiveAqours         int    `json:"unlock_random_live_aqours"`
}

// AfterUserInfo ...
type AfterUserInfo struct {
	Level                          int    `json:"level"`
	Exp                            int    `json:"exp"`
	PreviousExp                    int    `json:"previous_exp"`
	NextExp                        int    `json:"next_exp"`
	GameCoin                       int    `json:"game_coin"`
	SnsCoin                        int    `json:"sns_coin"`
	FreeSnsCoin                    int    `json:"free_sns_coin"`
	PaidSnsCoin                    int    `json:"paid_sns_coin"`
	SocialPoint                    int    `json:"social_point"`
	UnitMax                        int    `json:"unit_max"`
	WaitingUnitMax                 int    `json:"waiting_unit_max"`
	CurrentEnergy                  int    `json:"current_energy"`
	EnergyMax                      int    `json:"energy_max"`
	TrainingEnergy                 int    `json:"training_energy"`
	TrainingEnergyMax              int    `json:"training_energy_max"`
	EnergyFullTime                 string `json:"energy_full_time"`
	LicenseLiveEnergyRecoverlyTime int    `json:"license_live_energy_recoverly_time"`
	FriendMax                      int    `json:"friend_max"`
	TutorialState                  int    `json:"tutorial_state"`
	OverMaxEnergy                  int    `json:"over_max_energy"`
	UnlockRandomLiveMuse           int    `json:"unlock_random_live_muse"`
	UnlockRandomLiveAqours         int    `json:"unlock_random_live_aqours"`
}

// NextLevelInfo ...
type NextLevelInfo struct {
	Level   int `json:"level"`
	FromExp int `json:"from_exp"`
}

// GoalAccompInfo ...
type GoalAccompInfo struct {
	AchievedIds []any `json:"achieved_ids"`
	Rewards     []any `json:"rewards"`
}

// RewardRankInfo ...
type RewardRankInfo struct {
	BeforeClassRankID int    `json:"before_class_rank_id"`
	AfterClassRankID  int    `json:"after_class_rank_id"`
	RankUpDate        string `json:"rank_up_date"`
}

// ClassSystem ...
type ClassSystem struct {
	RankInfo     RewardRankInfo `json:"rank_info"`
	CompleteFlag bool           `json:"complete_flag"`
	IsOpened     bool           `json:"is_opened"`
	IsVisible    bool           `json:"is_visible"`
}

// PlayRewardList ...
type PlayRewardList struct {
	ItemID         int  `json:"item_id"`
	AddType        int  `json:"add_type"`
	Amount         int  `json:"amount"`
	ItemCategoryID int  `json:"item_category_id"`
	RewardBoxFlag  bool `json:"reward_box_flag"`
}

// AccomplishedAchievementList ...
type AccomplishedAchievementList struct {
	AchievementID       int              `json:"achievement_id"`
	Count               int              `json:"count"`
	IsAccomplished      bool             `json:"is_accomplished"`
	InsertDate          string           `json:"insert_date"`
	EndDate             string           `json:"end_date"`
	RemainingTime       string           `json:"remaining_time"`
	IsNew               bool             `json:"is_new"`
	ForDisplay          bool             `json:"for_display"`
	IsLocked            bool             `json:"is_locked"`
	OpenConditionString string           `json:"open_condition_string"`
	AccomplishID        string           `json:"accomplish_id"`
	RewardList          []PlayRewardList `json:"reward_list"`
}

// RewardUnitSupportList ...
type RewardUnitSupportList struct {
	UnitID int `json:"unit_id"`
	Amount int `json:"amount"`
}

// RewardRes ...
type RewardRes struct {
	LiveInfo                     []RewardLiveInfo              `json:"live_info"`
	Rank                         int                           `json:"rank"`
	ComboRank                    int                           `json:"combo_rank"`
	TotalLove                    int                           `json:"total_love"`
	IsHighScore                  bool                          `json:"is_high_score"`
	HiScore                      int                           `json:"hi_score"`
	BaseRewardInfo               BaseRewardInfo                `json:"base_reward_info"`
	RewardUnitList               RewardUnitList                `json:"reward_unit_list"`
	UnlockedSubscenarioIds       []any                         `json:"unlocked_subscenario_ids"`
	UnlockedMultiUnitScenarioIds []any                         `json:"unlocked_multi_unit_scenario_ids"`
	EffortPoint                  []EffortPoint                 `json:"effort_point"`
	IsEffortPointVisible         bool                          `json:"is_effort_point_visible"`
	LimitedEffortBox             []any                         `json:"limited_effort_box"`
	UnitList                     []PlayRewardUnitList          `json:"unit_list"`
	BeforeUserInfo               BeforeUserInfo                `json:"before_user_info"`
	AfterUserInfo                AfterUserInfo                 `json:"after_user_info"`
	NextLevelInfo                []NextLevelInfo               `json:"next_level_info"`
	GoalAccompInfo               GoalAccompInfo                `json:"goal_accomp_info"`
	SpecialRewardInfo            []any                         `json:"special_reward_info"`
	EventInfo                    []any                         `json:"event_info"`
	DailyRewardInfo              []any                         `json:"daily_reward_info"`
	CanSendFriendRequest         bool                          `json:"can_send_friend_request"`
	UsingBuffInfo                []any                         `json:"using_buff_info"`
	ClassSystem                  ClassSystem                   `json:"class_system"`
	AccomplishedAchievementList  []AccomplishedAchievementList `json:"accomplished_achievement_list"`
	UnaccomplishedAchievementCnt int                           `json:"unaccomplished_achievement_cnt"`
	AddedAchievementList         []any                         `json:"added_achievement_list"`
	MuseumInfo                   Museum                        `json:"museum_info"`
	UnitSupportList              []RewardUnitSupportList       `json:"unit_support_list"`
	ServerTimestamp              int64                         `json:"server_timestamp"`
	PresentCnt                   int                           `json:"present_cnt"`
}

// RewardResp ...
type RewardResp struct {
	ResponseData RewardRes `json:"response_data"`
	ReleaseInfo  []any     `json:"release_info"`
	StatusCode   int       `json:"status_code"`
}

// LiveSeInfoRes ...
type LiveSeInfoRes struct {
	LiveSeList []int `json:"live_se_list"`
}

// LiveSeInfoResp ...
type LiveSeInfoResp struct {
	Result     LiveSeInfoRes `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

// LiveIconInfoRes ...
type LiveIconInfoRes struct {
	LiveNotesIconList []int `json:"live_notes_icon_list"`
}

// LiveIconInfoResp ...
type LiveIconInfoResp struct {
	Result     LiveIconInfoRes `json:"result"`
	Status     int             `json:"status"`
	CommandNum bool            `json:"commandNum"`
	TimeStamp  int64           `json:"timeStamp"`
}
