#!/usr/bin/env bash

#!/bin/sh

SESSION=single_spa
PROJECT_ROOT="$HOME/Projects/private/multi-app/single-spa"

tmux new-session -d -s $SESSION
tmux split-window -h
tmux split-window -v

tmux select-pane -t 0
tmux send-keys "cd $PROJECT_ROOT/root" C-m
tmux send-keys "docker run -p 8080:80 --rm \$(docker build -q .)" C-m

tmux select-pane -t 1
tmux send-keys "cd $PROJECT_ROOT/the-3000" C-m
tmux send-keys " lsof -i :3000 | awk '{print \$2}' | tail -n 1 | xargs kill" C-m
tmux send-keys "npm start -- --https --port 3000" C-m

tmux select-pane -t 2
tmux send-keys "cd $PROJECT_ROOT/the-1337" C-m
tmux send-keys " lsof -i :1337 | awk '{print \$2}' | tail -n 1 | xargs kill" C-m
tmux send-keys "yarn serve" C-m

tmux attach-session -t $SESSION
