// Package forum provides the types that encapsulate the idea of a forum or
// discussion group and the data therein.
package forum

import "time"

const (
	AttachmentsPath   string = "attachments/"
	CommentsPath      string = "comments/"
	ConversationsPath string = "conversations/"
	FollowsPath       string = "follows/"
	ForumsPath        string = "forums/"
	MessagesPath      string = "messages/"
	ProfilesPath      string = "profiles/"
	RolesPath         string = "roles/"
)

// DirIndex provides a way of describing which items were exported, it spares
// the importer from walking a directory tree of potentially millions of items
// by describing them in one file.
// This file should be called `index.json` and found within each exported/type
// directory. i.e. if exported/ is your root, and exported/comments/ holds all
// comments and exported/comments/1.json is comment ID = 1, then we would hope
// to find exported/comments/index.json with a file reference that is:
// {"id":1, "path":"1.json"}
type DirIndex struct {
	Type  string    `json:"type"`
	Files []DirFile `json:"files"`
}

// DirFile describes a single exported item within a child of the exported
// directory
type DirFile struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`

	// Email is used by profiles to allow us to identify duplicates and skip
	// them during concurrent process of non-duplicates
	Email string `json:"email,omitempty"`
}

// ID represents an identifier of a thing within a forum
type ID struct {
	ID int64 `json:"id"`
}

// Profiles represents basic knowledge of a user within a forum. It is presumed
// that all forums have a concept of a username and that users have an email
// address and an identifier. It is not presumed that either the email or
// username is the unique identifier, systems that import users are free to
// merge duplicates based on email address or username.
type Profile struct {
	ID                        int64      `json:"id"`
	Name                      string     `json:"name"`
	Email                     string     `json:"email"`
	DateCreated               time.Time  `json:"dateCreated,omitempty"`
	LastActive                time.Time  `json:"lastActive,omitempty"`
	IPAddress                 string     `json:"ipAddress,omitempty"`
	ReceiveEmailFromAdmins    bool       `json:"receiveEmailFromAdmins,omitempty"`
	ReceiveEmailNotifications bool       `json:"receiveEmailNotifications,omitempty"`
	Banned                    bool       `json:"isBanned,omitempty"`
	Usergroups                []ID       `json:"usergroups,omitempty"`
	Avatar                    Attachment `json:"avatar,omitempty"`
}

/*
Role represents the concept of a group of users that share a set of
permissions. This maps very neatly onto the concept of usergroups that a lot of
forums have.

Users are included into a role either explicitly (this user is in this
group) or implicitly (anyone who has more than 1,500 comments is in this group).

Unless explicitly declared and documented in a source system, export systems
should assume that if a role has criteria for implicit inclusion in a
usergroup, then explicit inclusion does not apply. i.e. If a vBulletin promotion
puts users into a usergroup, then the promotion can be considered a criteria and
exporting systems shouldn't bother listing all users in the individually, as the
assumption is that they were put there by meeting the criteria.
*/
type Role struct {
	ID                int64            `json:"id"`
	Name              string           `json:"name,omitempty"`
	Text              string           `json:"text,omitempty"`
	Banned            bool             `json:"isBanned,omitempty"`
	Moderator         bool             `json:"isModerator,omitempty"`
	ForumPermissions  ForumPermissions `json:"forumPermissions,omitempty"`
	IncludeRegistered bool             `json:"includeRegisteredUsers,omitempty"`
	IncludeGuests     bool             `json:"includeGuests,omitempty"`
	Users             []ID             `json:"users,omitempty"`
	Criteria          []Criterion      `json:"criteria,omitempty"`

	// DefaultRole indicates that this is a role that will be applied to all
	// forums
	DefaultRole bool `json:"default,omitempty"`
}

// ForumPermissions describes the permissions that users within a usergroup have
// on a given forum.
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

/*
Criterion describes an implicity inclusion of a User into a Usergroup.

Criterion belonging to the same usergroup are applied according to the
OrGroup value, where like values are AND and other values are OR.

An example:
   Criterion{OrGroup: 0, Key: "comments", Predicate "ge", Value: 1500}
   Criterion{OrGroup: 0, Key: "is_member", Predicate "eq", Value: true}
   Criterion{OrGroup: 1, Key: "foo", Predicate "eq", Value: "bar"}

Should be equivalent to:
   All users where
        (user.comments >= 1500 AND user.is_member == true)
     OR user.foo == "bar"

It is the responsibility of an importing system to determine the meaning of
the Key field.
*/
type Criterion struct {
	OrGroup   int64       `json:"orGroup"`
	Key       string      `json:"key,omitempty"`
	Predicate string      `json:"predicate"`
	Value     interface{} `json:"value,omitempty"`
}

// PredicateEquals and the other predicates are the range of valid predicates
// for use in Criterion. Valid predicates are determined by the type of the
// Criterion Value, equality applies to all but substr predicates only apply to
// strings.
const (
	PredicateEquals              string = "eq"
	PredicateNotEquals           string = "ne"
	PredicateLessThan            string = "lt"      // Only on numbers and dates
	PredicateLessThanOrEquals    string = "le"      // Only on numbers and dates
	PredicateGreaterThanOrEquals string = "ge"      // Only on numbers and dates
	PredicateGreaterThan         string = "gt"      // Only on numbers and dates
	PredicateSubstring           string = "substr"  // Only on strings
	PredicateNotSubstring        string = "nsubstr" // Only on strings
)

// Forum represents a group/forum/section of a discussion site. This is the
// container for content. It is assumened that usergroup permissions are applied
// generally to the forum level and not to specific items within the forum.
type Forum struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Author       int64  `json:"author,omitempty"`
	Text         string `json:"text,omitempty"`
	DisplayOrder int64  `json:"displayOrder,omitempty"`
	Open         bool   `json:"isOpen,omitempty"`
	Sticky       bool   `json:"isSticky,omitempty"`
	Moderated    bool   `json:"isModerated,omitempty"`
	Deleted      bool   `json:"isDeleted,omitempty"`
	Usergroups   []Role `json:"usergroups,omitempty"`
	Moderators   []ID   `json:"moderators,omitempty"`
}

// Conversation represents a discussion/thread within a forum.
type Conversation struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	ForumID     int64     `json:"forumId, omitempty"`
	Author      int64     `json:"author,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	ViewCount   int64     `json:"viewCount,omitempty"`
	Open        bool      `json:"isOpen,omitempty"`
	Sticky      bool      `json:"isSticky,omitempty"`
	Moderated   bool      `json:"isModerated,omitempty"`
	Deleted     bool      `json:"isDeleted,omitempty"`
}

