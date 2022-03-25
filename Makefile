FUNC := Test
KEYLOG = cmd/keylog/main
DLLMAIN = cmd/dllmain/main
OUT := ${DLLMAIN}SRC

copy:
	go run addons/addons.go

keylog:
	env GOOS=windows go build -o ${KEYLOG} -ldflags="-s -w" -trimpath ${KEYLOG}.go

dllrun:
	rundll32.exe ${OUT}.dll ${FUNC}

dllbuild:
	set CC=i686-w64-mingw32-gcc
	set CXX=i686-w64-mingw32-g++
	set GOARCH=386
	set CGO_ENABLED=1
	go build -i -v -o ${OUT}.dll -buildmode=c-shared ${DLLMAIN}.go

dllexe:
	set CC=i686-w64-mingw32-gcc
	set CXX=i686-w64-mingw32-g++
	set GOARCH=386
	set CGO_ENABLED=1
	env GOOS=windows go build -o ${DLLMAIN} -ldflags="-s -w" -trimpath ${DLLMAIN}.go

bootstrapper:
	go run cmd/bootstrapper/main.go