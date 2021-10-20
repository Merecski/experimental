if [[ -f "client/client" ]]; then
    cd client
    go build
fi
./client 127.0.0.1:1234
cd ../