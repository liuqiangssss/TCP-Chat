package main

const (
	CMD_NICK int = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type commandID = int

type command struct {
	id     commandID
	client *client
	args   []string
}
