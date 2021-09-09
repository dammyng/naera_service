start_localy: 
	cd cmd && go build -o bin/naerarauth main.go && cd bin && ./naerarauth
build_linux_prod: 
	cd cmd && env GOOS=linux go build -o bin/naerarauth
push_linux_prod:
	cd cmd/bin && scp naerarauth kd@168.61.88.70:/var/www/naerarauth