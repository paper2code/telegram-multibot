FROM golang:1-alpine AS build
MAINTAINER paper2code <contact@paper2code.com>

ARG VERSION
ARG GIT_COMMIT
ARG BUILD_DATE

ARG CGO=1
ENV CGO_ENABLED=${CGO}
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR /go/src/github.com/paper2code/telegram-multibot

COPY . /go/src/github.com/paper2code/telegram-multibot

# gcc/g++ are required to build SASS libraries for extended version
RUN apk update && \
    apk add --no-cache gcc g++ musl-dev pkgconfig cmake git ca-certificates make boost-dev ragel pcre2-dev pcre-dev libpcap-dev

# install hyperscan
RUN git clone --depth=1 https://github.com/intel/hyperscan /tmp/hyperscan && \
    cd /tmp/hyperscan && \
    mkdir -p /tmp/hyperscan/build && \
    cd /tmp/hyperscan/build && \
    cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_STATIC_AND_SHARED=ON .. && \
    make -j12 && \
    make install

ENV PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib64/pkgconfig

RUN go build -ldflags "-extldflags=-static -extldflags=-lm" -o /go/bin/tg-multibot ./cmd/tg-multibot && \
    make deps && \
    make plugins

FROM alpine:3.12

RUN apk add --no-cache ca-certificates libc6-compat libstdc++ bash

COPY --from=build /go/bin/tg-multibot /usr/bin/tg-multibot
COPY --from=build /go/src/github.com/paper2code/telegram-multibot/shared/plugins /usr/bin/shared/plugins

# copy hyperscan dependecy
COPY --from=build /usr/local/lib64/pkgconfig/libhs.pc /usr/local/lib64/pkgconfig/libhs.pc
COPY --from=build /usr/local/include/hs/hs.h usr/local/include/hs/hs.h
COPY --from=build /usr/local/include/hs/hs_common.h /usr/local/include/hs/hs_common.h
COPY --from=build /usr/local/include/hs/hs_compile.h /usr/local/include/hs/hs_compile.h
COPY --from=build /usr/local/include/hs/hs_runtime.h /usr/local/include/hs/hs_runtime.h
COPY --from=build /usr/local/lib64/libhs_runtime.a /usr/local/lib64/libhs_runtime.a
COPY --from=build /usr/local/lib64/libhs_runtime.so.5.3.0 /usr/local/lib64/libhs_runtime.so.5.3.0
COPY --from=build /usr/local/lib64/libhs_runtime.so.5 /usr/local/lib64/libhs_runtime.so.5
COPY --from=build /usr/local/lib64/libhs_runtime.so /usr/local/lib64/libhs_runtime.so
COPY --from=build /usr/local/lib64/libhs.a /usr/local/lib64/libhs.a
COPY --from=build /usr/local/lib64/libhs.so.5.3.0 /usr/local/lib64/libhs.so.5.3.0
COPY --from=build /usr/local/lib64/libhs.so.5 /usr/local/lib64/libhs.so.5
COPY --from=build /usr/local/lib64/libhs.so /usr/local/lib64/libhs.so

VOLUME /data
WORKDIR /data

ENV PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib64/pkgconfig
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib64

# Expose port for live server
EXPOSE 8899

CMD ["/usr/bin/tg-multibot"]
