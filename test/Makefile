include ../Makefile.defs
include Makefile.defs

.PHONY: kind-init
kind-init: check_env prepare
	@echo -e "\033[35mInit kind for the cluster \033[0m"
	make setup_kind
	@if [ "$(INSTALL_CILIUM)" == "true" ] ; then \
  		echo -e "\033[35m install cilium for the clusters \033[0m" ;\
		make install_cilium_on_all_cluster ; \
	fi ;
	@echo "success init kind clusters"
	@echo ""
	@echo "-----------------------------------------------------------------------------------------------------"
	@echo "       ip family: $(E2E_IP_FAMILY)"
	@for CLUSTER_ID in $(KIND_CLUSTER_IDS); do \
		E2E_KIND_CLUSTER_NAME=$$(echo $(KIND_CLUSTER_NAMES) | cut -d ' ' -f $$CLUSTER_ID); \
		KIND_CLUSTER_KUBECONFIG=$$(echo $(KIND_CLUSTER_KUBECONFIGS) | cut -d ' ' -f $$CLUSTER_ID); \
		echo -e "       cluster name: " $$E2E_KIND_CLUSTER_NAME; \
		echo -e "       cluster kubeconfig:" $$KIND_CLUSTER_KUBECONFIG ; \
		echo "" ;\
	done
	@echo "-----------------------------------------------------------------------------------------------------"
	@echo ""