// Comment represents a post/comment that is attached to a conversation or other
// type of content within a forum or on the wider site (a user may have a
// comment as part of their profile to provide a bio). Comments are optionally
// threaded by indicating which comment they are in reply to. Comments may also
// have different versions, the latest version is presumed to be the live
// version.
type Comment struct {
	ID int64 `json:"id"`
	Association
	InReplyTo   int64            `json:"inReplyTo,omitempty"`
	Author      int64            `json:"author,omitempty"`
	DateCreated time.Time        `json:"dateCreated,omitempty"`
	IPAddress   string           `json:"ipAddress,omitempty"`
	Moderated   bool             `json:"isModerated,omitempty"`
	Deleted     bool             `json:"isDeleted,omitempty"`
	Versions    []CommentVersion `json:"versions"`
}

// CommentVersion encapsulates the body of a comment. A single comment may have
// many historical versions, and the latest version is presumed to be the live
// one. Text is the raw body of the comment, which may be in plain text, bbcode,
// Markdown or HTML or any combination thereof. It is assumed that any Text
// value only contains common markup and not custom markup, by which we mean a
// system exporting a comment is responsible for ensuring the exported value
// is stripped of custom bbcodes, or the applicable conversion to HTML is done
// by the export system.
type CommentVersion struct {
	Editor       int64     `json:"editor"`
	DateModified time.Time `json:"dateModified,omitempty"`
	Headline     string    `json:"headline,omitempty"`
	Text         string    `json:"text"`
	EditReason   string    `json:"editReason,omitempty"`
	IPAddress    string    `json:"ipAddress,omitempty"`
}

// Message describes a private message between one or more people. Different
// forum products handle this differently, for some private messages are no
// different from comments (and use comment identifiers), whereas other systems
// treat private messages as an entirely unique construct. The essence is the
// same though: Some text/markup from 1 person to 1 or many people, with some
// systems allowing a BCC to 1 or many people. Most systems treat private
// messages as if they were SMS messages to a one-time distribution list
type Message struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Author int64  `json:"author,omitempty"`

	// If deleted = true then the sender has deleted their copy.
	Deleted     bool               `json:"isDeleted,omitempty"`
	To          []MessageRecipient `json:"to,omitempty"`
	BCC         []MessageRecipient `json:"bcc,omitempty"`
	InReplyTo   int64              `json:"inReplyTo,omitempty"`
	DateCreated time.Time          `json:"dateCreated,omitempty"`
	IPAddress   string             `json:"ipAddress,omitempty"`
	Versions    []CommentVersion   `json:"versions"`
}

// MessageRecipient tracks recipients of a message and indicates whether the
// recipient still has a copy of the message
type MessageRecipient struct {
	ID int64 `json:"id"`

	// If deleted = true then the recipient has deleted their copy.
	Deleted bool `json:"isDeleted,omitempty"`
	Read    bool `json:"isRead,omitempty"`
}

// Attachment represents a file that may be attached to a comment or other type.
type Attachment struct {
	ID           int64         `json:"id"`
	Author       int64         `json:"author,omitempty"`
	DateCreated  time.Time     `json:"dateCreated,omitempty"`
	Associations []Association `json:"associations,omitempty"`
	Name         string        `json:"name,omitempty"`
	ContentSize  int32         `json:"contentSize,omitempty"`
	ContentURL   string        `json:"contentUrl"`
	MimeType     string        `json:"mimetype"`
	Width        int64         `json:"width,omitempty"`
	Height       int64         `json:"height,omitempty"`
}

// Follow represents a like/follow/subscribe relationship between a user and
// any content on the site
type Follow struct {
	Author               int64          `json:"author"`
	Users                []FollowNotify `json:"users"`
	UsersIgnored         []int64        `json:"usersIgnored"`
	Forums               []FollowNotify `json:"forums"`
	ForumsIgnored        []int64        `json:"forumsIgnored"`
	Conversations        []FollowNotify `json:"conversations"`
	ConversationsIgnored []int64        `json:"conversationsIgnored"`
}

// FollowNotify encapsulates a followed item and whether the user wants an
// explicit (email/SMS) notification when the item is updated.
type FollowNotify struct {
	ID     int64 `json:"id"`
	Notify bool  `json:"notify"`
}

// Association describes any content by type and ID.
// E.g. Association{OnType: "conversation", OnID: 123} means conversation 123
type Association struct {
	OnType string `json:"onType,omitempty"`
	OnID   int64  `json:"onId,omitempty"`
}
