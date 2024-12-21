# Makefile for running go app in the background

# Define the app command as a variable for easier reuse
GO_CMD = nohup go run . > app.log 2>&1 &

# Target to start the app
start:
    $(GO`_CMD)
    echo "go app started."

# Target to stop the app
stop:
    @pkill -f "go run ."
    echo "go app stopped."

# Target to restart the app
restart: stop start