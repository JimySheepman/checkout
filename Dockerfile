FROM golang:1.19.6-alpine AS BUILDER

# Set working directory
WORKDIR /app

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY go.mod ./
COPY go.sum ./
# Download dependencies
RUN go mod download
# Move source files
COPY . ./
# Set revision information
ARG revisionID
ARG commitID
ENV REVISION_ID ${revisionID}
ENV COMMIT_ID ${commitID}

# Build
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.Revision=`date -u +%Y.%-m.%-d`.${REVISION_ID} -X main.Commit=${COMMIT_ID}" -o app

FROM alpine

COPY --from=BUILDER /app/app /app

ENTRYPOINT ["/app"]
