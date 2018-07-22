# Simulate the environment used by Travis CI so that we can run local tests to
# find and resolve issues that are consistent with the Travis CI environment.
# This is helpful because Travis CI often finds issues that our own local tests
# do not.

# go vet ./...
# golint -set_exit_status `go list ./... | grep -Ev "(stackint/asm|vendor)"`
# golint `go list ./... | grep -Ev "(stackint/asm|vendor)"`

go build `go list ./... | grep -Ev "(env|vendor)"`
go vet `go list ./... | grep -Ev "(env|vendor)"`
golint `go list ./... | grep -Ev "(env|vendor)"`

GOMAXPROCS=1 CI=true ginkgo -v --race --cover --coverprofile cover.out ./...
covermerge httpadapter/cover.out ingress/cover.out > cover.out

sed -i '/.pb.go/d' cover.out
sed -i '/bindings/d' cover.out
sed -i '/cmd/d' cover.out