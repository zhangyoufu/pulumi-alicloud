// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cms.DynamicTagGroupArgs;
import com.pulumi.alicloud.cms.inputs.DynamicTagGroupState;
import com.pulumi.alicloud.cms.outputs.DynamicTagGroupMatchExpress;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Monitor Service Dynamic Tag Group resource.
 * 
 * For information about Cloud Monitor Service Dynamic Tag Group and how to use it, see [What is Dynamic Tag Group](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createdynamictaggroup).
 * 
 * &gt; **NOTE:** Available since v1.142.0.
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
 * import com.pulumi.alicloud.cms.AlarmContactGroup;
 * import com.pulumi.alicloud.cms.AlarmContactGroupArgs;
 * import com.pulumi.alicloud.cms.DynamicTagGroup;
 * import com.pulumi.alicloud.cms.DynamicTagGroupArgs;
 * import com.pulumi.alicloud.cms.inputs.DynamicTagGroupMatchExpressArgs;
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
 *         var default_ = new AlarmContactGroup("default", AlarmContactGroupArgs.builder()
 *             .alarmContactGroupName("example_value")
 *             .describe("example_value")
 *             .enableSubscribed(true)
 *             .build());
 * 
 *         var defaultDynamicTagGroup = new DynamicTagGroup("defaultDynamicTagGroup", DynamicTagGroupArgs.builder()
 *             .contactGroupLists(default_.id())
 *             .tagKey("your_tag_key")
 *             .matchExpresses(DynamicTagGroupMatchExpressArgs.builder()
 *                 .tagValue("your_tag_value")
 *                 .tagValueMatchFunction("all")
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
 * Cloud Monitor Service Dynamic Tag Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cms/dynamicTagGroup:DynamicTagGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cms/dynamicTagGroup:DynamicTagGroup")
public class DynamicTagGroup extends com.pulumi.resources.CustomResource {
    /**
     * Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
     * 
     */
    @Export(name="contactGroupLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> contactGroupLists;

    /**
     * @return Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.
     * 
     */
    public Output<List<String>> contactGroupLists() {
        return this.contactGroupLists;
    }
    /**
     * The relationship between conditional expressions. Valid values: `and`, `or`.
     * 
     */
    @Export(name="matchExpressFilterRelation", refs={String.class}, tree="[0]")
    private Output<String> matchExpressFilterRelation;

    /**
     * @return The relationship between conditional expressions. Valid values: `and`, `or`.
     * 
     */
    public Output<String> matchExpressFilterRelation() {
        return this.matchExpressFilterRelation;
    }
    /**
     * The label generates a matching expression that applies the grouping. See `match_express` below.
     * 
     */
    @Export(name="matchExpresses", refs={List.class,DynamicTagGroupMatchExpress.class}, tree="[0,1]")
    private Output<List<DynamicTagGroupMatchExpress>> matchExpresses;

    /**
     * @return The label generates a matching expression that applies the grouping. See `match_express` below.
     * 
     */
    public Output<List<DynamicTagGroupMatchExpress>> matchExpresses() {
        return this.matchExpresses;
    }
    /**
     * The status of the resource. Valid values: `RUNNING`, `FINISH`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `RUNNING`, `FINISH`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tag key of the tag.
     * 
     */
    @Export(name="tagKey", refs={String.class}, tree="[0]")
    private Output<String> tagKey;

    /**
     * @return The tag key of the tag.
     * 
     */
    public Output<String> tagKey() {
        return this.tagKey;
    }
    /**
     * Alarm template ID list.
     * 
     */
    @Export(name="templateIdLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> templateIdLists;

    /**
     * @return Alarm template ID list.
     * 
     */
    public Output<Optional<List<String>>> templateIdLists() {
        return Codegen.optional(this.templateIdLists);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DynamicTagGroup(java.lang.String name) {
        this(name, DynamicTagGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DynamicTagGroup(java.lang.String name, DynamicTagGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DynamicTagGroup(java.lang.String name, DynamicTagGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/dynamicTagGroup:DynamicTagGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DynamicTagGroup(java.lang.String name, Output<java.lang.String> id, @Nullable DynamicTagGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/dynamicTagGroup:DynamicTagGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static DynamicTagGroupArgs makeArgs(DynamicTagGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DynamicTagGroupArgs.Empty : args;
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
    public static DynamicTagGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable DynamicTagGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DynamicTagGroup(name, id, state, options);
    }
}
