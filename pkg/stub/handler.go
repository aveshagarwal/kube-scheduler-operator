package stub

import (
	"context"
	"fmt"

	"github.com/aveshagarwal/kube-scheduler-operator/pkg/apis/kube-scheduler/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.KubeScheduler:
		err := sdk.Create(newKubeSchedulerPod(o))
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create kube-scheduler pod : %v", err)
			return err
		}
	}
	return nil
}

// newKubeSchedulerPod creates kube-scheduler
func newKubeSchedulerPod(cr *v1alpha1.KubeScheduler) *corev1.Pod {
	/*labels := map[string]string{
		"app": "busy-box",
	}*/
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kube-scheduler-pod",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "KubeScheduler",
				}),
			},
			//Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "kubescheduler-container",
					Image:   "openshift/origin-hyperkube",
					Command: []string{"/usr/bin/hyperkube", "kube-scheduler", fmt.Sprintf("--scheduler-name=%s", "custom-scheduler")},
				},
			},
		},
	}
}
