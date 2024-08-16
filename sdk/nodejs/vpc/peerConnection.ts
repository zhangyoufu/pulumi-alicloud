// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Peer Connection resource.
 *
 * For information about VPC Peer Connection and how to use it, see [What is Peer Connection](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createvpcpeer).
 *
 * > **NOTE:** Available since v1.186.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.getAccount({});
 * const config = new pulumi.Config();
 * const acceptingRegion = config.get("acceptingRegion") || "cn-beijing";
 * const localVpc = new alicloud.vpc.Network("local_vpc", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const acceptingVpc = new alicloud.vpc.Network("accepting_vpc", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const defaultPeerConnection = new alicloud.vpc.PeerConnection("default", {
 *     peerConnectionName: "terraform-example",
 *     vpcId: localVpc.id,
 *     acceptingAliUid: _default.then(_default => _default.id),
 *     acceptingRegionId: acceptingRegion,
 *     acceptingVpcId: acceptingVpc.id,
 *     description: "terraform-example",
 * });
 * ```
 *
 * ## Import
 *
 * VPC Peer Connection can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/peerConnection:PeerConnection example <id>
 * ```
 */
export class PeerConnection extends pulumi.CustomResource {
    /**
     * Get an existing PeerConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PeerConnectionState, opts?: pulumi.CustomResourceOptions): PeerConnection {
        return new PeerConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/peerConnection:PeerConnection';

    /**
     * Returns true if the given object is an instance of PeerConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PeerConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PeerConnection.__pulumiType;
    }

    /**
     * The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
     * - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
     * - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
     * > **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
     */
    public readonly acceptingAliUid!: pulumi.Output<number | undefined>;
    /**
     * The region ID of the recipient of the VPC peering connection to be created.
     * - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
     * - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
     */
    public readonly acceptingRegionId!: pulumi.Output<string>;
    /**
     * The VPC ID of the receiving end of the VPC peer connection.
     */
    public readonly acceptingVpcId!: pulumi.Output<string>;
    /**
     * The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The creation time of the VPC peer connection. Use UTC time in the format `YYYY-MM-DDThh:mm:ssZ`.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to PreCheck only this request. Default value: `false`. Valid values:
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
     */
    public readonly peerConnectionName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The status of the VPC peer connection.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the requester VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a PeerConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PeerConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PeerConnectionArgs | PeerConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PeerConnectionState | undefined;
            resourceInputs["acceptingAliUid"] = state ? state.acceptingAliUid : undefined;
            resourceInputs["acceptingRegionId"] = state ? state.acceptingRegionId : undefined;
            resourceInputs["acceptingVpcId"] = state ? state.acceptingVpcId : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["peerConnectionName"] = state ? state.peerConnectionName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as PeerConnectionArgs | undefined;
            if ((!args || args.acceptingRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceptingRegionId'");
            }
            if ((!args || args.acceptingVpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceptingVpcId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["acceptingAliUid"] = args ? args.acceptingAliUid : undefined;
            resourceInputs["acceptingRegionId"] = args ? args.acceptingRegionId : undefined;
            resourceInputs["acceptingVpcId"] = args ? args.acceptingVpcId : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["peerConnectionName"] = args ? args.peerConnectionName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PeerConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PeerConnection resources.
 */
export interface PeerConnectionState {
    /**
     * The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
     * - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
     * - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
     * > **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
     */
    acceptingAliUid?: pulumi.Input<number>;
    /**
     * The region ID of the recipient of the VPC peering connection to be created.
     * - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
     * - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
     */
    acceptingRegionId?: pulumi.Input<string>;
    /**
     * The VPC ID of the receiving end of the VPC peer connection.
     */
    acceptingVpcId?: pulumi.Input<string>;
    /**
     * The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The creation time of the VPC peer connection. Use UTC time in the format `YYYY-MM-DDThh:mm:ssZ`.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to PreCheck only this request. Default value: `false`. Valid values:
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
     */
    peerConnectionName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the VPC peer connection.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the requester VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PeerConnection resource.
 */
export interface PeerConnectionArgs {
    /**
     * The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
     * - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
     * - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
     * > **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
     */
    acceptingAliUid?: pulumi.Input<number>;
    /**
     * The region ID of the recipient of the VPC peering connection to be created.
     * - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
     * - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
     */
    acceptingRegionId: pulumi.Input<string>;
    /**
     * The VPC ID of the receiving end of the VPC peer connection.
     */
    acceptingVpcId: pulumi.Input<string>;
    /**
     * The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to PreCheck only this request. Default value: `false`. Valid values:
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
     */
    peerConnectionName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the VPC peer connection.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the requester VPC.
     */
    vpcId: pulumi.Input<string>;
}
