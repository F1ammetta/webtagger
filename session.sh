#!/bin/bash

SESSION_NAME="webtagger"
WORKING_DIR="./"

tmux new-session -s $SESSION_NAME -c $WORKING_DIR -d
tmux send-keys -t $SESSION_NAME:0 "nix develop" C-m

# Create additional windows
tmux new-window -t $SESSION_NAME:1 -n 'window1' -c $WORKING_DIR
tmux send-keys -t $SESSION_NAME:1 "nix develop" C-m
tmux new-window -t $SESSION_NAME:2 -n 'window2' -c $WORKING_DIR
tmux send-keys -t $SESSION_NAME:2 "nix develop" C-m
tmux new-window -t $SESSION_NAME:3 -n 'window3' -c $WORKING_DIR
tmux send-keys -t $SESSION_NAME:3 "nix develop" C-m

tmux switch-client -t $SESSION_NAME
