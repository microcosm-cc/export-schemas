package forum

import "time"

type Id struct {
	Id int64 `json:"id"`
}

type Association struct {
	OnType string `json:"onType,omitempty"`
	OnId   int64  `json:"onId,omitempty"`
}

type User struct {
	Id                        int64     `json:"id"`
	Name                      string    `json:"name"`
	Email                     string    `json:"email"`
	DateCreated               time.Time `json:"dateCreated,omitempty"`
	LastActive                time.Time `json:"lastActive,omitempty"`
	IpAddress                 string    `json:"ipAddress,omitempty"`
	ReceiveEmailFromAdmins    bool      `json:"receiveEmailFromAdmins,omitempty"`
	ReceiveEmailNotifications bool      `json:"receiveEmailNotifications,omitempty"`
	Banned                    bool      `json:"isBanned,omitempty"`
	Usergroups                []Id      `json:"usergroups,omitempty"`
}

type Usergroup struct {
	Id               int64            `json:"id"`
	Name             string           `json:"name"`
	Text             string           `json:"text,omitempty"`
	Banned           bool             `json:"isBanned,omitempty"`
	Moderator        bool             `json:"isModerator,omitempty"`
	ForumPermissions ForumPermissions `json:"forumPermissions,omitempty"`
	Users            []Id             `json:"users,omitempty"`
}

type ForumPermissions struct {
	View         bool `json:"canView,omitempty"`
	PostNew      bool `json:"canPostNew,omitempty"`
	EditOwn      bool `json:"canEditOwn,omitempty"`
	EditOthers   bool `json:"canEditOthers,omitempty"`
	DeleteOwn    bool `json:"canDeleteOwn,omitempty"`
	DeleteOthers bool `json:"canDeleteOthers,omitempty"`
	CloseOwn     bool `json:"canCloseOwn,omitempty"`
	OpenOwn      bool `json:"canOpenOwn,omitempty"`
}

type Forum struct {
	Id           int64       `json:"id"`
	Name         string      `json:"name"`
	Text         string      `json:"text,omitempty"`
	DisplayOrder int32       `json:"displayOrder,omitempty"`
	Open         bool        `json:"isOpen,omitempty"`
	Sticky       bool        `json:"isSticky,omitempty"`
	Moderated    bool        `json:"isModerated,omitempty"`
	Deleted      bool        `json:"isDeleted,omitempty"`
	Usergroups   []Usergroup `json:"usergroups,omitempty"`
}

type Conversation struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	ForumId         int64     `json:"forumId, omitempty"`
	Author          int64     `json:"author,omitempty"`
	DateCreated     string    `json:"dateCreated,omitempty"`
	DateCreatedTime time.Time `json:"-"`
	ViewCount       int64     `json:"viewCount,omitempty"`
	Open            bool      `json:"isOpen,omitempty"`
	Sticky          bool      `json:"isSticky,omitempty"`
	Moderated       bool      `json:"isModerated,omitempty"`
	Deleted         bool      `json:"isDeleted,omitempty"`
}

type Comment struct {
	Id              int64            `json:"id"`
	OnType          string           `json:"onType,omitempty"`
	OnId            int64            `json:"onId,omitempty"`
	InReplyTo       int64            `json:"inReplyTo,omitempty"`
	Author          int64            `json:"author,omitempty"`
	DateCreated     string           `json:"dateCreated,omitempty"`
	DateCreatedTime time.Time        `json:"-"`
	IpAddress       string           `json:"ipAddress,omitempty"`
	Moderated       bool             `json:"isModerated,omitempty"`
	Deleted         bool             `json:"isDeleted,omitempty"`
	Versions        []CommentVersion `json:"versions"`
}

type CommentVersion struct {
	Editor           int64     `json:"editor"`
	DateModified     string    `json:"dateModified,omitempty"`
	DateModifiedTime time.Time `json:"-"`
	Headline         string    `json:"headline,omitempty"`
	Text             string    `json:"text"`
	EditReason       string    `json:"editReason,omitempty"`
	IpAddress        string    `json:"ipAddress,omitempty"`
}

type Message struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Users []Id   `json:"users,omitempty"`
}

type Attachment struct {
	Id              int64     `json:"id"`
	OnType          string    `json:"onType,omitempty"`
	OnId            int64     `json:"onId,omitempty"`
	Author          int64     `json:"author,omitempty"`
	DateCreated     string    `json:"dateCreated,omitempty"`
	DateCreatedTime time.Time `json:"-"`
	Name            int64     `json:"name,omitempty"`
	ContentSize     uint64    `json:"contentSize,omitempty"`
	ContentUrl      string    `json:"contentUrl"`
}

type Follow struct {
	Author  int64         `json:"author"`
	Follows []Association `json:"follows"`
}
