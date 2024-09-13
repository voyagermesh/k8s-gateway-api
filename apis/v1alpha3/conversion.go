package v1alpha3

import (
	"k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	"unsafe"
)

func autoConvert_v1alpha2_BackendTLSPolicy_To_v1alpha3_BackendTLSPolicy(in *v1alpha2.BackendTLSPolicy, out *BackendTLSPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_BackendTLSPolicySpec_To_v1alpha3_BackendTLSPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = in.Status
	return nil
}

// Convert_v1alpha2_BackendTLSPolicy_To_v1alpha3_BackendTLSPolicy is an autogenerated conversion function.
func Convert_v1alpha2_BackendTLSPolicy_To_v1alpha3_BackendTLSPolicy(in *v1alpha2.BackendTLSPolicy, out *BackendTLSPolicy, s conversion.Scope) error {
	return autoConvert_v1alpha2_BackendTLSPolicy_To_v1alpha3_BackendTLSPolicy(in, out, s)
}

func autoConvert_v1alpha3_BackendTLSPolicy_To_v1alpha2_BackendTLSPolicy(in *BackendTLSPolicy, out *v1alpha2.BackendTLSPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_BackendTLSPolicySpec_To_v1alpha2_BackendTLSPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = in.Status
	return nil
}

// Convert_v1alpha3_BackendTLSPolicy_To_v1alpha2_BackendTLSPolicy is an autogenerated conversion function.
func Convert_v1alpha3_BackendTLSPolicy_To_v1alpha2_BackendTLSPolicy(in *BackendTLSPolicy, out *v1alpha2.BackendTLSPolicy, s conversion.Scope) error {
	return autoConvert_v1alpha3_BackendTLSPolicy_To_v1alpha2_BackendTLSPolicy(in, out, s)
}

func autoConvert_v1alpha2_BackendTLSPolicyList_To_v1alpha3_BackendTLSPolicyList(in *v1alpha2.BackendTLSPolicyList, out *BackendTLSPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackendTLSPolicy, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_BackendTLSPolicy_To_v1alpha3_BackendTLSPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_BackendTLSPolicyList_To_v1alpha3_BackendTLSPolicyList is an autogenerated conversion function.
func Convert_v1alpha2_BackendTLSPolicyList_To_v1alpha3_BackendTLSPolicyList(in *v1alpha2.BackendTLSPolicyList, out *BackendTLSPolicyList, s conversion.Scope) error {
	return autoConvert_v1alpha2_BackendTLSPolicyList_To_v1alpha3_BackendTLSPolicyList(in, out, s)
}

func autoConvert_v1alpha3_BackendTLSPolicyList_To_v1alpha2_BackendTLSPolicyList(in *BackendTLSPolicyList, out *v1alpha2.BackendTLSPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha2.BackendTLSPolicy, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_BackendTLSPolicy_To_v1alpha2_BackendTLSPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_BackendTLSPolicyList_To_v1alpha2_BackendTLSPolicyList is an autogenerated conversion function.
func Convert_v1alpha3_BackendTLSPolicyList_To_v1alpha2_BackendTLSPolicyList(in *BackendTLSPolicyList, out *v1alpha2.BackendTLSPolicyList, s conversion.Scope) error {
	return autoConvert_v1alpha3_BackendTLSPolicyList_To_v1alpha2_BackendTLSPolicyList(in, out, s)
}

func Convert_v1alpha2_BackendTLSPolicySpec_To_v1alpha3_BackendTLSPolicySpec(in *v1alpha2.BackendTLSPolicySpec, out *BackendTLSPolicySpec, s conversion.Scope) error {
	out.TargetRefs = []v1alpha2.LocalPolicyTargetReferenceWithSectionName{
		{
			LocalPolicyTargetReference: v1alpha2.LocalPolicyTargetReference{
				Group: in.TargetRef.Group,
				Kind:  in.TargetRef.Kind,
				Name:  in.TargetRef.Name,
			},
			SectionName: in.TargetRef.SectionName,
		},
	}
	if err := Convert_v1alpha2_BackendTLSPolicyConfig_To_v1alpha3_BackendTLSPolicyValidation(&in.TLS, &out.Validation, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha2_BackendTLSPolicyConfig_To_v1alpha3_BackendTLSPolicyValidation(in *v1alpha2.BackendTLSPolicyConfig, out *BackendTLSPolicyValidation, s conversion.Scope) error {
	if in == nil {
		return nil
	}
	out.CACertificateRefs = in.CACertRefs
	out.WellKnownCACertificates = (*WellKnownCACertificatesType)(unsafe.Pointer(in.WellKnownCACerts))
	out.Hostname = in.Hostname
	return nil
}

func Convert_v1alpha3_BackendTLSPolicySpec_To_v1alpha2_BackendTLSPolicySpec(in *BackendTLSPolicySpec, out *v1alpha2.BackendTLSPolicySpec, s conversion.Scope) error {
	if len(in.TargetRefs) > 0 {
		targetRef := in.TargetRefs[0]
		out.TargetRef = v1alpha2.PolicyTargetReferenceWithSectionName{
			NamespacedPolicyTargetReference: v1alpha2.NamespacedPolicyTargetReference{
				Group:     targetRef.Group,
				Kind:      targetRef.Kind,
				Name:      targetRef.Name,
				Namespace: nil,
			},
			SectionName: targetRef.SectionName,
		}
	}
	if err := Convert_v1alpha3_BackendTLSPolicyValidation_To_v1alpha2_BackendTLSPolicyConfig(&in.Validation, &out.TLS, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha3_BackendTLSPolicyValidation_To_v1alpha2_BackendTLSPolicyConfig(in *BackendTLSPolicyValidation, out *v1alpha2.BackendTLSPolicyConfig, s conversion.Scope) error {
	if in == nil {
		return nil
	}
	out.CACertRefs = in.CACertificateRefs
	out.WellKnownCACerts = (*v1alpha2.WellKnownCACertType)(unsafe.Pointer(in.WellKnownCACertificates))
	out.Hostname = in.Hostname
	return nil
}
