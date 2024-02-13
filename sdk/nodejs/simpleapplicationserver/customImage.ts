// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Simple Application Server Custom Image resource.
 *
 * For information about Simple Application Server Custom Image and how to use it, see [What is Custom Image](https://www.alibabacloud.com/help/en/doc-detail/333535.htm).
 *
 * > **NOTE:** Available since v1.143.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf_example";
 * const defaultImages = alicloud.simpleapplicationserver.getImages({});
 * const defaultServerPlans = alicloud.simpleapplicationserver.getServerPlans({});
 * const defaultInstance = new alicloud.simpleapplicationserver.Instance("defaultInstance", {
 *     paymentType: "Subscription",
 *     planId: defaultServerPlans.then(defaultServerPlans => defaultServerPlans.plans?.[0]?.id),
 *     instanceName: name,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     period: 1,
 *     dataDiskSize: 100,
 * });
 * const defaultServerDisks = alicloud.simpleapplicationserver.getServerDisksOutput({
 *     instanceId: defaultInstance.id,
 * });
 * const defaultSnapshot = new alicloud.simpleapplicationserver.Snapshot("defaultSnapshot", {
 *     diskId: defaultServerDisks.apply(defaultServerDisks => defaultServerDisks.ids?.[0]),
 *     snapshotName: name,
 * });
 * const defaultCustomImage = new alicloud.simpleapplicationserver.CustomImage("defaultCustomImage", {
 *     customImageName: name,
 *     instanceId: defaultInstance.id,
 *     systemSnapshotId: defaultSnapshot.id,
 *     status: "Share",
 *     description: name,
 * });
 * ```
 *
 * ## Import
 *
 * Simple Application Server Custom Image can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:simpleapplicationserver/customImage:CustomImage example <id>
 * ```
 */
export class CustomImage extends pulumi.CustomResource {
    /**
     * Get an existing CustomImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomImageState, opts?: pulumi.CustomResourceOptions): CustomImage {
        return new CustomImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:simpleapplicationserver/customImage:CustomImage';

    /**
     * Returns true if the given object is an instance of CustomImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomImage.__pulumiType;
    }

    /**
     * The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
     */
    public readonly customImageName!: pulumi.Output<string>;
    /**
     * The description of the Custom Image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
     *
     * **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The ID of the system snapshot.
     */
    public readonly systemSnapshotId!: pulumi.Output<string>;

    /**
     * Create a CustomImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomImageArgs | CustomImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomImageState | undefined;
            resourceInputs["customImageName"] = state ? state.customImageName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["systemSnapshotId"] = state ? state.systemSnapshotId : undefined;
        } else {
            const args = argsOrState as CustomImageArgs | undefined;
            if ((!args || args.customImageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customImageName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.systemSnapshotId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'systemSnapshotId'");
            }
            resourceInputs["customImageName"] = args ? args.customImageName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["systemSnapshotId"] = args ? args.systemSnapshotId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomImage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomImage resources.
 */
export interface CustomImageState {
    /**
     * The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
     */
    customImageName?: pulumi.Input<string>;
    /**
     * The description of the Custom Image.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
     *
     * **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the system snapshot.
     */
    systemSnapshotId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomImage resource.
 */
export interface CustomImageArgs {
    /**
     * The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
     */
    customImageName: pulumi.Input<string>;
    /**
     * The description of the Custom Image.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
     *
     * **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the system snapshot.
     */
    systemSnapshotId: pulumi.Input<string>;
}
