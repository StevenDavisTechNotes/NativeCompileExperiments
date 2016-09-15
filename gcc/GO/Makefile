all: scimark2 

CFLAGS = -Ofast

scimark2 : *.go
	go build -compiler gccgo -gccgoflags '$(CFLAGS)' -o scimark2

clean:
	rm $(OBJS) scimark2
