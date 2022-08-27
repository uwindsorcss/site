build:
      FROM golang:1.19 # start with go
      WORKDIR /app # change workdir

      COPY main.go wire.go go.mod go.sum . # copy files
      COPY --dir pkg/ vendor/ ./ # copy dirs

      RUN go install github.com/google/wire/cmd/wire@latest; \
          wire ./wire.go;

      COPY default.hcl . # copy config
      RUN go generate ./...;\
          rm wire.go # gets in the way of building

      RUN go build -mod=vendor -o site . # build  with vendoring
      SAVE ARTIFACT site AS LOCAL site # save file
