# LEGO Workshop


1 - Introduction
================
LEGO Workshop is a web IDE to develop Python script for the Lego Mindstorm EV3 robot.
It is something between IPython (Python web IDE) and Processing/arduino (simplicity)
There are two main parts :
  * **server side** written in Golang to start and stop Python scripts remotely
  * **client side** written in HTML, CSS, Javascript to edit Python code
Both communicates with websockets.


2 - Functionalities
===================
Main functionalities :
  * edit Python code in a web page (lines, color syntax)
  * execute Python code remotely
  * stop execution
  * tell every client about robot's state (execution, clients connected)

Future functionalities :
  * use nickname
  * add library documentation and help page
  * print robot state (motors, sensors) with a SVG graphic
  * add other language (C++)


3 - Tools
==========
3.1 - Installation
------------------
Assuming you already install ev3dev Linux distribution on an SD card and buy
an USB wifi dongle, you have to install some golang packages and library (for
websockets)
```bash
$ ssh robot@1.2.3.4
$ sudo apt install golang-go golang-go-linux-arm golang-go.tools golang-src (packages de base)
$ go get github.com/gorilla/websocket
```

3.2 - Compilation and execution
-------------------------------
Assuming you already download server.go on the robot :
```bash
$ go build
$ ./lego-workshop
```
You can also execute **lego-workshop** program with Brickman interface and
buttons on the Lego robot :
  * File manager...
  * lego-workshop


4 - ToDo
========
ToDo list :
  * translate source code in english
  * refactor some parts of source code
  * add some unit tests
  * ...


5 - Webography
===============
Some useful links :
  * https://godoc.org/github.com/gorilla/websocket (websocket documentation)
  * https://gist.github.com/tmichel/7390690 (tutorial on gorilla webscoket in golang)
