// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eventbridge.ConnectionArgs;
import com.pulumi.alicloud.eventbridge.inputs.ConnectionState;
import com.pulumi.alicloud.eventbridge.outputs.ConnectionAuthParameters;
import com.pulumi.alicloud.eventbridge.outputs.ConnectionNetworkParameters;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Event Bridge Connection resource.
 * 
 * For information about Event Bridge Connection and how to use it, see [What is Connection](https://www.alibabacloud.com/help/en/eventbridge/latest/api-eventbridge-2020-04-01-createconnection).
 * 
 * &gt; **NOTE:** Available since v1.210.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.eventbridge.Connection;
 * import com.pulumi.alicloud.eventbridge.ConnectionArgs;
 * import com.pulumi.alicloud.eventbridge.inputs.ConnectionNetworkParametersArgs;
 * import com.pulumi.alicloud.eventbridge.inputs.ConnectionAuthParametersArgs;
 * import com.pulumi.alicloud.eventbridge.inputs.ConnectionAuthParametersApiKeyAuthParametersArgs;
 * import com.pulumi.alicloud.eventbridge.inputs.ConnectionAuthParametersBasicAuthParametersArgs;
 * import com.pulumi.alicloud.eventbridge.inputs.ConnectionAuthParametersOauthParametersArgs;
 * import com.pulumi.alicloud.eventbridge.inputs.ConnectionAuthParametersOauthParametersClientParametersArgs;
 * import com.pulumi.alicloud.eventbridge.inputs.ConnectionAuthParametersOauthParametersOauthHttpParametersArgs;
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
 *         final var region = config.get(&#34;region&#34;).orElse(&#34;cn-chengdu&#34;);
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones();
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultSwitch.vpcId())
 *             .build());
 * 
 *         var defaultConnection = new Connection(&#34;defaultConnection&#34;, ConnectionArgs.builder()        
 *             .connectionName(name)
 *             .description(&#34;test-connection-basic-pre&#34;)
 *             .networkParameters(ConnectionNetworkParametersArgs.builder()
 *                 .networkType(&#34;PublicNetwork&#34;)
 *                 .vpcId(defaultNetwork.id())
 *                 .vswitcheId(defaultSwitch.id())
 *                 .securityGroupId(defaultSecurityGroup.id())
 *                 .build())
 *             .authParameters(ConnectionAuthParametersArgs.builder()
 *                 .authorizationType(&#34;BASIC_AUTH&#34;)
 *                 .apiKeyAuthParameters(ConnectionAuthParametersApiKeyAuthParametersArgs.builder()
 *                     .apiKeyName(&#34;Token&#34;)
 *                     .apiKeyValue(&#34;Token-value&#34;)
 *                     .build())
 *                 .basicAuthParameters(ConnectionAuthParametersBasicAuthParametersArgs.builder()
 *                     .username(&#34;admin&#34;)
 *                     .password(&#34;admin&#34;)
 *                     .build())
 *                 .oauthParameters(ConnectionAuthParametersOauthParametersArgs.builder()
 *                     .authorizationEndpoint(&#34;http://127.0.0.1:8080&#34;)
 *                     .httpMethod(&#34;POST&#34;)
 *                     .clientParameters(ConnectionAuthParametersOauthParametersClientParametersArgs.builder()
 *                         .clientId(&#34;ClientId&#34;)
 *                         .clientSecret(&#34;ClientSecret&#34;)
 *                         .build())
 *                     .oauthHttpParameters(ConnectionAuthParametersOauthParametersOauthHttpParametersArgs.builder()
 *                         .headerParameters(ConnectionAuthParametersOauthParametersOauthHttpParametersHeaderParameterArgs.builder()
 *                             .key(&#34;name&#34;)
 *                             .value(&#34;name&#34;)
 *                             .isValueSecret(&#34;true&#34;)
 *                             .build())
 *                         .bodyParameters(ConnectionAuthParametersOauthParametersOauthHttpParametersBodyParameterArgs.builder()
 *                             .key(&#34;name&#34;)
 *                             .value(&#34;name&#34;)
 *                             .isValueSecret(&#34;true&#34;)
 *                             .build())
 *                         .queryStringParameters(ConnectionAuthParametersOauthParametersOauthHttpParametersQueryStringParameterArgs.builder()
 *                             .key(&#34;name&#34;)
 *                             .value(&#34;name&#34;)
 *                             .isValueSecret(&#34;true&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Event Bridge Connection can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eventbridge/connection:Connection example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eventbridge/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    /**
     * The parameters that are configured for authentication. See `auth_parameters` below.
     * 
     */
    @Export(name="authParameters", refs={ConnectionAuthParameters.class}, tree="[0]")
    private Output</* @Nullable */ ConnectionAuthParameters> authParameters;

    /**
     * @return The parameters that are configured for authentication. See `auth_parameters` below.
     * 
     */
    public Output<Optional<ConnectionAuthParameters>> authParameters() {
        return Codegen.optional(this.authParameters);
    }
    /**
     * The name of the connection.
     * 
     */
    @Export(name="connectionName", refs={String.class}, tree="[0]")
    private Output<String> connectionName;

    /**
     * @return The name of the connection.
     * 
     */
    public Output<String> connectionName() {
        return this.connectionName;
    }
    /**
     * The creation time of the Connection.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the Connection.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The description of the connection.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the connection.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The parameters that are configured for the network. See `network_parameters` below.
     * 
     */
    @Export(name="networkParameters", refs={ConnectionNetworkParameters.class}, tree="[0]")
    private Output<ConnectionNetworkParameters> networkParameters;

    /**
     * @return The parameters that are configured for the network. See `network_parameters` below.
     * 
     */
    public Output<ConnectionNetworkParameters> networkParameters() {
        return this.networkParameters;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eventbridge/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eventbridge/connection:Connection", name, state, makeResourceOptions(options, id));
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
    public static Connection get(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connection(name, id, state, options);
    }
}
