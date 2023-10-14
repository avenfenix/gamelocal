# Requisitos
sudo apt update
sudo apt install -y golang

# Compilar
go build -o gamelocal-client client.go
go build -o gamelocal-server server.go

# Desintalar version anterior
sudo rm /usr/local/bin/gamelocal-client
sudo rm /usr/local/bin/gamelocal-server

# Instalar
mkdir ~/.gamelocal/
sudo mv gamelocal-client /usr/local/bin/
sudo mv gamelocal-server /usr/local/bin/

