# Instalação
Para instalar a linguagem Go, siga os seguintes passos:
- Faça o download dos arquivos: https://golang.org/dl/
- Descompacte o arquivo em /usr/local: `tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`
- Adicione o PATH da linguagem `export PATH=$PATH:/usr/local/go/bin`
- Clone o projeto no repositório desejado e faça as seguintes exportações:
   - export GOPATH=$HOME/Local de clonagem/winner-team-hack
   - export GOBIN=$HOME/Local de clonagem/winner-team-hack/bin

# Instalando as dependências
Execute:
- go get gopkg.in/zabawaba99/firego.v1
- go get github.com/gorilla/mux
