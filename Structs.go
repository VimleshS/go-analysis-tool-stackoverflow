package main

import (
	"html/template"
)

/* Global Vars*/
var StackExchangeAccessToken AccessToken

/* Constants */
const (
	StackExchangeOauthUri       = "https://stackexchange.com/oauth"
	StackExchangeAccessTokenUri = "https://stackexchange.com/oauth/access_token"
)

/* Structs */

type ChartData struct {
	label []string
	value []int
}

type NewChartData struct {
	Label string `json:"key"`
	Value int    `json:"y"`
}

type NCD struct {
	Key    string          `json:"key"`
	Values []*NewChartData `json:"values"`
}

type QNCD struct {
	Key    string `json:"key"`
	Values NCD    `json:"values"`
}

type UnAnsQuestionData struct {
	UserId       int
	CreationDate template.HTML
	UserName     template.HTML
	Question     template.HTML
	Link         template.HTML
}

type UserData struct {
	Display_name template.HTML
	Website_url  template.HTML
	About_me     template.HTML
	Location     template.HTML
	Link         template.HTML
	ImageUrl     template.HTML
}

type authError struct {
	Error map[string]string
}

type AccessToken struct {
	access_token string
	expires      string
}

type Tag struct {
	Name               string
	Count              int
	Is_required        bool
	Is_moderator_only  bool
	User_id            int
	Has_synonyms       bool
	Last_activity_date int64
}

type Tags struct {
	Items           []Tag
	Error_id        int
	Error_name      string
	Error_message   string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Questions struct {
	Items         []Question
	Error_id      int
	Error_name    string
	Error_message string

	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type Question struct {
	Question_id          int
	Last_edit_date       int64
	Creation_date        int64
	Last_activity_date   int64
	Locked_date          int64
	Community_owned_date int64
	Score                int
	Answer_count         int
	Accepted_answer_id   int
	Migrated_to          Migration_info
	Migrated_from        Migration_info
	Bounty_closes_date   int64
	Bounty_amount        int
	Closed_date          int64
	Protected_date       int64
	Body                 string
	Title                string
	Tags                 []string
	Closed_reason        string
	Up_vote_count        int
	Down_vote_count      int
	Favorite_count       int
	View_count           int
	Owner                ShallowUser
	Comments             []Comment
	Answers              []Answer
	Link                 string
	Is_answered          bool
}
type Migration_info struct {
	Question_id int
	Other_site  Site
	On_date     int64
}

type ShallowUser struct {
	User_id       int
	Display_name  string
	Reputation    int
	User_type     string //one of unregistered, registered, moderator, or does_not_exist
	Profile_image string
	Link          string
}

type Site struct {
	Site_type           string
	Name                string
	Logo_url            string
	Api_site_parameter  string
	Site_url            string
	Audience            string
	Icon_url            string
	Aliases             []string
	Site_state          string //one of normal, closed_beta, open_beta, or linked_meta
	Styling             Styling
	Closed_beta_date    int64
	Open_beta_date      int64
	Launch_date         int64
	Favicon_url         string
	Related_sites       []RelatedSite
	Twitter_account     string
	Markdown_extensions []string
}

type Styling struct {
	Link_color           string
	Tag_foreground_color string
	Tag_background_color string
}

type RelatedSite struct {
	Name               string
	Site_url           string
	Relation           string //one of parent, meta, chat, or other
	Api_site_parameter string
}

type Answer struct {
	Question_id          int
	Answer_id            int
	Locked_date          int64
	Creation_date        int64
	Last_edit_date       int64
	Last_activity_date   int64
	Score                int
	Community_owned_date int64
	Is_accepted          bool
	Body                 string
	Owner                ShallowUser
	Title                string
	Up_vote_count        int
	Down_vote_count      int
	Comments             []Comment
	Link                 string
}

type Badge struct {
	Badge_id    int
	Rank        string
	Name        string
	Description string
	Award_count int
	Badge_type  string
	User        ShallowUser
	Link        string
}

type BadgeCount struct {
	Gold   int
	Silver int
	Bronze int
}

type Comment struct {
	Comment_id    int
	Post_id       int
	Creation_date int64
	Post_type     string //one of question, or answer
	Score         int
	Edited        bool
	Body          string
	Owner         ShallowUser
	Reply_to_user ShallowUser
	Link          string
}

type Users struct {
	Items           []User
	Error_id        int
	Error_name      string
	Error_message   string
	Backoff         int
	Has_more        bool
	Page            int
	Page_size       int
	Quota_max       int
	Quota_remaining int
	Total           int
	Type            string
}

type User struct {
	User_id                   int
	User_type                 string //one of unregistered, registered, moderator, or does_not_exist
	Creation_date             int64
	Display_name              string
	Profile_image             string
	Reputation                int
	Reputation_change_day     int
	Reputation_change_week    int
	Reputation_change_month   int
	Reputation_change_quarter int
	Reputation_change_year    int
	Age                       int
	Last_access_date          int64
	Last_modified_date        int64
	Is_employee               bool
	Link                      string
	Website_url               string
	Location                  string
	Account_id                int
	Timed_penalty_date        int64
	Badge_counts              BadgeCount
	Question_count            int
	Answer_count              int
	Up_vote_count             int
	Down_vote_count           int
	About_me                  string
	View_count                int
	Accept_rate               int
}
