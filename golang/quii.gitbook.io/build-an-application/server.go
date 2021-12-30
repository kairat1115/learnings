package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"
)

const (
	jsonContentType = "application/json"
	gameHTMLFile    = "game.html"
	markupsPath     = "./markups"
)

var (
	wsUpgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store    PlayerStore
	game     Game
	template *template.Template
	http.Handler
}

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	tmpl, err := template.ParseFiles(getMarkupFilePath(gameHTMLFile))
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", gameHTMLFile, err)
	}

	p := &PlayerServer{
		store:    store,
		game:     game,
		template: tmpl,
	}

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))

	p.Handler = router
	return p, nil
}

func (p *PlayerServer) leagueHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", jsonContentType)
	json.NewEncoder(rw).Encode(p.store.GetLeague())
	rw.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(rw http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(rw, player)
	case http.MethodGet:
		p.showScore(rw, player)
	}
}

func (p *PlayerServer) gameHandler(rw http.ResponseWriter, r *http.Request) {
	p.template.Execute(rw, nil)
}

func (p *PlayerServer) webSocket(rw http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(rw, r)
	// defer ws.Close()

	numberOfPlayersMsg := ws.waitForMsg()
	numberOfPlayers, err := strconv.Atoi(string(numberOfPlayersMsg))

	if err != nil || numberOfPlayers < 1 {
		fmt.Fprint(ws, BadPlayerInputErrMsg)
		return
	}

	p.game.Start(numberOfPlayers, ws)

	winner := ws.waitForMsg()
	p.game.Finish(winner)
}

func (p *PlayerServer) processWin(rw http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	rw.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(rw http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		rw.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(rw, score)
}

type playerServerWS struct {
	*websocket.Conn
}

func newPlayerServerWS(rw http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Printf("problem upgrading connection to WebSockets %v\n", err)
	}
	return &playerServerWS{conn}
}

func (ws *playerServerWS) waitForMsg() string {
	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}
	return string(msg)
}

func (ws *playerServerWS) Write(p []byte) (n int, err error) {
	err = ws.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, nil
	}
	return len(p), nil
}

func getMarkupFilePath(file string) string {
	return fmt.Sprintf("%s/%s", markupsPath, file)
}
