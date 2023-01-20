// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides CEN Transit Router VPC Attachments available to the user.[What is Cen Transit Router VPC Attachments](https://help.aliyun.com/document_detail/261222.html)
 *
 * > **NOTE:** Available in 1.126.0+
 */
export function getTransitRouterVpcAttachments(args: GetTransitRouterVpcAttachmentsArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitRouterVpcAttachmentsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cen/getTransitRouterVpcAttachments:getTransitRouterVpcAttachments", {
        "cenId": args.cenId,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
        "transitRouterId": args.transitRouterId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitRouterVpcAttachments.
 */
export interface GetTransitRouterVpcAttachmentsArgs {
    /**
     * ID of the CEN instance.
     */
    cenId: string;
    /**
     * A list of resource id. The element value is same as `transitRouterId`.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The status of the resource. Valid values `Attached`, `Attaching` and `Detaching`.
     */
    status?: string;
    /**
     * The transit router ID.
     */
    transitRouterId?: string;
}

/**
 * A collection of values returned by getTransitRouterVpcAttachments.
 */
export interface GetTransitRouterVpcAttachmentsResult {
    /**
     * A list of CEN Transit Router VPC Attachments. Each element contains the following attributes:
     */
    readonly attachments: outputs.cen.GetTransitRouterVpcAttachmentsAttachment[];
    readonly cenId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * The status of the transit router attachment.
     */
    readonly status?: string;
    /**
     * ID of the transit router.
     */
    readonly transitRouterId?: string;
}
/**
 * This data source provides CEN Transit Router VPC Attachments available to the user.[What is Cen Transit Router VPC Attachments](https://help.aliyun.com/document_detail/261222.html)
 *
 * > **NOTE:** Available in 1.126.0+
 */
export function getTransitRouterVpcAttachmentsOutput(args: GetTransitRouterVpcAttachmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitRouterVpcAttachmentsResult> {
    return pulumi.output(args).apply((a: any) => getTransitRouterVpcAttachments(a, opts))
}

/**
 * A collection of arguments for invoking getTransitRouterVpcAttachments.
 */
export interface GetTransitRouterVpcAttachmentsOutputArgs {
    /**
     * ID of the CEN instance.
     */
    cenId: pulumi.Input<string>;
    /**
     * A list of resource id. The element value is same as `transitRouterId`.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values `Attached`, `Attaching` and `Detaching`.
     */
    status?: pulumi.Input<string>;
    /**
     * The transit router ID.
     */
    transitRouterId?: pulumi.Input<string>;
}
