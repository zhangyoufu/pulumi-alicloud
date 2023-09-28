// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ALB Listener Acl Attachment resource. Associating ACL to listening.
 *
 * For information about ALB Listener Acl Attachment and how to use it, see [What is Listener Acl Attachment](https://www.alibabacloud.com/help/en/server-load-balancer/latest/associateaclswithlistener).
 *
 * > **NOTE:** Available since v1.163.0.
 *
 * > **NOTE:** You can associate at most three ACLs with a listener.
 *
 * > **NOTE:** You can only configure either a whitelist or a blacklist for listener, not at the same time.
 *
 * ## Import
 *
 * ALB Listener Acl Attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:alb/listenerAclAttachment:ListenerAclAttachment example <listener_id>:<acl_id>
 * ```
 */
export class ListenerAclAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ListenerAclAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerAclAttachmentState, opts?: pulumi.CustomResourceOptions): ListenerAclAttachment {
        return new ListenerAclAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alb/listenerAclAttachment:ListenerAclAttachment';

    /**
     * Returns true if the given object is an instance of ListenerAclAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ListenerAclAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ListenerAclAttachment.__pulumiType;
    }

    /**
     * The ID list of the access policy group bound by the listener.
     */
    public readonly aclId!: pulumi.Output<string>;
    /**
     * Access control type:
     * - **White**: only requests from IP addresses or address segments in the selected access control list are forwarded. The whitelist applies to scenarios where only specific IP addresses are allowed to access. There are certain business risks in setting up a whitelist. Once the whitelist is set, only the IP addresses in the whitelist can access the load balancer listener. If whitelist access is enabled but no IP addresses are added to the access policy group, the server load balancer listener forwards all requests.
     * - **Black**: All requests from IP addresses or address segments in the selected access control list are not forwarded. Blacklists are applicable to scenarios where only certain IP addresses are restricted. If blacklist access is enabled and no IP is added to the access policy group, the server load balancer listener forwards all requests.
     */
    public readonly aclType!: pulumi.Output<string>;
    /**
     * Listener instance ID.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * Listener Status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ListenerAclAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerAclAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerAclAttachmentArgs | ListenerAclAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerAclAttachmentState | undefined;
            resourceInputs["aclId"] = state ? state.aclId : undefined;
            resourceInputs["aclType"] = state ? state.aclType : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ListenerAclAttachmentArgs | undefined;
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
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ListenerAclAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerAclAttachment resources.
 */
export interface ListenerAclAttachmentState {
    /**
     * The ID list of the access policy group bound by the listener.
     */
    aclId?: pulumi.Input<string>;
    /**
     * Access control type:
     * - **White**: only requests from IP addresses or address segments in the selected access control list are forwarded. The whitelist applies to scenarios where only specific IP addresses are allowed to access. There are certain business risks in setting up a whitelist. Once the whitelist is set, only the IP addresses in the whitelist can access the load balancer listener. If whitelist access is enabled but no IP addresses are added to the access policy group, the server load balancer listener forwards all requests.
     * - **Black**: All requests from IP addresses or address segments in the selected access control list are not forwarded. Blacklists are applicable to scenarios where only certain IP addresses are restricted. If blacklist access is enabled and no IP is added to the access policy group, the server load balancer listener forwards all requests.
     */
    aclType?: pulumi.Input<string>;
    /**
     * Listener instance ID.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * Listener Status.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ListenerAclAttachment resource.
 */
export interface ListenerAclAttachmentArgs {
    /**
     * The ID list of the access policy group bound by the listener.
     */
    aclId: pulumi.Input<string>;
    /**
     * Access control type:
     * - **White**: only requests from IP addresses or address segments in the selected access control list are forwarded. The whitelist applies to scenarios where only specific IP addresses are allowed to access. There are certain business risks in setting up a whitelist. Once the whitelist is set, only the IP addresses in the whitelist can access the load balancer listener. If whitelist access is enabled but no IP addresses are added to the access policy group, the server load balancer listener forwards all requests.
     * - **Black**: All requests from IP addresses or address segments in the selected access control list are not forwarded. Blacklists are applicable to scenarios where only certain IP addresses are restricted. If blacklist access is enabled and no IP is added to the access policy group, the server load balancer listener forwards all requests.
     */
    aclType: pulumi.Input<string>;
    /**
     * Listener instance ID.
     */
    listenerId: pulumi.Input<string>;
}
