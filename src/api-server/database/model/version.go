package model

import (
    "gitlab.com/hs-api-go/util/timeutil"
    "gopkg.in/mgo.v2/bson"
)

const COLLECTION_VERSION = "versions"

type Version struct {
    Base `bson:",inline"`
    Code string `json:"code" bson:"code"`
    Name string `json:"name" bson:"name"`
}

func NewVersion(code string, name string) *Version {
    version := Version{
        Code: code,
        Name: name,
    }
    version.Base.Id = bson.NewObjectId()
    version.Base.CreatedAt = timeutil.CurrentMillis()
    version.Base.ModifiedAt = timeutil.CurrentMillis()
    return &version
}