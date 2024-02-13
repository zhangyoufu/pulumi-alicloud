// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.simpleapplicationserver;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.simpleapplicationserver.InstanceArgs;
import com.pulumi.alicloud.simpleapplicationserver.inputs.InstanceState;
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
 * Simple Application Server Instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:simpleapplicationserver/instance:Instance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:simpleapplicationserver/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Specifies whether to enable auto-renewal. Unit: months. Valid values: `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
     * 
     */
    @Export(name="autoRenewPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> autoRenewPeriod;

    /**
     * @return The auto renew period. Valid values: `1`,`3`, `6`, `12`, `24`, `36`. **NOTE:** The attribute `auto_renew` is valid when the attribute is `true`.
     * 
     */
    public Output<Optional<Integer>> autoRenewPeriod() {
        return Codegen.optional(this.autoRenewPeriod);
    }
    /**
     * The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
     * 
     */
    @Export(name="dataDiskSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> dataDiskSize;

    /**
     * @return The size of the data disk. Unit: GB. Valid values: `0` to `16380`.
     * 
     */
    public Output<Optional<Integer>> dataDiskSize() {
        return Codegen.optional(this.dataDiskSize);
    }
    /**
     * The ID of the image.  You can use the `alicloud.simpleapplicationserver.getImages` to query the available images in the specified region. The value must be an integral multiple of 20.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return The ID of the image.  You can use the `alicloud.simpleapplicationserver.getImages` to query the available images in the specified region. The value must be an integral multiple of 20.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The name of the simple application server.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceName;

    /**
     * @return The name of the simple application server.
     * 
     */
    public Output<Optional<String>> instanceName() {
        return Codegen.optional(this.instanceName);
    }
    /**
     * The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ &amp; * - + = | { } [ ] : ; &lt; &gt; , . ? /`.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include: `( ) ~ ! @ # $ % ^ &amp; * - + = | { } [ ] : ; &lt; &gt; , . ? /`.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The paymen type of the resource. Valid values: `Subscription`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The paymen type of the resource. Valid values: `Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output<Integer> period;

    /**
     * @return The period. Unit: months. Valid values: `1`,`3`, `6`, `12`, `24`, `36`.
     * 
     */
    public Output<Integer> period() {
        return this.period;
    }
    /**
     * The ID of the plan. You can use the `alicloud.simpleapplicationserver.getServerPlans`  to query all the plans provided by Simple Application Server in the specified region.
     * 
     */
    @Export(name="planId", refs={String.class}, tree="[0]")
    private Output<String> planId;

    /**
     * @return The ID of the plan. You can use the `alicloud.simpleapplicationserver.getServerPlans`  to query all the plans provided by Simple Application Server in the specified region.
     * 
     */
    public Output<String> planId() {
        return this.planId;
    }
    /**
     * The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the simple application server. Valid values: `Resetting`, `Running`, `Stopped`.
     * 
     */
    public Output<String> status() {
        return this.status;
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
        super("alicloud:simpleapplicationserver/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:simpleapplicationserver/instance:Instance", name, state, makeResourceOptions(options, id));
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
