package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

// BotPlugin struct for store one plugin
type BotPlugin struct {
	Name                string
	Description         string
	Commands            []string
	Keyboard            *tgbotapi.ReplyKeyboardMarkup
	EachUpdateHandler   func(tgbotapi.Update) error
	RunCommandHandler   func(string, tgbotapi.Update) error
	StartCommandHandler func(tgbotapi.Update) error
}

var (
	botPlugins          map[string]*BotPlugin
	botPluginsByCommand map[string]*BotPlugin
)

// LoadPlugins load all plugins from directory
func LoadPlugins() (err error) {
	var pluginFiles []os.FileInfo
	botPlugins = make(map[string]*BotPlugin)
	botPluginsByCommand = make(map[string]*BotPlugin)

	log.Debugf("Try to load plugins from: %s", options.PluginDir)

	// pluginFiles, err := filepath.Glob(fmt.Sprintf("%s/*.so", options.PluginDir))
	// if err != nil {
	// 	return
	// }
	if pluginFiles, err = ioutil.ReadDir(options.PluginDir); err != nil {
		return
	}

	for _, pluginFile := range pluginFiles {
		if pluginFile.IsDir() {
			log.Debugf("pluginFile.IsDir: %s", pluginFile.Name())
			continue
		}
		if !strings.HasSuffix(pluginFile.Name(), ".so") {
			log.Debugf("pluginFile.IsNotPlugin: %s", pluginFile.Name())
			continue
		}
		log.Debugf("Try to load plugin: %s", pluginFile.Name())
		fullPath := filepath.Join(options.PluginDir, pluginFile.Name())

		var (
			p                   *plugin.Plugin
			initPlugin          plugin.Symbol
			getName             plugin.Symbol
			getDescription      plugin.Symbol
			getCommands         plugin.Symbol
			getKeyboard         plugin.Symbol
			updateHandler       plugin.Symbol
			runCommandHandler   plugin.Symbol
			startCommandHandler plugin.Symbol
		)
		if p, err = plugin.Open(fullPath); err != nil {
			return
		}
		if initPlugin, err = p.Lookup("InitPlugin"); err != nil {
			return
		}
		if getName, err = p.Lookup("GetName"); err != nil {
			return
		}
		if getDescription, err = p.Lookup("GetDescription"); err != nil {
			return
		}
		if getCommands, err = p.Lookup("GetCommands"); err != nil {
			return
		}
		if getKeyboard, err = p.Lookup("GetKeyboard"); err != nil {
			return
		}
		if updateHandler, err = p.Lookup("UpdateHandler"); err != nil {
			return
		}
		if runCommandHandler, err = p.Lookup("RunCommand"); err != nil {
			return
		}
		if startCommandHandler, err = p.Lookup("StartCommand"); err != nil {
			return
		}

		botPlugin := &BotPlugin{
			Name:                getName.(func() string)(),
			Description:         getDescription.(func() string)(),
			Commands:            getCommands.(func() []string)(),
			Keyboard:            getKeyboard.(func() *tgbotapi.ReplyKeyboardMarkup)(),
			EachUpdateHandler:   updateHandler.(func(tgbotapi.Update) error),
			RunCommandHandler:   runCommandHandler.(func(string, tgbotapi.Update) error),
			StartCommandHandler: startCommandHandler.(func(tgbotapi.Update) error),
		}
		if _, ok := botPlugins[botPlugin.Name]; ok {
			return fmt.Errorf("plugin %s already exists", botPlugin.Name)
		}
		botPlugins[botPlugin.Name] = botPlugin

		for _, cmd := range botPlugin.Commands {
			rcmd := fmt.Sprintf("%s_%s", botPlugin.Name, cmd)
			if pl, ok := botPluginsByCommand[rcmd]; ok {
				return fmt.Errorf("we have command with name \"%s\" in plugin \"%s\"", rcmd, pl.Name)
			}
			botPluginsByCommand[rcmd] = botPlugin
			log.Debugf("Set command \"%s\" for plugin \"%s\"", rcmd, botPlugin.Name)
		}

		// try to get plugins settings
		loadPLuginConfig(botPlugin.Name)

		if err = initPlugin.(func(*context.MultiBotContext) error)(botContext); err != nil {
			return
		}
		log.Debugf("Loaded plugin: %s (%s)", botPlugin.Name, botPlugin.Description)
	}

	return
}
