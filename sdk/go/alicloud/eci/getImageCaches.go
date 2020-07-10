// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eci

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a collection of ECI Image Cache to the specified filters.
//
// > **NOTE:** Available in 1.90.0+.
func GetImageCaches(ctx *pulumi.Context, args *GetImageCachesArgs, opts ...pulumi.InvokeOption) (*GetImageCachesResult, error) {
	var rv GetImageCachesResult
	err := ctx.Invoke("alicloud:eci/getImageCaches:getImageCaches", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageCaches.
type GetImageCachesArgs struct {
	// A list ids of ECI Image Cache.
	Ids []string `pulumi:"ids"`
	// Find the mirror cache containing it according to the image name.
	Image *string `pulumi:"image"`
	// The name of ECI Image Cache.
	ImageCacheName *string `pulumi:"imageCacheName"`
	// A regex string to filter results by the image cache name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The id of snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// The status of ECI Image Cache.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getImageCaches.
type GetImageCachesResult struct {
	// A list of caches. Each element contains the following attributes:
	Caches []GetImageCachesCach `pulumi:"caches"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list ids of ECI Image Cache.
	Ids   []string `pulumi:"ids"`
	Image *string  `pulumi:"image"`
	// The name of the ECI Image Cache.
	ImageCacheName *string `pulumi:"imageCacheName"`
	NameRegex      *string `pulumi:"nameRegex"`
	// A list of ECI Image Cache names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The id of snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// The status of ECI Image Cache.
	Status *string `pulumi:"status"`
}
