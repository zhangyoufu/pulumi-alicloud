// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this data source can create Event Bridge service-linked roles(SLR). EventBridge may need to access another Alibaba Cloud service to implement a specific feature. In this case, EventBridge must assume a specific service-linked role, which is a Resource Access Management (RAM) role, to obtain permissions to access another Alibaba Cloud service.
 *
 * For information about Event Bridge service-linked roles(SLR) and how to use it, see [What is service-linked roles](https://www.alibabacloud.com/help/doc-detail/181425.htm).
 *
 * > **NOTE:** Available in v1.129.0+. After the version 1.142.0, the resource is renamed as `alicloud.eventbridge.ServiceLinkedRole`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const serviceLinkedRole = new alicloud.eventbridge.ServiceLinkedRole("serviceLinkedRole", {productName: "AliyunServiceRoleForEventBridgeSendToMNS"});
 * ```
 *
 * ## Import
 *
 * Event Bridge service-linked roles(SLR) can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eventbridge/serviceLinkedRole:ServiceLinkedRole example <product_name>
 * ```
 */
export class ServiceLinkedRole extends pulumi.CustomResource {
    /**
     * Get an existing ServiceLinkedRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceLinkedRoleState, opts?: pulumi.CustomResourceOptions): ServiceLinkedRole {
        return new ServiceLinkedRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eventbridge/serviceLinkedRole:ServiceLinkedRole';

    /**
     * Returns true if the given object is an instance of ServiceLinkedRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceLinkedRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceLinkedRole.__pulumiType;
    }

    /**
     * The product name for SLR. EventBridge can automatically create the following service-linked roles:
     * Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
     * Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
     */
    public readonly productName!: pulumi.Output<string>;

    /**
     * Create a ServiceLinkedRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceLinkedRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceLinkedRoleArgs | ServiceLinkedRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceLinkedRoleState | undefined;
            resourceInputs["productName"] = state ? state.productName : undefined;
        } else {
            const args = argsOrState as ServiceLinkedRoleArgs | undefined;
            if ((!args || args.productName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productName'");
            }
            resourceInputs["productName"] = args ? args.productName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceLinkedRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceLinkedRole resources.
 */
export interface ServiceLinkedRoleState {
    /**
     * The product name for SLR. EventBridge can automatically create the following service-linked roles:
     * Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
     * Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
     */
    productName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceLinkedRole resource.
 */
export interface ServiceLinkedRoleArgs {
    /**
     * The product name for SLR. EventBridge can automatically create the following service-linked roles:
     * Event source related: `AliyunServiceRoleForEventBridgeSendToMNS`,`AliyunServiceRoleForEventBridgeSourceRocketMQ`, `AliyunServiceRoleForEventBridgeSourceActionTrail`, `AliyunServiceRoleForEventBridgeSourceRabbitMQ`
     * Target related: `AliyunServiceRoleForEventBridgeConnectVPC`, `AliyunServiceRoleForEventBridgeSendToFC`, `AliyunServiceRoleForEventBridgeSendToSMS`, `AliyunServiceRoleForEventBridgeSendToDirectMail`, `AliyunServiceRoleForEventBridgeSendToRabbitMQ`, `AliyunServiceRoleForEventBridgeSendToRocketMQ`
     */
    productName: pulumi.Input<string>;
}
