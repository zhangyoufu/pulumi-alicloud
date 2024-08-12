// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.InstanceAttachmentArgs;
import com.pulumi.alicloud.cen.inputs.InstanceAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CEN child instance attachment resource that associate the network(VPC, CCN, VBR) with the CEN instance.
 * 
 * &gt; **NOTE:** Available since v1.42.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.InstanceAttachment;
 * import com.pulumi.alicloud.cen.InstanceAttachmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var default = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var example = new Network("example", NetworkArgs.builder()
 *             .vpcName("tf_example")
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()
 *             .cenInstanceName("tf_example")
 *             .description("an example for cen")
 *             .build());
 * 
 *         var exampleInstanceAttachment = new InstanceAttachment("exampleInstanceAttachment", InstanceAttachmentArgs.builder()
 *             .instanceId(exampleInstance.id())
 *             .childInstanceId(example.id())
 *             .childInstanceType("VPC")
 *             .childInstanceRegionId(default_.regions()[0].id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * CEN instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cen/instanceAttachment:InstanceAttachment example cen-m7i7pjmkon********:vpc-2ze2w07mcy9nz********:VPC:cn-beijing
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/instanceAttachment:InstanceAttachment")
public class InstanceAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The account ID to which the CEN instance belongs.
     * 
     * -&gt;**NOTE:** Ensure that the child instance is not used in Express Connect.
     * 
     */
    @Export(name="cenOwnerId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cenOwnerId;

    /**
     * @return The account ID to which the CEN instance belongs.
     * 
     * -&gt;**NOTE:** Ensure that the child instance is not used in Express Connect.
     * 
     */
    public Output<Optional<Integer>> cenOwnerId() {
        return Codegen.optional(this.cenOwnerId);
    }
    /**
     * The ID of the child instance to attach.
     * 
     */
    @Export(name="childInstanceId", refs={String.class}, tree="[0]")
    private Output<String> childInstanceId;

    /**
     * @return The ID of the child instance to attach.
     * 
     */
    public Output<String> childInstanceId() {
        return this.childInstanceId;
    }
    /**
     * The uid of the child instance. Only used when attach a child instance of other account.
     * 
     */
    @Export(name="childInstanceOwnerId", refs={Integer.class}, tree="[0]")
    private Output<Integer> childInstanceOwnerId;

    /**
     * @return The uid of the child instance. Only used when attach a child instance of other account.
     * 
     */
    public Output<Integer> childInstanceOwnerId() {
        return this.childInstanceOwnerId;
    }
    /**
     * The region ID of the child instance to attach.
     * 
     */
    @Export(name="childInstanceRegionId", refs={String.class}, tree="[0]")
    private Output<String> childInstanceRegionId;

    /**
     * @return The region ID of the child instance to attach.
     * 
     */
    public Output<String> childInstanceRegionId() {
        return this.childInstanceRegionId;
    }
    /**
     * The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
     * 
     */
    @Export(name="childInstanceType", refs={String.class}, tree="[0]")
    private Output<String> childInstanceType;

    /**
     * @return The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
     * 
     */
    public Output<String> childInstanceType() {
        return this.childInstanceType;
    }
    /**
     * The ID of the CEN.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the CEN.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The associating status of the network.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The associating status of the network.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceAttachment(java.lang.String name) {
        this(name, InstanceAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceAttachment(java.lang.String name, InstanceAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceAttachment(java.lang.String name, InstanceAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/instanceAttachment:InstanceAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private InstanceAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/instanceAttachment:InstanceAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static InstanceAttachmentArgs makeArgs(InstanceAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstanceAttachmentArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static InstanceAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceAttachment(name, id, state, options);
    }
}
