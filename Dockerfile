FROM golang
RUN mkdir /gtracker
ADD . /gtracker
WORKDIR /gtracker
RUN go build -p server .
CMD ["/gtracker/server"]