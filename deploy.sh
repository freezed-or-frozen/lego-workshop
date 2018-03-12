#!/bin/bash

# Robot IP Address
IP_ROBOT=192.168.0.165

# Uploading files with SSH
echo "Uploading files to Lego Mindstorm EV3 robot..."
echo "  => css directory"
sshpass -p 'maker' scp -r css robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => img directory"
sshpass -p 'maker' scp -r img robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => js directory"
sshpass -p 'maker' scp -r js robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => index.html file"
sshpass -p 'maker' scp index.html robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => help.html file"
sshpass -p 'maker' scp help.html robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => reset_robot.sh file"
sshpass -p 'maker' scp reset_robot.sh robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => start_lego-workshop.sh file"
sshpass -p 'maker' scp start_lego-workshop.sh robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => lego-workshop-arm file"
sshpass -p 'maker' scp lego-workshop-arm robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => todo.py file"
sshpass -p 'maker' scp todo.py robot@$IP_ROBOT:/home/robot/lego-workshop/
