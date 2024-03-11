// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the acls in the region.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const sampleDs = alicloud.slb.getAcls({});
 * export const firstSlbAclId = sampleDs.then(sampleDs => sampleDs.acls?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Entry Block
 *
 * The entry mapping supports the following:
 *
 * * `entry`   - An IP addresses or CIDR blocks.
 * * `comment` - the comment of the entry.
 *
 * ## Listener Block
 *
 * The Listener mapping supports the following:
 *
 * * `loadBalancerId` - the id of load balancer instance, the listener belongs to.
 * * `frontendPort` - the listener port.
 * * `protocol`      - the listener protocol (such as tcp/udp/http/https, etc).
 * * `aclType`      - the type of acl (such as white/black).
 */
export function getAcls(args?: GetAclsArgs, opts?: pulumi.InvokeOptions): Promise<GetAclsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:slb/getAcls:getAcls", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getAcls.
 */
export interface GetAclsArgs {
    /**
     * A list of acls IDs to filter results.
     */
    ids?: string[];
    /**
     * A regex string to filter results by acl name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The Id of resource group which acl belongs.
     */
    resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getAcls.
 */
export interface GetAclsResult {
    /**
     * A list of SLB  acls. Each element contains the following attributes:
     */
    readonly acls: outputs.slb.GetAclsAcl[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of SLB acls IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of SLB acls names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * Resource group ID.
     */
    readonly resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
}
/**
 * This data source provides the acls in the region.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const sampleDs = alicloud.slb.getAcls({});
 * export const firstSlbAclId = sampleDs.then(sampleDs => sampleDs.acls?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Entry Block
 *
 * The entry mapping supports the following:
 *
 * * `entry`   - An IP addresses or CIDR blocks.
 * * `comment` - the comment of the entry.
 *
 * ## Listener Block
 *
 * The Listener mapping supports the following:
 *
 * * `loadBalancerId` - the id of load balancer instance, the listener belongs to.
 * * `frontendPort` - the listener port.
 * * `protocol`      - the listener protocol (such as tcp/udp/http/https, etc).
 * * `aclType`      - the type of acl (such as white/black).
 */
export function getAclsOutput(args?: GetAclsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAclsResult> {
    return pulumi.output(args).apply((a: any) => getAcls(a, opts))
}

/**
 * A collection of arguments for invoking getAcls.
 */
export interface GetAclsOutputArgs {
    /**
     * A list of acls IDs to filter results.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by acl name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The Id of resource group which acl belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
