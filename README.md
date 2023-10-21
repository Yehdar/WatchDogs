# WatchDogs
## Abstract
Verifies if the messages have been transmitted through gRPC in Go 

<p align="center"><img src="https://github.com/Yehdar/watchdogs/blob/main/demo/demo.png" width="80%"></p>

## Notes
### Understanding gRPC
#### Unary RPC
Think of this as a simple request-response where the client sends a single request to the server and recieves a single response. As an analogy, think of requesting a food order at a restaurant and receiving that order as a response.
#### Server Streaming RPC
This allows the client to send a single request to the server and recieve a stream of responses. As an analogy, think of a music streaming service. You request a playlist and the server keeps providing songs from the playlist to your device without any extra input.
#### Client Streaming RPC
Lets the client send a stream of requests to the server and receive a single response. As an analogy, think of breaking down a really long message into short burts of text. Once your friend recieves all your individual texts that come together to explain your overall message, then they'll respond. 
#### Bidirectional Streaming RPC
Both the client and server send a stream of messages independently. As an analogy, think of multiple conversations at a social gathering where multiple people are talking and responding on top of each other.

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

### Missing packages
1. Run `go mod tidy` in the terminal to update go.mod with the missing packages


## Configuration
1. Git clone the repository and cd into the project directory
2. Set your network constants in the config file and then cd into the server directory to boot up the server-side by running `go run *`
3. Head to the client directory to turn on the client-side by running `go run *`
