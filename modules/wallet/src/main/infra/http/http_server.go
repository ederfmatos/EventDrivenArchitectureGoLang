package http

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type ServerHttpHandler interface {
	Handle(body io.ReadCloser, server Server, request *http.Request) (any, error)
}

type Server interface {
	Post(path string, handler ServerHttpHandler)
	ParseBody(body io.ReadCloser, data any) error
	Listen(port string)
}

type MuxHttpServer struct {
	server *http.ServeMux
}

func NewMuxHttpServer() *MuxHttpServer {
	return &MuxHttpServer{
		server: http.NewServeMux(),
	}
}

func (httpServer MuxHttpServer) ParseBody(body io.ReadCloser, data any) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(&data)
}

func (httpServer MuxHttpServer) Post(path string, handler ServerHttpHandler) {
	httpServer.server.HandleFunc("POST "+path, func(writer http.ResponseWriter, request *http.Request) {
		response, err := handler.Handle(request.Body, httpServer, request)
		writer.Header().Set("Content-Type", "application/json")
		var output interface{}
		if err != nil {
			writer.WriteHeader(422)
			output = map[string]interface{}{"error": err.Error()}
		} else {
			writer.WriteHeader(200)
			output = response
		}
		responseJson, _ := json.Marshal(output)
		_, _ = writer.Write(responseJson)
	})
}

func (httpServer MuxHttpServer) Listen(port string) {
	log.Info().Msg("Server running")
	err := http.ListenAndServe(":"+port, httpServer.server)
	if err != nil {
		panic(err)
	}
}
