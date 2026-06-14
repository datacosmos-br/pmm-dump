# syntax=docker/dockerfile:1.7
# Copyright 2023 Percona LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# ---- build stage ----------------------------------------------------------
FROM golang:1.26-alpine AS build

WORKDIR /src
ENV CGO_ENABLED=0 GOFLAGS=-mod=readonly

# Cache modules first for faster incremental builds.
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=unknown
ARG BRANCH=unknown
ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath \
      -ldflags "-s -w -X 'main.GitBranch=${BRANCH}' -X 'main.GitCommit=${COMMIT}' -X 'main.GitVersion=${VERSION}'" \
      -o /out/pmm-dump pmm-dump/cmd/pmm-dump

# ---- runtime stage --------------------------------------------------------
# Distroless static + nonroot: no shell, no package manager, minimal CVE surface.
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/pmm-dump /usr/local/bin/pmm-dump

USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/pmm-dump"]

ARG VERSION=dev
ARG COMMIT=unknown
LABEL org.opencontainers.image.title="pmm-dump" \
      org.opencontainers.image.description="Percona Monitoring and Management (PMM) data export/import tool" \
      org.opencontainers.image.source="https://github.com/datacosmos-br/pmm-dump" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.revision="${COMMIT}"
