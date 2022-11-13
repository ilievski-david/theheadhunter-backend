FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY ./main.go /app
COPY ./go.mod /app
COPY ./go.sum /app
#COPY ./.env /app
COPY ./models /app/models
COPY ./initializers /app/initializers

RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD [ "/app/main" ]