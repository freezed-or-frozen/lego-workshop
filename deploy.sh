#!/bin/bash

# Robot IP Address
$IP_ROBOT=192.168.0.165

# Uploading files with SSH
echo "Uploading files to Lego Mindstorm EV3 robot..."
sshpass -p 'maker' scp -r css robot@$(IP_ROBOT):/home/robot/lego-workshop/
sshpass -p 'maker' scp -r img robot@$(IP_ROBOT):/home/robot/lego-workshop/
sshpass -p 'maker' scp -r js robot@$(IP_ROBOT):/home/robot/lego-workshop/
sshpass -p 'maker' scp index.html robot@$(IP_ROBOT):/home/robot/lego-workshop/
sshpass -p 'maker' scp help.html robot@$(IP_ROBOT):/home/robot/lego-workshop/
sshpass -p 'maker' scp reset_robot.sh robot@$(IP_ROBOT):/home/robot/lego-workshop/
sshpass -p 'maker' scp start_lego-workshop.sh robot@$(IP_ROBOT):/home/robot/lego-workshop/
