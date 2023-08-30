FROM docker.io/golang:latest AS build-env
COPY . /build
WORKDIR /build
RUN go build ./cmd/runonany && ls -l

FROM ubuntu:latest
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=America/New_York
RUN apt-get update && apt-get install -y openssh-client
COPY --from=build-env /build/runonany /usr/bin
ADD testdata/runonany.yaml /etc
RUN groupadd -r user && useradd -r -g user user
RUN mkdir -p /home/user/.ssh
COPY testdata/keys/* /home/user/.ssh
RUN chown user:user /home/user/.ssh
COPY testdata/runonany-hostname-loop /
CMD /runonany-hostname-loop
