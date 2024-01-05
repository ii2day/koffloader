include Makefile.defs


.PHONY: e2e_init
e2e_init:
	$(QUIET)  make -C test kind-init


.PHONY: clean
clean:
	$(QUIET)  make -C test clean_e2e