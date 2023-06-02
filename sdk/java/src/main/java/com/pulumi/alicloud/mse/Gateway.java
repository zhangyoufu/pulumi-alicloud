// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mse;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mse.GatewayArgs;
import com.pulumi.alicloud.mse.inputs.GatewayState;
import com.pulumi.alicloud.mse.outputs.GatewaySlbList;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Microservice Engine (MSE) Gateway resource.
 * 
 * For information about Microservice Engine (MSE) Gateway and how to use it, see [What is Gateway](https://help.aliyun.com/document_detail/347638.html).
 * 
 * &gt; **NOTE:** Available in v1.157.0+.
 * 
 * ## Import
 * 
 * Microservice Engine (MSE) Gateway can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:mse/gateway:Gateway example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:mse/gateway:Gateway")
public class Gateway extends com.pulumi.resources.CustomResource {
    /**
     * The backup vswitch id.
     * 
     */
    @Export(name="backupVswitchId", type=String.class, parameters={})
    private Output</* @Nullable */ String> backupVswitchId;

    /**
     * @return The backup vswitch id.
     * 
     */
    public Output<Optional<String>> backupVswitchId() {
        return Codegen.optional(this.backupVswitchId);
    }
    /**
     * Whether to delete the SLB purchased on behalf of the gateway at the same time.
     * 
     */
    @Export(name="deleteSlb", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> deleteSlb;

    /**
     * @return Whether to delete the SLB purchased on behalf of the gateway at the same time.
     * 
     */
    public Output<Optional<Boolean>> deleteSlb() {
        return Codegen.optional(this.deleteSlb);
    }
    /**
     * Whether the enterprise security group type.
     * 
     */
    @Export(name="enterpriseSecurityGroup", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enterpriseSecurityGroup;

    /**
     * @return Whether the enterprise security group type.
     * 
     */
    public Output<Optional<Boolean>> enterpriseSecurityGroup() {
        return Codegen.optional(this.enterpriseSecurityGroup);
    }
    /**
     * The name of the Gateway .
     * 
     */
    @Export(name="gatewayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> gatewayName;

    /**
     * @return The name of the Gateway .
     * 
     */
    public Output<Optional<String>> gatewayName() {
        return Codegen.optional(this.gatewayName);
    }
    /**
     * Public network SLB specifications.
     * 
     */
    @Export(name="internetSlbSpec", type=String.class, parameters={})
    private Output</* @Nullable */ String> internetSlbSpec;

    /**
     * @return Public network SLB specifications.
     * 
     */
    public Output<Optional<String>> internetSlbSpec() {
        return Codegen.optional(this.internetSlbSpec);
    }
    /**
     * Number of Gateway Nodes.
     * 
     */
    @Export(name="replica", type=Integer.class, parameters={})
    private Output<Integer> replica;

    /**
     * @return Number of Gateway Nodes.
     * 
     */
    public Output<Integer> replica() {
        return this.replica;
    }
    /**
     * A list of gateway Slb.
     * 
     */
    @Export(name="slbLists", type=List.class, parameters={GatewaySlbList.class})
    private Output<List<GatewaySlbList>> slbLists;

    /**
     * @return A list of gateway Slb.
     * 
     */
    public Output<List<GatewaySlbList>> slbLists() {
        return this.slbLists;
    }
    /**
     * Private network SLB specifications.
     * 
     */
    @Export(name="slbSpec", type=String.class, parameters={})
    private Output</* @Nullable */ String> slbSpec;

    /**
     * @return Private network SLB specifications.
     * 
     */
    public Output<Optional<String>> slbSpec() {
        return Codegen.optional(this.slbSpec);
    }
    /**
     * Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
     * 
     */
    @Export(name="spec", type=String.class, parameters={})
    private Output<String> spec;

    /**
     * @return Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
     * 
     */
    public Output<String> spec() {
        return this.spec;
    }
    /**
     * The status of the gateway.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the gateway.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the vpc.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the vpc.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The ID of the vswitch.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output<String> vswitchId;

    /**
     * @return The ID of the vswitch.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Gateway(String name) {
        this(name, GatewayArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Gateway(String name, GatewayArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Gateway(String name, GatewayArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mse/gateway:Gateway", name, args == null ? GatewayArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Gateway(String name, Output<String> id, @Nullable GatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mse/gateway:Gateway", name, state, makeResourceOptions(options, id));
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
    public static Gateway get(String name, Output<String> id, @Nullable GatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Gateway(name, id, state, options);
    }
}
