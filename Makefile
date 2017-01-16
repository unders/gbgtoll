VERSION=v0.0.2

.PHONY: help
help:
	cat Makefile

.PHONY: install
install:
	go install github.com/unders/gbgtoll

.PHONY: spec
spec:
	go test github.com/unders/gbgtoll/prog -v

.PHONY: test
test:
	go test github.com/unders/gbgtoll/...

.PHONY: push
push: test
	git push

.PHONY: release
release: test
	@mkdir -p release
	@./bin/release $(VERSION)

.PHONY: clean
clean:
	@rm -rf release

.PHONY: run
run: install
	gbgtoll car -d=2017-01-01 -t=05:56,06:10
	@echo ""
	gbgtoll emergency -d=2017-01-01 -t=05:56,06:10
	@echo ""
	gbgtoll bus -d=2017-01-01 -t=05:56,06:10
	@echo ""
	gbgtoll diplomat -d=2017-01-01 -t=05:56,06:10
	@echo ""
	gbgtoll car -d=2017-01-02 -t=05:56,06:10
	@echo ""
	gbgtoll pickup -d=2017-01-02 -t=05:56,06:10,17:50
	@echo ""
	gbgtoll pickup -d=2016-01-02 -t=05:56,06:10,17:50
	@echo ""
	gbgtoll pickup -d=2017-01-02 -t=07:00,08:29,15:30,18:00
	@echo ""
