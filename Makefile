# General
PROJECT_NAME:=tg-multibot
PLUGINS:=log_messages save_messages filer

all: multibot subdirs

multibot:
	@echo "Building ${PROJECT_NAME}"
	@go build -o ./bin/${PROJECT_NAME} ./cmd/${PROJECT_NAME}

subdirs: plugins $(PLUGINS)

$(PLUGINS):
	@$(MAKE) -C plugins/$@
	@$(MAKE) -C plugins/$@ install

plugins:
	@install -m 0755 -d ./shared/plugins

clean:
	@rm -rf plugins
	@rm -rf bin/${PROJECT_NAME}
