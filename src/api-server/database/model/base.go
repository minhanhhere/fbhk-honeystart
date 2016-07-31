package model

import (
    "gopkg.in/mgo.v2/bson"
    "gitlab.com/hs-api-go/util/timeutil"
    "time"
)

type Base struct {
    Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
    CreatedAt  int64 `json:"createdAt" bson:"createdAt"`
    ModifiedAt int64 `json:"modifiedAt" bson:"modifiedAt"`
    Deleted    bool `json:"deleted" bson:"deleted"`

    exists     bool
}

func (base *Base) construct() {
    base.Id = bson.NewObjectId()
    base.CreatedAt = timeutil.CurrentMillis()
    base.ModifiedAt = timeutil.CurrentMillis()
}

func (base *Base) Construct() {
    base.Id = bson.NewObjectId()
    base.CreatedAt = timeutil.CurrentMillis()
    base.ModifiedAt = timeutil.CurrentMillis()
}

// Satisfy the new tracker interface
func (d *Base) SetIsNew(isNew bool) {
    d.exists = !isNew
}

func (d *Base) IsNew() bool {
    return !d.exists
}

// Satisfy the document interface
func (d *Base) GetId() bson.ObjectId {
    return d.Id
}

func (d *Base) SetId(id bson.ObjectId) {
    d.Id = id
}

func (d *Base) SetCreated(t time.Time) {
    d.CreatedAt = timeutil.GetMillis(t)
}

func (d *Base) SetModified(t time.Time) {
    d.ModifiedAt = timeutil.GetMillis(t)
}
