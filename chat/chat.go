package chat

import (
	chat "Zabawa_z_gRPC"
	"log"

	"golang.org/x/net/context"
)

type Server struct { // Pusta struktura serwera
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) { // Funckja która przyjmuję kontekst i wiadomość i zwraca wiadomość i error
	log.Printf("Messages: $V", message.Body) // Wypisywanie wiadomości
	return &Message{Body: "Hello from the server!"}, nil // Zwracanie wiadomości i błędu ( a raczej jego brak )
}
