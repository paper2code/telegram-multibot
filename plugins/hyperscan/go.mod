module github.com/paper2code/golang-telegram-multibot/v2/plugins/hyperscan

replace github.com/paper2code/golang-telegram-multibot/v2 => ../..

go 1.15

require (
	github.com/flier/gohs v1.0.0
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/paper2code/golang-telegram-multibot/v2 v2.0.0-00010101000000-000000000000
	gopkg.in/telegram-bot-api.v4 v4.6.4
)
