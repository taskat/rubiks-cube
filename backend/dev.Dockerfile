FROM taskat/devserver:latest AS devmode
WORKDIR /src/devmode

FROM golang:1.20 AS final
WORKDIR /app/dev
COPY --from=devmode /bin/dev_server /bin/dev_server
COPY ./scripts/run_server_background.sh /app/scripts/run_server_background.sh
COPY ./scripts/kill_server.sh /app/scripts/kill_server.sh

ENV WATCH_FOLDER="/app/dev"
ENV INCLUDE_FILES="^.*(\.go|go\.sum|go\.mod)$"
ENV START_SERVER_SCRIPT="/app/scripts/run_server_background.sh"
ENV KILL_SERVER_SCRIPT="/app/scripts/kill_server.sh"
ENTRYPOINT ["/bin/dev_server"]