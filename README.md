# GRPC
Pierwszy kod + wskazówki

<h1> Kod przerżnięty z kursu https://tutorialedge.net/golang/go-grpc-beginners-tutorial/ </h1> </br>
Coś ogarnąłem, ale chyba nie wystarczająco.  </br> </br>

Teraz wiem jak: kompilować za pomocą protoc, stworzyć prosty serwer, stworzyć prostego klienta.  </br> </br>

Dalej nie wiem: czego kontekst bierze "context" / "context.Background()"; </br>
Za co odpowiadają pliki: chat.pb.go i chat_grpc.pb.go.
  
  <h2> Wskazówki </h2>
  Jak kompilowac proto: protoc -I $GOPATH/src/ --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative /Users/nintyswinty/go/src/Zabawa_z_gRPC/pb/chat.proto </br> </br>
  Nie uruchamiać Go z terminala tylko z GoLand'a. ( Przeczytałęm że go run żadko wychodzi - lepiej używać go build + w GoLand wykrywa mi GoPath ).  </br> </br>
  Robić folder pb.  </br> </br>
  Dodawać "option go_package = "Nazwa_Projektu;Plik";" do .proto </br> </br>
  Struktura serwera zawsze musi zawierać funkcję: "UnimplementedChatServiceServer"
  
  
