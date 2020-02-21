FROM golang
RUN mkdir /gtracker
ADD . /gtracker
WORKDIR /gtracker
RUN go build -o server .
CMD ["/gtracker/server"]
