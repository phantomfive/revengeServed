revengeServed: main.go
	go build

run: revengeServed
	./revengeServed 8080 &
	sudo ./revengeServed 80 &
	sudo ./revengeServed 23 &
	sudo ./revengeServed 22 &


