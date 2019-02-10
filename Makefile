psi_collector.out: ./main.go ./collector.go
	go build -o $@ -v -i .
