# General
PROJECT_NAME:=tg-multibot
PLUGINS:=core-log_messages core-save_messages core-filer core-ssh paper2code-mashup paper2code-search paper2code-trends 

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
