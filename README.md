# LocalControlleDevToolsCollection

[Link_Documentação_gin](https://gin-gonic.com/en/docs/)

[Local_Repositorio_Projeto](https://github.com/Furaxis/LocalControlleDevToolsCollection)


### Comandos inicialização do projeto
* go mod init github.com/Furaxis/LocalControlleDevToolsCollection
* cat go.mod
* go build -o exec .
* go mod tigy


### Comandos de Instalação de dependencias, Projeto ( seco, seco )
_________________________________________________________________________________
```
go get -u golang.org/x/tools/gopls
go get -u github.com/go-delve/delve/cmd/dlv
go get -u golang.org/x/tools/cmd/gofmt
go get -u golang.org/x/tools/cmd/goimports
go get -u golang.org/x/lint/golint
go get -u github.com/mdempsky/gocode
go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs
go get -u github.com/fatih/gomodifytags
go get -u github.com/josharian/impl
go get -u golang.org/x/tools/cmd/guru
go get -u golang.org/x/tools/cmd/gorename
go install golang.org/x/tools/gopls@latest
go install github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
go install github.com/fatih/gomodifytags@latest
go install github.com/josharian/impl@latest
```
_________________________________________________________________________________

Infelismente quando o W11 começa a atrapalhar tive que começar a usar a WSL, 

caso alguem queira terar que instalar os pacotes necessarios 
```
 sudo snap install go --classic
```
_________________________________________________________________________________

```
wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.25.1.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```
_________________________________________________________________________________
por fim para baixar todas as dependencias do GO

```
sudo apt install golang-go
```

estou pensando em começar a usar p https://alacritty.org/, sei la
[proxies](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies)