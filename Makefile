all: run

run:
	go run example/main.go

push:
	git push git@github.com:RB-PRO/KadTG.git

pull:
	git pull git@github.com:RB-PRO/KadTG.git