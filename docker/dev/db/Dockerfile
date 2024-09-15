FROM mysql:8.0.26

RUN apt-get update \
  ; apt-get upgrade -y \
  && apt-get install -y mycli \
  && apt-get -y clean \
  && rm -rf /var/lib/apt/lists/*
