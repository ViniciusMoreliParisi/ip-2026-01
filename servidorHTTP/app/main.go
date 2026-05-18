package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"servidorHTTP/app/handlers"
	"servidorHTTP/app/utils"
)

func main() {
	utils.ConnectToDB()

	// Rotas da aplicação (registradas antes do file server)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/updateAccount", handlers.UpdateAccountHandler)
	http.HandleFunc("/deleteAccount", handlers.DeleteAccountHandler)

	// Health Management
	http.HandleFunc("/forms/patient.html", handlers.PatientFormHandler)
	http.HandleFunc("/patients/new", handlers.PatientFormHandler)
	http.HandleFunc("/patients", handlers.PatientHandler)
	http.HandleFunc("/appointments/new", handlers.AppointmentFormHandler)
	http.HandleFunc("/appointments", handlers.AppointmentHandler)
	http.HandleFunc("/schedule", handlers.ScheduleHandler)

	// Arquivos estáticos por último
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	var localIP string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			localIP = ipNet.IP.String()
			break
		}
	}

	port := "3000"
	if localIP == "" {
		localIP = "127.0.0.1"
	}

	fmt.Printf("Servidor rodando em: http://%s:%s/\n", localIP, port)

	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatal(err)
	}
}
