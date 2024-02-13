// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.amqp;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.amqp.InstanceArgs;
import com.pulumi.alicloud.amqp.inputs.InstanceState;
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
 * ## Import
 * 
 * RabbitMQ (AMQP) Instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:amqp/instance:Instance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:amqp/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * The instance name.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return The instance name.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    /**
     * @return The Instance Type. Valid values: `professional`, `enterprise`, `vip`.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The logistic information This parameter is not required when you create a ApsaraMQ for RabbitMQ instance. You do not need to specify this parameter.
     * 
     */
    @Export(name="logistics", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logistics;

    /**
     * @return The logistic information This parameter is not required when you create a ApsaraMQ for RabbitMQ instance. You do not need to specify this parameter.
     * 
     */
    public Output<Optional<String>> logistics() {
        return Codegen.optional(this.logistics);
    }
    /**
     * The max eip tps. It is valid when `support_eip` is true. The valid value is [128, 45000] with the step size 128.
     * 
     */
    @Export(name="maxEipTps", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> maxEipTps;

    /**
     * @return The max eip tps. It is valid when `support_eip` is true. The valid value is [128, 45000] with the step size 128.
     * 
     */
    public Output<Optional<String>> maxEipTps() {
        return Codegen.optional(this.maxEipTps);
    }
    /**
     * The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
     * 
     */
    @Export(name="maxTps", refs={String.class}, tree="[0]")
    private Output<String> maxTps;

    /**
     * @return The peak TPS traffic. The smallest valid value is 1000 and the largest value is 100,000.
     * 
     */
    public Output<String> maxTps() {
        return this.maxTps;
    }
    /**
     * The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
     * 
     */
    @Export(name="modifyType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> modifyType;

    /**
     * @return The modify type. Valid values: `Downgrade`, `Upgrade`. It is required when updating other attributes.
     * 
     */
    public Output<Optional<String>> modifyType() {
        return Codegen.optional(this.modifyType);
    }
    /**
     * The payment type. Valid values: `Subscription`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The payment type. Valid values: `Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The period. Valid values: `1`, `12`, `2`, `24`, `3`, `6`.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The queue capacity. The smallest value is 50 and the step size 5.
     * 
     */
    @Export(name="queueCapacity", refs={String.class}, tree="[0]")
    private Output<String> queueCapacity;

    /**
     * @return The queue capacity. The smallest value is 50 and the step size 5.
     * 
     */
    public Output<String> queueCapacity() {
        return this.queueCapacity;
    }
    /**
     * RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
     * 
     */
    @Export(name="renewalDuration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> renewalDuration;

    /**
     * @return RenewalDuration. Valid values: `1`, `12`, `2`, `3`, `6`.
     * 
     */
    public Output<Optional<Integer>> renewalDuration() {
        return Codegen.optional(this.renewalDuration);
    }
    /**
     * Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     * 
     */
    @Export(name="renewalDurationUnit", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> renewalDurationUnit;

    /**
     * @return Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
     * 
     */
    public Output<Optional<String>> renewalDurationUnit() {
        return Codegen.optional(this.renewalDurationUnit);
    }
    /**
     * Whether to renew an instance automatically or not. Default to &#34;ManualRenewal&#34;.
     * 
     */
    @Export(name="renewalStatus", refs={String.class}, tree="[0]")
    private Output<String> renewalStatus;

    /**
     * @return Whether to renew an instance automatically or not. Default to &#34;ManualRenewal&#34;.
     * 
     */
    public Output<String> renewalStatus() {
        return this.renewalStatus;
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
     * The storage size. It is valid when `instance_type` is vip.
     * 
     */
    @Export(name="storageSize", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> storageSize;

    /**
     * @return The storage size. It is valid when `instance_type` is vip.
     * 
     */
    public Output<Optional<String>> storageSize() {
        return Codegen.optional(this.storageSize);
    }
    /**
     * Whether to support EIP.
     * 
     */
    @Export(name="supportEip", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> supportEip;

    /**
     * @return Whether to support EIP.
     * 
     */
    public Output<Boolean> supportEip() {
        return this.supportEip;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:amqp/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:amqp/instance:Instance", name, state, makeResourceOptions(options, id));
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
