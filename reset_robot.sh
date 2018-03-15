#!/bin/bash

# Process killing
killall python3
killall aplay
killall espeak
killall beep
killall amixer

# Motors reset
echo reset > /sys/class/tacho-motor/motor0/command
echo reset > /sys/class/tacho-motor/motor1/command
echo reset > /sys/class/tacho-motor/motor2/command

# Speaker reset
echo 0 > /sys/devices/platform/snd-legoev3/tone
