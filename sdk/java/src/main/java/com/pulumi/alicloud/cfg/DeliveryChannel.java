// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cfg.DeliveryChannelArgs;
import com.pulumi.alicloud.cfg.inputs.DeliveryChannelState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ## Import
 * 
 * Alicloud Config Delivery Channel can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cfg/deliveryChannel:DeliveryChannel example cdc-49a2ad756057********
 * ```
 * 
 */
@ResourceType(type="alicloud:cfg/deliveryChannel:DeliveryChannel")
public class DeliveryChannel extends com.pulumi.resources.CustomResource {
    /**
     * The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
     * 
     */
    @Export(name="deliveryChannelAssumeRoleArn", refs={String.class}, tree="[0]")
    private Output<String> deliveryChannelAssumeRoleArn;

    /**
     * @return The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
     * 
     */
    public Output<String> deliveryChannelAssumeRoleArn() {
        return this.deliveryChannelAssumeRoleArn;
    }
    /**
     * The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
     * 
     */
    @Export(name="deliveryChannelCondition", refs={String.class}, tree="[0]")
    private Output<String> deliveryChannelCondition;

    /**
     * @return The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
     * 
     */
    public Output<String> deliveryChannelCondition() {
        return this.deliveryChannelCondition;
    }
    /**
     * The name of the delivery channel.
     * 
     */
    @Export(name="deliveryChannelName", refs={String.class}, tree="[0]")
    private Output<String> deliveryChannelName;

    /**
     * @return The name of the delivery channel.
     * 
     */
    public Output<String> deliveryChannelName() {
        return this.deliveryChannelName;
    }
    /**
     * The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
     * - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
     * - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
     * - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
     * 
     */
    @Export(name="deliveryChannelTargetArn", refs={String.class}, tree="[0]")
    private Output<String> deliveryChannelTargetArn;

    /**
     * @return The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
     * - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
     * - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
     * - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
     * 
     */
    public Output<String> deliveryChannelTargetArn() {
        return this.deliveryChannelTargetArn;
    }
    /**
     * The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
     * 
     */
    @Export(name="deliveryChannelType", refs={String.class}, tree="[0]")
    private Output<String> deliveryChannelType;

    /**
     * @return The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
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
    private Output<String> description;

    /**
     * @return The description of the delivery method.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
     * 
     */
    @Export(name="status", refs={Integer.class}, tree="[0]")
    private Output<Integer> status;

    /**
     * @return The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DeliveryChannel(String name) {
        this(name, DeliveryChannelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DeliveryChannel(String name, DeliveryChannelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DeliveryChannel(String name, DeliveryChannelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/deliveryChannel:DeliveryChannel", name, args == null ? DeliveryChannelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DeliveryChannel(String name, Output<String> id, @Nullable DeliveryChannelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/deliveryChannel:DeliveryChannel", name, state, makeResourceOptions(options, id));
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
    public static DeliveryChannel get(String name, Output<String> id, @Nullable DeliveryChannelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DeliveryChannel(name, id, state, options);
    }
}
