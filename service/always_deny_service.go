package service

import (
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

type AlwaysDenyService struct{}

func NewAlwaysDenyService() *AlwaysDenyService {
	return &AlwaysDenyService{}
}

func (s *AlwaysDenyService) AlwaysDeny(ar v1.AdmissionReview) *v1.AdmissionResponse {
	klog.V(2).Info("calling always-deny")
	klog.V(2).Infoln(ar)
	reviewResponse := v1.AdmissionResponse{}
	reviewResponse.Allowed = false
	reviewResponse.Result = &metav1.Status{Message: "this webhook denies all requests"}
	return &reviewResponse
}
