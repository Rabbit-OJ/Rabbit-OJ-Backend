package websocket

import (
	"Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/submission"
)

type Hub struct {
	JudgeHub   *submission.Hub
	ContestHub *contest.Hub
}

func newHub() *Hub {
	hub := Hub{
		JudgeHub:   submission.NewJudgeHub(),
		ContestHub: contest.NewContestHub(),
	}
	return &hub
}
