FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

ENV PORT 8605
EXPOSE 8605

CMD [ "./main" ]