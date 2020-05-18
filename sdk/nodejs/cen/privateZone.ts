// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This topic describes how to configure PrivateZone access. 
 * PrivateZone is a VPC-based resolution and management service for private domain names. 
 * After you set a PrivateZone access, the Cloud Connect Network (CCN) and Virtual Border Router (VBR) attached to a CEN instance can access the PrivateZone service through CEN.
 * 
 * For information about CEN Private Zone and how to use it, see [Manage CEN Private Zone](https://www.alibabacloud.com/help/en/doc-detail/106693.htm).
 * 
 * > **NOTE:** Available in 1.83.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const defaultInstance = new alicloud.cen.Instance("default", {});
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const defaultInstanceAttachment = new alicloud.cen.InstanceAttachment("default", {
 *     childInstanceId: defaultNetwork.id,
 *     childInstanceRegionId: "cn-hangzhou",
 *     instanceId: defaultInstance.id,
 * }, { dependsOn: [defaultInstance, defaultNetwork] });
 * const defaultPrivateZone = new alicloud.cen.PrivateZone("default", {
 *     accessRegionId: "cn-hangzhou",
 *     cenId: defaultInstance.id,
 *     hostRegionId: "cn-hangzhou",
 *     hostVpcId: defaultNetwork.id,
 * }, { dependsOn: [defaultInstanceAttachment] });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_private_zone.html.markdown.
 */
export class PrivateZone extends pulumi.CustomResource {
    /**
     * Get an existing PrivateZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateZoneState, opts?: pulumi.CustomResourceOptions): PrivateZone {
        return new PrivateZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/privateZone:PrivateZone';

    /**
     * Returns true if the given object is an instance of PrivateZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateZone.__pulumiType;
    }

    /**
     * The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
     */
    public readonly accessRegionId!: pulumi.Output<string>;
    /**
     * The ID of the CEN instance.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The service region. The service region is the target region of the PrivateZone service to be accessed through CEN. 
     */
    public readonly hostRegionId!: pulumi.Output<string>;
    /**
     * The VPC that belongs to the service region.
     */
    public readonly hostVpcId!: pulumi.Output<string>;
    /**
     * The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"].
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a PrivateZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateZoneArgs | PrivateZoneState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PrivateZoneState | undefined;
            inputs["accessRegionId"] = state ? state.accessRegionId : undefined;
            inputs["cenId"] = state ? state.cenId : undefined;
            inputs["hostRegionId"] = state ? state.hostRegionId : undefined;
            inputs["hostVpcId"] = state ? state.hostVpcId : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as PrivateZoneArgs | undefined;
            if (!args || args.accessRegionId === undefined) {
                throw new Error("Missing required property 'accessRegionId'");
            }
            if (!args || args.cenId === undefined) {
                throw new Error("Missing required property 'cenId'");
            }
            if (!args || args.hostRegionId === undefined) {
                throw new Error("Missing required property 'hostRegionId'");
            }
            if (!args || args.hostVpcId === undefined) {
                throw new Error("Missing required property 'hostVpcId'");
            }
            inputs["accessRegionId"] = args ? args.accessRegionId : undefined;
            inputs["cenId"] = args ? args.cenId : undefined;
            inputs["hostRegionId"] = args ? args.hostRegionId : undefined;
            inputs["hostVpcId"] = args ? args.hostVpcId : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PrivateZone.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateZone resources.
 */
export interface PrivateZoneState {
    /**
     * The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
     */
    readonly accessRegionId?: pulumi.Input<string>;
    /**
     * The ID of the CEN instance.
     */
    readonly cenId?: pulumi.Input<string>;
    /**
     * The service region. The service region is the target region of the PrivateZone service to be accessed through CEN. 
     */
    readonly hostRegionId?: pulumi.Input<string>;
    /**
     * The VPC that belongs to the service region.
     */
    readonly hostVpcId?: pulumi.Input<string>;
    /**
     * The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"].
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateZone resource.
 */
export interface PrivateZoneArgs {
    /**
     * The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.
     */
    readonly accessRegionId: pulumi.Input<string>;
    /**
     * The ID of the CEN instance.
     */
    readonly cenId: pulumi.Input<string>;
    /**
     * The service region. The service region is the target region of the PrivateZone service to be accessed through CEN. 
     */
    readonly hostRegionId: pulumi.Input<string>;
    /**
     * The VPC that belongs to the service region.
     */
    readonly hostVpcId: pulumi.Input<string>;
}
