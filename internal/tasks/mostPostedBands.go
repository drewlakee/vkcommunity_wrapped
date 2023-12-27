package tasks

import (
	"fmt"
	"sort"
	"vkcommunity_wrapped/internal/models"
	"vkcommunity_wrapped/internal/utils"
)

type MostPostedBandsTask struct{}

func (task *MostPostedBandsTask) Run(context models.Context) {
	type band struct {
		bandName string
		count    int32
	}

	data := make(map[string]int32)
	for _, wallPost := range context.WallPosts {
		artists := make(map[string]bool)
		for _, attachment := range wallPost.Attachments {
			if attachment.Type == models.Audio {
				artists[attachment.Audio.Artist] = true
			}
		}

		for artist, _ := range artists {
			_, ok := data[artist]
			if ok {
				data[artist] = data[artist] + 1
			} else {
				data[artist] = 1
			}
		}
	}

	var bands []band
	for key, value := range data {
		bands = append(bands, band{key, value})
	}

	sort.Slice(bands, func(left, right int) bool {
		return bands[left].count > bands[right].count
	})

	file := utils.CreateNewFile(context, "most_posted_bands.txt")
	defer file.Close()

	file.Write("// Package tasks группы, которые чаще всего попадали в посты\n")
	for i, band := range bands {
		file.Write(fmt.Sprintf("%d: %s - %d\n", i, band.bandName, band.count))
	}
}
