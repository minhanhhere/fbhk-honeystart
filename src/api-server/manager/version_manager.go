package manager

import (
    "gitlab.com/hs-api-go/database/model"
    "gitlab.com/hs-api-go/database"
    "gopkg.in/mgo.v2/bson"
)

type versionManager struct {
}

var VersionManager = &versionManager{}

func (m *versionManager) GetAll() []model.Version {
    var result []model.Version
    database.Versions().Find(bson.M{}).Query.All(&result)
    return result
}

func (m *versionManager) Save(model *model.Version) error {
    return database.Versions().Save(model)
}