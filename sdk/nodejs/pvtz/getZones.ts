// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source lists a number of Private Zones resource information owned by an Alibaba Cloud account.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const pvtzZonesDs = alicloud_pvtz_zone_basic.zoneName.apply(zoneName => alicloud.pvtz.getZones({
 *     keyword: zoneName,
 * }, { async: true }));
 * 
 * export const firstZoneId = pvtzZonesDs.zones[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/pvtz_zones.html.markdown.
 */
export function getZones(args?: GetZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:pvtz/getZones:getZones", {
        "ids": args.ids,
        "keyword": args.keyword,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getZones.
 */
export interface GetZonesArgs {
    /**
     * A list of zone IDs. 
     */
    readonly ids?: string[];
    /**
     * keyword for zone name.
     */
    readonly keyword?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getZones.
 */
export interface GetZonesResult {
    /**
     * A list of zone IDs. 
     */
    readonly ids: string[];
    readonly keyword?: string;
    /**
     * A list of zone names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of zones. Each element contains the following attributes:
     */
    readonly zones: outputs.pvtz.GetZonesZone[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
