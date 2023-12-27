package tasks

import (
	"fmt"
	"sort"
	"vkcommunity_wrapped/internal/models"
	"vkcommunity_wrapped/internal/utils"
)

type MostLikedPostsTask struct{}

func (task *MostLikedPostsTask) Run(context models.Context) {
	sorted := append([]models.WallPost{}, context.WallPosts...)
	sort.Slice(sorted, func(left, right int) bool {
		return sorted[left].Likes.Count > sorted[right].Likes.Count
	})

	file := utils.CreateNewFile(context, "most_liked_posts.txt")
	defer file.Close()

	file.Write("// Package tasks больше всего лайков на постах\n")
	for i, post := range sorted {
		file.Write(fmt.Sprintf("%d: https://vk.com/wall%d_%d - %d\n", i, post.OwnerId, post.ID, post.Likes.Count))
	}
}
