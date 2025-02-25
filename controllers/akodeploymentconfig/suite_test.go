// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package akodeploymentconfig_test

import (
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	akoov1alpha1 "github.com/vmware-samples/load-balancer-operator-for-kubernetes/api/v1alpha1"
	"github.com/vmware-samples/load-balancer-operator-for-kubernetes/controllers/akodeploymentconfig"
	"github.com/vmware-samples/load-balancer-operator-for-kubernetes/pkg/aviclient"
	"github.com/vmware-samples/load-balancer-operator-for-kubernetes/pkg/test/builder"
	testutil "github.com/vmware-samples/load-balancer-operator-for-kubernetes/pkg/test/util"
	akov1alpha1 "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/apis/ako/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	networkv1alpha1 "github.com/vmware-samples/load-balancer-operator-for-kubernetes/api/v1alpha1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"

	adccluster "github.com/vmware-samples/load-balancer-operator-for-kubernetes/controllers/akodeploymentconfig/cluster"
	cluster "github.com/vmware-samples/load-balancer-operator-for-kubernetes/controllers/cluster"
	ctrlmgr "sigs.k8s.io/controller-runtime/pkg/manager"
)

// suite is used for unit and integration testing this controller.
var suite = builder.NewTestSuiteForController(
	func(mgr ctrlmgr.Manager) error {
		rec := &akodeploymentconfig.AKODeploymentConfigReconciler{
			Client: mgr.GetClient(),
			Log:    ctrl.Log.WithName("controllers").WithName("AKODeploymentConfig"),
			Scheme: mgr.GetScheme(),
		}
		builder.FakeAvi = aviclient.NewFakeAviClient()
		rec.SetAviClient(builder.FakeAvi)

		adcClusterReconciler := adccluster.NewReconciler(rec.Client, rec.Log, rec.Scheme)
		adcClusterReconciler.GetRemoteClient = adccluster.GetFakeRemoteClient
		rec.ClusterReconciler = adcClusterReconciler

		if err := rec.SetupWithManager(mgr); err != nil {
			return err
		}

		// involve the cluster controller as well for the resetting skip-default-adc label test
		if err := (&cluster.ClusterReconciler{
			Client: mgr.GetClient(),
			Log:    ctrl.Log.WithName("controllers").WithName("Cluster"),
			Scheme: mgr.GetScheme(),
		}).SetupWithManager(mgr); err != nil {
			return err
		}

		return nil
	},
	func(scheme *runtime.Scheme) (err error) {
		err = networkv1alpha1.AddToScheme(scheme)
		if err != nil {
			return err
		}
		err = akoov1alpha1.AddToScheme(scheme)
		if err != nil {
			return err
		}
		err = akov1alpha1.AddToScheme(scheme)
		if err != nil {
			return err
		}
		err = corev1.AddToScheme(scheme)
		if err != nil {
			return err
		}
		err = clusterv1.AddToScheme(scheme)
		if err != nil {
			return err
		}
		return nil
	},
	filepath.Join(testutil.FindModuleDir("sigs.k8s.io/cluster-api"), "config", "crd", "bases"),
	filepath.Join(testutil.FindModuleDir("github.com/vmware/load-balancer-and-ingress-services-for-kubernetes"), "helm", "ako", "crds"),
)

func TestController(t *testing.T) {
	suite.Register(t, "AKO Operator", intgTests, unitTests)
}

var _ = BeforeSuite(suite.BeforeSuite)

var _ = AfterSuite(suite.AfterSuite)

func intgTests() {
	Describe("AkoDeploymentConfigController Test", intgTestAkoDeploymentConfigController)
}

func unitTests() {
	Describe("Ensure static ranges Test", unitTestEnsureStaticRanges)
}
