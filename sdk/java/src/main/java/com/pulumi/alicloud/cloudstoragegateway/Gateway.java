// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudstoragegateway;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudstoragegateway.GatewayArgs;
import com.pulumi.alicloud.cloudstoragegateway.inputs.GatewayState;
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
 * Provides a Cloud Storage Gateway Gateway resource.
 * 
 * For information about Cloud Storage Gateway Gateway and how to use it, see [What is Gateway](https://www.alibabacloud.com/help/en/cloud-storage-gateway/latest/deploygateway).
 * 
 * &gt; **NOTE:** Available since v1.132.0.
 * 
 * ## Import
 * 
 * Cloud Storage Gateway Gateway can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudstoragegateway/gateway:Gateway example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudstoragegateway/gateway:Gateway")
public class Gateway extends com.pulumi.resources.CustomResource {
    /**
     * The description of the gateway.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the gateway.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The specification of the gateway. Valid values: `Basic`, `Standard`,`Enhanced`,`Advanced`.
     * 
     */
    @Export(name="gatewayClass", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> gatewayClass;

    /**
     * @return The specification of the gateway. Valid values: `Basic`, `Standard`,`Enhanced`,`Advanced`.
     * 
     */
    public Output<Optional<String>> gatewayClass() {
        return Codegen.optional(this.gatewayClass);
    }
    /**
     * The name of the gateway.
     * 
     */
    @Export(name="gatewayName", refs={String.class}, tree="[0]")
    private Output<String> gatewayName;

    /**
     * @return The name of the gateway.
     * 
     */
    public Output<String> gatewayName() {
        return this.gatewayName;
    }
    /**
     * The location of the gateway. Valid values: `Cloud`, `On_Premise`.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location of the gateway. Valid values: `Cloud`, `On_Premise`.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The Payment type of gateway. Valid values: `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> paymentType;

    /**
     * @return The Payment type of gateway. Valid values: `PayAsYouGo`.
     * 
     */
    public Output<Optional<String>> paymentType() {
        return Codegen.optional(this.paymentType);
    }
    /**
     * The public network bandwidth of gateway. Default value: `5`. Valid values: `5` to `200`.
     * 
     */
    @Export(name="publicNetworkBandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> publicNetworkBandwidth;

    /**
     * @return The public network bandwidth of gateway. Default value: `5`. Valid values: `5` to `200`.
     * 
     */
    public Output<Integer> publicNetworkBandwidth() {
        return this.publicNetworkBandwidth;
    }
    /**
     * The reason detail of gateway.
     * 
     */
    @Export(name="reasonDetail", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> reasonDetail;

    /**
     * @return The reason detail of gateway.
     * 
     */
    public Output<Optional<String>> reasonDetail() {
        return Codegen.optional(this.reasonDetail);
    }
    /**
     * The reason type when user deletes the gateway.
     * 
     */
    @Export(name="reasonType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> reasonType;

    /**
     * @return The reason type when user deletes the gateway.
     * 
     */
    public Output<Optional<String>> reasonType() {
        return Codegen.optional(this.reasonType);
    }
    /**
     * Whether to release the gateway due to expiration.
     * 
     */
    @Export(name="releaseAfterExpiration", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> releaseAfterExpiration;

    /**
     * @return Whether to release the gateway due to expiration.
     * 
     */
    public Output<Optional<Boolean>> releaseAfterExpiration() {
        return Codegen.optional(this.releaseAfterExpiration);
    }
    /**
     * The status of the Gateway.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Gateway.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the gateway cluster.
     * 
     */
    @Export(name="storageBundleId", refs={String.class}, tree="[0]")
    private Output<String> storageBundleId;

    /**
     * @return The ID of the gateway cluster.
     * 
     */
    public Output<String> storageBundleId() {
        return this.storageBundleId;
    }
    /**
     * The type of the gateway. Valid values: `File`, `Iscsi`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the gateway. Valid values: `File`, `Iscsi`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The ID of the vSwitch.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The ID of the vSwitch.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Gateway(java.lang.String name) {
        this(name, GatewayArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Gateway(java.lang.String name, GatewayArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Gateway(java.lang.String name, GatewayArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudstoragegateway/gateway:Gateway", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Gateway(java.lang.String name, Output<java.lang.String> id, @Nullable GatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudstoragegateway/gateway:Gateway", name, state, makeResourceOptions(options, id), false);
    }

    private static GatewayArgs makeArgs(GatewayArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GatewayArgs.Empty : args;
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
    public static Gateway get(java.lang.String name, Output<java.lang.String> id, @Nullable GatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Gateway(name, id, state, options);
    }
}
