FROM golang:1.23.4-bookworm
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o stage_one
EXPOSE 8080
CMD [ "./stage_one"]