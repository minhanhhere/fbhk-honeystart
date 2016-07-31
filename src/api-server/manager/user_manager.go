package manager

import (
    "gitlab.com/hs-api-go/database/model"
    "gitlab.com/hs-api-go/database"
    "gopkg.in/mgo.v2/bson"
    "gitlab.com/hs-api-go/secure/password"
)

type userManager struct {
}

var UserManager = &userManager{}

func (m *userManager) GetById(id string) *model.User {
    result := &model.User{}
    database.Users().FindById(bson.ObjectIdHex(id), result)
    if !password.IsHashed(result.Password) {
        result.Password, _ = password.CreateHash(result.Password)
        m.Save(result)
    }
    return result
}

func (m *userManager) GetByEmail(email string) *model.User {
    result := &model.User{}
    if err := database.Users().FindOne(bson.M{"email": email}, result); err != nil {
        return nil
    }
    if !password.IsHashed(result.Password) {
        result.Password, _ = password.CreateHash(result.Password)
        m.Save(result)
    }
    return result
}

func (m *userManager) GetByFacebookId(facebookId string) *model.User {
    result := &model.User{}
    if err := database.Users().FindOne(bson.M{"facebookId": facebookId}, result); err != nil {
        return nil
    }
    return result
}

func (m *userManager) GetAll() (result []model.User) {
    database.Users().Find(bson.M{"deleted": bson.M{"$ne": true}}).Query.All(&result)
    return
}

func (m *userManager) CountAll() int {
    count, err := database.Users().Find(bson.M{"deleted": bson.M{"$ne": true}}).Query.Count()
    if err != nil {
        return 0
    }
    return count
}

func (m *userManager) Save(model *model.User) error {
    return database.Users().Save(model)
}