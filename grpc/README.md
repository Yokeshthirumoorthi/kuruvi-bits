Follow instruction from here - https://github.com/grpc/grpc-go/tree/master/examples


To generate proto
```
chmod +x ./protogen.sh
./protogen.sh
```

#### Installation Help

Use this to install protoc

```
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.10.0/protoc-3.10.0-linux-x86_64.zip

# Unzip
unzip protoc-3.10.0-linux-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

# Optional: change owner
sudo chwon [user] /usr/local/bin/protoc
sudo chwon -R [user] /usr/local/include/google
```