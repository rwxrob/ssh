FROM ubuntu:latest
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=America/New_York
RUN apt-get update && apt-get install -y openssh-server
RUN service ssh start
RUN addgroup user && adduser user --ingroup user --disabled-password
RUN mkdir -p /home/user/.ssh
COPY testdata/keys/* /home/user/.ssh
RUN chown user:user /home/user/.ssh
COPY testdata/entrypoint /
CMD /entrypoint