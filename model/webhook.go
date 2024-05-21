package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PushReport struct {
	ProjectID string `bson:"projectid"`
	Username  string `bson:"username"`
	Email     string `bson:"email,omitempty"`
	Repo      string `bson:"repo"`
	Ref       string `bson:"ref"`
	Message   string `bson:"message"`
	Modified  string `bson:"modified,omitempty"`
}

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Secret      string             `bson:"secret"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Owner       Userdomyikado      `bson:"owner"`
	Member      Userdomyikado      `bson:"member"`
	Closed      bool               `bson:"closed,omitempty"`
}

type Userdomyikado struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name            string             `bson:"name" json:"name"`
	PhoneNumber     string             `bson:"phonenumber" json:"phonenumber"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	GithubUsername  string             `bson:"githubusername,omitempty" json:"githubusername,omitempty"`
	GitlabUsername  string             `bson:"gitlabusername,omitempty" json:"gitlabusername,omitempty"`
	GitHostUsername string             `bson:"githostusername,omitempty" json:"githostusername,omitempty"`
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID string             `bson:"projectid"`
	Name      string             `bson:"name"`
	PIC       Userdomyikado      `bson:"pic"`
	Done      bool               `bson:"done,omitempty"`
}
