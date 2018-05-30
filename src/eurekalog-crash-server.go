package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(0)
		index := 0
		for {
			name := fmt.Sprintf("el_upload_file_%d", index)
			file, handler, err := r.FormFile(name)
			if err != nil {
				//log.Printf("name = %s. Error:%s", name, err)
				break
			}
			file.Close()
			f, err := os.OpenFile("./report/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Printf("Error open file %s Error:%s", handler.Filename, err)
				break
			}
			defer f.Close()
			io.Copy(f, file)

			log.Printf("Store File:%s", handler.Filename)
			index++
		}
	} else {
		log.Printf("Error Metod:%s", r.Method)
	}
}

func main() {
	os.Mkdir("./log", 0777)
	os.Mkdir("./report", 0777)
	logFile, err := os.OpenFile("./log/eurekalog-server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":5050", nil)
}

// TODO - linux
//	logwriter, err := syslog.New(syslog.LOG_ERR, "GoExample")
//	logwriter, e := syslog.New(syslog.LOG_NOTICE, "EurekaLog-http-server")
//	if err == nil {
//        log.SetOutput(logwriter)
//	}
