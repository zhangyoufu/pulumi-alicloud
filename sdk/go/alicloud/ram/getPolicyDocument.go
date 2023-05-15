// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source Generates a RAM policy document of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.184.0+.
//
// ## Example Usage
// ### Basic Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			basicExample, err := ram.GetPolicyDocument(ctx, &ram.GetPolicyDocumentArgs{
//				Version: pulumi.StringRef("1"),
//				Statements: []ram.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"oss:*",
//						},
//						Resources: []string{
//							"acs:oss:*:*:myphotos",
//							"acs:oss:*:*:myphotos/*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewPolicy(ctx, "default", &ram.PolicyArgs{
//				PolicyName:     pulumi.String("tf-test"),
//				PolicyDocument: *pulumi.String(basicExample.Document),
//				Force:          pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// `data.alicloud_ram_policy_document.basic_example.document` will evaluate to:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// ### Example Multiple Condition Keys and Values
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			multipleCondition, err := ram.GetPolicyDocument(ctx, &ram.GetPolicyDocumentArgs{
//				Version: pulumi.StringRef("1"),
//				Statements: pulumi.Array{
//					ram.GetPolicyDocumentStatement{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"oss:ListBuckets",
//							"oss:GetBucketStat",
//							"oss:GetBucketInfo",
//							"oss:GetBucketTagging",
//							"oss:GetBucketAcl",
//						},
//						Resources: []string{
//							"acs:oss:*:*:*",
//						},
//					},
//					ram.GetPolicyDocumentStatement{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"oss:GetObject",
//							"oss:GetObjectAcl",
//						},
//						Resources: []string{
//							"acs:oss:*:*:myphotos/hangzhou/2015/*",
//						},
//					},
//					ram.GetPolicyDocumentStatement{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"oss:ListObjects",
//						},
//						Resources: []string{
//							"acs:oss:*:*:myphotos",
//						},
//						Conditions: []ram.GetPolicyDocumentStatementCondition{
//							{
//								Operator: "StringLike",
//								Variable: "oss:Delimiter",
//								Values: []string{
//									"/",
//								},
//							},
//							{
//								Operator: "StringLike",
//								Variable: "oss:Prefix",
//								Values: []string{
//									"",
//									"hangzhou/",
//									"hangzhou/2015/*",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewPolicy(ctx, "policy", &ram.PolicyArgs{
//				PolicyName:     pulumi.String("tf-test-condition"),
//				PolicyDocument: *pulumi.String(multipleCondition.Document),
//				Force:          pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// `data.alicloud_ram_policy_document.multiple_condition.document` will evaluate to:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// ### Example Assume-Role Policy with RAM Principal
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ramExample, err := ram.GetPolicyDocument(ctx, &ram.GetPolicyDocumentArgs{
//				Statements: []ram.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//						Principals: []ram.GetPolicyDocumentStatementPrincipal{
//							{
//								Entity: "RAM",
//								Identifiers: []string{
//									"acs:ram::123456789012****:root",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewRole(ctx, "role", &ram.RoleArgs{
//				Document: *pulumi.String(ramExample.Document),
//				Force:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// `data.alicloud_ram_policy_document.ram_example.document` will evaluate to:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// ### Example Assume-Role Policy with Service Principal
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			serviceExample, err := ram.GetPolicyDocument(ctx, &ram.GetPolicyDocumentArgs{
//				Statements: []ram.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//						Principals: []ram.GetPolicyDocumentStatementPrincipal{
//							{
//								Entity: "Service",
//								Identifiers: []string{
//									"ecs.aliyuncs.com",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewRole(ctx, "role", &ram.RoleArgs{
//				Document: *pulumi.String(serviceExample.Document),
//				Force:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// `data.alicloud_ram_policy_document.service_example.document` will evaluate to:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// ### Example Assume-Role Policy with Federated Principal
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			federatedExample, err := ram.GetPolicyDocument(ctx, &ram.GetPolicyDocumentArgs{
//				Statements: []ram.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//						Principals: []ram.GetPolicyDocumentStatementPrincipal{
//							{
//								Entity: "Federated",
//								Identifiers: []string{
//									"acs:ram::123456789012****:saml-provider/testprovider",
//								},
//							},
//						},
//						Conditions: []ram.GetPolicyDocumentStatementCondition{
//							{
//								Operator: "StringEquals",
//								Variable: "saml:recipient",
//								Values: []string{
//									"https://signin.aliyun.com/saml-role/sso",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewRole(ctx, "role", &ram.RoleArgs{
//				Document: *pulumi.String(federatedExample.Document),
//				Force:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// `data.alicloud_ram_policy_document.federated_example.document` will evaluate to:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
func GetPolicyDocument(ctx *pulumi.Context, args *GetPolicyDocumentArgs, opts ...pulumi.InvokeOption) (*GetPolicyDocumentResult, error) {
	var rv GetPolicyDocumentResult
	err := ctx.Invoke("alicloud:ram/getPolicyDocument:getPolicyDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyDocument.
type GetPolicyDocumentArgs struct {
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Statement of the RAM policy document. See the following `Block statement`.
	Statements []GetPolicyDocumentStatement `pulumi:"statements"`
	// Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getPolicyDocument.
type GetPolicyDocumentResult struct {
	// Standard policy document rendered based on the arguments above.
	Document string `pulumi:"document"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                       `pulumi:"id"`
	OutputFile *string                      `pulumi:"outputFile"`
	Statements []GetPolicyDocumentStatement `pulumi:"statements"`
	Version    *string                      `pulumi:"version"`
}

func GetPolicyDocumentOutput(ctx *pulumi.Context, args GetPolicyDocumentOutputArgs, opts ...pulumi.InvokeOption) GetPolicyDocumentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicyDocumentResult, error) {
			args := v.(GetPolicyDocumentArgs)
			r, err := GetPolicyDocument(ctx, &args, opts...)
			var s GetPolicyDocumentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicyDocumentResultOutput)
}

// A collection of arguments for invoking getPolicyDocument.
type GetPolicyDocumentOutputArgs struct {
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Statement of the RAM policy document. See the following `Block statement`.
	Statements GetPolicyDocumentStatementArrayInput `pulumi:"statements"`
	// Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (GetPolicyDocumentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentArgs)(nil)).Elem()
}

// A collection of values returned by getPolicyDocument.
type GetPolicyDocumentResultOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentResult)(nil)).Elem()
}

func (o GetPolicyDocumentResultOutput) ToGetPolicyDocumentResultOutput() GetPolicyDocumentResultOutput {
	return o
}

func (o GetPolicyDocumentResultOutput) ToGetPolicyDocumentResultOutputWithContext(ctx context.Context) GetPolicyDocumentResultOutput {
	return o
}

// Standard policy document rendered based on the arguments above.
func (o GetPolicyDocumentResultOutput) Document() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) string { return v.Document }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicyDocumentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPolicyDocumentResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetPolicyDocumentResultOutput) Statements() GetPolicyDocumentStatementArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) []GetPolicyDocumentStatement { return v.Statements }).(GetPolicyDocumentStatementArrayOutput)
}

func (o GetPolicyDocumentResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyDocumentResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyDocumentResultOutput{})
}
