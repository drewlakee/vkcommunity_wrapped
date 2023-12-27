package tasks

import "vkcommunity_wrapped/internal/models"

type Task interface {
	Run(context models.Context)
}
