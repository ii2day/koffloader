// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1"
	"github.com/koffloader-io/koffloader/pkg/k8s/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type koffloaderV1Interface interface {
	RESTClient() rest.Interface
	MybooksGetter
}

// koffloaderV1Client is used to interact with features provided by the koffloader.koffloader.io group.
type koffloaderV1Client struct {
	restClient rest.Interface
}

func (c *koffloaderV1Client) Mybooks() MybookInterface {
	return newMybooks(c)
}

// NewForConfig creates a new koffloaderV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*koffloaderV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new koffloaderV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*koffloaderV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &koffloaderV1Client{client}, nil
}

// NewForConfigOrDie creates a new koffloaderV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *koffloaderV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new koffloaderV1Client for the given RESTClient.
func New(c rest.Interface) *koffloaderV1Client {
	return &koffloaderV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *koffloaderV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
