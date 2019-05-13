NANOMSG_DIR=nanomsg
NANOMSG4PY_DIR=nanomsg4py

nanomsg4py:
	$(MAKE) -C $(NANOMSG4PY_DIR)

nanomsg:
	cd $(NANOMSG_DIR) \
		&& mkdir build \
		&& cd build \
		&& cmake .. -DCMAKE_INSTALL_PREFIX=./ \
		&& cmake --build . \
		&& cmake --build . --target install

clean:
	$(MAKE) -C $(NANOMSG4PY_DIR) clean

.PHONY: nanomsg nanomsg4py
