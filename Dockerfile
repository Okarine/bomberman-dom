FROM golang:1.21.0
LABEL Developers="Avengerist, aivanenk, MartaTino"
LABEL version="1.0"
COPY . /bomberman-dom
WORKDIR /bomberman-dom
RUN go mod download
RUN make build
EXPOSE 8080
CMD ./bomberman-dom
