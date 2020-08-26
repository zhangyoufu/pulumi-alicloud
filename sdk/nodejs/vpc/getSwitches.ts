// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of VSwitches owned by an Alibaba Cloud account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "vswitchDatasourceName";
 *
 * const defaultZones = pulumi.output(alicloud.getZones({ async: true }));
 * const vpc = new alicloud.vpc.Network("vpc", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const vswitch = new alicloud.vpc.Switch("vswitch", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: vpc.id,
 * });
 * const defaultSwitches = vswitch.name.apply(name => alicloud.vpc.getSwitches({
 *     nameRegex: name,
 * }, { async: true }));
 * ```
 */
export function getSwitches(args?: GetSwitchesArgs, opts?: pulumi.InvokeOptions): Promise<GetSwitchesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:vpc/getSwitches:getSwitches", {
        "cidrBlock": args.cidrBlock,
        "ids": args.ids,
        "isDefault": args.isDefault,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceGroupId": args.resourceGroupId,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSwitches.
 */
export interface GetSwitchesArgs {
    /**
     * Filter results by a specific CIDR block. For example: "172.16.0.0/12".
     */
    readonly cidrBlock?: string;
    /**
     * A list of VSwitch IDs.
     */
    readonly ids?: string[];
    /**
     * Indicate whether the VSwitch is created by the system.
     */
    readonly isDefault?: boolean;
    /**
     * A regex string to filter results by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The Id of resource group which VSWitch belongs.
     */
    readonly resourceGroupId?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: {[key: string]: any};
    /**
     * ID of the VPC that owns the VSwitch.
     */
    readonly vpcId?: string;
    /**
     * The availability zone of the VSwitch.
     */
    readonly zoneId?: string;
}

/**
 * A collection of values returned by getSwitches.
 */
export interface GetSwitchesResult {
    /**
     * CIDR block of the VSwitch.
     */
    readonly cidrBlock?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of VSwitch IDs.
     */
    readonly ids: string[];
    /**
     * Whether the VSwitch is the default one in the region.
     */
    readonly isDefault?: boolean;
    readonly nameRegex?: string;
    /**
     * A list of VSwitch names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly resourceGroupId?: string;
    readonly tags?: {[key: string]: any};
    /**
     * ID of the VPC that owns the VSwitch.
     */
    readonly vpcId?: string;
    /**
     * A list of VSwitches. Each element contains the following attributes:
     */
    readonly vswitches: outputs.vpc.GetSwitchesVswitch[];
    /**
     * ID of the availability zone where the VSwitch is located.
     */
    readonly zoneId?: string;
}
