package models

type UserLikes struct {
	// postId to list of users
	PostLikes map[int32][]int64
}

func (userLikes *UserLikes) GetLikes(postId int32) *[]int64 {
	val, ok := userLikes.PostLikes[postId]
	if ok {
		return &val
	} else {
		return &[]int64{}
	}
}

type CommentLikes struct {
	Count int32
}

type Comment struct {
	FromId int64 `json:"from_id"`
	Text   string
	Likes  CommentLikes
}

type UserComments struct {
	// postId to list of comments
	PostComments map[int32][]Comment
}

func (userComments *UserComments) GetComments(postId int32) *[]Comment {
	val, ok := userComments.PostComments[postId]
	if ok {
		return &val
	} else {
		return &[]Comment{}
	}
}

type WallPostType string

const (
	Photo WallPostType = "photo"
	Audio WallPostType = "audio"
)

type WallPostPhoto struct {
	OwnerId int64 `json:"owner_id"`
	ID      int64 `json:"id"`
}

type WallPostAudio struct {
	ID     int64 `json:"id"`
	Artist string
	Title  string
}

type WallPostAttachment struct {
	Type WallPostType

	// nullable
	Photo WallPostPhoto
	// nullable
	Audio WallPostAudio
}

type WallPostLikes struct {
	Count int32
}

type WallPostReposts struct {
	Count int32
}

type WallPostViews struct {
	Count int32
}

type WallPost struct {
	ID          int64
	OwnerId     int64 `json:"owner_id"`
	CreatedBy   int64 `json:"created_by"`
	SignerID    int64 `json:"signer_id"`
	Likes       WallPostLikes
	Reposts     WallPostReposts
	Views       WallPostViews
	Attachments []WallPostAttachment
}

type Context struct {
	UserLikes    UserLikes
	UserComments UserComments
	WallPosts    []WallPost
	OutputDir    string
}
