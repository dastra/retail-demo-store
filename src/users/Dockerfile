FROM public.ecr.aws/s5z3t2n9/golang:1.15-alpine AS build
ARG GOPROXY_OVERRIDE=https://proxy.golang.org
WORKDIR /src/
COPY src/users-service/*.* /src/
COPY src/users-service/data/*.* /src/data/
RUN apk add --no-cache git

RUN echo "Setting GOPROXY to $GOPROXY_OVERRIDE"
RUN go env -w GOPROXY=$GOPROXY_OVERRIDE

# Download dependencies.  Running as a separate step from build to make use of docker layer caching
RUN go mod download -x
# Build
RUN CGO_ENABLED=0 go build -o /bin/users-service

RUN apk add ca-certificates

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/users-service /bin/users-service
COPY --from=build /src/data/*.* /bin/data/
EXPOSE 80
ENTRYPOINT ["/bin/users-service"]