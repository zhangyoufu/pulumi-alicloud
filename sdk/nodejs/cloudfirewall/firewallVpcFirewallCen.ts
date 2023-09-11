// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Firewall Vpc Firewall Cen resource.
 *
 * For information about Cloud Firewall Vpc Firewall Cen and how to use it, see [What is Vpc Firewall Cen](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure).
 *
 * > **NOTE:** Available since v1.194.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.cloudfirewall.FirewallVpcFirewallCen("default", {
 *     cenId: "cen-cjok7uyb5w2b27573v",
 *     localVpc: {
 *         networkInstanceId: "vpc-a2d4wzzfuumzuq6uog5w4",
 *     },
 *     memberUid: "1415189284827022",
 *     status: "open",
 *     vpcFirewallName: "tf-vpc-firewall-name",
 *     vpcRegion: "ap-south-1",
 * });
 * ```
 *
 * ## Import
 *
 * Cloud Firewall Vpc Firewall Cen can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen example <id>
 * ```
 */
export class FirewallVpcFirewallCen extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVpcFirewallCen resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVpcFirewallCenState, opts?: pulumi.CustomResourceOptions): FirewallVpcFirewallCen {
        return new FirewallVpcFirewallCen(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen';

    /**
     * Returns true if the given object is an instance of FirewallVpcFirewallCen.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVpcFirewallCen {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVpcFirewallCen.__pulumiType;
    }

    /**
     * The ID of the CEN instance.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * Intercommunication type, value: expressconnect: Express Channel cen: Cloud Enterprise Network
     */
    public /*out*/ readonly connectType!: pulumi.Output<string>;
    /**
     * The language type of the requested and received messages. Valid values:
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * The details of the VPC. See `localVpc` below.
     */
    public readonly localVpc!: pulumi.Output<outputs.cloudfirewall.FirewallVpcFirewallCenLocalVpc>;
    /**
     * The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
     */
    public readonly memberUid!: pulumi.Output<string | undefined>;
    /**
     * Firewall switch status.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * VPC firewall ID
     */
    public /*out*/ readonly vpcFirewallId!: pulumi.Output<string>;
    /**
     * The name of the VPC firewall instance.
     */
    public readonly vpcFirewallName!: pulumi.Output<string>;
    /**
     * The ID of the region to which the VPC is created.
     */
    public readonly vpcRegion!: pulumi.Output<string>;

    /**
     * Create a FirewallVpcFirewallCen resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallVpcFirewallCenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVpcFirewallCenArgs | FirewallVpcFirewallCenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVpcFirewallCenState | undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["connectType"] = state ? state.connectType : undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
            resourceInputs["localVpc"] = state ? state.localVpc : undefined;
            resourceInputs["memberUid"] = state ? state.memberUid : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpcFirewallId"] = state ? state.vpcFirewallId : undefined;
            resourceInputs["vpcFirewallName"] = state ? state.vpcFirewallName : undefined;
            resourceInputs["vpcRegion"] = state ? state.vpcRegion : undefined;
        } else {
            const args = argsOrState as FirewallVpcFirewallCenArgs | undefined;
            if ((!args || args.cenId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cenId'");
            }
            if ((!args || args.localVpc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localVpc'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.vpcFirewallName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcFirewallName'");
            }
            if ((!args || args.vpcRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcRegion'");
            }
            resourceInputs["cenId"] = args ? args.cenId : undefined;
            resourceInputs["lang"] = args ? args.lang : undefined;
            resourceInputs["localVpc"] = args ? args.localVpc : undefined;
            resourceInputs["memberUid"] = args ? args.memberUid : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vpcFirewallName"] = args ? args.vpcFirewallName : undefined;
            resourceInputs["vpcRegion"] = args ? args.vpcRegion : undefined;
            resourceInputs["connectType"] = undefined /*out*/;
            resourceInputs["vpcFirewallId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallVpcFirewallCen.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVpcFirewallCen resources.
 */
export interface FirewallVpcFirewallCenState {
    /**
     * The ID of the CEN instance.
     */
    cenId?: pulumi.Input<string>;
    /**
     * Intercommunication type, value: expressconnect: Express Channel cen: Cloud Enterprise Network
     */
    connectType?: pulumi.Input<string>;
    /**
     * The language type of the requested and received messages. Valid values:
     */
    lang?: pulumi.Input<string>;
    /**
     * The details of the VPC. See `localVpc` below.
     */
    localVpc?: pulumi.Input<inputs.cloudfirewall.FirewallVpcFirewallCenLocalVpc>;
    /**
     * The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
     */
    memberUid?: pulumi.Input<string>;
    /**
     * Firewall switch status.
     */
    status?: pulumi.Input<string>;
    /**
     * VPC firewall ID
     */
    vpcFirewallId?: pulumi.Input<string>;
    /**
     * The name of the VPC firewall instance.
     */
    vpcFirewallName?: pulumi.Input<string>;
    /**
     * The ID of the region to which the VPC is created.
     */
    vpcRegion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallVpcFirewallCen resource.
 */
export interface FirewallVpcFirewallCenArgs {
    /**
     * The ID of the CEN instance.
     */
    cenId: pulumi.Input<string>;
    /**
     * The language type of the requested and received messages. Valid values:
     */
    lang?: pulumi.Input<string>;
    /**
     * The details of the VPC. See `localVpc` below.
     */
    localVpc: pulumi.Input<inputs.cloudfirewall.FirewallVpcFirewallCenLocalVpc>;
    /**
     * The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
     */
    memberUid?: pulumi.Input<string>;
    /**
     * Firewall switch status.
     */
    status: pulumi.Input<string>;
    /**
     * The name of the VPC firewall instance.
     */
    vpcFirewallName: pulumi.Input<string>;
    /**
     * The ID of the region to which the VPC is created.
     */
    vpcRegion: pulumi.Input<string>;
}
