/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e_test

import (
	"kubeform.dev/kfc/test/e2e/framework"
	aws "kubeform.dev/kubeform/apis/aws/v1alpha1"
	azure "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	digitalocean "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	google "kubeform.dev/kubeform/apis/google/v1alpha1"
	linode "kubeform.dev/kubeform/apis/linode/v1alpha1"
	modules "kubeform.dev/kubeform/apis/modules/v1alpha1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
)

var _ = Describe("KFC", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
	})
	Describe("Test", func() {
		Context("Google", func() {
			var (
				providerRef        *core.Secret
				serviceAccountName string
				secretName         string
				serviceAccount     *google.ServiceAccount
			)

			BeforeEach(func() {
				serviceAccountName = f.GetRandomName("")
				secretName = f.GetRandomName("secret")
				providerRef = f.GoogleProviderRef(secretName)
				serviceAccount = f.ServiceAccount(serviceAccountName, secretName)
			})

			It("should create and delete service account successfully", func() {
				By("Creating GoogleProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating ServiceAccount")
				err = f.CreateServiceAccount(serviceAccount)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Running ServiceAccount")
				f.EventuallyServiceAccountRunning(serviceAccount.ObjectMeta).Should(BeTrue())

				By("Deleting service account")
				err = f.DeleteServiceAccount(serviceAccount.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Deleting ServiceAccount")
				f.EventuallyServiceAccountDeleted(serviceAccount.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})

		Context("AWS", func() {
			var (
				providerRef    *core.Secret
				secretName     string
				dbInstanceName string
				s3Bucket       *aws.S3Bucket
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				dbInstanceName = f.GetRandomName("")
				providerRef = f.AwsProviderRef(secretName)
				s3Bucket = f.S3Bucket(dbInstanceName, secretName)
			})

			It("should create and delete s3 bucket successfully", func() {
				By("Creating AwsProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating S3Bucket")
				err = f.CreateS3Bucket(s3Bucket)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Running S3 Bucket")
				f.EventuallyS3BucketRunning(s3Bucket.ObjectMeta).Should(BeTrue())

				By("Deleting S3 Bucket")
				err = f.DeleteS3Bucket(s3Bucket.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting S3 Bucket")
				f.EventuallyS3BucketDeleted(s3Bucket.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})

		Context("DigitalOcean", func() {
			var (
				providerRef *core.Secret
				secretName  string
				dropletName string
				droplet     *digitalocean.Droplet
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				dropletName = f.GetRandomName("")
				providerRef = f.DigitalOceanProviderRef(secretName)
				droplet = f.Droplets(dropletName, secretName)
			})

			It("should create and delete Droplet successfully", func() {
				By("Creating DigitalOceanProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating droplet")
				err = f.CreateDroplet(droplet)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running droplet")
				f.EventuallyDropletRunning(droplet.ObjectMeta).Should(BeTrue())

				By("Deleting droplet")
				err = f.DeleteDroplet(droplet.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting droplet")
				f.EventuallyDropletDeleted(droplet.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})

		Context("Linode", func() {
			var (
				providerRef   *core.Secret
				sensitiveData *core.Secret
				secretName    string
				instanceName  string
				instance      *linode.Instance
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				instanceName = f.GetRandomName("")
				providerRef = f.LinodeProviderRef(secretName)
				sensitiveData = f.InstanceSensitiveData()
				instance = f.Instance(instanceName, secretName)
			})

			It("should create and delete instance successfully", func() {
				By("Creating LinodeProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Secret")
				err = f.CreateSecret(sensitiveData)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Instance")
				err = f.CreateInstance(instance)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Instance")
				f.EventuallyInstanceRunning(instance.ObjectMeta).Should(BeTrue())

				By("Deleting Instance")
				err = f.DeleteInstance(instance.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting instance")
				f.EventuallyInstanceDeleted(instance.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})

		Context("Azure", func() {
			var (
				providerRef       *core.Secret
				resourceGroupName string
				secretName        string
				resourceGroup     *azure.ResourceGroup
			)

			BeforeEach(func() {
				resourceGroupName = f.GetRandomName("")
				secretName = f.GetRandomName("secret")
				providerRef = f.AzureProviderRef(secretName)
				resourceGroup = f.ResourceGroup(resourceGroupName, secretName)
			})

			It("should create and delete ResourceGroup successfully", func() {
				By("Creating AzureProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating ResourceGroup")
				err = f.CreateResourceGroup(resourceGroup)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running ResourceGroup")
				f.EventuallyResourceGroupRunning(resourceGroup.ObjectMeta).Should(BeTrue())

				By("Deleting ResourceGroup")
				err = f.DeleteResourceGroup(resourceGroup.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting ResourceGroup")
				f.EventuallyResourceGroupDeleted(resourceGroup.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})

		Context("Modules", func() {
			var (
				providerRef        *core.Secret
				serviceAccountName string
				secretName         string
				serviceAccount     *modules.GoogleServiceAccount
			)

			BeforeEach(func() {
				serviceAccountName = f.GetRandomName("")
				secretName = f.GetRandomName("secret")
				providerRef = f.GoogleProviderRef(secretName)
				serviceAccount = f.ModuleServiceAccount(serviceAccountName, secretName)
			})

			It("should create and delete service account successfully", func() {
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating ServiceAccount")
				err = f.CreateModuleServiceAccount(serviceAccount)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running ServiceAccount")
				f.EventuallyModuleServiceAccountRunning(serviceAccount.ObjectMeta).Should(BeTrue())

				By("Deleting service account")
				err = f.DeleteModuleServiceAccount(serviceAccount.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting service account")
				f.EventuallyModuleServiceAccountDeleted(serviceAccount.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})
	})
})
