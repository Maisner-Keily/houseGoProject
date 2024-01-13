FROM golang:1.21.1

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o house_go_project ./main.go
#RUN chmod +x ./

CMD ["./house_go_project"]