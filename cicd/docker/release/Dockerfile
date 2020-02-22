ARG GIT_COMMIT_SHORT_HASH
FROM hobbyfarm/gargantua-sdk:${GIT_COMMIT_SHORT_HASH} AS sdk

RUN RUN go install -v ./...


FROM alpine:latest

COPY --from=sdk /go/bin/gargantua /usr/local/bin/

ENTRYPOINT ["gargantua"] 
CMD ["-v=9", "-logtostderr"] 
