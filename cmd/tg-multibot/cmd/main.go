package cmd

import (
	"flag"
	"sync"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"
	"gorm.io/gorm"

	"github.com/paper2code/golang-telegram-multibot/v2/pkg/context"
)

var (
	configName string
	wg         sync.WaitGroup
	botContext *context.MultiBotContext
	db         *gorm.DB
	options    *context.Options
)

/*
func init() {
	flag.StringVar(&configName, "config", "multibot", "configuration file name")
}

func main() {
	var err error
	flag.Parse()

	if options, err = LoadConfig(configName); err != nil {
		log.Fatalf("Unable to load configuration file %s: %s", configName, err)
	}
	if level, err := log.ParseLevel(options.LogLevel); err != nil {
		log.Warnf("Unable to parse log level %s: %s", options.LogLevel, err)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(level)
	}

	if db, err = InitDatabase(); err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	if bot, err = tgbotapi.NewBotAPI(options.APIKey); err != nil {
		log.Fatalf("Unable to initialize telegram bot: %s", err)
	}
	log.Debug("Telegram bot initialized sucessful")
	botContext = context.InitContext(db, bot, options, log.StandardLogger())

	if err = LoadPlugins(); err != nil {
		log.Fatalf("Unable to load plugins: %s", err)
	}

	wg.Add(1)
	if err = BotServe(); err != nil {
		log.Fatalf("Unable to server bot: %s", err)
	}

	log.Warnf("Application started...")
}
*/
