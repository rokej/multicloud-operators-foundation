package agentdeploy

import (
	"context"
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/openshift/library-go/pkg/operator/events"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"open-cluster-management.io/addon-framework/pkg/addonmanager/constants"
	addonapiv1alpha1 "open-cluster-management.io/api/addon/v1alpha1"
	addonv1alpha1client "open-cluster-management.io/api/client/addon/clientset/versioned"
	workv1client "open-cluster-management.io/api/client/work/clientset/versioned"
	worklister "open-cluster-management.io/api/client/work/listers/work/v1"
	workapiv1 "open-cluster-management.io/api/work/v1"
)

func preDeleteHookWorkName(addonName string) string {
	return fmt.Sprintf("addon-%s-pre-delete", addonName)
}

func hasFinalizer(existingFinalizers []string, finalizer string) bool {
	for _, f := range existingFinalizers {
		if f == finalizer {
			return true
		}
	}
	return false
}

func removeFinalizer(existingFinalizers []string, finalizer string) []string {
	var rst []string
	for _, f := range existingFinalizers {
		if f != finalizer {
			rst = append(rst, f)
		}
	}
	return rst
}

func manifestsEqual(new, old []workapiv1.Manifest) bool {
	if len(new) != len(old) {
		return false
	}

	for i := range new {
		if !equality.Semantic.DeepEqual(new[i].Raw, old[i].Raw) {
			return false
		}
	}
	return true
}

func manifestWorkSpecEqual(new, old workapiv1.ManifestWorkSpec) bool {
	if !manifestsEqual(new.Workload.Manifests, old.Workload.Manifests) {
		return false
	}
	if !equality.Semantic.DeepEqual(new.ManifestConfigs, old.ManifestConfigs) {
		return false
	}
	if !equality.Semantic.DeepEqual(new.DeleteOption, old.DeleteOption) {
		return false
	}
	return true
}

func newManifestWork(workName, addonName, clusterName string, manifests []workapiv1.Manifest) *workapiv1.ManifestWork {
	if len(manifests) == 0 {
		return nil
	}

	return &workapiv1.ManifestWork{
		ObjectMeta: metav1.ObjectMeta{
			Name:      workName,
			Namespace: clusterName,
			Labels: map[string]string{
				constants.AddonLabel: addonName,
			},
		},
		Spec: workapiv1.ManifestWorkSpec{
			Workload: workapiv1.ManifestsTemplate{
				Manifests: manifests,
			},
		},
	}
}

// isPreDeleteHookObject check the object is a pre-delete hook resources.
// currently, we only support job and pod as hook resources.
// we use WellKnownStatus here to get the job/pad status fields to check if the job/pod is completed.
func isPreDeleteHookObject(obj *unstructured.Unstructured) (bool, *workapiv1.ManifestConfigOption) {
	var resource string
	gvk := obj.GroupVersionKind()
	switch gvk.Kind {
	case "Job":
		resource = "jobs"
	case "Pod":
		resource = "pods"
	default:
		return false, nil
	}
	labels := obj.GetLabels()
	if _, ok := labels[constants.PreDeleteHookLabel]; !ok {
		return false, nil
	}

	return true, &workapiv1.ManifestConfigOption{
		ResourceIdentifier: workapiv1.ResourceIdentifier{
			Group:     gvk.Group,
			Resource:  resource,
			Name:      obj.GetName(),
			Namespace: obj.GetNamespace(),
		},
		FeedbackRules: []workapiv1.FeedbackRule{
			{
				Type: workapiv1.WellKnownStatusType,
			},
		},
	}
}

func buildManifestWorkFromObject(
	cluster, addonName string,
	objects []runtime.Object) (deployManifestWork, hookManifestWork *workapiv1.ManifestWork, err error) {
	var deployManifests []workapiv1.Manifest
	var hookManifests []workapiv1.Manifest
	var manifestConfigs []workapiv1.ManifestConfigOption

	for _, object := range objects {
		rawObject, err := runtime.Encode(unstructured.UnstructuredJSONScheme, object)
		if err != nil {
			return nil, nil, err
		}
		unstructuredObj := &unstructured.Unstructured{}
		err = unstructuredObj.UnmarshalJSON(rawObject)
		if err != nil {
			return nil, nil, err
		}
		isHookObject, manifestConfig := isPreDeleteHookObject(unstructuredObj)
		if isHookObject {
			hookManifests = append(hookManifests, workapiv1.Manifest{
				RawExtension: runtime.RawExtension{Raw: rawObject},
			})
			manifestConfigs = append(manifestConfigs, *manifestConfig)
		} else {
			deployManifests = append(deployManifests, workapiv1.Manifest{
				RawExtension: runtime.RawExtension{Raw: rawObject},
			})
		}
	}

	deployManifestWork = newManifestWork(constants.DeployWorkName(addonName), addonName, cluster, deployManifests)
	hookManifestWork = newManifestWork(preDeleteHookWorkName(addonName), addonName, cluster, hookManifests)
	if hookManifestWork != nil {
		hookManifestWork.Spec.ManifestConfigs = manifestConfigs
	}

	return deployManifestWork, hookManifestWork, nil
}

