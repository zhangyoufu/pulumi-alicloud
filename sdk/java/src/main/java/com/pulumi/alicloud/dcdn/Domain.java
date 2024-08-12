// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dcdn.DomainArgs;
import com.pulumi.alicloud.dcdn.inputs.DomainState;
import com.pulumi.alicloud.dcdn.outputs.DomainSource;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DCDN Domain resource.
 * 
 * Full station accelerated domain name.
 * 
 * For information about DCDN Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/en/doc-detail/130628.htm).
 * 
 * &gt; **NOTE:** Available since v1.94.0.
 * 
 * &gt; **NOTE:** Field `force_set`, `security_token` has been removed from provider version 1.227.1.
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.dcdn.Domain;
 * import com.pulumi.alicloud.dcdn.DomainArgs;
 * import com.pulumi.alicloud.dcdn.inputs.DomainSourceArgs;
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
 *         final var domainName = config.get("domainName").orElse("tf-example.com");
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var example = new Domain("example", DomainArgs.builder()
 *             .domainName(String.format("%s-%s", domainName,default_.result()))
 *             .scope("overseas")
 *             .sources(DomainSourceArgs.builder()
 *                 .content("1.1.1.1")
 *                 .port("80")
 *                 .priority("20")
 *                 .type("ipaddr")
 *                 .weight("10")
 *                 .build())
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
 * DCDN Domain can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dcdn/domain:Domain example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dcdn/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
     * 
     */
    @Export(name="certId", refs={String.class}, tree="[0]")
    private Output<String> certId;

    /**
     * @return The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
     * 
     */
    public Output<String> certId() {
        return this.certId;
    }
    /**
     * The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
     * 
     */
    @Export(name="certName", refs={String.class}, tree="[0]")
    private Output<String> certName;

    /**
     * @return The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
     * 
     */
    public Output<String> certName() {
        return this.certName;
    }
    /**
     * The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
     * 
     */
    @Export(name="certRegion", refs={String.class}, tree="[0]")
    private Output<String> certRegion;

    /**
     * @return The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
     * 
     */
    public Output<String> certRegion() {
        return this.certRegion;
    }
    /**
     * The certificate type.
     * 
     */
    @Export(name="certType", refs={String.class}, tree="[0]")
    private Output<String> certType;

    /**
     * @return The certificate type.
     * 
     */
    public Output<String> certType() {
        return this.certType;
    }
    /**
     * The URL that is used for health checks.
     * 
     */
    @Export(name="checkUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> checkUrl;

    /**
     * @return The URL that is used for health checks.
     * 
     */
    public Output<Optional<String>> checkUrl() {
        return Codegen.optional(this.checkUrl);
    }
    /**
     * The CNAME domain name corresponding to the accelerated domain name.
     * 
     */
    @Export(name="cname", refs={String.class}, tree="[0]")
    private Output<String> cname;

    /**
     * @return The CNAME domain name corresponding to the accelerated domain name.
     * 
     */
    public Output<String> cname() {
        return this.cname;
    }
    /**
     * The time when the accelerated domain name was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time when the accelerated domain name was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
     * 
     */
    @Export(name="env", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> env;

    /**
     * @return Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
     * 
     */
    public Output<Optional<String>> env() {
        return Codegen.optional(this.env);
    }
    /**
     * Computing service type. Valid values:
     * 
     */
    @Export(name="functionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> functionType;

    /**
     * @return Computing service type. Valid values:
     * 
     */
    public Output<Optional<String>> functionType() {
        return Codegen.optional(this.functionType);
    }
    /**
     * The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The Acceleration scen. Supported:
     * 
     */
    @Export(name="scene", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scene;

    /**
     * @return The Acceleration scen. Supported:
     * 
     */
    public Output<Optional<String>> scene() {
        return Codegen.optional(this.scene);
    }
    /**
     * The region where the acceleration service is deployed. Valid values:
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scope;

    /**
     * @return The region where the acceleration service is deployed. Valid values:
     * 
     */
    public Output<Optional<String>> scope() {
        return Codegen.optional(this.scope);
    }
    /**
     * Source  See `sources` below.
     * 
     */
    @Export(name="sources", refs={List.class,DomainSource.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DomainSource>> sources;

    /**
     * @return Source  See `sources` below.
     * 
     */
    public Output<Optional<List<DomainSource>>> sources() {
        return Codegen.optional(this.sources);
    }
    /**
     * The private key. Specify the private key only if you want to enable the SSL certificate.
     * 
     */
    @Export(name="sslPri", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sslPri;

    /**
     * @return The private key. Specify the private key only if you want to enable the SSL certificate.
     * 
     */
    public Output<Optional<String>> sslPri() {
        return Codegen.optional(this.sslPri);
    }
    /**
     * Specifies whether to enable the SSL certificate. Valid values:
     * 
     */
    @Export(name="sslProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sslProtocol;

    /**
     * @return Specifies whether to enable the SSL certificate. Valid values:
     * 
     */
    public Output<Optional<String>> sslProtocol() {
        return Codegen.optional(this.sslProtocol);
    }
    /**
     * The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
     * 
     */
    @Export(name="sslPub", refs={String.class}, tree="[0]")
    private Output<String> sslPub;

    /**
     * @return The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
     * 
     */
    public Output<String> sslPub() {
        return this.sslPub;
    }
    /**
     * The status of the domain name. Valid values:
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the domain name. Valid values:
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tag of the resource
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tag of the resource
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The top-level domain.
     * 
     */
    @Export(name="topLevelDomain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> topLevelDomain;

    /**
     * @return The top-level domain.
     * 
     */
    public Output<Optional<String>> topLevelDomain() {
        return Codegen.optional(this.topLevelDomain);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Domain(java.lang.String name) {
        this(name, DomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Domain(java.lang.String name, DomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Domain(java.lang.String name, DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dcdn/domain:Domain", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Domain(java.lang.String name, Output<java.lang.String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dcdn/domain:Domain", name, state, makeResourceOptions(options, id), false);
    }

    private static DomainArgs makeArgs(DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DomainArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "sslPri"
            ))
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
    public static Domain get(java.lang.String name, Output<java.lang.String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Domain(name, id, state, options);
    }
}
