package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"server/service"
	"strconv"
)

type Text struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

type TopWordController struct {
	StopCh  chan os.Signal
	Service service.Service
}

func (t *TopWordController) Stat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	count, err := strconv.Atoi(vars["count"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Cant parse int", 400)
	}
	_, err = w.Write([]byte(t.Service.GetTop(count)))
	if err != nil {
		log.Println(err)
		http.Error(w, "Write error", 500)
	}
}

func (t *TopWordController) Text(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var text Text
	err := json.Unmarshal(reqBody, &text)
	if err != nil {
		log.Println(err)
		http.Error(w, "Cant parse json", 400)
	}

	t.Service.ProcessText(text.Text)
}

func (t *TopWordController) Stop(w http.ResponseWriter, r *http.Request) {
	t.StopCh <- os.Interrupt
}
