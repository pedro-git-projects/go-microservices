OBJ_NAME = 
LDFLAGS = 
front:
	$(eval OBJ_NAME += front)
	$(eval LDFLAGS += "-w -s")
	cd ./cmd/web/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin 
run:
	$(eval OBJ_NAME += front)
	./bin/$(OBJ_NAME)
