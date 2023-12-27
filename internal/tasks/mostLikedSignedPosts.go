package tasks

import (
	"fmt"
	"sort"
	"vkcommunity_wrapped/internal/models"
	"vkcommunity_wrapped/internal/utils"
)

type MostLikedSignedPostsTask struct{}

func (task *MostLikedSignedPostsTask) Run(context models.Context) {
	type post struct {
		postID int64
		count  int32
	}

	var ownerID int64
	data := make(map[int64]int32)
	for _, wallPost := range context.WallPosts {
		ownerID = wallPost.OwnerId
		if wallPost.SignerID == 0 {
			continue
		}

		data[wallPost.ID] = wallPost.Likes.Count
	}

	var posts []post
	for key, value := range data {
		posts = append(posts, post{key, value})
	}

	sort.Slice(posts, func(left, right int) bool {
		return posts[left].count > posts[right].count
	})

	file := utils.CreateNewFile(context, "most_liked_signed_posts.txt")
	defer file.Close()

	file.Write("// Package tasks самые залайканые посты из предложки\n")
	for i, wallPost := range posts {
		file.Write(fmt.Sprintf("%d: https://vk.com/wall%d_%d - %d\n", i, ownerID, wallPost.postID, wallPost.count))
	}
}
