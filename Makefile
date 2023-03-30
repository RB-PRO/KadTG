all: run

run:
	go run example/main.go

push:
	git push git@github.com:RB-PRO/KadArbitr.git

pull:
	git pull git@github.com:RB-PRO/KadArbitr.git