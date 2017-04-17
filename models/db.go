package models

import (
	"errors"
)

// This file is used to abstract a lot of database manipulations, to ease future migration to other databases.
// All function for adding, modifiying and removing items from the database are all in this file.
// Currently this file uses Sqlite as database and Beego orm as a layer for accessing the database

// GetNewsItems takes the offset and limit and returns an array of NewsItem type and error
func GetNewsItems(offset, limit int) ([]*NewsItem, error) {
	news_items := []*NewsItem{}
	var err error
	if offset == 0 && limit == 0 {
		_, err = o.Raw("SELECT * FROM news_item").QueryRows(&news_items)
	} else if offset == 0 {
		_, err = o.Raw("SELECT * FROM news_item LIMIT ?", limit).QueryRows(&news_items)
	} else if limit == 0 {
		_, err = o.Raw("SELECT * FROM news_item OFFSET ? ", offset).QueryRows(&news_items)
	} else {
		_, err = o.Raw("SELECT * FROM news_item LIMIT ? OFFSET ? ", limit, offset).QueryRows(&news_items)
	}

	if err != nil {
		Verbose("GetNewsItems: %v", err)
		return news_items, err
	}
	return news_items, nil
}

// GetNewsItem fetchs news item by its id
func GetNewsItem(item_id int) (NewsItem, error) {
	Verbose("Get news items")
	news_item := NewsItem{}
	var err error
	if item_id == 0 {
		return news_item, errors.New("invalid news item id")
	}
	_, err = o.Raw("SELECT * FROM news_item WHERE id = ? ", item_id).QueryRows(&news_item)

	if err != nil {
		return news_item, err
	}
	return news_item, nil
}

// GetNewsItemsByWebsiteId fetchs a news_items by its website_id
func GetNewsItemsByWebsiteId(wesite_id int) (error, *NewsItem) {
	news_item := NewsItem{}
	err := o.Raw("SELECT * FROM news_item WHERE website_id = ? ", wesite_id).QueryRow(&news_item)
	if err != nil {
		return err, &news_item
	}
	return nil, &news_item
}

// GetUserByUsername fetchs a user by its username field
func GetUserByUsername(username string) (error, *User) {
	user := User{Username: username}
	err := o.Read(&user, "username")
	if err != nil {
		return err, &user
	}
	return nil, &user
}

// GetWebsites fetchs websites
func GetWebsites() ([]*Website, error) {
	sites := []*Website{}
	_, err := o.Raw("SELECT * FROM website").QueryRows(&sites)
	if err != nil {
		return sites, err
	}
	return sites, nil
}

// GetWebsite fetchs a website
func GetWebsite(site_id int) (*Website, error) {
	site := Website{}
	err := o.Raw("SELECT * FROM website WHERE id = ? ", site_id).QueryRow(&site)
	if err != nil {
		return &site, err
	}
	return &site, nil
}
