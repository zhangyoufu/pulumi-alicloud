// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Service Mesh Extension Provider resource.
 *
 * For information about Service Mesh Extension Provider and how to use it, see [What is Extension Provider](https://help.aliyun.com/document_detail/461549.html).
 *
 * > **NOTE:** Available in v1.191.0+.
 *
 * ## Import
 *
 * Service Mesh Extension Provider can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:servicemesh/extensionProvider:ExtensionProvider example <service_mesh_id>:<type>:<extension_provider_name>
 * ```
 */
export class ExtensionProvider extends pulumi.CustomResource {
    /**
     * Get an existing ExtensionProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionProviderState, opts?: pulumi.CustomResourceOptions): ExtensionProvider {
        return new ExtensionProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:servicemesh/extensionProvider:ExtensionProvider';

    /**
     * Returns true if the given object is an instance of ExtensionProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExtensionProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExtensionProvider.__pulumiType;
    }

    /**
     * The config of the Service Mesh Extension Provider. The `config` format is json.
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The name of the Service Mesh Extension Provider. It must be prefixed with `$type-`, for example `httpextauth-xxx`, `grpcextauth-xxx`.
     */
    public readonly extensionProviderName!: pulumi.Output<string>;
    /**
     * The ID of the Service Mesh.
     */
    public readonly serviceMeshId!: pulumi.Output<string>;
    /**
     * The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ExtensionProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtensionProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionProviderArgs | ExtensionProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtensionProviderState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["extensionProviderName"] = state ? state.extensionProviderName : undefined;
            resourceInputs["serviceMeshId"] = state ? state.serviceMeshId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ExtensionProviderArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.extensionProviderName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extensionProviderName'");
            }
            if ((!args || args.serviceMeshId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceMeshId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["extensionProviderName"] = args ? args.extensionProviderName : undefined;
            resourceInputs["serviceMeshId"] = args ? args.serviceMeshId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExtensionProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExtensionProvider resources.
 */
export interface ExtensionProviderState {
    /**
     * The config of the Service Mesh Extension Provider. The `config` format is json.
     */
    config?: pulumi.Input<string>;
    /**
     * The name of the Service Mesh Extension Provider. It must be prefixed with `$type-`, for example `httpextauth-xxx`, `grpcextauth-xxx`.
     */
    extensionProviderName?: pulumi.Input<string>;
    /**
     * The ID of the Service Mesh.
     */
    serviceMeshId?: pulumi.Input<string>;
    /**
     * The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExtensionProvider resource.
 */
export interface ExtensionProviderArgs {
    /**
     * The config of the Service Mesh Extension Provider. The `config` format is json.
     */
    config: pulumi.Input<string>;
    /**
     * The name of the Service Mesh Extension Provider. It must be prefixed with `$type-`, for example `httpextauth-xxx`, `grpcextauth-xxx`.
     */
    extensionProviderName: pulumi.Input<string>;
    /**
     * The ID of the Service Mesh.
     */
    serviceMeshId: pulumi.Input<string>;
    /**
     * The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
     */
    type: pulumi.Input<string>;
}
