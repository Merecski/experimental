if [[ -f "server/server" ]]; then
    cd server
    go build
fi
./server 1234
cd ../