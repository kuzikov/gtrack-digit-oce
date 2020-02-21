all:
	@echo "building image..."
	@docker build -t gtracker . >> log.txt
	@echo "OK!"
	@echo "Finding image...\n"
	@echo "- - - - - - - - -"
	@docker images | grep gtracker
	@echo "- - - - - - - - -"
	@echo "Running container in daemon mode..."
	@docker run -d -p 80:6600 gtracker | tee -a log.txt
	@echo "Done."