PROJECT_NAME:=multibot
GO_GET:=go get
PLUGINS:=log_messages save_messages filer

all: multibot subdirs

multibot:
	@echo "Building ${PROJECT_NAME}"
	@go build

subdirs: bin_plugins $(PLUGINS)

$(PLUGINS):
	@$(MAKE) -C plugins/$@
	@$(MAKE) -C plugins/$@ install

bin_plugins:
	@install -m 0755 -d bin_plugins

clean:
	@rm -rf bin_plugins
	@rm -rf ${PROJECT_NAME}
