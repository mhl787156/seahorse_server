package models

import (
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionUsers = "users"
        CollectionSessions = "session"
)

type User struct {
  Id        string   `gorethink:"id,omitempty" json:"id"`
  FId       string   `gorethink:"fid" json:"-"`
  Name      string   `gorethink:"name" json:"name"`
  Avatar    string   `gorethink:"avatar" json:"avatar"`
  Email     string   `gorethink:"email" json:"email"`
  Private   bool     `gorethink:"private" json:"private"`
  Projects  []string `gorethink:"projects" json:"-"`
  Following []string `gorethink:"following" json:"-"`
}

type Session struct {
  SessionKey string   `gorethink:"id"`
  UID        string   `gorethink:"uid"`
  Expires    int64    `gorethink:"expires_at"`
}

