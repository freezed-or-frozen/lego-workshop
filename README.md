# LEGO Workshop


1 - Introduction
================
LEGO Workshop is a web IDE to develop Python script for the Lego Mindstorm EV3 robot.
It is something between IPython (Python web IDE) and Processing/arduino (simplicity)
There are two main parts :
  * **server side** written in Golang to start and stop Python scripts remotely
  * **client side** written in HTML, CSS, Javascript to edit Python code
Both communicates with websockets.


2 - Goals and functionalities
=============================
Main functionalities :
  - [x] edit Python code in a web page (lines, color syntax)
  - [x] execute Python code remotely
  - [x] stop execution
  - [X] init robot state after stopping execution (motors...)
  - [x] tell every client about robot's state (execution, clients connected)
  - [X] add an help page with code example and documentation

Future functionalities :
  - [X] use nickname
  - [ ] print robot state (motors, sensors) with a SVG graphic
  - [ ] add other language (C++)


3 - Installation
================
3.1 - From binary version
-------------------------
Assuming you already install ev3dev Linux distribution on an SD card and buy
an USB wifi dongle, you just need to upload **lego-workshop** binary to
the robot with SSH
```bash
$ git clone https://github.com/freezed-or-frozen/lego-workshop.git
$ cd lego-workshop
$ scp lego-workshop robot@1.2.3.4:/home/robot/ (password=maker)
```
Now you can execute **lego-workshop** program with Brickman interface and
buttons on the Lego robot :
  * File manager...
  * lego-workshop

3.2 - From source version
-------------------------
On your computer, you need to install some golang packages and library (for websockets)
```bash
$ sudo apt install golang-go golang-go-linux-arm golang-go.tools golang-src
$ go get github.com/gorilla/websocket
```

Clone this repository and compile/package a binary
```bash
$ git clone https://github.com/freezed-or-frozen/lego-workshop.git
$ cd lego-workshop
$ GOOS=linux GOARCH=arm GOARM=5 go build
$ sudo GOPATH=$HOME/golang GOOS=linux GOARCH=arm GOARM=5 go install
```

Upload the binary to the robot
```bash
$ cd $HOME/golang/bin/linux_arm
$ scp lego-workshop robot@1.2.3.4:/home/robot/ (password=maker)
```
Now you can execute **lego-workshop** program with Brickman interface and
buttons on the Lego robot :
  * File manager...
  * lego-workshop


4 - ToDo
========
ToDo list :
  * fix bugs
  * translate source code in english
  * refactor some parts of source code
  * add some unit tests
  * ...


5 - Webography
===============
Some useful links :
  * https://godoc.org/github.com/gorilla/websocket (websocket documentation)
  * https://gist.github.com/tmichel/7390690 (tutorial on gorilla webscoket in golang)
  * https://mattrajca.github.io/GoEV3/ (robot programming with golang)
  * http://www.ev3dev.org/docs/tutorials/tacho-motors/
