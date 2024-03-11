// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Firewall Vpc Firewall Control Policy resource.
 *
 * For information about Cloud Firewall Vpc Firewall Control Policy and how to use it, see [What is Vpc Firewall Control Policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcontrolpolicy).
 *
 * > **NOTE:** Available since v1.194.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultAccount = alicloud.getAccount({});
 * const defaultInstance = new alicloud.cen.Instance("defaultInstance", {
 *     cenInstanceName: name,
 *     description: "example_value",
 *     tags: {
 *         Created: "TF",
 *         For: "acceptance test",
 *     },
 * });
 * const defaultFirewallVpcFirewallControlPolicy = new alicloud.cloudfirewall.FirewallVpcFirewallControlPolicy("defaultFirewallVpcFirewallControlPolicy", {
 *     order: 1,
 *     destination: "127.0.0.2/32",
 *     applicationName: "ANY",
 *     description: "example_value",
 *     sourceType: "net",
 *     destPort: "80/88",
 *     aclAction: "accept",
 *     lang: "zh",
 *     destinationType: "net",
 *     source: "127.0.0.1/32",
 *     destPortType: "port",
 *     proto: "TCP",
 *     release: true,
 *     memberUid: defaultAccount.then(defaultAccount => defaultAccount.id),
 *     vpcFirewallId: defaultInstance.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Cloud Firewall Vpc Firewall Control Policy can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy example <vpc_firewall_id>:<acl_uuid>
 * ```
 */
export class FirewallVpcFirewallControlPolicy extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVpcFirewallControlPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVpcFirewallControlPolicyState, opts?: pulumi.CustomResourceOptions): FirewallVpcFirewallControlPolicy {
        return new FirewallVpcFirewallControlPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy';

    /**
     * Returns true if the given object is an instance of FirewallVpcFirewallControlPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVpcFirewallControlPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVpcFirewallControlPolicy.__pulumiType;
    }

    /**
     * The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
     */
    public readonly aclAction!: pulumi.Output<string>;
    /**
     * Access control over VPC firewalls strategy unique identifier.
     */
    public /*out*/ readonly aclUuid!: pulumi.Output<string>;
    /**
     * Policy specifies the application ID.
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
     */
    public readonly applicationName!: pulumi.Output<string>;
    /**
     * Access control over VPC firewalls description of the strategy information.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
     */
    public readonly destPort!: pulumi.Output<string>;
    /**
     * Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
     */
    public readonly destPortGroup!: pulumi.Output<string | undefined>;
    /**
     * Port Address Book port list.
     */
    public /*out*/ readonly destPortGroupPorts!: pulumi.Output<string[]>;
    /**
     * The type of the destination port in the access control policy. Valid values: `port`, `group`.
     */
    public readonly destPortType!: pulumi.Output<string>;
    /**
     * The destination address in the access control policy. Valid values:
     * - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
     * - If `destinationType` is set to `group`, the value of `destination` must be an address book.
     * - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Destination address book defined in the address list.
     */
    public /*out*/ readonly destinationGroupCidrs!: pulumi.Output<string[]>;
    /**
     * The destination address book type in the access control policy.
     */
    public /*out*/ readonly destinationGroupType!: pulumi.Output<string>;
    /**
     * The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
     */
    public readonly destinationType!: pulumi.Output<string>;
    /**
     * Control strategy of hits per second.
     */
    public /*out*/ readonly hitTimes!: pulumi.Output<number>;
    /**
     * The language of the content within the request and response. Valid values: `zh`, `en`.
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * The UID of the member account of the current Alibaba cloud account.
     */
    public readonly memberUid!: pulumi.Output<string>;
    /**
     * The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
     */
    public readonly proto!: pulumi.Output<string>;
    /**
     * The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
     */
    public readonly release!: pulumi.Output<boolean>;
    /**
     * Access control over VPC firewalls strategy in the source address.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * SOURCE address of the address list.
     */
    public /*out*/ readonly sourceGroupCidrs!: pulumi.Output<string[]>;
    /**
     * The source address type in the access control policy.
     */
    public /*out*/ readonly sourceGroupType!: pulumi.Output<string>;
    /**
     * The type of the source address in the access control policy. Valid values: `net`, `group`.
     */
    public readonly sourceType!: pulumi.Output<string>;
    /**
     * The ID of the VPC firewall instance. Valid values:
     * - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
     * - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
     */
    public readonly vpcFirewallId!: pulumi.Output<string>;

    /**
     * Create a FirewallVpcFirewallControlPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallVpcFirewallControlPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVpcFirewallControlPolicyArgs | FirewallVpcFirewallControlPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVpcFirewallControlPolicyState | undefined;
            resourceInputs["aclAction"] = state ? state.aclAction : undefined;
            resourceInputs["aclUuid"] = state ? state.aclUuid : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["applicationName"] = state ? state.applicationName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destPort"] = state ? state.destPort : undefined;
            resourceInputs["destPortGroup"] = state ? state.destPortGroup : undefined;
            resourceInputs["destPortGroupPorts"] = state ? state.destPortGroupPorts : undefined;
            resourceInputs["destPortType"] = state ? state.destPortType : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["destinationGroupCidrs"] = state ? state.destinationGroupCidrs : undefined;
            resourceInputs["destinationGroupType"] = state ? state.destinationGroupType : undefined;
            resourceInputs["destinationType"] = state ? state.destinationType : undefined;
            resourceInputs["hitTimes"] = state ? state.hitTimes : undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
            resourceInputs["memberUid"] = state ? state.memberUid : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["proto"] = state ? state.proto : undefined;
            resourceInputs["release"] = state ? state.release : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourceGroupCidrs"] = state ? state.sourceGroupCidrs : undefined;
            resourceInputs["sourceGroupType"] = state ? state.sourceGroupType : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["vpcFirewallId"] = state ? state.vpcFirewallId : undefined;
        } else {
            const args = argsOrState as FirewallVpcFirewallControlPolicyArgs | undefined;
            if ((!args || args.aclAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclAction'");
            }
            if ((!args || args.applicationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationName'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.destinationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationType'");
            }
            if ((!args || args.order === undefined) && !opts.urn) {
                throw new Error("Missing required property 'order'");
            }
            if ((!args || args.proto === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proto'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            if ((!args || args.vpcFirewallId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcFirewallId'");
            }
            resourceInputs["aclAction"] = args ? args.aclAction : undefined;
            resourceInputs["applicationName"] = args ? args.applicationName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destPort"] = args ? args.destPort : undefined;
            resourceInputs["destPortGroup"] = args ? args.destPortGroup : undefined;
            resourceInputs["destPortType"] = args ? args.destPortType : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["destinationType"] = args ? args.destinationType : undefined;
            resourceInputs["lang"] = args ? args.lang : undefined;
            resourceInputs["memberUid"] = args ? args.memberUid : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["proto"] = args ? args.proto : undefined;
            resourceInputs["release"] = args ? args.release : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["vpcFirewallId"] = args ? args.vpcFirewallId : undefined;
            resourceInputs["aclUuid"] = undefined /*out*/;
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["destPortGroupPorts"] = undefined /*out*/;
            resourceInputs["destinationGroupCidrs"] = undefined /*out*/;
            resourceInputs["destinationGroupType"] = undefined /*out*/;
            resourceInputs["hitTimes"] = undefined /*out*/;
            resourceInputs["sourceGroupCidrs"] = undefined /*out*/;
            resourceInputs["sourceGroupType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallVpcFirewallControlPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVpcFirewallControlPolicy resources.
 */
export interface FirewallVpcFirewallControlPolicyState {
    /**
     * The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
     */
    aclAction?: pulumi.Input<string>;
    /**
     * Access control over VPC firewalls strategy unique identifier.
     */
    aclUuid?: pulumi.Input<string>;
    /**
     * Policy specifies the application ID.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
     */
    applicationName?: pulumi.Input<string>;
    /**
     * Access control over VPC firewalls description of the strategy information.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
     */
    destPort?: pulumi.Input<string>;
    /**
     * Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
     */
    destPortGroup?: pulumi.Input<string>;
    /**
     * Port Address Book port list.
     */
    destPortGroupPorts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the destination port in the access control policy. Valid values: `port`, `group`.
     */
    destPortType?: pulumi.Input<string>;
    /**
     * The destination address in the access control policy. Valid values:
     * - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
     * - If `destinationType` is set to `group`, the value of `destination` must be an address book.
     * - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
     */
    destination?: pulumi.Input<string>;
    /**
     * Destination address book defined in the address list.
     */
    destinationGroupCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The destination address book type in the access control policy.
     */
    destinationGroupType?: pulumi.Input<string>;
    /**
     * The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
     */
    destinationType?: pulumi.Input<string>;
    /**
     * Control strategy of hits per second.
     */
    hitTimes?: pulumi.Input<number>;
    /**
     * The language of the content within the request and response. Valid values: `zh`, `en`.
     */
    lang?: pulumi.Input<string>;
    /**
     * The UID of the member account of the current Alibaba cloud account.
     */
    memberUid?: pulumi.Input<string>;
    /**
     * The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
     */
    order?: pulumi.Input<number>;
    /**
     * The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
     */
    proto?: pulumi.Input<string>;
    /**
     * The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
     */
    release?: pulumi.Input<boolean>;
    /**
     * Access control over VPC firewalls strategy in the source address.
     */
    source?: pulumi.Input<string>;
    /**
     * SOURCE address of the address list.
     */
    sourceGroupCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The source address type in the access control policy.
     */
    sourceGroupType?: pulumi.Input<string>;
    /**
     * The type of the source address in the access control policy. Valid values: `net`, `group`.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * The ID of the VPC firewall instance. Valid values:
     * - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
     * - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
     */
    vpcFirewallId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallVpcFirewallControlPolicy resource.
 */
export interface FirewallVpcFirewallControlPolicyArgs {
    /**
     * The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
     */
    aclAction: pulumi.Input<string>;
    /**
     * The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
     */
    applicationName: pulumi.Input<string>;
    /**
     * Access control over VPC firewalls description of the strategy information.
     */
    description: pulumi.Input<string>;
    /**
     * The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
     */
    destPort?: pulumi.Input<string>;
    /**
     * Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
     */
    destPortGroup?: pulumi.Input<string>;
    /**
     * The type of the destination port in the access control policy. Valid values: `port`, `group`.
     */
    destPortType?: pulumi.Input<string>;
    /**
     * The destination address in the access control policy. Valid values:
     * - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
     * - If `destinationType` is set to `group`, the value of `destination` must be an address book.
     * - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
     */
    destination: pulumi.Input<string>;
    /**
     * The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
     */
    destinationType: pulumi.Input<string>;
    /**
     * The language of the content within the request and response. Valid values: `zh`, `en`.
     */
    lang?: pulumi.Input<string>;
    /**
     * The UID of the member account of the current Alibaba cloud account.
     */
    memberUid?: pulumi.Input<string>;
    /**
     * The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
     */
    order: pulumi.Input<number>;
    /**
     * The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
     */
    proto: pulumi.Input<string>;
    /**
     * The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
     */
    release?: pulumi.Input<boolean>;
    /**
     * Access control over VPC firewalls strategy in the source address.
     */
    source: pulumi.Input<string>;
    /**
     * The type of the source address in the access control policy. Valid values: `net`, `group`.
     */
    sourceType: pulumi.Input<string>;
    /**
     * The ID of the VPC firewall instance. Valid values:
     * - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
     * - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
     */
    vpcFirewallId: pulumi.Input<string>;
}
