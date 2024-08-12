// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rocketmq.InstanceArgs;
import com.pulumi.alicloud.rocketmq.inputs.InstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an ONS instance resource.
 * 
 * For more information about how to use it, see [RocketMQ Instance Management API](https://www.alibabacloud.com/help/doc-detail/106354.html).
 * 
 * &gt; **NOTE:** The number of instances in the same region cannot exceed 8. At present, the resource does not support region &#34;mq-internet-access&#34; and &#34;ap-southeast-5&#34;.
 * 
 * &gt; **NOTE:** Available in 1.51.0+
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.rocketmq.Instance;
 * import com.pulumi.alicloud.rocketmq.InstanceArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var example = new Instance("example", InstanceArgs.builder()
 *             .instanceName(String.format("%s-%s", name,default_.result()))
 *             .remark(name)
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
 * ONS INSTANCE can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:rocketmq/instance:Instance instance MQ_INST_1234567890_Baso1234567
 * ```
 * 
 */
@ResourceType(type="alicloud:rocketmq/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
     * 
     */
    @Export(name="instanceStatus", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceStatus;

    /**
     * @return The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
     * 
     */
    public Output<Integer> instanceStatus() {
        return this.instanceStatus;
    }
    /**
     * The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
     * 
     */
    @Export(name="instanceType", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceType;

    /**
     * @return The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
     * 
     */
    public Output<Integer> instanceType() {
        return this.instanceType;
    }
    /**
     * Replaced by `instance_name` after version 1.97.0.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.97.0. Use &#39;instance_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.97.0. Use 'instance_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Replaced by `instance_name` after version 1.97.0.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Platinum edition instance expiration time.
     * 
     */
    @Export(name="releaseTime", refs={String.class}, tree="[0]")
    private Output<String> releaseTime;

    /**
     * @return Platinum edition instance expiration time.
     * 
     */
    public Output<String> releaseTime() {
        return this.releaseTime;
    }
    /**
     * This attribute is a concise description of instance. The length cannot exceed 128.
     * 
     */
    @Export(name="remark", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remark;

    /**
     * @return This attribute is a concise description of instance. The length cannot exceed 128.
     * 
     */
    public Output<Optional<String>> remark() {
        return Codegen.optional(this.remark);
    }
    /**
     * The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
     * 
     */
    @Export(name="status", refs={Integer.class}, tree="[0]")
    private Output<Integer> status;

    /**
     * @return The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(java.lang.String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(java.lang.String name, @Nullable InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(java.lang.String name, @Nullable InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/instance:Instance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Instance(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/instance:Instance", name, state, makeResourceOptions(options, id), false);
    }

    private static InstanceArgs makeArgs(@Nullable InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstanceArgs.Empty : args;
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
    public static Instance get(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
