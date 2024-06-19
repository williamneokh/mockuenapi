FROM golang:1.22
RUN apt-get update -y
RUN apt-get install -y tzdata

ENV TZ=Asia/Singapore

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 3300

CMD [ "/app/mockuenapi/main"]