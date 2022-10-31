FROM golang:bullseye

RUN apt update
RUN apt install ffmpeg -y

WORKDIR /app

COPY . /app

RUN go build

ENTRYPOINT [ "./youtrim" ]
