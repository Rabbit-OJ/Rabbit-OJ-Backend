package websocket

import (
	"Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/judger"
)

type Hub struct {
	JudgeHub   *judger.Hub
	ContestHub *contest.Hub
}

func newHub() *Hub {
	hub := Hub{
		JudgeHub:   judger.NewJudgeHub(),
		ContestHub: contest.NewContestHub(),
	}
	return &hub
}
