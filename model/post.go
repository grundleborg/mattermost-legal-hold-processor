package model

import "time"

// Post represents one post and its associated data as required for a legal hold record.
// It must be kept in sync with model.LegalHoldPost from mattermost-plugin-legal-hold
type Post struct {

	// From Team
	TeamID          string `csv:"TeamId"`
	TeamName        string `csv:"TeamName"`
	TeamDisplayName string `csv:"TeamDisplayName"`

	// From Channel
	ChannelName        string `csv:"ChannelName"`
	ChannelDisplayName string `csv:"ChannelDisplayName"`
	ChannelType        string `csv:"ChannelType"`

	// From User
	UserUsername string `csv:"UserUsername"`
	UserEmail    string `csv:"UserEmail"`
	UserNickname string `csv:"UserNickname"`

	// From Post
	PostID         string `csv:"PostId"`
	PostCreateAt   int64  `csv:"PostCreateAt"`
	PostUpdateAt   int64  `csv:"PostUpdateAt"`
	PostDeleteAt   int64  `csv:"PostDeleteAt"`
	PostRootID     string `csv:"PostRootId"`
	PostOriginalID string `csv:"PostOriginalId"`
	PostMessage    string `csv:"PostMessage"`
	PostType       string `csv:"PostType"`
	PostProps      string `csv:"PostProps"`
	PostHashtags   string `csv:"PostHashtags"`
	PostFileIDs    string `csv:"PostFileIds"`

	IsBot bool `csv:"IsBot"`
}

// PrintCreateAt prints the CreateAt time in a human-readable format.
func (p Post) PrintCreateAt() string {
	t := time.Unix(0, p.PostCreateAt*int64(time.Millisecond))
	return t.Format("15:04 on 2006-01-02")
}
