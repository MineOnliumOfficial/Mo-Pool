FROM golang:1.17-buster AS builder
ENV CGO_ENABLED 0
WORKDIR /project
COPY . ./
RUN cd /project/cmd/be && go build -o /project/bin/

FROM registry.access.redhat.com/ubi8/ubi-minimal
EXPOSE 8080
ENV KO_DATA_PATH /kodata
COPY --from=builder /project/cmd/be/kodata/ ${KO_DATA_PATH}/
COPY --from=builder /project/bin/be /be

ENTRYPOINT ["/be"]
