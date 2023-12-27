package tasks

import (
	"fmt"
	"sort"
	"vkcommunity_wrapped/internal/models"
	"vkcommunity_wrapped/internal/utils"
)

type MostLikingUsersTask struct{}

func (task *MostLikingUsersTask) Run(context models.Context) {
	type user struct {
		userID int64
		count  int32
	}

	data := make(map[int64]int32)
	for _, likesFrom := range context.UserLikes.PostLikes {
		for _, userID := range likesFrom {
			_, ok := data[userID]
			if ok {
				data[userID] = data[userID] + 1
			} else {
				data[userID] = 1
			}
		}
	}

	var users []user
	for key, value := range data {
		users = append(users, user{key, value})
	}

	sort.Slice(users, func(left, right int) bool {
		return users[left].count > users[right].count
	})

	file := utils.CreateNewFile(context, "most_liking_users.txt")
	defer file.Close()

	file.Write("// Package tasks больше всего поставили лайков\n")
	for i, user := range users {
		file.Write(fmt.Sprintf("%d: https://vk.com/id%d - %d\n", i, user.userID, user.count))
	}
}
