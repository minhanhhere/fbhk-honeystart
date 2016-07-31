package model

const COLLECTION_TEAM string = "teams"

type Team struct {
    Base `bson:",inline"`
    Name     string `json:"name" bson:"name"`
    Location string `json:"location" bson:"location"`
    Logo     string `json:"logo" bson:"logo"`
    Email    string `json:"email" bson:"email"`
    Website  string `json:"website" bson:"website"`
    About    string `json:"about" bson:"about"`
}

func NewTeam() *Team {
    team := Team{}
    team.Base.construct()
    return &team
}