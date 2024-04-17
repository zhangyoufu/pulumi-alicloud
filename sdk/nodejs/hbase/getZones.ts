// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
 *
 * > **NOTE:** Available in v1.73.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const zonesIds = alicloud.hbase.getZones({});
 * const hbase = new alicloud.hbase.Instance("hbase", {zoneId: zonesIds.then(zonesIds => zonesIds.zones?.[0]?.id)});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getZones(args?: GetZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetZonesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:hbase/getZones:getZones", {
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesArgs {
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getZones.
 */
export interface GetZonesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of zone IDs.
     */
    readonly ids: string[];
    readonly outputFile?: string;
    /**
     * A list of availability zones. Each element contains the following attributes:
     */
    readonly zones: outputs.hbase.GetZonesZone[];
}
/**
 * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
 *
 * > **NOTE:** Available in v1.73.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const zonesIds = alicloud.hbase.getZones({});
 * const hbase = new alicloud.hbase.Instance("hbase", {zoneId: zonesIds.then(zonesIds => zonesIds.zones?.[0]?.id)});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getZonesOutput(args?: GetZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZonesResult> {
    return pulumi.output(args).apply((a: any) => getZones(a, opts))
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesOutputArgs {
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
