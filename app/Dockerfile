FROM golang:1.19.2-alpine3.15 as builder

WORKDIR /build
COPY go.mod go.sum ./
RUN ls -l
RUN go mod download && go mod tidy
COPY . .
RUN go build -o notes-api .


# Input parameters for the Dockerfile expected in os.Getenv

ARG dbport="5432"
ARG dbname="notes"
ENV DATABASE_PORT=$dbport
ENV DB_NAME=$dbname


FROM golang:1.19.2-alpine3.15
# Add maintainer info
LABEL maintainer="Danail Surudzhiyski"
RUN addgroup -S pipeline && adduser -S k8s-pipeline -G pipeline

WORKDIR /home/k8s-pipeline
COPY --from=builder /build/notes-api .
COPY --from=builder /build/templates/* templates/
COPY --from=builder /build/static/* static/
RUN ls -l

USER k8s-pipeline
EXPOSE 8001


CMD ["./notes-api"]