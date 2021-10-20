# Instalação das ferramentas gRPC

Para usar as ferramentas que trabalharemos na aula seguinte para compilar os protofiles será necessário instalar alguns pacotes.

Execute estes comandos no seu terminal Linux/MAC:

```bash
sudo apt install protobuf-compiler 
brew install protobuf #Mac, requer Homebrew instalado
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
Para finalizar, temos que adicionar a pasta "/go/bin" no PATH do Linux para que tudo que seja instalado nesta pasta esteja disponível como comandos no terminal. Adicione no final do seu ~/.bash

PATH="/go/bin:$PATH"
Execute o comando abaixo para atualizar seu terminal:

source ~/.bashrc
Pronto, todos os executáveis usados na aula anterior já estão disponíveis no seu terminal.
 ```

```bash
 iniciando modulo go
$ go mod init github.com/wallacemachado/estudos-grpc

instalando go protocol buffers plugin

https://developers.google.com/protocol-buffers/docs/gotutorial
$ go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
 ```