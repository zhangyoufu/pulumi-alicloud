// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cfg.DeliveryArgs;
import com.pulumi.alicloud.cfg.inputs.DeliveryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Config Delivery resource.
 * 
 * For information about Cloud Config Delivery and how to use it, see [What is Delivery](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createconfigdeliverychannel).
 * 
 * &gt; **NOTE:** Available since v1.171.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.Store;
 * import com.pulumi.alicloud.log.StoreArgs;
 * import com.pulumi.alicloud.cfg.Delivery;
 * import com.pulumi.alicloud.cfg.DeliveryArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example-sls&#34;);
 *         final var thisAccount = AlicloudFunctions.getAccount();
 * 
 *         final var thisRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultProject = new Project(&#34;defaultProject&#34;);
 * 
 *         var defaultStore = new Store(&#34;defaultStore&#34;, StoreArgs.builder()        
 *             .project(defaultProject.name())
 *             .build());
 * 
 *         var defaultDelivery = new Delivery(&#34;defaultDelivery&#34;, DeliveryArgs.builder()        
 *             .configurationItemChangeNotification(true)
 *             .nonCompliantNotification(true)
 *             .deliveryChannelName(name)
 *             .deliveryChannelTargetArn(Output.tuple(defaultProject.name(), defaultStore.name()).applyValue(values -&gt; {
 *                 var defaultProjectName = values.t1;
 *                 var defaultStoreName = values.t2;
 *                 return String.format(&#34;acs:log:%s:%s:project/%s/logstore/%s&#34;, thisRegions.applyValue(getRegionsResult -&gt; getRegionsResult.ids()[0]),thisAccount.applyValue(getAccountResult -&gt; getAccountResult.id()),defaultProjectName,defaultStoreName);
 *             }))
 *             .deliveryChannelType(&#34;SLS&#34;)
 *             .description(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud Config Delivery can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cfg/delivery:Delivery example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cfg/delivery:Delivery")
public class Delivery extends com.pulumi.resources.CustomResource {
    /**
     * Open or close delivery configuration change history. true: open, false: close.
     * 
     */
    @Export(name="configurationItemChangeNotification", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> configurationItemChangeNotification;

    /**
     * @return Open or close delivery configuration change history. true: open, false: close.
     * 
     */
    public Output<Boolean> configurationItemChangeNotification() {
        return this.configurationItemChangeNotification;
    }
    /**
     * Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
     * 
     */
    @Export(name="configurationSnapshot", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> configurationSnapshot;

    /**
     * @return Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
     * 
     */
    public Output<Boolean> configurationSnapshot() {
        return this.configurationSnapshot;
    }
    /**
     * The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
     * 
     */
    @Export(name="deliveryChannelCondition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deliveryChannelCondition;

    /**
     * @return The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
     * 
     */
    public Output<Optional<String>> deliveryChannelCondition() {
        return Codegen.optional(this.deliveryChannelCondition);
    }
    /**
     * The name of the delivery method.
     * 
     */
    @Export(name="deliveryChannelName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deliveryChannelName;

    /**
     * @return The name of the delivery method.
     * 
     */
    public Output<Optional<String>> deliveryChannelName() {
        return Codegen.optional(this.deliveryChannelName);
    }
    /**
     * The ARN of the delivery destination. The value must be in one of the following formats:
     * * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
     * * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
     * * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
     * 
     */
    @Export(name="deliveryChannelTargetArn", refs={String.class}, tree="[0]")
    private Output<String> deliveryChannelTargetArn;

    /**
     * @return The ARN of the delivery destination. The value must be in one of the following formats:
     * * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
     * * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
     * * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
     * 
     */
    public Output<String> deliveryChannelTargetArn() {
        return this.deliveryChannelTargetArn;
    }
    /**
     * The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
     * 
     */
    @Export(name="deliveryChannelType", refs={String.class}, tree="[0]")
    private Output<String> deliveryChannelType;

    /**
     * @return The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
     * 
     */
    public Output<String> deliveryChannelType() {
        return this.deliveryChannelType;
    }
    /**
     * The description of the delivery method.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the delivery method.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
     * 
     */
    @Export(name="nonCompliantNotification", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> nonCompliantNotification;

    /**
     * @return Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
     * 
     */
    public Output<Boolean> nonCompliantNotification() {
        return this.nonCompliantNotification;
    }
    /**
     * The oss ARN of the delivery channel when the value data oversized limit.
     * * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
     * * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
     * 
     */
    @Export(name="oversizedDataOssTargetArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oversizedDataOssTargetArn;

    /**
     * @return The oss ARN of the delivery channel when the value data oversized limit.
     * * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
     * * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
     * 
     */
    public Output<Optional<String>> oversizedDataOssTargetArn() {
        return Codegen.optional(this.oversizedDataOssTargetArn);
    }
    /**
     * The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
     * 
     */
    @Export(name="status", refs={Integer.class}, tree="[0]")
    private Output<Integer> status;

    /**
     * @return The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Delivery(String name) {
        this(name, DeliveryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Delivery(String name, DeliveryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Delivery(String name, DeliveryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/delivery:Delivery", name, args == null ? DeliveryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Delivery(String name, Output<String> id, @Nullable DeliveryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/delivery:Delivery", name, state, makeResourceOptions(options, id));
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
    public static Delivery get(String name, Output<String> id, @Nullable DeliveryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Delivery(name, id, state, options);
    }
}
