fuzz:
	go install github.com/dvyukov/go-fuzz/go-fuzz-build
	go-fuzz-build github.com/gogo/fuzztests
	go install github.com/dvyukov/go-fuzz/go-fuzz
	go-fuzz -bin=./fuzztests-fuzz.zip -workdir=.

regenerate:
	go install ./protoc-gen-gogomsglist
	protoc --go_out=./golang/ *.proto
	protoc --gogomsglist_out=./golang/ *proto
	protoc --gofast_out=./gofast/ *.proto
	protoc --gogomsglist_out=./gofast/ *proto
	protoc --gogo_out=./gogo/ *.proto
	protoc --gogomsglist_out=./gogo/ *proto
	protoc --gogofast_out=./gogofast/ *.proto
	protoc --gogomsglist_out=./gogofast/ *proto
	go install ./protoc-gen-gogopop
	protoc --gogopop_out=./gengogofuzztests/gogopop/ *.proto
	protoc --gogomsglist_out=./gengogofuzztests/gogopop/ *.proto
	rm -rf corpus || true
	rm -rf crashers || true
	rm -rf suppressions || true
	mkdir corpus
	go install ./gengogofuzztests
	gengogofuzztests ./corpus
	make gofmt

clean:
	go clean ./...

nuke:
	go clean -i ./...

gofmt:
	gofmt -l -s -w .