func applyWork(
	ctx context.Context,
	workClient workv1client.Interface,
	workLister worklister.ManifestWorkLister,
	cache *workCache,
	eventRecorder events.Recorder,
	required *workapiv1.ManifestWork) (*workapiv1.ManifestWork, error) {
	existingWork, err := workLister.ManifestWorks(required.Namespace).Get(required.Name)
	existingWork = existingWork.DeepCopy()
	if err != nil {
		if errors.IsNotFound(err) {
			existingWork, err = workClient.WorkV1().ManifestWorks(required.Namespace).Create(ctx, required, metav1.CreateOptions{})
			if err == nil {
				eventRecorder.Eventf("ManifestWorkCreated", "Created %s/%s because it was missing", required.Namespace, required.Name)
				cache.updateCache(required, existingWork)
				return existingWork, nil
			}
			eventRecorder.Warningf("ManifestWorkCreateFailed", "Failed to create ManifestWork %s/%s: %v", required.Namespace, required.Name, err)
			return nil, err
		}
		return nil, err
	}

	if cache.safeToSkipApply(required, existingWork) {
		return existingWork, nil
	}

	if manifestWorkSpecEqual(required.Spec, existingWork.Spec) {
		return existingWork, nil
	}

	existingWork.Spec = required.Spec
	existingWork, err = workClient.WorkV1().ManifestWorks(existingWork.Namespace).Update(ctx, existingWork, metav1.UpdateOptions{})
	if err == nil {
		cache.updateCache(required, existingWork)
		eventRecorder.Eventf("ManifestWorkUpdate", "Updated %s/%s because it was changing", required.Namespace, required.Name)
		return existingWork, nil
	}
	eventRecorder.Warningf("ManifestWorkUpdateFailed", "Failed to update ManifestWork %s/%s: %v", required.Namespace, required.Name, err)
	return nil, err
}

func FindManifestValue(
	resourceStatus workapiv1.ManifestResourceStatus,
	identifier workapiv1.ResourceIdentifier,
	valueName string) workapiv1.FieldValue {
	for _, manifest := range resourceStatus.Manifests {
		values := manifest.StatusFeedbacks.Values
		if len(values) == 0 {
			return workapiv1.FieldValue{}
		}
		resourceMeta := manifest.ResourceMeta
		if identifier.Group == resourceMeta.Group &&
			identifier.Resource == resourceMeta.Resource &&
			identifier.Name == resourceMeta.Name &&
			identifier.Namespace == resourceMeta.Namespace {
			for _, v := range values {
				if v.Name == valueName {
					return v.Value
				}
			}
		}
	}
	return workapiv1.FieldValue{}
}

func patchCondition(ctx context.Context, addonClient addonv1alpha1client.Interface, new, old *addonapiv1alpha1.ManagedClusterAddOn) error {
	if equality.Semantic.DeepEqual(new.Status.Conditions, old.Status.Conditions) {
		return nil
	}

	oldData, err := json.Marshal(&addonapiv1alpha1.ManagedClusterAddOn{
		Status: addonapiv1alpha1.ManagedClusterAddOnStatus{
			Conditions: old.Status.Conditions,
		},
	})
	if err != nil {
		return err
	}

	newData, err := json.Marshal(&addonapiv1alpha1.ManagedClusterAddOn{
		ObjectMeta: metav1.ObjectMeta{
			UID:             new.UID,
			ResourceVersion: new.ResourceVersion,
		},
		Status: addonapiv1alpha1.ManagedClusterAddOnStatus{
			Conditions: new.Status.Conditions,
		},
	})
	if err != nil {
		return err
	}

	patchBytes, err := jsonpatch.CreateMergePatch(oldData, newData)
	if err != nil {
		return fmt.Errorf("failed to create patch for addon %s: %w", new.Name, err)
	}
	if err != nil {
		return err
	}

	_, err = addonClient.AddonV1alpha1().ManagedClusterAddOns(new.Namespace).Patch(ctx, new.Name, types.MergePatchType, patchBytes, metav1.PatchOptions{}, "status")
	return err
}