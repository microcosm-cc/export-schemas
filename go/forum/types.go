// Package forum provides the types that encapsulate the idea of a forum or
// discussion group and the data therein.
package forum

import "time"

// ID represents an identifier of a thing within a forum
type ID struct {
	ID int64 `json:"id"`
}

// User represents basic knowledge of a user within a forum. It is presumed that
// all forums have a concept of a username and that users have an email address
// and an identifier
type User struct {
	ID                        int64     `json:"id"`
	Name                      string    `json:"name"`
	Email                     string    `json:"email"`
	DateCreated               time.Time `json:"dateCreated,omitempty"`
	LastActive                time.Time `json:"lastActive,omitempty"`
	IPAddress                 string    `json:"ipAddress,omitempty"`
	ReceiveEmailFromAdmins    bool      `json:"receiveEmailFromAdmins,omitempty"`
	ReceiveEmailNotifications bool      `json:"receiveEmailNotifications,omitempty"`
	Banned                    bool      `json:"isBanned,omitempty"`
	Usergroups                []ID      `json:"usergroups,omitempty"`
}

/*
Usergroup represents the concept of a group of users that share a set of
permissions.

Users are included into a usergroup either explicitly (this user is in this
group) or implicitly (anyone who has more than 1,500 comments is in this group).

Unless explicitly declared and documented in a source system, exports systems
should assume that if a usergroup has criteria for implicit inclusion in a
usergroup, that explicit inclusion does not apply.

An example:
   If vBulletin has usergroup promotions defined to automatically move users
   into a usergroup, then an export system need only define the promotions that
   would move someone into this usergroup as criteria of this usergroup to
   implicitly include all of those users. You do not need to explicitly
   list all of the users within a usergroup if criteria takes care of it.
*/
type Usergroup struct {
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
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	Text         string      `json:"text,omitempty"`
	DisplayOrder int64       `json:"displayOrder,omitempty"`
	Open         bool        `json:"isOpen,omitempty"`
	Sticky       bool        `json:"isSticky,omitempty"`
	Moderated    bool        `json:"isModerated,omitempty"`
	Deleted      bool        `json:"isDeleted,omitempty"`
	Usergroups   []Usergroup `json:"usergroups,omitempty"`
	Moderators   []ID        `json:"moderators,omitempty"`
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

// Message describes a private message between one or more people
type Message struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Users []ID   `json:"users,omitempty"`
}

// Attachment represents a file that may be attached to a comment or other type.
type Attachment struct {
	ID int64 `json:"id"`
	Association
	Author      int64     `json:"author,omitempty"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	Name        int64     `json:"name,omitempty"`
	ContentSize uint64    `json:"contentSize,omitempty"`
	ContentURL  string    `json:"contentUrl"`
}

// Follow represents a like/follow/subscribe relationship between a user and
// any content on the site
type Follow struct {
	Author  int64         `json:"author"`
	Follows []Association `json:"follows"`
}

// Association describes any content by type and ID.
// E.g. Association{OnType: "conversation", OnID: 123} means conversation 123
type Association struct {
	OnType string `json:"onType,omitempty"`
	OnID   int64  `json:"onId,omitempty"`
}
