-include Makefile.include

CFLAGS += -fPIC
CFLAGS += -I.
CFLAGS += $(PYTHON_CFLAGS)
CFLAGS += $(NANOMSG_CFLAGS)
LDFLAGS += $(PYTHON_LDFLAGS)
LDFLAGS += $(NANOMSG_LDFLAGS)


all: py2
#echo "Specify target py2 or py3"

py2: nanomsg2module.so

nanomsg2module.so: nanomsg2.c
	$(CC) $(CFLAGS) -shared -o $@ $< $(LDFLAGS)

clean:
	rm -f *.o
	rm -f *.so
	rm -f $(TARGETS)
	
.PHONY: all clean
