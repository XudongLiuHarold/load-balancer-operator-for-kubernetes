// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package ako

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	akoov1alpha1 "github.com/vmware-samples/load-balancer-operator-for-kubernetes/api/v1alpha1"
	appv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeClient "sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var _ = Describe("AKO", func() {
	var (
		ctx      context.Context
		fclient  client.Client
		logger   logr.Logger
		ss       *appv1.StatefulSet
		finished bool
		err      error
		createSS bool
	)
	BeforeEach(func() {
		ctx = context.Background()
		scheme := runtime.NewScheme()
		Expect(appv1.AddToScheme(scheme)).NotTo(HaveOccurred())
		fclient = fakeClient.NewClientBuilder().WithScheme(scheme).Build()
		logger = log.Log
		log.SetLogger(zap.New())
		ss = &appv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      akoStatefulSetName,
				Namespace: akoov1alpha1.AviNamespace,
			},
		}
		createSS = true
	})

	JustBeforeEach(func() {
		if createSS {
			Expect(fclient.Create(ctx, ss)).ToNot(HaveOccurred())
		}
		finished, err = CleanupFinished(ctx, fclient, logger)
	})

	When("StatefulSet does not exist", func() {
		BeforeEach(func() {
			createSS = false
		})
		It("should claim finished", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(finished).To(BeTrue())
		})
	})

	When("no annotation is in StatefulSet", func() {
		It("should not claim finished", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(finished).To(BeFalse())
		})
	})
	When("Clean up annotation is in progress", func() {
		BeforeEach(func() {
			ss.Annotations = map[string]string{
				akoCleanUpAnnotationKey: akoCleanUpInProgressStatus,
			}
		})
		It("should not claim finished", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(finished).To(BeFalse())
		})
	})
	When("Clean up annotation is done", func() {
		BeforeEach(func() {
			ss.Annotations = map[string]string{
				akoCleanUpAnnotationKey: akoCleanUpFinishedStatus,
			}
		})
		It("should claim finished", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(finished).To(BeTrue())
		})
	})
	When("Clean up annotation is timeout", func() {
		BeforeEach(func() {
			ss.Annotations = map[string]string{
				akoCleanUpAnnotationKey: akoCleanUpTimeoutStatus,
			}
		})
		It("should claim finished", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(finished).To(BeTrue())
		})
	})
})
