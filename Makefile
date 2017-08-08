.PHONY: test build dev start

test:
	go test -cover -v `glide novendor`
