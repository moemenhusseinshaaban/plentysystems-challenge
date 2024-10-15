FROM golang:1.23

WORKDIR /app

COPY ./command ./command
COPY ./telemetry ./telemetry

CMD ["tail", "-f", "/dev/null"]
