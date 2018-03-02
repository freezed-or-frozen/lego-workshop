// Package du programme
package main

// Librairies à utiliser
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"bytes"
	"github.com/gorilla/websocket" // option 3
)


// Pour représenter une trame
type Trame struct {
	Source string `json:"source"`
	Action string `json:"action"`
	Details string `json:"details"`
	ExecPID int `json:"exec_pid"`
	ExecUID string `json:"exec_uid"`
	Clients int `json:"clients"`
}


// Configuration de la webscoket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


// Fonction pour lancer un script Python dans un nouveau processus
func writePythonScript(code string) {
	fmt.Println(" => Ecriture du script : ", code)

	// Conversion en octets
	octets := []byte(code)

	// Enregistrement du code source dans un script
	err := ioutil.WriteFile("todo.py", octets, 0644)
	if err != nil {
		fmt.Println("ERROR with WriteFile() => ", err)
	}
}


// Fonction pour lancer un script Python dans un nouveau processus
func executePythonScript() string {
	fmt.Println(" => Lancement du script todo.py : ")

	// Exécution du script
	cmd := exec.Command("python3", "./todo.py")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		errStr := string(stderr.Bytes())
		fmt.Println(errStr)
		return errStr
	}
	outStr := string(stdout.Bytes())
	return outStr
}


// Fonction pour gérer les websockets
func WebSockethandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// Lecture et décodage des données JSON sur la websocket
		requete := Trame{}
		err := conn.ReadJSON(&requete)
		if err != nil {
			fmt.Println("ERROR with ReadJSON() => ", err)
			break
		}
		fmt.Println(" => WebSockethandler :")
		fmt.Println("  + Source  :", requete.Source)
		fmt.Println("  + Action  :", requete.Action)
		fmt.Println("  + Details :", requete.Details)
		fmt.Println("  + PID     :", requete.ExecPID)
		fmt.Println("  + UID     :", requete.ExecUID)
		fmt.Println("  + Clients :", requete.Clients)

		// Choix selon l'action à effectuer
		if (requete.Action == "lancer") {
			// Ecriture du code dans un script Python
			writePythonScript(requete.Details)

			// Exécution du script fraichement créé
			resultat := executePythonScript()

			// Réponse renvoyée sur la websocket
			var reponse Trame
			reponse.Source = "robot"
			reponse.Action = "informer"
			reponse.Details = resultat
			reponse.ExecPID = 0
			reponse.ExecUID = ""
			reponse.Clients = 0
			err = conn.WriteJSON(reponse)
			if err != nil {
				fmt.Println("ERROR with WriteJSON() => ", err)
				break
			}
		} else if (requete.Action == "arreter") {

		}
	}
}


/*
// Fonction pour servir la page gérant l'IDE
func IdeFileHandler(response http.ResponseWriter, request *http.Request) {
    fmt.Println(" => Requête pour : ", request.URL.Path[1:])

	tmpl := template.Must(template.ParseFiles("ide.html"))
    data :=
    tmpl.Execute(response, data)
}
*/


// Fonction pour servir les pages statiques (html, css et js)
func StaticFileHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(" => StaticFileHandler :", request.URL.Path[1:])
	http.ServeFile(response, request, request.URL.Path[1:])
}


// Fonction principale du programme
func main() {
	// Bannière d'accueil du serveur
	fmt.Printf("*********************************\n")
	fmt.Printf("*** LEGO Workshop server v0.2 ***\n")
	fmt.Printf("*********************************\n")
	fmt.Println(" => Lancement en cours...")

	// Enregistrement de la fonction gérant les websockets
	http.HandleFunc("/ws", WebSockethandler)

	// Enregistrement de la fonction gérant les pages statiques (IMG, CSS, JS...)
	http.HandleFunc("/", StaticFileHandler)

	// Lancement du serveur web
	http.ListenAndServe(":1337", nil)
}
