#!/bin/bash
while true; do
    clear
    echo "Monitoring open files for PID 37543"
    echo "Press [Ctrl+C] to stop..."
    echo ""
    echo "Number of open files:"
    sudo lsof -p 37543
    sleep 0.5
done