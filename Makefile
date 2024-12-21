APP_NAME=myapp       # A name to represent the application
PID_FILE=app.pid     # File to store the PID of the running process
LOG_FILE=app.log     # Log file for the application

# Run the application in the background
start:
	@if [ -f $(PID_FILE) ]; then \
		echo "Application is already running (PID=$$(cat $(PID_FILE)))"; \
	else \
		echo "Starting application..."; \
		nohup go run . > $(LOG_FILE) 2>&1 & echo $$! > $(PID_FILE); \
		echo "Application started (PID=$$(cat $(PID_FILE)))"; \
	fi

# Stop the running application
stop:
	@if [ -f $(PID_FILE) ]; then \
		echo "Stopping application..."; \
		kill -9 $$(cat $(PID_FILE)) && rm -f $(PID_FILE); \
		echo "Application stopped."; \
	else \
		echo "Application is not running."; \
	fi

# Clean up logs and PID file
clean: stop
	rm -f $(PID_FILE) $(LOG_FILE)

# Show logs
logs:
	@tail -f $(LOG_FILE)
