// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RabbitMQ (AMQP) Binding resource to bind tha exchange with another exchange or queue.
 *
 * For information about RabbitMQ (AMQP) Binding and how to use it, see [What is Binding](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/createbinding).
 *
 * > **NOTE:** Available since v1.135.0.
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
 * const defaultInstance = new alicloud.amqp.Instance("defaultInstance", {
 *     instanceType: "enterprise",
 *     maxTps: "3000",
 *     queueCapacity: "200",
 *     storageSize: "700",
 *     supportEip: false,
 *     maxEipTps: "128",
 *     paymentType: "Subscription",
 *     period: 1,
 * });
 * const defaultVirtualHost = new alicloud.amqp.VirtualHost("defaultVirtualHost", {
 *     instanceId: defaultInstance.id,
 *     virtualHostName: "tf-example",
 * });
 * const defaultExchange = new alicloud.amqp.Exchange("defaultExchange", {
 *     autoDeleteState: false,
 *     exchangeName: "tf-example",
 *     exchangeType: "HEADERS",
 *     instanceId: defaultInstance.id,
 *     internal: false,
 *     virtualHostName: defaultVirtualHost.virtualHostName,
 * });
 * const defaultQueue = new alicloud.amqp.Queue("defaultQueue", {
 *     instanceId: defaultInstance.id,
 *     queueName: "tf-example",
 *     virtualHostName: defaultVirtualHost.virtualHostName,
 * });
 * const defaultBinding = new alicloud.amqp.Binding("defaultBinding", {
 *     argument: "x-match:all",
 *     bindingKey: defaultQueue.queueName,
 *     bindingType: "QUEUE",
 *     destinationName: "tf-example",
 *     instanceId: defaultInstance.id,
 *     sourceExchange: defaultExchange.exchangeName,
 *     virtualHostName: defaultVirtualHost.virtualHostName,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * RabbitMQ (AMQP) Binding can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:amqp/binding:Binding example <instance_id>:<virtual_host_name>:<source_exchange>:<destination_name>
 * ```
 */
export class Binding extends pulumi.CustomResource {
    /**
     * Get an existing Binding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BindingState, opts?: pulumi.CustomResourceOptions): Binding {
        return new Binding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:amqp/binding:Binding';

    /**
     * Returns true if the given object is an instance of Binding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Binding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Binding.__pulumiType;
    }

    /**
     * X-match Attributes. Valid Values: 
     * * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
     * * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
     *
     * **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
     */
    public readonly argument!: pulumi.Output<string>;
    /**
     * The Binding Key.
     * * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     * The binding key must be 1 to 255 characters in length.
     * * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     * If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
     * The binding key must be 1 to 255 characters in length.
     */
    public readonly bindingKey!: pulumi.Output<string>;
    /**
     * The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
     */
    public readonly bindingType!: pulumi.Output<string>;
    /**
     * The Target Queue Or Exchange of the Name.
     */
    public readonly destinationName!: pulumi.Output<string>;
    /**
     * Instance Id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The Source Exchange Name.
     */
    public readonly sourceExchange!: pulumi.Output<string>;
    /**
     * Virtualhost Name.
     */
    public readonly virtualHostName!: pulumi.Output<string>;

    /**
     * Create a Binding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BindingArgs | BindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BindingState | undefined;
            resourceInputs["argument"] = state ? state.argument : undefined;
            resourceInputs["bindingKey"] = state ? state.bindingKey : undefined;
            resourceInputs["bindingType"] = state ? state.bindingType : undefined;
            resourceInputs["destinationName"] = state ? state.destinationName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["sourceExchange"] = state ? state.sourceExchange : undefined;
            resourceInputs["virtualHostName"] = state ? state.virtualHostName : undefined;
        } else {
            const args = argsOrState as BindingArgs | undefined;
            if ((!args || args.bindingKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindingKey'");
            }
            if ((!args || args.bindingType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindingType'");
            }
            if ((!args || args.destinationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.sourceExchange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceExchange'");
            }
            if ((!args || args.virtualHostName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualHostName'");
            }
            resourceInputs["argument"] = args ? args.argument : undefined;
            resourceInputs["bindingKey"] = args ? args.bindingKey : undefined;
            resourceInputs["bindingType"] = args ? args.bindingType : undefined;
            resourceInputs["destinationName"] = args ? args.destinationName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["sourceExchange"] = args ? args.sourceExchange : undefined;
            resourceInputs["virtualHostName"] = args ? args.virtualHostName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Binding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Binding resources.
 */
export interface BindingState {
    /**
     * X-match Attributes. Valid Values: 
     * * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
     * * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
     *
     * **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
     */
    argument?: pulumi.Input<string>;
    /**
     * The Binding Key.
     * * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     * The binding key must be 1 to 255 characters in length.
     * * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     * If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
     * The binding key must be 1 to 255 characters in length.
     */
    bindingKey?: pulumi.Input<string>;
    /**
     * The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
     */
    bindingType?: pulumi.Input<string>;
    /**
     * The Target Queue Or Exchange of the Name.
     */
    destinationName?: pulumi.Input<string>;
    /**
     * Instance Id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The Source Exchange Name.
     */
    sourceExchange?: pulumi.Input<string>;
    /**
     * Virtualhost Name.
     */
    virtualHostName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Binding resource.
 */
export interface BindingArgs {
    /**
     * X-match Attributes. Valid Values: 
     * * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
     * * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
     *
     * **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
     */
    argument?: pulumi.Input<string>;
    /**
     * The Binding Key.
     * * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     * The binding key must be 1 to 255 characters in length.
     * * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
     * If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
     * The binding key must be 1 to 255 characters in length.
     */
    bindingKey: pulumi.Input<string>;
    /**
     * The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
     */
    bindingType: pulumi.Input<string>;
    /**
     * The Target Queue Or Exchange of the Name.
     */
    destinationName: pulumi.Input<string>;
    /**
     * Instance Id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The Source Exchange Name.
     */
    sourceExchange: pulumi.Input<string>;
    /**
     * Virtualhost Name.
     */
    virtualHostName: pulumi.Input<string>;
}
