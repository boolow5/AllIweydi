package models

import (
	"fmt"
	"time"
)

type Website struct {
	WebsiteId   int    `json:"website_id" orm:"auto"`
	RootUrl     string `json:"root_url" orm:"unique"`
	Name        string `json:"name" orm:"size(50)"`
	Description string `json:"description" orm:"size(200)"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *Website) TableName() string {
	return "website"
}

func (this *Website) Valid() bool {
	return (len(this.Name) > 0 && len(this.RootUrl) > 0)
}

func (this *Website) SetId(id int) {
	this.WebsiteId = id
}

func (this *Website) GetId() int {
	return this.WebsiteId
}

func (this *Website) String() string {
	return fmt.Sprintf("%s", this.Name)
}

// CRUD operations
