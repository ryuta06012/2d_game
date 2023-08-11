FROM golang:1.20 as local

WORKDIR /app

RUN apt-get update \
    && apt-get install -y pkg-config gcc libgl1-mesa-dev xorg-dev

# install air (hot reload tool)
RUN go install github.com/cosmtrek/air@v1.40.2

COPY . .

CMD [ "air" ]