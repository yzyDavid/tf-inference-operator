package tfinference

import (
	"strconv"

	"context"

	servingv1 "github.com/yzyDavid/tf-inference-operator/pkg/apis/serving/v1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_tfinference")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new TfInference Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileTfInference{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("tfinference-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource TfInference
	err = c.Watch(&source.Kind{Type: &servingv1.TfInference{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner TfInference
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &servingv1.TfInference{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTfInference{}

// ReconcileTfInference reconciles a TfInference object
type ReconcileTfInference struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a TfInference object and makes changes based on the state read
// and what is in the TfInference.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileTfInference) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling TfInference")

	// Fetch the TfInference instance
	instance := &servingv1.TfInference{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define Deployment objects given by algorithm
	deploymentMetas := getDeploymentMetas(instance.Spec.Models, instance.Spec.Nodes,instance.Spec.Groups)
	deployments := make([]*appsv1.Deployment,0)
    for _, meta := range deploymentMetas {
        deployments = append(deployments, r.newDeployment(instance, &meta))
        reqLogger.Info("New Deployment with config:", "Hash", meta.Hash, "Models", meta.Models, "Replicas", meta.Replicas)
	}

	// Check all Deployment already exist, if not, create it
	for _, deployment := range deployments {
		found := &appsv1.Deployment{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, found)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
			err = r.client.Create(context.TODO(), deployment)
			if err != nil {
				return reconcile.Result{}, err
			}
		} else if err != nil {
			return reconcile.Result{}, err
		}
	}

	// TODO: remove not yet used Deployments

	// Pod created successfully - don't requeue
	return reconcile.Result{}, nil
}

// new One Deployment for TfInference
func (r *ReconcileTfInference) newDeployment(cr *servingv1.TfInference, meta *DeploymentMeta) *appsv1.Deployment {
	name := cr.Name + "-" + strconv.FormatUint(uint64(meta.Hash), 16)
	labels := map[string]string{
		"app":             "tf_inference",
		"tf_inference_cr": name,
	}
	replicas := meta.Replicas

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "tensorflow/serving",
						Name:  "tf-inference",
						VolumeMounts: []corev1.VolumeMount{{
							Name: "models-volume",
							MountPath: "/models",
						}},
					}},
					Volumes: []corev1.Volume{{
						Name: "models-volume",
						VolumeSource: corev1.VolumeSource{
							PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
								ClaimName: "local-nfs-pvc",
							},
						},
					}},
				},
			},
		},
	}

	err := controllerutil.SetControllerReference(cr, deployment, r.scheme)
	if err != nil {
		// TODO(Zhenyun Yu): should logging or return error.
		return &appsv1.Deployment{}
	}

	return deployment
}
