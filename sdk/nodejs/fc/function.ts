// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Function Compute function can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fc/function:Function foo my-fc-service:hello-world
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
     */
    public readonly caPort!: pulumi.Output<number | undefined>;
    /**
     * The checksum (crc64) of the function code.
     */
    public readonly codeChecksum!: pulumi.Output<string>;
    /**
     * The configuration for custom container runtime.
     */
    public readonly customContainerConfig!: pulumi.Output<outputs.fc.FunctionCustomContainerConfig | undefined>;
    /**
     * The Function Compute function description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A map that defines environment variables for the function.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
     */
    public readonly filename!: pulumi.Output<string | undefined>;
    /**
     * The Function Compute service ID.
     */
    public /*out*/ readonly functionId!: pulumi.Output<string>;
    /**
     * The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
     */
    public readonly handler!: pulumi.Output<string>;
    /**
     * The maximum length of time, in seconds, that the function's initialization should be run for.
     */
    public readonly initializationTimeout!: pulumi.Output<number | undefined>;
    /**
     * The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
     */
    public readonly initializer!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of requests can be executed concurrently within the single function instance.
     */
    public readonly instanceConcurrency!: pulumi.Output<number | undefined>;
    /**
     * The instance type of the function.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The configuration for layers.
     */
    public readonly layers!: pulumi.Output<string[] | undefined>;
    /**
     * Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
     */
    public readonly memorySize!: pulumi.Output<number | undefined>;
    /**
     * The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Setting a prefix to get a only function name. It is conflict with "name".
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
     */
    public readonly ossBucket!: pulumi.Output<string | undefined>;
    /**
     * The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    public readonly ossKey!: pulumi.Output<string | undefined>;
    /**
     * See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * The Function Compute service name.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * The amount of time your function has to run in seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["caPort"] = state ? state.caPort : undefined;
            resourceInputs["codeChecksum"] = state ? state.codeChecksum : undefined;
            resourceInputs["customContainerConfig"] = state ? state.customContainerConfig : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            resourceInputs["filename"] = state ? state.filename : undefined;
            resourceInputs["functionId"] = state ? state.functionId : undefined;
            resourceInputs["handler"] = state ? state.handler : undefined;
            resourceInputs["initializationTimeout"] = state ? state.initializationTimeout : undefined;
            resourceInputs["initializer"] = state ? state.initializer : undefined;
            resourceInputs["instanceConcurrency"] = state ? state.instanceConcurrency : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["lastModified"] = state ? state.lastModified : undefined;
            resourceInputs["layers"] = state ? state.layers : undefined;
            resourceInputs["memorySize"] = state ? state.memorySize : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["ossBucket"] = state ? state.ossBucket : undefined;
            resourceInputs["ossKey"] = state ? state.ossKey : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.handler === undefined) && !opts.urn) {
                throw new Error("Missing required property 'handler'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            resourceInputs["caPort"] = args ? args.caPort : undefined;
            resourceInputs["codeChecksum"] = args ? args.codeChecksum : undefined;
            resourceInputs["customContainerConfig"] = args ? args.customContainerConfig : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["filename"] = args ? args.filename : undefined;
            resourceInputs["handler"] = args ? args.handler : undefined;
            resourceInputs["initializationTimeout"] = args ? args.initializationTimeout : undefined;
            resourceInputs["initializer"] = args ? args.initializer : undefined;
            resourceInputs["instanceConcurrency"] = args ? args.instanceConcurrency : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["layers"] = args ? args.layers : undefined;
            resourceInputs["memorySize"] = args ? args.memorySize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["ossBucket"] = args ? args.ossBucket : undefined;
            resourceInputs["ossKey"] = args ? args.ossKey : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["functionId"] = undefined /*out*/;
            resourceInputs["lastModified"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
     */
    caPort?: pulumi.Input<number>;
    /**
     * The checksum (crc64) of the function code.
     */
    codeChecksum?: pulumi.Input<string>;
    /**
     * The configuration for custom container runtime.
     */
    customContainerConfig?: pulumi.Input<inputs.fc.FunctionCustomContainerConfig>;
    /**
     * The Function Compute function description.
     */
    description?: pulumi.Input<string>;
    /**
     * A map that defines environment variables for the function.
     */
    environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
     */
    filename?: pulumi.Input<string>;
    /**
     * The Function Compute service ID.
     */
    functionId?: pulumi.Input<string>;
    /**
     * The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
     */
    handler?: pulumi.Input<string>;
    /**
     * The maximum length of time, in seconds, that the function's initialization should be run for.
     */
    initializationTimeout?: pulumi.Input<number>;
    /**
     * The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
     */
    initializer?: pulumi.Input<string>;
    /**
     * The maximum number of requests can be executed concurrently within the single function instance.
     */
    instanceConcurrency?: pulumi.Input<number>;
    /**
     * The instance type of the function.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    lastModified?: pulumi.Input<string>;
    /**
     * The configuration for layers.
     */
    layers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
     */
    memorySize?: pulumi.Input<number>;
    /**
     * The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only function name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
     */
    ossBucket?: pulumi.Input<string>;
    /**
     * The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    ossKey?: pulumi.Input<string>;
    /**
     * See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
     */
    runtime?: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service?: pulumi.Input<string>;
    /**
     * The amount of time your function has to run in seconds.
     */
    timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * The port that the function listen to, only valid for [custom runtime](https://www.alibabacloud.com/help/doc-detail/132044.htm) and [custom container runtime](https://www.alibabacloud.com/help/doc-detail/179368.htm).
     */
    caPort?: pulumi.Input<number>;
    /**
     * The checksum (crc64) of the function code.
     */
    codeChecksum?: pulumi.Input<string>;
    /**
     * The configuration for custom container runtime.
     */
    customContainerConfig?: pulumi.Input<inputs.fc.FunctionCustomContainerConfig>;
    /**
     * The Function Compute function description.
     */
    description?: pulumi.Input<string>;
    /**
     * A map that defines environment variables for the function.
     */
    environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
     */
    filename?: pulumi.Input<string>;
    /**
     * The function [entry point](https://www.alibabacloud.com/help/doc-detail/157704.htm) in your code.
     */
    handler: pulumi.Input<string>;
    /**
     * The maximum length of time, in seconds, that the function's initialization should be run for.
     */
    initializationTimeout?: pulumi.Input<number>;
    /**
     * The entry point of the function's [initialization](https://www.alibabacloud.com/help/doc-detail/157704.htm).
     */
    initializer?: pulumi.Input<string>;
    /**
     * The maximum number of requests can be executed concurrently within the single function instance.
     */
    instanceConcurrency?: pulumi.Input<number>;
    /**
     * The instance type of the function.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The configuration for layers.
     */
    layers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amount of memory in MB your function can use at runtime. Defaults to `128`. Limits to [128, 32768].
     */
    memorySize?: pulumi.Input<number>;
    /**
     * The Function Compute function name. It is the only in one service and is conflict with "namePrefix".
     */
    name?: pulumi.Input<string>;
    /**
     * Setting a prefix to get a only function name. It is conflict with "name".
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
     */
    ossBucket?: pulumi.Input<string>;
    /**
     * The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
     */
    ossKey?: pulumi.Input<string>;
    /**
     * See [Runtimes][https://www.alibabacloud.com/help/zh/function-compute/latest/manage-functions#multiTask3514] for valid values.
     */
    runtime: pulumi.Input<string>;
    /**
     * The Function Compute service name.
     */
    service: pulumi.Input<string>;
    /**
     * The amount of time your function has to run in seconds.
     */
    timeout?: pulumi.Input<number>;
}
