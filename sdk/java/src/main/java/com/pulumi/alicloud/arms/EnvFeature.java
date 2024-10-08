// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.EnvFeatureArgs;
import com.pulumi.alicloud.arms.inputs.EnvFeatureState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a ARMS Env Feature resource. Feature of the arms environment.
 * 
 * For information about ARMS Env Feature and how to use it, see [What is Env Feature](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-installenvironmentfeature).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
 * 
 * ## Import
 * 
 * ARMS Env Feature can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:arms/envFeature:EnvFeature example &lt;environment_id&gt;:&lt;env_feature_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:arms/envFeature:EnvFeature")
public class EnvFeature extends com.pulumi.resources.CustomResource {
    /**
     * The name of the resource.
     * 
     */
    @Export(name="envFeatureName", refs={String.class}, tree="[0]")
    private Output<String> envFeatureName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> envFeatureName() {
        return this.envFeatureName;
    }
    /**
     * The first ID of the resource.
     * 
     */
    @Export(name="environmentId", refs={String.class}, tree="[0]")
    private Output<String> environmentId;

    /**
     * @return The first ID of the resource.
     * 
     */
    public Output<String> environmentId() {
        return this.environmentId;
    }
    /**
     * Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
     * 
     */
    @Export(name="featureVersion", refs={String.class}, tree="[0]")
    private Output<String> featureVersion;

    /**
     * @return Version information of the Feature. You can query Feature information by using ListEnvironmentFeatures.
     * 
     */
    public Output<String> featureVersion() {
        return this.featureVersion;
    }
    /**
     * Namespace.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output<String> namespace;

    /**
     * @return Namespace.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
    }
    /**
     * Status.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnvFeature(java.lang.String name) {
        this(name, EnvFeatureArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnvFeature(java.lang.String name, EnvFeatureArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnvFeature(java.lang.String name, EnvFeatureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/envFeature:EnvFeature", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EnvFeature(java.lang.String name, Output<java.lang.String> id, @Nullable EnvFeatureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/envFeature:EnvFeature", name, state, makeResourceOptions(options, id), false);
    }

    private static EnvFeatureArgs makeArgs(EnvFeatureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EnvFeatureArgs.Empty : args;
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
    public static EnvFeature get(java.lang.String name, Output<java.lang.String> id, @Nullable EnvFeatureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnvFeature(name, id, state, options);
    }
}
