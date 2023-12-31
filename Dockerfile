FROM golang:1.21-bullseye AS build

WORKDIR /app 

# We optimize our path to discovery, selecting only the files required to install dependencies. 🧭
# With this choice, we unlock the potential of better layer caching, improving our image's efficiency.
COPY go.mod go.sum ./

# Cache mounts speed up the installation of existing dependencies,
# empowering our image to sail smoothly through vast dependency seas.
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download


FROM build AS build-production

COPY . .

# During this stage, we compile our application ahead of time, avoiding any runtime surprises.
# The resulting binary, web-app-golang, will be our steadfast companion in the final leg of our journey.
# We strategically add flags to statically link our binary.
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
   go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o app

# The scratch base image welcomes us as a blank canvas for our prod stage.
FROM scratch

WORKDIR /app

# We transport the binary to our deployable image
COPY --from=build-production /app/ /app/

ENV CONTAINER="true"
# By exposing port 3000, we signal to the Docker environment the intended entry point for our application.
EXPOSE 8000

CMD ["/app/app"]
