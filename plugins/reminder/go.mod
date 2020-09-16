module github.com/paper2code/golang-telegram-multibot/v2/plugins/reminder

replace github.com/paper2code/golang-telegram-multibot/v2 => ../..

go 1.15

require (
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/paper2code/golang-telegram-multibot/v2 v2.0.0-00010101000000-000000000000
	gopkg.in/telegram-bot-api.v4 v4.6.4
)
