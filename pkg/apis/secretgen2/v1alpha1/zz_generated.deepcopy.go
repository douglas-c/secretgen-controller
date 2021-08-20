// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretExport) DeepCopyInto(out *SecretExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretExport.
func (in *SecretExport) DeepCopy() *SecretExport {
	if in == nil {
		return nil
	}
	out := new(SecretExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretExportList) DeepCopyInto(out *SecretExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretExportList.
func (in *SecretExportList) DeepCopy() *SecretExportList {
	if in == nil {
		return nil
	}
	out := new(SecretExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretExportSpec) DeepCopyInto(out *SecretExportSpec) {
	*out = *in
	if in.ToNamespaces != nil {
		in, out := &in.ToNamespaces, &out.ToNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretExportSpec.
func (in *SecretExportSpec) DeepCopy() *SecretExportSpec {
	if in == nil {
		return nil
	}
	out := new(SecretExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretExportStatus) DeepCopyInto(out *SecretExportStatus) {
	*out = *in
	in.GenericStatus.DeepCopyInto(&out.GenericStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretExportStatus.
func (in *SecretExportStatus) DeepCopy() *SecretExportStatus {
	if in == nil {
		return nil
	}
	out := new(SecretExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretImport) DeepCopyInto(out *SecretImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretImport.
func (in *SecretImport) DeepCopy() *SecretImport {
	if in == nil {
		return nil
	}
	out := new(SecretImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretImportList) DeepCopyInto(out *SecretImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretImportList.
func (in *SecretImportList) DeepCopy() *SecretImportList {
	if in == nil {
		return nil
	}
	out := new(SecretImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretImportSpec) DeepCopyInto(out *SecretImportSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretImportSpec.
func (in *SecretImportSpec) DeepCopy() *SecretImportSpec {
	if in == nil {
		return nil
	}
	out := new(SecretImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretImportStatus) DeepCopyInto(out *SecretImportStatus) {
	*out = *in
	in.GenericStatus.DeepCopyInto(&out.GenericStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretImportStatus.
func (in *SecretImportStatus) DeepCopy() *SecretImportStatus {
	if in == nil {
		return nil
	}
	out := new(SecretImportStatus)
	in.DeepCopyInto(out)
	return out
}
