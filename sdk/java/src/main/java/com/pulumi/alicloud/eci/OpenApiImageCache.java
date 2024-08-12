// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eci.OpenApiImageCacheArgs;
import com.pulumi.alicloud.eci.inputs.OpenApiImageCacheState;
import com.pulumi.alicloud.eci.outputs.OpenApiImageCacheImageRegistryCredential;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="alicloud:eci/openApiImageCache:OpenApiImageCache")
public class OpenApiImageCache extends com.pulumi.resources.CustomResource {
    @Export(name="containerGroupId", refs={String.class}, tree="[0]")
    private Output<String> containerGroupId;

    public Output<String> containerGroupId() {
        return this.containerGroupId;
    }
    @Export(name="eipInstanceId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eipInstanceId;

    public Output<Optional<String>> eipInstanceId() {
        return Codegen.optional(this.eipInstanceId);
    }
    @Export(name="imageCacheName", refs={String.class}, tree="[0]")
    private Output<String> imageCacheName;

    public Output<String> imageCacheName() {
        return this.imageCacheName;
    }
    @Export(name="imageCacheSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> imageCacheSize;

    public Output<Optional<Integer>> imageCacheSize() {
        return Codegen.optional(this.imageCacheSize);
    }
    @Export(name="imageRegistryCredentials", refs={List.class,OpenApiImageCacheImageRegistryCredential.class}, tree="[0,1]")
    private Output</* @Nullable */ List<OpenApiImageCacheImageRegistryCredential>> imageRegistryCredentials;

    public Output<Optional<List<OpenApiImageCacheImageRegistryCredential>>> imageRegistryCredentials() {
        return Codegen.optional(this.imageRegistryCredentials);
    }
    @Export(name="images", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> images;

    public Output<List<String>> images() {
        return this.images;
    }
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    @Export(name="retentionDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> retentionDays;

    public Output<Optional<Integer>> retentionDays() {
        return Codegen.optional(this.retentionDays);
    }
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> zoneId;

    public Output<Optional<String>> zoneId() {
        return Codegen.optional(this.zoneId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OpenApiImageCache(java.lang.String name) {
        this(name, OpenApiImageCacheArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OpenApiImageCache(java.lang.String name, OpenApiImageCacheArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OpenApiImageCache(java.lang.String name, OpenApiImageCacheArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eci/openApiImageCache:OpenApiImageCache", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private OpenApiImageCache(java.lang.String name, Output<java.lang.String> id, @Nullable OpenApiImageCacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eci/openApiImageCache:OpenApiImageCache", name, state, makeResourceOptions(options, id), false);
    }

    private static OpenApiImageCacheArgs makeArgs(OpenApiImageCacheArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? OpenApiImageCacheArgs.Empty : args;
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
    public static OpenApiImageCache get(java.lang.String name, Output<java.lang.String> id, @Nullable OpenApiImageCacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OpenApiImageCache(name, id, state, options);
    }
}
