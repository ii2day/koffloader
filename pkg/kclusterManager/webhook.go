// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterManager

import (
	"context"
	"fmt"
	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1"
	"go.uber.org/zap"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"time"
)

// --------------------

type webhookhander struct {
	logger *zap.Logger
}

var _ webhook.CustomValidator = (*webhookhander)(nil)

// mutating webhook
func (s *webhookhander) Default(ctx context.Context, obj runtime.Object) error {
	logger := s.logger.Named("mutating wehbook")

	r, ok := obj.(*crd.Kcluster)
	if !ok {
		s := "failed to get obj"
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("obj: %+v", r)
	r.Annotations["test"] = "add-by-mutating-webhook"

	finalizerName := "koffloader.koffloader.io"
	if dt := r.GetDeletionTimestamp(); dt.IsZero() && !controllerutil.ContainsFinalizer(client.Object(r), finalizerName) {
		controllerutil.AddFinalizer(client.Object(r), finalizerName)
	}

	return nil

}

func (s *webhookhander) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	logger := s.logger.Named("validating create webhook")

	r, ok := obj.(*crd.Kcluster)
	if !ok {
		s := "failed to get obj"
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil
}

func (s *webhookhander) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	logger := s.logger.Named("validating update webhook")

	old, ok := oldObj.(*crd.Kcluster)
	if !ok {
		s := "failed to get oldObj"
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	new, ok := newObj.(*crd.Kcluster)
	if !ok {
		s := "failed to get newObj"
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("oldObj: %+v", old)
	logger.Sugar().Infof("newObj: %+v", new)

	return nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (s *webhookhander) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	logger := s.logger.Named("validating delete webhook")

	r, ok := obj.(*crd.Kcluster)
	if !ok {
		s := "failed to get obj"
		logger.Error(s)
		return apierrors.NewBadRequest(s)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil
}

// --------------------

// https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/builder/example_webhook_test.go
// https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/builder/webhook_test.go
func (s *kclusterManager) RunWebhookServer(webhookPort int, tlsDir string) {
	logger := s.logger
	r := &webhookhander{
		logger: logger,
	}
	s.webhook = r

	logger.Sugar().Infof("setup webhook on port %v, with tls under %v", webhookPort, tlsDir)

	scheme := runtime.NewScheme()
	if e := crd.AddToScheme(scheme); e != nil {
		logger.Sugar().Fatalf("failed to add crd scheme, reason=%v", e)
	}
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		LeaderElection:         false,
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
		// webhook port
		Port: webhookPort,
		// directory that contains the webhook server key and certificate, The server key and certificate must be named tls.key and tls.crt
		CertDir: tlsDir,
	})
	if err != nil {
		logger.Sugar().Fatalf("failed to NewManager, reason=%v", err)
	}

	// the mutating route path : "/mutate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	// the validate route path : "/validate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	e := ctrl.NewWebhookManagedBy(mgr).
		For(&crd.Kcluster{}).
		WithDefaulter(r).
		WithValidator(r).
		Complete()
	if e != nil {
		logger.Sugar().Fatalf("failed to NewWebhookManagedBy, reason=%v", e)
	}

	go func() {
		s := "webhook down"

		// mgr.Start()
		if err := mgr.GetWebhookServer().Start(context.Background()); err != nil {
			s += fmt.Sprintf(", reason=%v", err)
		}
		logger.Error(s)
		time.Sleep(time.Second)
	}()

}
