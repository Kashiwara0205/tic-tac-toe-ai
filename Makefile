in:
	sudo docker exec -ti code /bin/bash

gorun:
	sudo docker exec -ti code go run main.go

gotest:
	sudo docker exec -ti code go test ./game -v
	sudo docker exec -ti code go test ./random_player -v
	sudo docker exec -ti code go test ./minmax_player -v
	sudo docker exec -ti code go test ./negmax_player -v
	sudo docker exec -ti code go test ./utils -v