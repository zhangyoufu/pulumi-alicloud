// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.actiontrail;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.actiontrail.TrailArgs;
import com.pulumi.alicloud.actiontrail.inputs.TrailState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ActionTrail Trail resource. For information about alicloud actiontrail trail and how to use it, see [What is Resource Alicloud ActionTrail Trail](https://www.alibabacloud.com/help/en/actiontrail/latest/api-actiontrail-2020-07-06-createtrail).
 * 
 * &gt; **NOTE:** Available since v1.95.0.
 * 
 * &gt; **NOTE:** You can create a trail to deliver events to Log Service, Object Storage Service (OSS), or both. Before you call this operation to create a trail, make sure that the following requirements are met.
 * - Deliver events to Log Service: A project is created in Log Service.
 * - Deliver events to OSS: A bucket is created in OSS.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.ram.RamFunctions;
 * import com.pulumi.alicloud.ram.inputs.GetRolesArgs;
 * import com.pulumi.alicloud.actiontrail.Trail;
 * import com.pulumi.alicloud.actiontrail.TrailArgs;
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
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var exampleRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         final var exampleAccount = AlicloudFunctions.getAccount();
 * 
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;tf actiontrail example&#34;)
 *             .build());
 * 
 *         final var exampleRoles = RamFunctions.getRoles(GetRolesArgs.builder()
 *             .nameRegex(&#34;AliyunServiceRoleForActionTrail&#34;)
 *             .build());
 * 
 *         var exampleTrail = new Trail(&#34;exampleTrail&#34;, TrailArgs.builder()        
 *             .trailName(name)
 *             .slsWriteRoleArn(exampleRoles.applyValue(getRolesResult -&gt; getRolesResult.roles()[0].arn()))
 *             .slsProjectArn(exampleProject.name().applyValue(name -&gt; String.format(&#34;acs:log:%s:%s:project/%s&#34;, exampleRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()),exampleAccount.applyValue(getAccountResult -&gt; getAccountResult.id()),name)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Action trail can be imported using the id or trail_name, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:actiontrail/trail:Trail default abc12345678
 * ```
 * 
 */
@ResourceType(type="alicloud:actiontrail/trail:Trail")
public class Trail extends com.pulumi.resources.CustomResource {
    /**
     * Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
     * 
     */
    @Export(name="eventRw", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eventRw;

    /**
     * @return Indicates whether the event is a read or a write event. Valid values: `Read`, `Write`, and `All`. Default to `Write`.
     * 
     */
    public Output<Optional<String>> eventRw() {
        return Codegen.optional(this.eventRw);
    }
    /**
     * Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
     * 
     */
    @Export(name="isOrganizationTrail", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isOrganizationTrail;

    /**
     * @return Specifies whether to create a multi-account trail. Valid values:`true`: Create a multi-account trail.`false`: Create a single-account trail. It is the default value.
     * 
     */
    public Output<Optional<Boolean>> isOrganizationTrail() {
        return Codegen.optional(this.isOrganizationTrail);
    }
    /**
     * Field `mns_topic_arn` has been deprecated from version 1.118.0.
     * 
     * @deprecated
     * Field &#39;mns_topic_arn&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'mns_topic_arn' has been deprecated from version 1.118.0 */
    @Export(name="mnsTopicArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mnsTopicArn;

    /**
     * @return Field `mns_topic_arn` has been deprecated from version 1.118.0.
     * 
     */
    public Output<Optional<String>> mnsTopicArn() {
        return Codegen.optional(this.mnsTopicArn);
    }
    /**
     * Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.95.0. Use &#39;trail_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Field `name` has been deprecated from version 1.95.0. Use `trail_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     * 
     */
    @Export(name="ossBucketName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ossBucketName;

    /**
     * @return The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     * 
     */
    public Output<Optional<String>> ossBucketName() {
        return Codegen.optional(this.ossBucketName);
    }
    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     * 
     */
    @Export(name="ossKeyPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ossKeyPrefix;

    /**
     * @return The prefix of the specified OSS bucket name. This parameter can be left empty.
     * 
     */
    public Output<Optional<String>> ossKeyPrefix() {
        return Codegen.optional(this.ossKeyPrefix);
    }
    /**
     * The unique ARN of the Oss role.
     * 
     */
    @Export(name="ossWriteRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ossWriteRoleArn;

    /**
     * @return The unique ARN of the Oss role.
     * 
     */
    public Output<Optional<String>> ossWriteRoleArn() {
        return Codegen.optional(this.ossWriteRoleArn);
    }
    /**
     * Field `name` has been deprecated from version 1.118.0.
     * 
     * @deprecated
     * Field &#39;role_name&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'role_name' has been deprecated from version 1.118.0 */
    @Export(name="roleName", refs={String.class}, tree="[0]")
    private Output<String> roleName;

    /**
     * @return Field `name` has been deprecated from version 1.118.0.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }
    /**
     * The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
     * 
     */
    @Export(name="slsProjectArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> slsProjectArn;

    /**
     * @return The unique ARN of the Log Service project. Ensure that `sls_project_arn` is valid .
     * 
     */
    public Output<Optional<String>> slsProjectArn() {
        return Codegen.optional(this.slsProjectArn);
    }
    /**
     * The unique ARN of the Log Service role.
     * 
     */
    @Export(name="slsWriteRoleArn", refs={String.class}, tree="[0]")
    private Output<String> slsWriteRoleArn;

    /**
     * @return The unique ARN of the Log Service role.
     * 
     */
    public Output<String> slsWriteRoleArn() {
        return this.slsWriteRoleArn;
    }
    /**
     * The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    /**
     * @return The status of ActionTrail Trail. After creation, tracking is turned on by default, and you can set the status value to `Disable` to turn off tracking. Valid values: `Enable`, `Disable`. Default to `Enable`.
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    /**
     * The name of the trail to be created, which must be unique for an account.
     * 
     */
    @Export(name="trailName", refs={String.class}, tree="[0]")
    private Output<String> trailName;

    /**
     * @return The name of the trail to be created, which must be unique for an account.
     * 
     */
    public Output<String> trailName() {
        return this.trailName;
    }
    /**
     * The regions to which the trail is applied. Default to `All`.
     * 
     */
    @Export(name="trailRegion", refs={String.class}, tree="[0]")
    private Output<String> trailRegion;

    /**
     * @return The regions to which the trail is applied. Default to `All`.
     * 
     */
    public Output<String> trailRegion() {
        return this.trailRegion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Trail(String name) {
        this(name, TrailArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Trail(String name, @Nullable TrailArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Trail(String name, @Nullable TrailArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/trail:Trail", name, args == null ? TrailArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Trail(String name, Output<String> id, @Nullable TrailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/trail:Trail", name, state, makeResourceOptions(options, id));
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
    public static Trail get(String name, Output<String> id, @Nullable TrailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Trail(name, id, state, options);
    }
}
