FROM golang:1.21

WORKDIR /offline_project/backend/service_gateway

COPY . .

EXPOSE 8080

# This is not works if vscode stats container. But may work on production server. Untested yet
ENTRYPOINT [ "/scripts/prepare_for_work.sh" ] 
