# LEGO Workshop


1 - Introduction
================
Application Golang pour proposer un IDE web de développement pour les robots
Lego Mindstorm EV3.


2 - Cahier des charges
======================
Liste des fonctionnalités principales :
  * générer une page web pour l'ide avec l'adresse IP du robot LEGO
  * récupérer le code source
  * écrire le code source récupéré dans un script Python
  * éxecuter ce script Python
  * gérer plusieurs websocket en même temps
  * arrêter un script Python 
  * informer tous les clients de l'état du robot

Liste des fonctionnalités secondaires :
  * fg,fg,fg,
  * ...


3 - Outils
==========
Installation des outils sur le robot Lego Mindstorm EV3
```bash
$ sudo apt install golang-go golang-go-linux-arm golang-go.tools golang-src (packages de base)
```

Installation d'une librairie gérant les websockets
```bash
$ sudo apt install golang-websocket-dev (option 1)
$ sudo apt install golang-golang-x-net-dev (option 2)
$ go get github.com/gorilla/websocket (option 3)
```

Options 3 retenue
  * https://godoc.org/github.com/gorilla/websocket


4 - Webographie
===============
Quelques liens pour aller plus loin
  * https://gist.github.com/tmichel/7390690 (démonstration de l'utilisation des websocket gorilla en golang)
