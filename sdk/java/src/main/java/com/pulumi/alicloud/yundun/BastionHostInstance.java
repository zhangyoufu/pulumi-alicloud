// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.yundun;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.yundun.BastionHostInstanceArgs;
import com.pulumi.alicloud.yundun.inputs.BastionHostInstanceState;
import com.pulumi.alicloud.yundun.outputs.BastionHostInstanceAdAuthServer;
import com.pulumi.alicloud.yundun.outputs.BastionHostInstanceLdapAuthServer;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="alicloud:yundun/bastionHostInstance:BastionHostInstance")
public class BastionHostInstance extends com.pulumi.resources.CustomResource {
    @Export(name="adAuthServers", refs={List.class,BastionHostInstanceAdAuthServer.class}, tree="[0,1]")
    private Output<List<BastionHostInstanceAdAuthServer>> adAuthServers;

    public Output<List<BastionHostInstanceAdAuthServer>> adAuthServers() {
        return this.adAuthServers;
    }
    @Export(name="bandwidth", refs={String.class}, tree="[0]")
    private Output<String> bandwidth;

    public Output<String> bandwidth() {
        return this.bandwidth;
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    public Output<String> description() {
        return this.description;
    }
    @Export(name="enablePublicAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enablePublicAccess;

    public Output<Boolean> enablePublicAccess() {
        return this.enablePublicAccess;
    }
    @Export(name="ldapAuthServers", refs={List.class,BastionHostInstanceLdapAuthServer.class}, tree="[0,1]")
    private Output<List<BastionHostInstanceLdapAuthServer>> ldapAuthServers;

    public Output<List<BastionHostInstanceLdapAuthServer>> ldapAuthServers() {
        return this.ldapAuthServers;
    }
    @Export(name="licenseCode", refs={String.class}, tree="[0]")
    private Output<String> licenseCode;

    public Output<String> licenseCode() {
        return this.licenseCode;
    }
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    @Export(name="planCode", refs={String.class}, tree="[0]")
    private Output<String> planCode;

    public Output<String> planCode() {
        return this.planCode;
    }
    @Export(name="publicWhiteLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> publicWhiteLists;

    public Output<Optional<List<String>>> publicWhiteLists() {
        return Codegen.optional(this.publicWhiteLists);
    }
    @Export(name="renewPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> renewPeriod;

    public Output<Optional<Integer>> renewPeriod() {
        return Codegen.optional(this.renewPeriod);
    }
    @Export(name="renewalPeriodUnit", refs={String.class}, tree="[0]")
    private Output<String> renewalPeriodUnit;

    public Output<String> renewalPeriodUnit() {
        return this.renewalPeriodUnit;
    }
    @Export(name="renewalStatus", refs={String.class}, tree="[0]")
    private Output<String> renewalStatus;

    public Output<String> renewalStatus() {
        return this.renewalStatus;
    }
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroupIds;

    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    @Export(name="storage", refs={String.class}, tree="[0]")
    private Output<String> storage;

    public Output<String> storage() {
        return this.storage;
    }
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BastionHostInstance(java.lang.String name) {
        this(name, BastionHostInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BastionHostInstance(java.lang.String name, BastionHostInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BastionHostInstance(java.lang.String name, BastionHostInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:yundun/bastionHostInstance:BastionHostInstance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BastionHostInstance(java.lang.String name, Output<java.lang.String> id, @Nullable BastionHostInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:yundun/bastionHostInstance:BastionHostInstance", name, state, makeResourceOptions(options, id), false);
    }

    private static BastionHostInstanceArgs makeArgs(BastionHostInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BastionHostInstanceArgs.Empty : args;
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
    public static BastionHostInstance get(java.lang.String name, Output<java.lang.String> id, @Nullable BastionHostInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BastionHostInstance(name, id, state, options);
    }
}
