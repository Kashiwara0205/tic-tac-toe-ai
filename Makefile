in:
	sudo docker exec -ti code /bin/bash

gorun:
	sudo docker exec -ti code go run main.go

gotest:
	sudo docker exec -ti code go test ./game -v
	sudo docker exec -ti code go test ./random_player -v