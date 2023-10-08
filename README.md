# WatchDogs
## Abstract
Verifies if the messages have been transmitted through gRPC in Go 

## Notes
### Generating Go and gRPC code from greet.proto
1. Install protobuf compilter: `sudo apt install protobuf-compiler`
1. Install the protobuf plugins: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` and \
    `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`. Fyi, do not install the protobuf plugins from github because the link is depracated and will not work. Anyways, make sure they are installed!  Like seriously go to your ~/.go to see if there is a /bin directory that contains the plugins. If not, run the commands from step one again in the ~/.go/bin
3. Ensure your $GOPATH, $GOBIN, and $PATH environment variables were correctly set in your shell profile file (e.g., ~/.bashrc or ~/.zshrc). You should see \
`export GOPATH=$HOME/go`\
`export GOBIN=$GOPATH/bin`\
`export PATH=$PATH:$GOBIN`\
If not, then add them.
4. Run `protoc --go_out=. --go-grpc_out=. --plugin=protoc-gen-go=$GOBIN/protoc-gen-go --plugin=protoc-gen-go-grpc=$GOBIN/protoc-gen-go-grpc {pwd}`. For example. you can replace pwd with "/home/computer-name/projects-directory/project-name/proto-folder/proto-file"
