package main

import (
	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

const (
	commandsAddCommand  = "ssh_commands_add"
	commandsDelCommand  = "ssh_commands_del"
	commandsListCommand = "ssh_commands_list"

	commandsAddKeyboard  = "Add SSH command"
	commandsDelKeyboard  = "Remove SSH command"
	commandsListKeyboard = "List of SSH Commands"
)

var (
	ctx *context.MultiBotContext
)

func InitPlugin(mbc *context.MultiBotContext) (err error) {
	ctx = mbc
	if err = ctx.DBCreateTable(&UserCommand{}); err != nil {
		ctx.Log().Errorf("Unable to create table for ssh user commands: %s", err)
		return
	}

	return
}

func GetName() string {
	return "ssh"
}

func GetDescription() string {
	return "Plugin for run command over SSH on remote hosts periodicaly and notification about it"
}

func GetCommands() []string {
	return []string{
		commandsAddCommand,
		commandsDelCommand,
		commandsListCommand,
	}
}
