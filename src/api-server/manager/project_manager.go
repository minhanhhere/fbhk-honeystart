package manager

import (
    "gitlab.com/hs-api-go/database/model"
    "gitlab.com/hs-api-go/database"
    "gopkg.in/mgo.v2/bson"
    . "github.com/ahmetalpbalkan/go-linq"
    "github.com/labstack/gommon/log"
)

type projectManager struct {
}

var ProjectManager = &projectManager{}

func (m *projectManager) FindById(id string) *model.Project {
    result := &model.Project{}
    err := database.Projects().FindById(bson.ObjectIdHex(id), result)
    if err != nil {
        log.Error(err)
        return nil
    }
    fillDetail(result)
    return result
}

func (m *projectManager) GetAll() (result []model.Project) {
    database.Projects().Find(bson.M{"deleted": bson.M{"$ne": true}}).Query.All(&result)
    for i, _ := range result {
        fillDetail(&result[i])
    }
    return result
}

func (m *projectManager) SaveProject(project *model.Project) {
    project.OwnerId = bson.NewObjectId()
    err := database.Projects().Save(project)
    if err != nil {
        log.Error(err)
    }
}

func (m *projectManager) GetUpdates(id string) (result []model.ProjectUpdate) {
    err := database.ProjectUpdates().Find(bson.M{"projectId" : bson.ObjectIdHex(id)}).Query.All(&result)
    if err != nil {
        log.Error(err)
        return nil
    }
    //TODO: load all owner in 1 query
    for i, item := range result {
        result[i].Owner = UserManager.GetById(item.OwnerId.Hex())
    }
    return
}

func (m *projectManager) GetBackers(id string) (result []model.ProjectBacker) {
    project := m.FindById(id)
    database.ProjectBackers().Find(bson.M{"projectId" : bson.ObjectIdHex(id)}).Query.All(&result)
    //TODO: load all owner in 1 query
    for i, _ := range result {
        result[i].Owner = UserManager.GetById(result[i].OwnerId.Hex())
        p, _ := From(project.Packages).Single(func(in T) (bool, error) {
            return in.(model.ProjectPackage).Id == result[i].PackageId, nil
        })
        pkg := p.(model.ProjectPackage)
        result[i].Package = &pkg
    }
    return
}

func (m *projectManager) SaveBacker(backer *model.ProjectBacker) {
    database.ProjectBackers().Save(backer)
}

func (m *projectManager) FindBackerById(id string) *model.ProjectBacker {
    result := &model.ProjectBacker{}
    err := database.ProjectBackers().FindById(bson.ObjectIdHex(id), result)
    if err != nil {
        log.Error(err)
        return nil
    }
    return result
}

func fillDetail(p *model.Project) {
    fillUpdatesCount(p)
    fillBackersCount(p)
}

func fillUpdatesCount(p *model.Project) {
    updatesCount, err := database.ProjectUpdates().Find(bson.M{"projectId" : p.Id}).Query.Count()
    if err == nil {
        p.UpdatesCount = updatesCount
    }
}

func fillBackersCount(p *model.Project) {
    backersCount, err := database.ProjectBackers().Find(bson.M{"projectId" : p.Id}).Query.Count()
    if err == nil {
        p.BackersCount = backersCount
    }
}