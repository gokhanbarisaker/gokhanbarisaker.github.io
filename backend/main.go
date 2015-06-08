package main

import (
    "net/http"
    "resume"
)

func init() {
    http.HandleFunc("/resume", resume.Handler)
    http.HandleFunc("/dummy", resume.DummyInjector)
}
