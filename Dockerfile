FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/ducky-api

FROM alpine:latest
RUN apk update && apk upgrade
RUN apk add ffmpeg python3 py3-pip
COPY --from=build /app/bin/ducky-api /usr/local/bin
COPY --from=build /app/cleanup.sh /usr/local/bin
RUN chmod +x /usr/local/bin/cleanup.sh
RUN echo "*/15 * * * * /usr/local/bin/cleanup.sh" >> /etc/crontabs/root
RUN wget -P /usr/local/bin https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp
RUN chmod +x /usr/local/bin/yt-dlp
EXPOSE 10000
CMD /bin/sh -c crond && ducky-api
