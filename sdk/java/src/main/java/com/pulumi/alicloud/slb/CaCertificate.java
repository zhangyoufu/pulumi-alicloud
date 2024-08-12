// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.slb.CaCertificateArgs;
import com.pulumi.alicloud.slb.inputs.CaCertificateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Load Balancer CA Certificate is used by the listener of the protocol https.
 * 
 * For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
 * 
 * For information about CA Certificate and how to use it, see [Configure CA Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
 * 
 * ## Example Usage
 * 
 * * using CA certificate content
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.slb.CaCertificate;
 * import com.pulumi.alicloud.slb.CaCertificateArgs;
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
 *         var foo = new CaCertificate("foo", CaCertificateArgs.builder()
 *             .caCertificateName("tf-testAccSlbCACertificate")
 *             .caCertificate("""
 * -----BEGIN CERTIFICATE-----
 * MIIDRjCCAq+gAwIBAgIJAJn3ox4K13PoMA0GCSqGSIb3DQEBBQUAMHYxCzAJBgNV
 * BAYTAkNOMQswCQYDVQQIEwJCSjELMAkGA1UEBxMCQkoxDDAKBgNVBAoTA0FMSTEP
 * MA0GA1UECxMGQUxJWVVOMQ0wCwYDVQQDEwR0ZXN0MR8wHQYJKoZIhvcNAQkBFhB0
 * ZXN0QGhvdG1haWwuY29tMB4XDTE0MTEyNDA2MDQyNVoXDTI0MTEyMTA2MDQyNVow
 * djELMAkGA1UEBhMCQ04xCzAJBgNVBAgTAkJKMQswCQYDVQQHEwJCSjEMMAoGA1UE
 * ChMDQUxJMQ8wDQYDVQQLEwZBTElZVU4xDTALBgNVBAMTBHRlc3QxHzAdBgkqhkiG
 * 9w0BCQEWEHRlc3RAaG90bWFpbC5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJ
 * AoGBAM7SS3e9+Nj0HKAsRuIDNSsS3UK6b+62YQb2uuhKrp1HMrOx61WSDR2qkAnB
 * coG00Uz38EE+9DLYNUVQBK7aSgLP5M1Ak4wr4GqGyCgjejzzh3DshUzLCCy2rook
 * KOyRTlPX+Q5l7rE1fcSNzgepcae5i2sE1XXXzLRIDIvQxcspAgMBAAGjgdswgdgw
 * HQYDVR0OBBYEFBdy+OuMsvbkV7R14f0OyoLoh2z4MIGoBgNVHSMEgaAwgZ2AFBdy
 * +OuMsvbkV7R14f0OyoLoh2z4oXqkeDB2MQswCQYDVQQGEwJDTjELMAkGA1UECBMC
 * QkoxCzAJBgNVBAcTAkJKMQwwCgYDVQQKEwNBTEkxDzANBgNVBAsTBkFMSVlVTjEN
 * MAsGA1UEAxMEdGVzdDEfMB0GCSqGSIb3DQEJARYQdGVzdEBob3RtYWlsLmNvbYIJ
 * AJn3ox4K13PoMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADgYEAY7KOsnyT
 * cQzfhiiG7ASjiPakw5wXoycHt5GCvLG5htp2TKVzgv9QTliA3gtfv6oV4zRZx7X1
 * Ofi6hVgErtHaXJheuPVeW6eAW8mHBoEfvDAfU3y9waYrtUevSl07643bzKL6v+Qd
 * DUBTxOAvSYfXTtI90EAxEG/bJJyOm5LqoiA=
 * -----END CERTIFICATE-----            """)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * * using CA certificate file
 * 
 * ## Import
 * 
 * Server Load balancer CA Certificate can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:slb/caCertificate:CaCertificate example abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:slb/caCertificate:CaCertificate")
public class CaCertificate extends com.pulumi.resources.CustomResource {
    /**
     * the content of the CA certificate.
     * 
     */
    @Export(name="caCertificate", refs={String.class}, tree="[0]")
    private Output<String> caCertificate;

    /**
     * @return the content of the CA certificate.
     * 
     */
    public Output<String> caCertificate() {
        return this.caCertificate;
    }
    /**
     * Name of the CA Certificate.
     * 
     */
    @Export(name="caCertificateName", refs={String.class}, tree="[0]")
    private Output<String> caCertificateName;

    /**
     * @return Name of the CA Certificate.
     * 
     */
    public Output<String> caCertificateName() {
        return this.caCertificateName;
    }
    /**
     * Field `name` has been deprecated from provider version 1.123.1. New field `ca_certificate_name` instead
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;ca_certificate_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'ca_certificate_name' instead */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Field `name` has been deprecated from provider version 1.123.1. New field `ca_certificate_name` instead
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Id of resource group which the slb_ca certificate belongs.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the slb_ca certificate belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CaCertificate(java.lang.String name) {
        this(name, CaCertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CaCertificate(java.lang.String name, CaCertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CaCertificate(java.lang.String name, CaCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/caCertificate:CaCertificate", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CaCertificate(java.lang.String name, Output<java.lang.String> id, @Nullable CaCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/caCertificate:CaCertificate", name, state, makeResourceOptions(options, id), false);
    }

    private static CaCertificateArgs makeArgs(CaCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CaCertificateArgs.Empty : args;
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
    public static CaCertificate get(java.lang.String name, Output<java.lang.String> id, @Nullable CaCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CaCertificate(name, id, state, options);
    }
}
