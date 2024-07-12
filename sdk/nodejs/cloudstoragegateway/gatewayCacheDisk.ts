// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Storage Gateway Gateway Cache Disk resource.
 *
 * For information about Cloud Storage Gateway Gateway Cache Disk and how to use it, see [What is Gateway Cache Disk](https://www.alibabacloud.com/help/en/cloud-storage-gateway/latest/creategatewaycachedisk).
 *
 * > **NOTE:** Available since v1.144.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const default = alicloud.cloudstoragegateway.getStocks({
 *     gatewayClass: "Standard",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: _default.then(_default => _default.stocks?.[0]?.zoneId),
 *     vswitchName: name,
 * });
 * const defaultStorageBundle = new alicloud.cloudstoragegateway.StorageBundle("default", {storageBundleName: name});
 * const defaultGateway = new alicloud.cloudstoragegateway.Gateway("default", {
 *     description: name,
 *     gatewayClass: "Standard",
 *     type: "File",
 *     paymentType: "PayAsYouGo",
 *     vswitchId: defaultSwitch.id,
 *     releaseAfterExpiration: true,
 *     storageBundleId: defaultStorageBundle.id,
 *     location: "Cloud",
 *     gatewayName: name,
 * });
 * const defaultGatewayCacheDisk = new alicloud.cloudstoragegateway.GatewayCacheDisk("default", {
 *     gatewayId: defaultGateway.id,
 *     cacheDiskSizeInGb: 50,
 *     cacheDiskCategory: "cloud_efficiency",
 * });
 * ```
 *
 * ## Import
 *
 * Cloud Storage Gateway Gateway Cache Disk can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk example <gateway_id>:<cache_id>:<local_file_path>
 * ```
 */
export class GatewayCacheDisk extends pulumi.CustomResource {
    /**
     * Get an existing GatewayCacheDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayCacheDiskState, opts?: pulumi.CustomResourceOptions): GatewayCacheDisk {
        return new GatewayCacheDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudstoragegateway/gatewayCacheDisk:GatewayCacheDisk';

    /**
     * Returns true if the given object is an instance of GatewayCacheDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayCacheDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayCacheDisk.__pulumiType;
    }

    /**
     * The type of the cache disk. Valid values: `cloudEfficiency`, `cloudSsd`, `cloudEssd`. **NOTE:** From version 1.227.0, `cacheDiskCategory` can be set to `cloudEssd`.
     */
    public readonly cacheDiskCategory!: pulumi.Output<string>;
    /**
     * The capacity of the cache disk.
     */
    public readonly cacheDiskSizeInGb!: pulumi.Output<number>;
    /**
     * The ID of the cache disk.
     */
    public /*out*/ readonly cacheId!: pulumi.Output<string>;
    /**
     * The ID of the gateway.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * The path of the cache disk.
     */
    public /*out*/ readonly localFilePath!: pulumi.Output<string>;
    /**
     * The performance level (PL) of the Enterprise SSD (ESSD). Valid values: `PL1`, `PL2`, `PL3`. **NOTE:** If `cacheDiskCategory` is set to `cloudEssd`, `performanceLevel` is required.
     */
    public readonly performanceLevel!: pulumi.Output<string | undefined>;
    /**
     * The status of the Gateway Cache Disk.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;

    /**
     * Create a GatewayCacheDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayCacheDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayCacheDiskArgs | GatewayCacheDiskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayCacheDiskState | undefined;
            resourceInputs["cacheDiskCategory"] = state ? state.cacheDiskCategory : undefined;
            resourceInputs["cacheDiskSizeInGb"] = state ? state.cacheDiskSizeInGb : undefined;
            resourceInputs["cacheId"] = state ? state.cacheId : undefined;
            resourceInputs["gatewayId"] = state ? state.gatewayId : undefined;
            resourceInputs["localFilePath"] = state ? state.localFilePath : undefined;
            resourceInputs["performanceLevel"] = state ? state.performanceLevel : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as GatewayCacheDiskArgs | undefined;
            if ((!args || args.cacheDiskSizeInGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cacheDiskSizeInGb'");
            }
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            resourceInputs["cacheDiskCategory"] = args ? args.cacheDiskCategory : undefined;
            resourceInputs["cacheDiskSizeInGb"] = args ? args.cacheDiskSizeInGb : undefined;
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["performanceLevel"] = args ? args.performanceLevel : undefined;
            resourceInputs["cacheId"] = undefined /*out*/;
            resourceInputs["localFilePath"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GatewayCacheDisk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayCacheDisk resources.
 */
export interface GatewayCacheDiskState {
    /**
     * The type of the cache disk. Valid values: `cloudEfficiency`, `cloudSsd`, `cloudEssd`. **NOTE:** From version 1.227.0, `cacheDiskCategory` can be set to `cloudEssd`.
     */
    cacheDiskCategory?: pulumi.Input<string>;
    /**
     * The capacity of the cache disk.
     */
    cacheDiskSizeInGb?: pulumi.Input<number>;
    /**
     * The ID of the cache disk.
     */
    cacheId?: pulumi.Input<string>;
    /**
     * The ID of the gateway.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The path of the cache disk.
     */
    localFilePath?: pulumi.Input<string>;
    /**
     * The performance level (PL) of the Enterprise SSD (ESSD). Valid values: `PL1`, `PL2`, `PL3`. **NOTE:** If `cacheDiskCategory` is set to `cloudEssd`, `performanceLevel` is required.
     */
    performanceLevel?: pulumi.Input<string>;
    /**
     * The status of the Gateway Cache Disk.
     */
    status?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GatewayCacheDisk resource.
 */
export interface GatewayCacheDiskArgs {
    /**
     * The type of the cache disk. Valid values: `cloudEfficiency`, `cloudSsd`, `cloudEssd`. **NOTE:** From version 1.227.0, `cacheDiskCategory` can be set to `cloudEssd`.
     */
    cacheDiskCategory?: pulumi.Input<string>;
    /**
     * The capacity of the cache disk.
     */
    cacheDiskSizeInGb: pulumi.Input<number>;
    /**
     * The ID of the gateway.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * The performance level (PL) of the Enterprise SSD (ESSD). Valid values: `PL1`, `PL2`, `PL3`. **NOTE:** If `cacheDiskCategory` is set to `cloudEssd`, `performanceLevel` is required.
     */
    performanceLevel?: pulumi.Input<string>;
}
