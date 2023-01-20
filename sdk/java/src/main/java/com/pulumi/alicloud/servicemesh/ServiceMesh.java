// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.servicemesh.ServiceMeshArgs;
import com.pulumi.alicloud.servicemesh.inputs.ServiceMeshState;
import com.pulumi.alicloud.servicemesh.outputs.ServiceMeshExtraConfiguration;
import com.pulumi.alicloud.servicemesh.outputs.ServiceMeshLoadBalancer;
import com.pulumi.alicloud.servicemesh.outputs.ServiceMeshMeshConfig;
import com.pulumi.alicloud.servicemesh.outputs.ServiceMeshNetwork;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Service Mesh Service Mesh resource.
 * 
 * For information about Service Mesh Service Mesh and how to use it, see [What is Service Mesh](https://help.aliyun.com/document_detail/171559.html).
 * 
 * &gt; **NOTE:** Available in v1.138.0+.
 * 
 * ## Import
 * 
 * Service Mesh Service Mesh can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:servicemesh/serviceMesh:ServiceMesh example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:servicemesh/serviceMesh:ServiceMesh")
public class ServiceMesh extends com.pulumi.resources.CustomResource {
    /**
     * The array of the cluster ids.
     * 
     */
    @Export(name="clusterIds", type=List.class, parameters={String.class})
    private Output<List<String>> clusterIds;

    /**
     * @return The array of the cluster ids.
     * 
     */
    public Output<List<String>> clusterIds() {
        return this.clusterIds;
    }
    /**
     * The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
     * 
     */
    @Export(name="clusterSpec", type=String.class, parameters={})
    private Output<String> clusterSpec;

    /**
     * @return The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
     * 
     */
    public Output<String> clusterSpec() {
        return this.clusterSpec;
    }
    /**
     * The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
     * 
     */
    @Export(name="edition", type=String.class, parameters={})
    private Output<String> edition;

    /**
     * @return The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
     * 
     */
    public Output<String> edition() {
        return this.edition;
    }
    /**
     * The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
     * 
     */
    @Export(name="extraConfiguration", type=ServiceMeshExtraConfiguration.class, parameters={})
    private Output<ServiceMeshExtraConfiguration> extraConfiguration;

    /**
     * @return The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
     * 
     */
    public Output<ServiceMeshExtraConfiguration> extraConfiguration() {
        return this.extraConfiguration;
    }
    /**
     * This parameter is used for resource destroy. Default value is `false`.
     * 
     */
    @Export(name="force", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return This parameter is used for resource destroy. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * The configuration of the Load Balancer. See the following `Block load_balancer`.
     * 
     */
    @Export(name="loadBalancer", type=ServiceMeshLoadBalancer.class, parameters={})
    private Output<ServiceMeshLoadBalancer> loadBalancer;

    /**
     * @return The configuration of the Load Balancer. See the following `Block load_balancer`.
     * 
     */
    public Output<ServiceMeshLoadBalancer> loadBalancer() {
        return this.loadBalancer;
    }
    /**
     * The configuration of the Service grid. See the following `Block mesh_config`.
     * 
     */
    @Export(name="meshConfig", type=ServiceMeshMeshConfig.class, parameters={})
    private Output<ServiceMeshMeshConfig> meshConfig;

    /**
     * @return The configuration of the Service grid. See the following `Block mesh_config`.
     * 
     */
    public Output<ServiceMeshMeshConfig> meshConfig() {
        return this.meshConfig;
    }
    /**
     * The network configuration of the Service grid. See the following `Block network`.
     * 
     */
    @Export(name="network", type=ServiceMeshNetwork.class, parameters={})
    private Output<ServiceMeshNetwork> network;

    /**
     * @return The network configuration of the Service grid. See the following `Block network`.
     * 
     */
    public Output<ServiceMeshNetwork> network() {
        return this.network;
    }
    /**
     * The name of the resource.
     * 
     */
    @Export(name="serviceMeshName", type=String.class, parameters={})
    private Output<String> serviceMeshName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> serviceMeshName() {
        return this.serviceMeshName;
    }
    /**
     * The status of the resource. Valid values: `running` or `initial`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `running` or `initial`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The version of the resource. you can look up the version using `alicloud.servicemesh.getVersions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `alicloud.servicemesh.getServiceMeshes`.
     * 
     */
    @Export(name="version", type=String.class, parameters={})
    private Output<String> version;

    /**
     * @return The version of the resource. you can look up the version using `alicloud.servicemesh.getVersions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `alicloud.servicemesh.getServiceMeshes`.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceMesh(String name) {
        this(name, ServiceMeshArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceMesh(String name, ServiceMeshArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceMesh(String name, ServiceMeshArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:servicemesh/serviceMesh:ServiceMesh", name, args == null ? ServiceMeshArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceMesh(String name, Output<String> id, @Nullable ServiceMeshState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:servicemesh/serviceMesh:ServiceMesh", name, state, makeResourceOptions(options, id));
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
    public static ServiceMesh get(String name, Output<String> id, @Nullable ServiceMeshState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceMesh(name, id, state, options);
    }
}
