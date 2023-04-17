all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/KadTG.git

pull:
	git pull git@github.com:RB-PRO/KadTG.git

push-car:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o main ./cmd/main/main.go
	scp main token root@194.87.107.129:go/KadTG/