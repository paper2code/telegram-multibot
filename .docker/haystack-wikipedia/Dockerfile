FROM python:3-slim
MAINTAINER paper2code <contact@paper2code.com>

RUN apt-get update && \
    apt-get install -y nano bash jq

WORKDIR /opt/service

COPY requirements.txt .
RUN pip3 install -r requirements.txt

COPY . .

VOLUME ["/opt/data"]

CMD ["/bin/bash"]
