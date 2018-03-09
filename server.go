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


// Variables globales
var PID = 0
var UID = ""
var NbClients = 0
var Clients [3]*websocket.Conn
var PythonScript *exec.Cmd


// Fonction pour lancer un script Python dans un nouveau processus
func writePythonScript(code string) {
	fmt.Println(" => Ecriture du code source dans todo.py :\n", code)

	// Conversion en octets
	octets := []byte(code)

	// Enregistrement du code source dans un script
	err := ioutil.WriteFile("todo.py", octets, 0644)
	if err != nil {
		fmt.Println(" => ERROR with WriteFile() :", err)
	}
}


// Fonction pour lancer un script Python dans un nouveau processus
func executePythonScript(code string, conn *websocket.Conn) {
	// Ecriture du code dans un script Python
	writePythonScript(code)

	// Exécution du script
	fmt.Println(" => Lancement du script todo.py : ")
	cmd := exec.Command("python3", "./todo.py")
	PythonScript = cmd

	// Redirections des sorties stdout et stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Lancement du script Python
	err := cmd.Start()
	PID = cmd.Process.Pid
	fmt.Println("  + lancement avec PID : ", PID)
	sendToAllResponse("robot", "informer", "executing", PID, "", NbClients)

	// Attente de la fin du script
	err = cmd.Wait()
	fmt.Println("  + arrêt du processus : ", err)
	etat := 0

	resultat := ""
	if err != nil {
		if (err.Error() == "exit status 1") {
			etat += 100
		}
		if (err.Error() == "signal: killed") {
			etat += 566
		}
		errStr := string(stderr.Bytes())
		resultat = errStr
	} else {
		outStr := string(stdout.Bytes())
		resultat = outStr
	}
	// On retourne le résultat de l'exécution au client
	fmt.Println(etat)
	fmt.Println(resultat)
	sendToOneResponse(conn, "robot", "retourner", resultat, etat, "", NbClients)

	// On réinitialiser le robot
	resetRobot()

	// De nouveau libre pour une nouvelle exécution
	PID = 0
	sendToAllResponse("robot", "informer", "executing", 0, "", NbClients)
/*
	// Exécution du script
	cmd := exec.Command("python3", "./todo.py")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	fmt.Println("   + PID : ", cmd.Process.Pid)
	if err != nil {
		errStr := string(stderr.Bytes())
		fmt.Println(errStr)
		return errStr
	}
	outStr := string(stdout.Bytes())
	return outStr
*/
}


// Fonction pour ré-initialiser le robot (arrêt moteurs, beeper...)
func resetRobot() {
	// Message
	fmt.Println(" => resetRobot() :")

	// Lancement du script de réinitialisation du robot	
	cmd := exec.Command("bash", "./reset_robot.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(" => ERROR with Run() :", err)
	}
}



// Fonction pour envoyer une reponse sur la websocket
func sendToOneResponse(conn *websocket.Conn, s string, a string, d string, p int, u string, c int) bool {
	// Préparation de la trame à envoyer
	var reponse Trame
	reponse.Source = s
	reponse.Action = a
	reponse.Details = d
	reponse.ExecPID = p
	reponse.ExecUID = u
	reponse.Clients = c

	// Envoi de la trame préparée
	err := conn.WriteJSON(reponse)
	if err != nil {
		fmt.Println(" => ERROR with WriteJSON() :", err)
		return false
	}

	return true
}


// Fonction pour envoyer une reponse sur la websocket
func sendToAllResponse(s string, a string, d string, p int, u string, c int) {
	for _, conn := range Clients {
		if (conn != nil) {
        	sendToOneResponse(conn, s, a, d, p, u, c)
		}
    }
}


// Fonction pour ajouter un client à la liste
func addClient(conn *websocket.Conn) {
	// Ajout de la connexion au tableau de clients
	for i, _ := range Clients {
		if (Clients[i] == nil) {
			Clients[i] = conn
			break
		}
	}
	NbClients++
	fmt.Println(" => addClient :", Clients)

	// Envoi d'un message d'information à tout le monde
	sendToAllResponse("robot", "informer", "clients", 0, "", NbClients)
}


// Fonction pour supprimer un client de la liste
func removeClient(conn *websocket.Conn) {
	// Suppression de la connexion du tableau de clients
	for i, _ := range Clients {
		if (Clients[i] == conn) {
			Clients[i] = nil
		}
	}
	fmt.Println(" => removeClient :", Clients)

	// Fermeture de la connexion
	conn.Close()
	NbClients--

	// Envoi d'un message d'information à tout le monde
	sendToAllResponse("robot", "informer", "clients", 0, "", NbClients)
}


// Fonction pour gérer individuellement chaque connexion des clients
func ClientHandler(conn *websocket.Conn) {
	for {
		// Lecture et décodage des données JSON sur la websocket
		requete := Trame{}
		err := conn.ReadJSON(&requete)
		if err != nil {
			fmt.Println(" => ERROR with ReadJSON() :", err)
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
			// On teste si le robot est libre d'exécution
			if (PID == 0) {
				// On confirme le lancement de l'exécution
				//sendToOneResponse(conn, "robot", "informer", "writing", 0, "", NbClients)

				// Ecriture du code dans un script Python
				//writePythonScript(requete.Details)

				// On confirme le lancement de l'exécution
				//sendToOneResponse(conn, "robot", "informer", "executing", 0, "", NbClients)

				// Exécution du script fraichement créé
				go executePythonScript(requete.Details, conn)


			}
		} else if (requete.Action == "arreter") {
			fmt.Println("Demande d'arret en cours avec PID=", PID)
			if (PID != 0) {
				err := PythonScript.Process.Kill()
				if err != nil {
					fmt.Println(" => ERROR with Kill() :", err)
					break
				}
			}

			// Quoiqu'il arrive on réinitialiser le robot
			resetRobot()
		}
	}

	// Fermeture de la connexion
	//conn.Close()
	removeClient(conn)
}


// Fonction pour gérer les websockets (https://gist.github.com/tmichel/7390690)
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Attente de connexion d'un client
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer removeClient(conn)

	// Ajout de la connexion du client au tableau
	addClient(conn)

	// Lancement d'un thread pour gérer la connexion
	go ClientHandler(conn)
}


// Fonction pour servir les pages statiques (html, css et js)
func StaticFileHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(" => StaticFileHandler :", request.URL.Path[1:])
	http.ServeFile(response, request, request.URL.Path[1:])
}


// Fonction principale du programme
func main() {
	// Bannière d'accueil du serveur
	fmt.Printf("*********************************\n")
	fmt.Printf("*** LEGO Workshop server v0.3 ***\n")
	fmt.Printf("*********************************\n")
	fmt.Println(" => Lancement en cours...")

	// Enregistrement de la fonction gérant les websockets
	http.HandleFunc("/ws", WebsocketHandler)

	// Enregistrement de la fonction gérant les pages statiques (IMG, CSS, JS...)
	http.HandleFunc("/", StaticFileHandler)

	// Lancement du serveur web
	http.ListenAndServe(":1337", nil)
}
