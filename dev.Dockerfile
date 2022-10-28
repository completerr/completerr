FROM cosmtrek/air

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /server
COPY server .


EXPOSE 2345
