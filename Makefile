.PHONY: build install clean

build:
	@echo "Building..."
	mkdir -p build
	cp -R resources/haSystray.app build/haSystray.app
	mkdir -p build/haSystray.app/Contents/MacOS
	go build -o build/haSystray.app/Contents/MacOS/haSystray .
	@echo "Done Building..."

install:
	@echo "Installing..."
	cp -R build/haSystray.app ~/Applications/haSystray.app
	@echo "Done Installing..."

clean:
	@echo "Cleaning..."
	rm -rf build
	@echo "Done Cleaning..."

update: clean build install
	make clean