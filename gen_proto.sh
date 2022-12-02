export PATH="`pwd`/:$PATH"

rm -rf ./golang/*

ls ./pb/*.proto | xargs protoc -I=./pb/ --go_out=./golang --go_opt=paths=source_relative

ls ./golang/*.pb.go | xargs -n 1 -I {} protoc-go-inject-tag -input={} #-verbose

#sed -i "" -e "s/,omitempty//g" ./golang/*.go