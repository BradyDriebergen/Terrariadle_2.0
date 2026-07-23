## Building Terrariadle

This is the current build command contained within the MakeFile:

```makefile
build:
	cd frontend && npm run build
	rm -rf internal/web/build
	cp -r frontend/build internal/web/build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-trimpath \
		-ldflags "-s -w -X main.version=$$(git describe --tags --always --dirty)" \
		-o bin/terrariadle .
```

- **CGO_ENABLED=0**: produces a fully static binary with no libc dependency. Means I don't have to worry about glibc version mismatches between my build machine and whatever's on the Oracle image, and it plays nicely with a minimal systemd unit (no dynamic linker surprises).
- **GOOS=linux GOARCH=amd64**: explicit cross-compile target for the E2.1.Micro shape.
- **trimpath**: strips local filesystem paths (like /Users/me/...) from the compiled binary. Minor security/hygiene win with no runtime cost.
- **ldflags "-s -w"**: strips the symbol table and DWARF debug info. Cuts binary size by roughly a third, which matters more for the instance's limited boot volume than for runtime RAM (Go binaries don't page in symbol tables at runtime), but smaller is still better for scp transfer time and disk headroom.
