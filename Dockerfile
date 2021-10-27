FROM golang:1.17.2

WORKDIR /go/src/github.com/schultet/goa/
COPY cmd ./cmd
COPY pkg ./pkg
COPY scripts ./scripts
COPY translate ./translate
COPY test ./test
COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY search.sh ./search.sh


RUN go get -d -v ./...
RUN go install -v ./...

# Install python3
RUN apt-get update 
RUN apt-get -y install python3
#RUN apt-get -y install python3-setuptools
#RUN apt-get -y install python3-pip
# Install requirements
#COPY requirements.txt ./requirements.txt
#RUN pip3 install --user -r ./requirements.txt

#CMD ["search.sh"]
CMD ["bash", "search.sh", "-t", "test/", "-c", "mafs"]
