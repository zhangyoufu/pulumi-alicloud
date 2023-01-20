// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Firewall Vpc Firewall resource.
 *
 * For information about Cloud Firewall Vpc Firewall and how to use it, see [What is Vpc Firewall](https://help.aliyun.com/document_detail/342893.html).
 *
 * > **NOTE:** Available in v1.194.0+.
 *
 * ## Import
 *
 * Cloud Firewall Vpc Firewall can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cloudfirewall/firewallVpcFirewall:FirewallVpcFirewall example <id>
 * ```
 */
export class FirewallVpcFirewall extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVpcFirewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVpcFirewallState, opts?: pulumi.CustomResourceOptions): FirewallVpcFirewall {
        return new FirewallVpcFirewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudfirewall/firewallVpcFirewall:FirewallVpcFirewall';

    /**
     * Returns true if the given object is an instance of FirewallVpcFirewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVpcFirewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVpcFirewall.__pulumiType;
    }

    /**
     * Bandwidth specifications for high-speed channels. Unit: Mbps.
     */
    public /*out*/ readonly bandwidth!: pulumi.Output<number>;
    /**
     * The communication type of the VPC firewall. Valid value: **expressconnect**, which indicates Express Connect.
     */
    public /*out*/ readonly connectType!: pulumi.Output<string>;
    /**
     * The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
     */
    public readonly lang!: pulumi.Output<string>;
    /**
     * The details of the local VPC. See the following `Block LocalVpc`.
     */
    public readonly localVpc!: pulumi.Output<outputs.cloudfirewall.FirewallVpcFirewallLocalVpc>;
    /**
     * The UID of the Alibaba Cloud member account.
     */
    public readonly memberUid!: pulumi.Output<string | undefined>;
    /**
     * The details of the peer VPC. See the following `Block PeerVpc`.
     */
    public readonly peerVpc!: pulumi.Output<outputs.cloudfirewall.FirewallVpcFirewallPeerVpc>;
    /**
     * The region is open. Value:-**enable**: is enabled, indicating that VPC firewall can be configured in this region.-**disable**: indicates that VPC firewall cannot be configured in this region.
     */
    public /*out*/ readonly regionStatus!: pulumi.Output<string>;
    /**
     * The status of the resource
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The ID of the VPC firewall instance.
     */
    public /*out*/ readonly vpcFirewallId!: pulumi.Output<string>;
    /**
     * The name of the VPC firewall instance.
     */
    public readonly vpcFirewallName!: pulumi.Output<string>;

    /**
     * Create a FirewallVpcFirewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallVpcFirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVpcFirewallArgs | FirewallVpcFirewallState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVpcFirewallState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["connectType"] = state ? state.connectType : undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
            resourceInputs["localVpc"] = state ? state.localVpc : undefined;
            resourceInputs["memberUid"] = state ? state.memberUid : undefined;
            resourceInputs["peerVpc"] = state ? state.peerVpc : undefined;
            resourceInputs["regionStatus"] = state ? state.regionStatus : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpcFirewallId"] = state ? state.vpcFirewallId : undefined;
            resourceInputs["vpcFirewallName"] = state ? state.vpcFirewallName : undefined;
        } else {
            const args = argsOrState as FirewallVpcFirewallArgs | undefined;
            if ((!args || args.localVpc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localVpc'");
            }
            if ((!args || args.peerVpc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerVpc'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.vpcFirewallName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcFirewallName'");
            }
            resourceInputs["lang"] = args ? args.lang : undefined;
            resourceInputs["localVpc"] = args ? args.localVpc : undefined;
            resourceInputs["memberUid"] = args ? args.memberUid : undefined;
            resourceInputs["peerVpc"] = args ? args.peerVpc : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vpcFirewallName"] = args ? args.vpcFirewallName : undefined;
            resourceInputs["bandwidth"] = undefined /*out*/;
            resourceInputs["connectType"] = undefined /*out*/;
            resourceInputs["regionStatus"] = undefined /*out*/;
            resourceInputs["vpcFirewallId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallVpcFirewall.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVpcFirewall resources.
 */
export interface FirewallVpcFirewallState {
    /**
     * Bandwidth specifications for high-speed channels. Unit: Mbps.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The communication type of the VPC firewall. Valid value: **expressconnect**, which indicates Express Connect.
     */
    connectType?: pulumi.Input<string>;
    /**
     * The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
     */
    lang?: pulumi.Input<string>;
    /**
     * The details of the local VPC. See the following `Block LocalVpc`.
     */
    localVpc?: pulumi.Input<inputs.cloudfirewall.FirewallVpcFirewallLocalVpc>;
    /**
     * The UID of the Alibaba Cloud member account.
     */
    memberUid?: pulumi.Input<string>;
    /**
     * The details of the peer VPC. See the following `Block PeerVpc`.
     */
    peerVpc?: pulumi.Input<inputs.cloudfirewall.FirewallVpcFirewallPeerVpc>;
    /**
     * The region is open. Value:-**enable**: is enabled, indicating that VPC firewall can be configured in this region.-**disable**: indicates that VPC firewall cannot be configured in this region.
     */
    regionStatus?: pulumi.Input<string>;
    /**
     * The status of the resource
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the VPC firewall instance.
     */
    vpcFirewallId?: pulumi.Input<string>;
    /**
     * The name of the VPC firewall instance.
     */
    vpcFirewallName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallVpcFirewall resource.
 */
export interface FirewallVpcFirewallArgs {
    /**
     * The language type of the requested and received messages. Value:**zh** (default): Chinese.**en**: English.
     */
    lang?: pulumi.Input<string>;
    /**
     * The details of the local VPC. See the following `Block LocalVpc`.
     */
    localVpc: pulumi.Input<inputs.cloudfirewall.FirewallVpcFirewallLocalVpc>;
    /**
     * The UID of the Alibaba Cloud member account.
     */
    memberUid?: pulumi.Input<string>;
    /**
     * The details of the peer VPC. See the following `Block PeerVpc`.
     */
    peerVpc: pulumi.Input<inputs.cloudfirewall.FirewallVpcFirewallPeerVpc>;
    /**
     * The status of the resource
     */
    status: pulumi.Input<string>;
    /**
     * The name of the VPC firewall instance.
     */
    vpcFirewallName: pulumi.Input<string>;
}
