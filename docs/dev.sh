DIR=$(PWD)
cd $DIR/../../ && \
snippetgo -pkg=examples > ./x/docs/examples/examples-generated.go && \
cd $DIR && go run ./docsmain/

