FROM alpine:3.8
RUN mkdir -p /go/src/vote
ADD ./build /go/src/vote/
WORKDIR /go/src/vote
CMD ["/go/src/vote/main"]

EXPOSE 5110