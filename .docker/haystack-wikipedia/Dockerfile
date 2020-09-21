FROM python:3-slim
MAINTAINER paper2code <contact@paper2code.com>

RUN apt-get update && \
    apt-get install -y nano bash jq git ca-certificates gcc

WORKDIR /opt/service

COPY requirements.txt .

RUN python3 -m pip install -U pip && \
    pip3 install -r requirements.txt && \
    pip3 install git+https://github.com/deepset-ai/haystack.git@bde33ddaaac48d3f302e536d3155aed3d4cbd285

COPY . .

VOLUME ["/opt/data"]

CMD ["/bin/bash"]
