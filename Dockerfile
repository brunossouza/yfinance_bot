FROM golang:1.14 as build
# All these steps will be cached
RUN mkdir -p /app/src
WORKDIR /app/src
COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/yfinance


FROM scratch
COPY --from=build /app/yfinance /go/bin/yfinance
ENTRYPOINT ["/go/bin/yfinance"]