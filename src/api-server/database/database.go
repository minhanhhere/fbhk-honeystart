package database

import (
    "github.com/maxwellhealth/bongo"
    "gitlab.com/hs-api-go/database/model"
    "gopkg.in/mgo.v2"
    "log"
    "os"
)

const DB_SERVER string = "mongodb://localhost"

var connection *bongo.Connection

func init() {
}

func enableMongoLogger() {
    mgo.SetDebug(true)
    mgo.SetLogger(log.New(os.Stdout, "mgo", log.LstdFlags))
}

func GetConnection() *bongo.Connection {

    var CONFIG *bongo.Config = &bongo.Config{
        ConnectionString: "mongodb://localhost",
        Database:         "bk_database",
    }

    if connection == nil {
        conn, err := bongo.Connect(CONFIG)
        if err != nil {
            panic(err)
        }
        connection = conn
    }

    return connection
}

func Projects() *bongo.Collection {
    return GetConnection().Collection(model.COLLECTION_PROJECT)
}

func ProjectUpdates() *bongo.Collection {
    return GetConnection().Collection(model.COLLECTION_PROJECT_UPDATE)
}

func ProjectBackers() *bongo.Collection {
    return GetConnection().Collection(model.COLLECTION_PROJECT_BACKER)
}

func Teams() *bongo.Collection {
    return GetConnection().Collection(model.COLLECTION_TEAM)
}

func Users() *bongo.Collection {
    return GetConnection().Collection(model.COLLECTION_USER)
}

func Versions() *bongo.Collection {
    return GetConnection().Collection(model.COLLECTION_VERSION)
}