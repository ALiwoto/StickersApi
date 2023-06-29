FROM golang:1.20-alpine3.18

WORKDIR /app

COPY . .

RUN apk add --no-cache ca-certificates \
    bash \
    curl \
    ffmpeg \
    gcc \
    musl-dev && \
    go mod tidy && \
    go build -ldflags "-s -w" -o StickersAPI .

CMD [ "./StickersAPI" ]
