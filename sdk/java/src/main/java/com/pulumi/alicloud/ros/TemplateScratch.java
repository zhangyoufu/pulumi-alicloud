// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ros.TemplateScratchArgs;
import com.pulumi.alicloud.ros.inputs.TemplateScratchState;
import com.pulumi.alicloud.ros.outputs.TemplateScratchPreferenceParameter;
import com.pulumi.alicloud.ros.outputs.TemplateScratchSourceResource;
import com.pulumi.alicloud.ros.outputs.TemplateScratchSourceResourceGroup;
import com.pulumi.alicloud.ros.outputs.TemplateScratchSourceTag;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ROS Template Scratch resource.
 * 
 * For information about ROS Template Scratch and how to use it, see [What is Template Scratch](https://www.alibabacloud.com/help/zh/doc-detail/352074.html).
 * 
 * &gt; **NOTE:** Available in v1.151.0+.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.ros.TemplateScratch;
 * import com.pulumi.alicloud.ros.TemplateScratchArgs;
 * import com.pulumi.alicloud.ros.inputs.TemplateScratchPreferenceParameterArgs;
 * import com.pulumi.alicloud.ros.inputs.TemplateScratchSourceResourceGroupArgs;
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
 *         final var default = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var example = new TemplateScratch("example", TemplateScratchArgs.builder()        
 *             .description("tf_testacc")
 *             .templateScratchType("ResourceImport")
 *             .preferenceParameters(TemplateScratchPreferenceParameterArgs.builder()
 *                 .parameterKey("DeletionPolicy")
 *                 .parameterValue("Retain")
 *                 .build())
 *             .sourceResourceGroup(TemplateScratchSourceResourceGroupArgs.builder()
 *                 .resourceGroupId(default_.ids()[0])
 *                 .resourceTypeFilters("ALIYUN::ECS::VPC")
 *                 .build())
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
 * ROS Template Scratch can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ros/templateScratch:TemplateScratch example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ros/templateScratch:TemplateScratch")
public class TemplateScratch extends com.pulumi.resources.CustomResource {
    /**
     * The Description of the Template Scratch.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The Description of the Template Scratch.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The execution mode. Valid Values: `Async` or `Sync`.
     * 
     */
    @Export(name="executionMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> executionMode;

    /**
     * @return The execution mode. Valid Values: `Async` or `Sync`.
     * 
     */
    public Output<Optional<String>> executionMode() {
        return Codegen.optional(this.executionMode);
    }
    /**
     * Logical ID generation strategy. Valid Values: `LongTypePrefixAndIndexSuffix`, `LongTypePrefixAndHashSuffix` and `ShortTypePrefixAndHashSuffix`.
     * 
     */
    @Export(name="logicalIdStrategy", refs={String.class}, tree="[0]")
    private Output<String> logicalIdStrategy;

    /**
     * @return Logical ID generation strategy. Valid Values: `LongTypePrefixAndIndexSuffix`, `LongTypePrefixAndHashSuffix` and `ShortTypePrefixAndHashSuffix`.
     * 
     */
    public Output<String> logicalIdStrategy() {
        return this.logicalIdStrategy;
    }
    /**
     * Priority parameter. See the following `Block preference_parameters`.
     * 
     */
    @Export(name="preferenceParameters", refs={List.class,TemplateScratchPreferenceParameter.class}, tree="[0,1]")
    private Output<List<TemplateScratchPreferenceParameter>> preferenceParameters;

    /**
     * @return Priority parameter. See the following `Block preference_parameters`.
     * 
     */
    public Output<List<TemplateScratchPreferenceParameter>> preferenceParameters() {
        return this.preferenceParameters;
    }
    /**
     * Source resource grouping. See the following `Block source_resource_group`.
     * 
     */
    @Export(name="sourceResourceGroup", refs={TemplateScratchSourceResourceGroup.class}, tree="[0]")
    private Output</* @Nullable */ TemplateScratchSourceResourceGroup> sourceResourceGroup;

    /**
     * @return Source resource grouping. See the following `Block source_resource_group`.
     * 
     */
    public Output<Optional<TemplateScratchSourceResourceGroup>> sourceResourceGroup() {
        return Codegen.optional(this.sourceResourceGroup);
    }
    /**
     * Source resource. See the following `Block source_resources`.
     * 
     */
    @Export(name="sourceResources", refs={List.class,TemplateScratchSourceResource.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TemplateScratchSourceResource>> sourceResources;

    /**
     * @return Source resource. See the following `Block source_resources`.
     * 
     */
    public Output<Optional<List<TemplateScratchSourceResource>>> sourceResources() {
        return Codegen.optional(this.sourceResources);
    }
    /**
     * Source tag. See the following `Block source_tag`.
     * 
     */
    @Export(name="sourceTag", refs={TemplateScratchSourceTag.class}, tree="[0]")
    private Output</* @Nullable */ TemplateScratchSourceTag> sourceTag;

    /**
     * @return Source tag. See the following `Block source_tag`.
     * 
     */
    public Output<Optional<TemplateScratchSourceTag>> sourceTag() {
        return Codegen.optional(this.sourceTag);
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The type of the Template scan. Valid Values: `ResourceImport` or `ArchitectureReplication`.
     * 
     */
    @Export(name="templateScratchType", refs={String.class}, tree="[0]")
    private Output<String> templateScratchType;

    /**
     * @return The type of the Template scan. Valid Values: `ResourceImport` or `ArchitectureReplication`.
     * 
     */
    public Output<String> templateScratchType() {
        return this.templateScratchType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TemplateScratch(String name) {
        this(name, TemplateScratchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TemplateScratch(String name, TemplateScratchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TemplateScratch(String name, TemplateScratchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ros/templateScratch:TemplateScratch", name, args == null ? TemplateScratchArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TemplateScratch(String name, Output<String> id, @Nullable TemplateScratchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ros/templateScratch:TemplateScratch", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static TemplateScratch get(String name, Output<String> id, @Nullable TemplateScratchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TemplateScratch(name, id, state, options);
    }
}
