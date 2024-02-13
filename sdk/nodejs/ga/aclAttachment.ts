// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Acl Attachment resource.
 *
 * For information about Global Accelerator (GA) Acl Attachment and how to use it, see [What is Acl Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-associateaclswithlistener).
 *
 * > **NOTE:** Available since v1.150.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAccelerator = new alicloud.ga.Accelerator("defaultAccelerator", {
 *     duration: 1,
 *     autoUseCoupon: true,
 *     spec: "1",
 * });
 * const defaultBandwidthPackage = new alicloud.ga.BandwidthPackage("defaultBandwidthPackage", {
 *     bandwidth: 100,
 *     type: "Basic",
 *     bandwidthType: "Basic",
 *     paymentType: "PayAsYouGo",
 *     billingType: "PayBy95",
 *     ratio: 30,
 * });
 * const defaultBandwidthPackageAttachment = new alicloud.ga.BandwidthPackageAttachment("defaultBandwidthPackageAttachment", {
 *     acceleratorId: defaultAccelerator.id,
 *     bandwidthPackageId: defaultBandwidthPackage.id,
 * });
 * const defaultListener = new alicloud.ga.Listener("defaultListener", {
 *     acceleratorId: defaultBandwidthPackageAttachment.acceleratorId,
 *     portRanges: [{
 *         fromPort: 80,
 *         toPort: 80,
 *     }],
 * });
 * const defaultAcl = new alicloud.ga.Acl("defaultAcl", {
 *     aclName: "terraform-example",
 *     addressIpVersion: "IPv4",
 * });
 * const defaultAclEntryAttachment = new alicloud.ga.AclEntryAttachment("defaultAclEntryAttachment", {
 *     aclId: defaultAcl.id,
 *     entry: "192.168.1.1/32",
 *     entryDescription: "terraform-example",
 * });
 * const defaultAclAttachment = new alicloud.ga.AclAttachment("defaultAclAttachment", {
 *     listenerId: defaultListener.id,
 *     aclId: defaultAcl.id,
 *     aclType: "white",
 * });
 * ```
 *
 * ## Import
 *
 * Global Accelerator (GA) Acl Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ga/aclAttachment:AclAttachment example <listener_id>:<acl_id>
 * ```
 */
export class AclAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AclAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclAttachmentState, opts?: pulumi.CustomResourceOptions): AclAttachment {
        return new AclAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/aclAttachment:AclAttachment';

    /**
     * Returns true if the given object is an instance of AclAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AclAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AclAttachment.__pulumiType;
    }

    /**
     * The ID of an ACL.
     */
    public readonly aclId!: pulumi.Output<string>;
    /**
     * The type of the ACL. Valid values:
     */
    public readonly aclType!: pulumi.Output<string>;
    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * The status of the Acl Attachment.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AclAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclAttachmentArgs | AclAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclAttachmentState | undefined;
            resourceInputs["aclId"] = state ? state.aclId : undefined;
            resourceInputs["aclType"] = state ? state.aclType : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AclAttachmentArgs | undefined;
            if ((!args || args.aclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclId'");
            }
            if ((!args || args.aclType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclType'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            resourceInputs["aclId"] = args ? args.aclId : undefined;
            resourceInputs["aclType"] = args ? args.aclType : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AclAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AclAttachment resources.
 */
export interface AclAttachmentState {
    /**
     * The ID of an ACL.
     */
    aclId?: pulumi.Input<string>;
    /**
     * The type of the ACL. Valid values:
     */
    aclType?: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The status of the Acl Attachment.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AclAttachment resource.
 */
export interface AclAttachmentArgs {
    /**
     * The ID of an ACL.
     */
    aclId: pulumi.Input<string>;
    /**
     * The type of the ACL. Valid values:
     */
    aclType: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the listener.
     */
    listenerId: pulumi.Input<string>;
}
