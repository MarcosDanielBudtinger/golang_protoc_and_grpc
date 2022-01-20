# Golang, GRPC e Protocol Buffers
Estudos com Golang, GRPC e Protocol Buffers

Projeto feito para fins de estudos.

Para rodar basta seguir os passos abaixo:

Faça clone do projeto.

Execute o server com o comando: go run cmd/server/server.go  

Execute o client com o comando: go run cmd/client/client.go

Na func main do arquivo client.go temos 4 métodos, cada uma faz 
as requests de forma diferente.

Para testar os outros métodos basta descomentar o qual deseja e deixar os outros comentados.

Em alguns métodos trabalhamos com stream apenas do lado do client, outro apenas do lado do server e no 
último com stream bi-direcional.
