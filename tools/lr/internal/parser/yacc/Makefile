MAKEFLAGS = -s

parser.go: grammar.y
	go install golang.org/x/tools/cmd/goyacc
	goyacc -o grammar.go -v grammar.output -p Lr grammar.y

clean:
	rm -f grammar.output grammar.go