.PHONY: setup_kind
setup_kind:
	@for CLUSTER_ID in $(KIND_CLUSTER_IDS); do \
		export E2E_IP_FAMILY=$(E2E_IP_FAMILY) ; \
    	echo E2E_IP_FAMILY=$(E2E_IP_FAMILY) ; \
    	export CLUSTER_DIR=$(CLUSTER_DIR) ; \
    	echo CLUSTER_DIR=$(CLUSTER_DIR) ; \
    	export E2E_KUBE_PROXY_MODE=$(E2E_KUBE_PROXY_MODE) ; \
    	echo E2E_KUBE_PROXY_MODE=$(E2E_KUBE_PROXY_MODE) ; \
    	export E2E_KIND_NODE_IMAGE=$(E2E_KIND_NODE_IMAGE) ; \
    	echo E2E_KIND_NODE_IMAGE=$(E2E_KIND_NODE_IMAGE) ; \
		export GLOBAL_KIND_CONFIG_PATH=$(GLOBAL_KIND_CONFIG_PATH) ; \
		echo GLOBAL_KIND_CONFIG_PATH=$(GLOBAL_KIND_CONFIG_PATH) ; \
		export E2E_DISABLE_DEFAULT_CNI=$(E2E_DISABLE_DEFAULT_CNI) ; \
		echo E2E_DISABLE_DEFAULT_CNI=$(E2E_DISABLE_DEFAULT_CNI) ; \
  		# cluster id \
  		CILIUM_CLUSTER_ID=$$CLUSTER_ID ; \
  		export CILIUM_CLUSTER_ID=$$CILIUM_CLUSTER_ID ; \
    	echo CILIUM_CLUSTER_ID=$$CILIUM_CLUSTER_ID ; \
    	# cluster name \
		E2E_KIND_CLUSTER_NAME=$$(echo $(KIND_CLUSTER_NAMES) | cut -d ' ' -f $$CLUSTER_ID); \
		export E2E_KIND_CLUSTER_NAME=$$E2E_KIND_CLUSTER_NAME ; \
		echo E2E_KIND_CLUSTER_NAME=$$E2E_KIND_CLUSTER_NAME ; \
		# cluster service and pod cidr \
		if [ "$(E2E_IP_FAMILY)" == "ipv4" ] ; then \
			E2E_KIND_SERVICE_CIDR=$$(echo $(KIND_CLUSTER_SERVICE_IPV4_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
			E2E_KIND_POD_CIDR=$$(echo $(KIND_CLUSTER_POD_IPV4_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		elif [ "$(E2E_IP_FAMILY)" == "ipv6" ] ; then \
			E2E_KIND_SERVICE_CIDR=$$(echo $(KIND_CLUSTER_SERVICE_IPV6_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
			E2E_KIND_POD_CIDR=$$(echo $(KIND_CLUSTER_POD_IPV6_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		else \
			E2E_KIND_SERVICE_CIDR=$$(echo $(KIND_CLUSTER_SERVICE_DUAL_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
			E2E_KIND_POD_CIDR=$$(echo $(KIND_CLUSTER_POD_DUAL_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		fi ; \
		export E2E_KIND_SERVICE_CIDR=$$E2E_KIND_SERVICE_CIDR ; \
		echo E2E_KIND_SERVICE_CIDR=$$E2E_KIND_SERVICE_CIDR ; \
		export E2E_KIND_POD_CIDR=$$E2E_KIND_POD_CIDR ; \
		echo E2E_KIND_POD_CIDR=$$E2E_KIND_POD_CIDR ; \
		KIND_CLUSTER_KUBECONFIG=$$(echo $(KIND_CLUSTER_KUBECONFIGS) | cut -d ' ' -f $$CLUSTER_ID); \
		export KIND_CLUSTER_KUBECONFIG=$$KIND_CLUSTER_KUBECONFIG ; \
		echo KIND_CLUSTER_KUBECONFIG=$$KIND_CLUSTER_KUBECONFIG ; \
		echo setup $$E2E_KIND_CLUSTER_NAME kind cluster; \
		bash $(ROOT_DIR)/test/scripts/init-kind.sh ; \
	done


.PHONY: install_cilium
install_cilium: E2E_KIND_CLUSTER_NAME ?=
install_cilium: CILIUM_CLUSTER_ID ?=
install_cilium: E2E_IP_FAMILY ?= $E2E_IP_FAMILY
install_cilium: KIND_CLUSTER_KUBECONFIG ?=
install_cilium: HTTP_PROXY ?=
install_cilium: CILIUM_CLUSTER_IPV4_CIDR ?=
install_cilium: CILIUM_CLUSTER_IPV6_CIDR ?=
install_cilium: E2E_CILIUM_IMAGE_REPO ?= $(E2E_CILIUM_IMAGE_REPO)
install_cilium:
	echo "install cilium on" $(E2E_KIND_CLUSTER_NAME) ; \
	export E2E_KIND_CLUSTER_NAME=$(E2E_KIND_CLUSTER_NAME); \
	echo E2E_KIND_CLUSTER_NAME=$(E2E_KIND_CLUSTER_NAME); \
	export CILIUM_CLUSTER_ID=$(CILIUM_CLUSTER_ID); \
	echo CILIUM_CLUSTER_ID=$(CILIUM_CLUSTER_ID); \
	export E2E_IP_FAMILY=$(E2E_IP_FAMILY); \
	echo E2E_IP_FAMILY=$(E2E_IP_FAMILY); \
	export KIND_CLUSTER_KUBECONFIG=$(KIND_CLUSTER_KUBECONFIG); \
	echo KIND_CLUSTER_KUBECONFIG=$(KIND_CLUSTER_KUBECONFIG); \
	export HTTP_PROXY=$(HTTP_PROXY); \
	echo HTTP_PROXY=$(HTTP_PROXY); \
	export E2E_CILIUM_IMAGE_REPO=$(E2E_CILIUM_IMAGE_REPO); \
	echo E2E_CILIUM_IMAGE_REPO=$(E2E_CILIUM_IMAGE_REPO); \
	export CILIUM_CLUSTER_IPV4_CIDR=$(CILIUM_CLUSTER_IPV4_CIDR); \
	echo CILIUM_CLUSTER_IPV4_CIDR=$(CILIUM_CLUSTER_IPV4_CIDR); \
	export CILIUM_CLUSTER_IPV6_CIDR=$(CILIUM_CLUSTER_IPV6_CIDR); \
	echo CILIUM_CLUSTER_IPV6_CIDR=$(CILIUM_CLUSTER_IPV6_CIDR); \
	bash $(ROOT_DIR)/test/scripts/install-cilium.sh ;


.PHONY: install_cilium_on_all_cluster
install_cilium_on_all_cluster:
	@for CLUSTER_ID in $(KIND_CLUSTER_IDS); do \
		E2E_KIND_CLUSTER_NAME=$$(echo $(KIND_CLUSTER_NAMES) | cut -d ' ' -f $$CLUSTER_ID); \
  		CILIUM_CLUSTER_ID=$$CLUSTER_ID ; \
		KIND_CLUSTER_KUBECONFIG=$$(echo $(KIND_CLUSTER_KUBECONFIGS) | cut -d ' ' -f $$CLUSTER_ID); \
		CILIUM_CLUSTER_IPV4_CIDR=$$(echo $(CILIUM_CLUSTER_IPV4_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		CILIUM_CLUSTER_IPV6_CIDR=$$(echo $(CILIUM_CLUSTER_IPV6_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		make install_cilium -e E2E_KIND_CLUSTER_NAME=$$E2E_KIND_CLUSTER_NAME -e CILIUM_CLUSTER_ID=$$CILIUM_CLUSTER_ID \
			-e E2E_IP_FAMILY=$(E2E_IP_FAMILY) -e KIND_CLUSTER_KUBECONFIG=$$KIND_CLUSTER_KUBECONFIG \
			-e HTTP_PROXY=$(HTTP_PROXY) -e E2E_CILIUM_IMAGE_REPO=$(E2E_CILIUM_IMAGE_REPO) \
			-e CILIUM_CLUSTER_IPV4_CIDR=$$CILIUM_CLUSTER_IPV4_CIDR -e CILIUM_CLUSTER_IPV6_CIDR=$$CILIUM_CLUSTER_IPV6_CIDR; \
	done


.PHONY: check_env
check_env:
	$(QUIET) [ -n "$(E2E_CLUSTER_NAME)" ] || { echo "error, miss E2E_CLUSTER_NAME " ; false ; }
	$(QUIET) ( [ "$(E2E_IP_FAMILY)" == "ipv4" ] || [ "$(E2E_IP_FAMILY)" == "ipv6" ]  || [ "$(E2E_IP_FAMILY)" == "dual" ] ) \
			|| { echo "error, E2E_IP_FAMILY=$(E2E_IP_FAMILY) must be ipv4/ipv6/dual" ;  exit 1 ; }


.PHONY: prepare
prepare:
	echo -e "\033[35m Check The Tools For Ready: \033[0m"
	$(QUIET) JUST_CLI_CHECK=true scripts/install-tools.sh
	$(QUIET) mkdir -p $(DOWNLOAD_DIR)
	$(QUIET) IMAGE_LIST="" ; \
		 [ -n "$(PYROSCOPE_LOCAL_PORT)" ] && IMAGE_LIST+=" $(IMAGE_PYROSCOPE_NAME) " ; \
		 IMAGE_LIST+=" $(TEST_IMAGE_NAME) " ; \
		 ARCH=$(ARCH) IMAGE_LIST="$${IMAGE_LIST}" \
		 	CNI_PACKAGE_VERSION=$(CNI_PACKAGE_VERSION)  scripts/prepare.sh $(DOWNLOAD_DIR)
	- sudo sysctl -w fs.inotify.max_user_watches=524288 || true
	- sudo sysctl -w fs.inotify.max_user_instances=8192 || true
	- sudo sysctl -w net.ipv6.conf.all.disable_ipv6=0

.PHONY: clean_e2e
clean_e2e:
	@rm -rf $(CLUSTER_DIR)
	-@  kind get clusters | xargs -n1  kind delete cluster --name
	-@ docker stop $(PYROSCOPE_CONTAINER_NAME) &>/dev/null
	-@ docker rm $(PYROSCOPE_CONTAINER_NAME) &>/dev/null