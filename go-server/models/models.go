package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ToDoList is from the old code base
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

// new models

// Representative ... object of representative
type Representative struct {
	GUID                  string  `json:"guid" binding:"required" csv:"id"`
	Office                string  `json:"office" csv:"title"`
	Name                  string  `json:"name" binding:"required" csv:"first_name"`
	LastName              string  `csv:"last_name"`
	Location              string  `json:"location" csv:"state"`
	Division              string  `json:"division" csv:"ocd_id"`
	GovWebsite            string  `json:"gov_web" csv:"url"`
	Twitter               string  `json:"twitter" csv:"twitter_account"`
	TotalVotes            int     `json:"total_votes" csv:"total_votes"`
	MissedVotes           int     `json:"missed_votes" csv:"missed_votes"`
	PresentVotes          int     `json:"present_votes" csv:"present_votes"`
	PercentMissedVotes    float64 `json:"percent_missed_votes"`
	PercentPresentVotes   float64 `json:"percent_present_votes"`
	PercentVotesWithParty float64 `json:"percent_votes_with_party" csv:"votes_with_party_pct"`
}

// UserRepMap maps users to their representatives
type UserRepMap struct {
	UserGUID string `csv:"user_guid"`
	RepGUID  string `csv:"rep_guid"`
}
