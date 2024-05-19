package model

type PushReport struct {
	Username string `bson:"username"`
	Email    string `bson:"email,omitempty"`
	Repo     string `bson:"repo"`
	Ref      string `bson:"ref"`
	Message  string `bson:"message"`
	Modified string `bson:"modified,omitempty"`
}
