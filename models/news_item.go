package models

import (
	"fmt"
	"time"
)

type NewsItem struct {
	Id          int    `json:"id" orm:"auto"`
	Title       string `json:"title" orm:"size(500)"`
	Link        string `json:"link" orm:"unique"`
	WebsiteName string `json:"website_name"`
	WebsiteUrl  string `json:"website_url"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *NewsItem) TableName() string {
	return "news_item"
}

func (this *NewsItem) Valid() bool {
	return (len(this.Title) > 0 && len(this.WebsiteName) > 0 && len(this.WebsiteUrl) > 0)
}

func (this *NewsItem) SetId(id int) {
	this.Id = id
}

func (this *NewsItem) GetId() int {
	return this.Id
}

func (this *NewsItem) String() string {
	return fmt.Sprintf("%s", this.Title)
}

// CRUD operations
