#!/bin/bash

# Reset des moteurs
echo reset > /sys/class/tacho-motor/motor0/command
echo reset > /sys/class/tacho-motor/motor1/command
echo reset > /sys/class/tacho-motor/motor2/command

# Reset du speaker
echo 0 > /sys/devices/platform/snd-legoev3/tone
