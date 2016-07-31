package model

import "gopkg.in/mgo.v2/bson"

const (
    COLLECTION_PROJECT = "projects"
    COLLECTION_PROJECT_UPDATE = "projects_updates"
    COLLECTION_PROJECT_BACKER = "projects_backers"
)

type (

    Project struct {
        Base `bson:",inline"`
        Name           string `json:"name" bson:"name"`
        Summary        string `json:"summary" bson:"summary"`
        Avatar         string `json:"avatar" bson:"avatar"`
        Bride          *ProjectTarget `json:"bride" bson:"bride"`
        Groom          *ProjectTarget `json:"groom" bson:"groom"`
        Goal           float64 `json:"goal" bson:"goal"`
        Currency       string `json:"currency" bson:"currency"`
        CurrencySymbol string `json:"currencySymbol" bson:"currencySymbol"`
        TimestampStart int64 `json:"timestampStart" bson:"timestampStart"`
        TimestampEnd   int64 `json:"timestampEnd" bson:"timestampEnd"`
        Website        string `json:"website" bson:"website"`
        Story          string `json:"story" bson:"story"`
        Photos         []string `json:"photos" bson:"photos"`
        Videos         []string `json:"videos" bson:"videos"`
        Categories     []string `json:"categories" bson:"categories"`
        Owner          *Team `json:"owner" bson:"-"`
        Packages       []ProjectPackage `json:"packages" bson:"packages"`
        Updates        []ProjectUpdate `json:"updates" bson:"-"`
        Backers        []ProjectBacker `json:"backers" bson:"-"`
        UpdatesCount   int `json:"updatesCount" bson:"-"`
        BackersCount   int `json:"backersCount" bson:"-"`

        OwnerId        bson.ObjectId `json:"ownerId" bson:"ownerId"`
    }

    ProjectTarget struct {
        FirstName string `json:"firstName" bson:"firstName"`
        LastName  string `json:"lastName" bson:"lastName"`
        Avatar    string `json:"avatar" bson:"avatar"`
    }

    ProjectPackage struct {
        Id                bson.ObjectId `json:"id" bson:"id"`
        Name              string `json:"name" bson:"name"`
        Price             float64 `json:"price" bson:"price"`
        Limit             int `json:"limit" bson:"limit"`
        Feature           bool `json:"feature" bson:"feature"`
        Photo             []string `json:"photo" bson:"photo"`
        Description       string `json:"description" bson:"description"`
        TimestampEstimate int64 `json:"timestampEstimate" bson:"timestampEstimate"`
        ShippingIncluded  bool `json:"shippingIncluded" bson:"shippingIncluded"`
        ShippingCountries []string `json:"shippingCountries" bson:"shippingCountries"`
    }

    ProjectUpdate struct {
        Base `bson:",inline"`
        Content   string `json:"content" bson:"content"`
        Owner     *User `json:"owner" bson:"-"`

        OwnerId   bson.ObjectId `json:"ownerId" bson:"ownerId"`
        ProjectId bson.ObjectId `json:"projectId" bson:"projectId"`
    }

    ProjectBacker struct {
        Base `bson:",inline"`
        Package       *ProjectPackage `json:"package" bson:"-"`
        Owner         *User `json:"owner" bson:"-"`
        Quantity      int `json:"quantity" bson:"quantity"`
        Type          string `json:"type" bson:"type"`
        Address       *Address `json:"address" bson:"address"`
        TotalAmount   float64 `json:"totalAmount" bson:"totalAmount"`
        Verified      bool `json:"verified" bson:"verified"`
        TransactionId string `json:"transactionId" bson:"transactionId"`

        OwnerId       bson.ObjectId `json:"ownerId" bson:"ownerId"`
        PackageId     bson.ObjectId `json:"packageId" bson:"packageId"`
        ProjectId     bson.ObjectId `json:"projectId" bson:"projectId"`
    }
)

func NewProjectBacker() *ProjectBacker {
    backer := ProjectBacker{}
    backer.Base.construct()
    return &backer
}

func NewProject() *Project {
    project := Project{}
    project.Base.construct()
    return &project
}