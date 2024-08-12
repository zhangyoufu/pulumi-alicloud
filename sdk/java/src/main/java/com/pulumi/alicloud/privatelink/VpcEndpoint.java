// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.privatelink.VpcEndpointArgs;
import com.pulumi.alicloud.privatelink.inputs.VpcEndpointState;
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

/**
 * Provides a Private Link Vpc Endpoint resource.
 * 
 * For information about Private Link Vpc Endpoint and how to use it, see [What is Vpc Endpoint](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-createvpcendpoint).
 * 
 * &gt; **NOTE:** Available since v1.109.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.privatelink.VpcEndpoint;
 * import com.pulumi.alicloud.privatelink.VpcEndpointArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         final var default = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultbFzA4a = new Network("defaultbFzA4a", NetworkArgs.builder()
 *             .description("example-terraform")
 *             .cidrBlock("172.16.0.0/12")
 *             .vpcName(name)
 *             .build());
 * 
 *         var default1FTFrP = new SecurityGroup("default1FTFrP", SecurityGroupArgs.builder()
 *             .name(name)
 *             .vpcId(defaultbFzA4a.id())
 *             .build());
 * 
 *         var defaultjljY5S = new SecurityGroup("defaultjljY5S", SecurityGroupArgs.builder()
 *             .name(name)
 *             .vpcId(defaultbFzA4a.id())
 *             .build());
 * 
 *         var defaultVpcEndpoint = new VpcEndpoint("defaultVpcEndpoint", VpcEndpointArgs.builder()
 *             .endpointDescription(name)
 *             .vpcEndpointName(name)
 *             .resourceGroupId(default_.ids()[0])
 *             .endpointType("Interface")
 *             .vpcId(defaultbFzA4a.id())
 *             .serviceName("com.aliyuncs.privatelink.ap-southeast-5.oss")
 *             .dryRun("false")
 *             .zonePrivateIpAddressCount("1")
 *             .policyDocument(serializeJson(
 *                 jsonObject(
 *                     jsonProperty("Version", "1"),
 *                     jsonProperty("Statement", jsonArray(jsonObject(
 *                         jsonProperty("Effect", "Allow"),
 *                         jsonProperty("Action", jsonArray("*")),
 *                         jsonProperty("Resource", jsonArray("*")),
 *                         jsonProperty("Principal", "*")
 *                     )))
 *                 )))
 *             .securityGroupIds(default1FTFrP.id())
 *             .serviceId("epsrv-k1apjysze8u1l9t6uyg9")
 *             .protectedEnabled("false")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Private Link Vpc Endpoint can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:privatelink/vpcEndpoint:VpcEndpoint example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:privatelink/vpcEndpoint:VpcEndpoint")
public class VpcEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * The bandwidth of the endpoint connection.  1024 to 10240. Unit: Mbit/s. Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
     * 
     */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidth;

    /**
     * @return The bandwidth of the endpoint connection.  1024 to 10240. Unit: Mbit/s. Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * The state of the endpoint connection.
     * 
     */
    @Export(name="connectionStatus", refs={String.class}, tree="[0]")
    private Output<String> connectionStatus;

    /**
     * @return The state of the endpoint connection.
     * 
     */
    public Output<String> connectionStatus() {
        return this.connectionStatus;
    }
    /**
     * The time when the endpoint was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time when the endpoint was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to perform only a dry run, without performing the actual request. Valid values:
     * - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
     * - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The service state of the endpoint.
     * 
     */
    @Export(name="endpointBusinessStatus", refs={String.class}, tree="[0]")
    private Output<String> endpointBusinessStatus;

    /**
     * @return The service state of the endpoint.
     * 
     */
    public Output<String> endpointBusinessStatus() {
        return this.endpointBusinessStatus;
    }
    /**
     * The description of the endpoint.
     * 
     */
    @Export(name="endpointDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endpointDescription;

    /**
     * @return The description of the endpoint.
     * 
     */
    public Output<Optional<String>> endpointDescription() {
        return Codegen.optional(this.endpointDescription);
    }
    /**
     * The domain name of the endpoint.
     * 
     */
    @Export(name="endpointDomain", refs={String.class}, tree="[0]")
    private Output<String> endpointDomain;

    /**
     * @return The domain name of the endpoint.
     * 
     */
    public Output<String> endpointDomain() {
        return this.endpointDomain;
    }
    /**
     * The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
     * 
     */
    @Export(name="endpointType", refs={String.class}, tree="[0]")
    private Output<String> endpointType;

    /**
     * @return The endpoint type. Only the value: Interface, indicating the Interface endpoint. You can add the service resource types of Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
     * 
     */
    public Output<String> endpointType() {
        return this.endpointType;
    }
    /**
     * RAM access policies.
     * 
     */
    @Export(name="policyDocument", refs={String.class}, tree="[0]")
    private Output<String> policyDocument;

    /**
     * @return RAM access policies.
     * 
     */
    public Output<String> policyDocument() {
        return this.policyDocument;
    }
    /**
     * Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
     * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
     * - **false (default)**: disables user authentication.
     * 
     */
    @Export(name="protectedEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> protectedEnabled;

    /**
     * @return Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
     * - **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
     * - **false (default)**: disables user authentication.
     * 
     */
    public Output<Optional<Boolean>> protectedEnabled() {
        return Codegen.optional(this.protectedEnabled);
    }
    /**
     * The resource group ID.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The resource group ID.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroupIds;

    /**
     * @return The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.The endpoint can be associated with up to 10 security groups.
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * The ID of the endpoint service with which the endpoint is associated.
     * 
     */
    @Export(name="serviceId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceId;

    /**
     * @return The ID of the endpoint service with which the endpoint is associated.
     * 
     */
    public Output<Optional<String>> serviceId() {
        return Codegen.optional(this.serviceId);
    }
    /**
     * The name of the endpoint service with which the endpoint is associated.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The name of the endpoint service with which the endpoint is associated.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * The state of the endpoint.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The state of the endpoint.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The list of tags.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The list of tags.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The name of the endpoint.
     * 
     */
    @Export(name="vpcEndpointName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcEndpointName;

    /**
     * @return The name of the endpoint.
     * 
     */
    public Output<Optional<String>> vpcEndpointName() {
        return Codegen.optional(this.vpcEndpointName);
    }
    /**
     * The ID of the VPC to which the endpoint belongs.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC to which the endpoint belongs.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
     * 
     */
    @Export(name="zonePrivateIpAddressCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> zonePrivateIpAddressCount;

    /**
     * @return The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.
     * 
     */
    public Output<Integer> zonePrivateIpAddressCount() {
        return this.zonePrivateIpAddressCount;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpoint(java.lang.String name) {
        this(name, VpcEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpoint(java.lang.String name, VpcEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpoint(java.lang.String name, VpcEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpoint:VpcEndpoint", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VpcEndpoint(java.lang.String name, Output<java.lang.String> id, @Nullable VpcEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpoint:VpcEndpoint", name, state, makeResourceOptions(options, id), false);
    }

    private static VpcEndpointArgs makeArgs(VpcEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VpcEndpointArgs.Empty : args;
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
    public static VpcEndpoint get(java.lang.String name, Output<java.lang.String> id, @Nullable VpcEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpoint(name, id, state, options);
    }
}
