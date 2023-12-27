package tasks

import (
	"fmt"
	"sort"
	"vkcommunity_wrapped/internal/models"
	"vkcommunity_wrapped/internal/utils"
)

type MostRepostedPostsTask struct{}

func (task *MostRepostedPostsTask) Run(context models.Context) {
	sorted := append([]models.WallPost{}, context.WallPosts...)
	sort.Slice(sorted, func(left, right int) bool {
		return sorted[left].Reposts.Count > sorted[right].Reposts.Count
	})

	file := utils.CreateNewFile(context, "most_reposted_posts.txt")
	defer file.Close()

	file.Write("// Package tasks больше всего затащили себе на стену\n")
	for i, post := range sorted {
		file.Write(fmt.Sprintf("%d: https://vk.com/wall%d_%d - %d\n", i, post.OwnerId, post.ID, post.Reposts.Count))
	}
}
