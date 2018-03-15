#!/bin/bash

# Check number of arguments
if [ "$#" -ne 1 ]; then
    echo "/!\\ IP address of robot missing /!\\"
    echo "Usage : ./deploy.sh 1.2.3.4"
    exit 1
fi

# Robot IP Address
IP_ROBOT=$1

# Uploading files with SSH
echo "Uploading files to Lego Mindstorm EV3 robot..."
echo "  => delete old directory"
sshpass -p 'maker' ssh robot@$IP_ROBOT 'rm -Rf /home/robot/lego-workshop'
echo "  => create a new directory"
sshpass -p 'maker' ssh robot@$IP_ROBOT 'mkdir /home/robot/lego-workshop'
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
cp ../../bin/linux_arm/lego-workshop lego-workshop-arm
sshpass -p 'maker' scp lego-workshop-arm robot@$IP_ROBOT:/home/robot/lego-workshop/
echo "  => todo.py file"
sshpass -p 'maker' scp todo.py robot@$IP_ROBOT:/home/robot/lego-workshop/
