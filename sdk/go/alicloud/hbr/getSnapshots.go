// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Hbr Snapshots of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.133.0+.
func GetSnapshots(ctx *pulumi.Context, args *GetSnapshotsArgs, opts ...pulumi.InvokeOption) (*GetSnapshotsResult, error) {
	var rv GetSnapshotsResult
	err := ctx.Invoke("alicloud:hbr/getSnapshots:getSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshots.
type GetSnapshotsArgs struct {
	// The name of OSS bucket.
	Bucket *string `pulumi:"bucket"`
	// The time when the snapshot was completed. UNIX time in seconds.
	CompleteTime *string `pulumi:"completeTime"`
	// Complete time filter operator. Optional values: `MATCH_TERM`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `BETWEEN`.
	CompleteTimeChecker *string `pulumi:"completeTimeChecker"`
	// File System Creation Time of Nas. Unix Time Seconds.
	CreateTime *string `pulumi:"createTime"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// The ID of NAS File system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// A list of Snapshot IDs.
	Ids []string `pulumi:"ids"`
	// The ID of ECS instance.
	InstanceId *string `pulumi:"instanceId"`
	Limit      *int    `pulumi:"limit"`
	OutputFile *string `pulumi:"outputFile"`
	Query      *string `pulumi:"query"`
	// Data source type, optional values: `ECS_FILE`, `OSS`, `NAS`.
	SourceType string `pulumi:"sourceType"`
	// The status of snapshot execution. Possible values: `COMPLETE`, `PARTIAL_COMPLETE`, `FAILED`.
	Status *string `pulumi:"status"`
	// The ID of Vault.
	VaultId string `pulumi:"vaultId"`
}

// A collection of values returned by getSnapshots.
type GetSnapshotsResult struct {
	Bucket              *string `pulumi:"bucket"`
	CompleteTime        *string `pulumi:"completeTime"`
	CompleteTimeChecker *string `pulumi:"completeTimeChecker"`
	CreateTime          *string `pulumi:"createTime"`
	EnableDetails       *bool   `pulumi:"enableDetails"`
	FileSystemId        *string `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                 `pulumi:"id"`
	Ids        []string               `pulumi:"ids"`
	InstanceId *string                `pulumi:"instanceId"`
	Limit      *int                   `pulumi:"limit"`
	OutputFile *string                `pulumi:"outputFile"`
	Query      *string                `pulumi:"query"`
	Snapshots  []GetSnapshotsSnapshot `pulumi:"snapshots"`
	SourceType string                 `pulumi:"sourceType"`
	Status     *string                `pulumi:"status"`
	VaultId    string                 `pulumi:"vaultId"`
}
