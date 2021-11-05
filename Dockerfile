FROM golang:1.17.2

WORKDIR /go/src/github.com/schultet/goa/
COPY . ./

RUN go get -d -v ./...
RUN go install -v ./...

# Install python3
RUN apt-get update && apt-get -y install python3-pip
RUN pip3 install --user -r ./requirements.txt

CMD ["bash"] 
