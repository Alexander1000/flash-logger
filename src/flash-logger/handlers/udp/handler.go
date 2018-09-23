package udp

import (
	"log"

	"flash-logger/storage"
	"flash-logger/config"
	"flash-logger/udp"
)

type Handler struct {
	storage storage.Repository
	projects []config.Project
	relpParser *udp.RelpParser
}

func New(storage storage.Repository, projects []config.Project) *Handler {
	return &Handler{
		storage: storage,
		projects: projects,
		relpParser: udp.NewRelpParser(),
	}
}

// @example message
// https://tools.ietf.org/html/rfc3164
// <142>Sep 22 11:28:28 cs330699 nginx: 31.130.148.145 - - [22/Sep/2018:11:28:28 +0300] "GET / HTTP/1.1" 200 444 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.34
//  97.100 Safari/537.36"
func (h *Handler) Handle(buffer []byte) {
	relp, err := h.relpParser.Scan(buffer)
	if err != nil {
		log.Printf("error in parse udp: %v", err)
		return
	}

	projectID := 0
	prjByName := 0
	for _, prj := range h.projects {
		if prj.Token == relp.Process {
			projectID = prj.ID
			break
		} else if prj.Name == relp.Process {
			prjByName = prj.ID
		}
	}

	if projectID == 0 {
		if prjByName == 0 {
			// not found project
			log.Printf("not found project for process: %s", relp.Process)
			return
		}
		projectID = prjByName
	}

	if err := h.storage.SaveMessage(projectID, relp.GetSeverity(), string(relp.Message), nil, nil); err != nil {
		return
	}

	log.Println(string(relp.Message))
}
