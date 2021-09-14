FROM golang:1.14.3

COPY . /app
WORKDIR /app

ENV DB_HOST=sahabat-utd.id
ENV DB_NAME=sahabat_utd
ENV DB_USER=kic-utd
ENV DB_PASS=P@ssw0rdKIC!!
ENV DB_PORT=3306
ENV PORT=6008

RUN go mod vendor
RUN go build
ENTRYPOINT [ "./ms-simaba-riwayat-donor" ]