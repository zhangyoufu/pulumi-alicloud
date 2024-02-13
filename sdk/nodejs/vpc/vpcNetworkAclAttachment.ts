// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Network Acl Attachment resource. Resources associated with network Acl.
 *
 * For information about VPC Network Acl Attachment and how to use it, see [What is Network Acl Attachment](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/associatenetworkacl).
 *
 * > **NOTE:** Available since v1.193.0.
 *
 * ## Import
 *
 * VPC Network Acl Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/vpcNetworkAclAttachment:VpcNetworkAclAttachment example <network_acl_id>:<resource_id>
 * ```
 */
export class VpcNetworkAclAttachment extends pulumi.CustomResource {
    /**
     * Get an existing VpcNetworkAclAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcNetworkAclAttachmentState, opts?: pulumi.CustomResourceOptions): VpcNetworkAclAttachment {
        return new VpcNetworkAclAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/vpcNetworkAclAttachment:VpcNetworkAclAttachment';

    /**
     * Returns true if the given object is an instance of VpcNetworkAclAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcNetworkAclAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcNetworkAclAttachment.__pulumiType;
    }

    /**
     * The ID of the network ACL.
     */
    public readonly networkAclId!: pulumi.Output<string>;
    /**
     * The ID of the associated resource.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The type of the associated resource. Valid values: `VSwitch`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * The status of the Network Acl Attachment.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a VpcNetworkAclAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcNetworkAclAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcNetworkAclAttachmentArgs | VpcNetworkAclAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcNetworkAclAttachmentState | undefined;
            resourceInputs["networkAclId"] = state ? state.networkAclId : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as VpcNetworkAclAttachmentArgs | undefined;
            if ((!args || args.networkAclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkAclId'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            resourceInputs["networkAclId"] = args ? args.networkAclId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcNetworkAclAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcNetworkAclAttachment resources.
 */
export interface VpcNetworkAclAttachmentState {
    /**
     * The ID of the network ACL.
     */
    networkAclId?: pulumi.Input<string>;
    /**
     * The ID of the associated resource.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * The type of the associated resource. Valid values: `VSwitch`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * The status of the Network Acl Attachment.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcNetworkAclAttachment resource.
 */
export interface VpcNetworkAclAttachmentArgs {
    /**
     * The ID of the network ACL.
     */
    networkAclId: pulumi.Input<string>;
    /**
     * The ID of the associated resource.
     */
    resourceId: pulumi.Input<string>;
    /**
     * The type of the associated resource. Valid values: `VSwitch`.
     */
    resourceType: pulumi.Input<string>;
}
