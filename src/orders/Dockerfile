FROM public.ecr.aws/s5z3t2n9/golang:1.15-alpine AS build
ARG GOPROXY_OVERRIDE=https://proxy.golang.org
WORKDIR /src/
COPY src/orders-service/*.* /src/
RUN apk add --no-cache git bash

RUN echo "Setting GOPROXY to $GOPROXY_OVERRIDE"
RUN go env -w GOPROXY=$GOPROXY_OVERRIDE

# Download dependencies.  Running as a separate step from build to make use of docker layer caching
RUN go mod download -x
# Build
RUN CGO_ENABLED=0 go build -o /bin/orders-service

FROM scratch
COPY --from=build /bin/orders-service /bin/orders-service
EXPOSE 80
ENTRYPOINT ["/bin/orders-service"]