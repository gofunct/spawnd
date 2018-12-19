FROM golang
COPY    . "github.com/gofunct/grpcgen/example/cmd"
WORKDIR "github.com/gofunct/grpcgen/example/cmd"
CMD     ["example"]
EXPOSE  8000 9000
RUN     make install
