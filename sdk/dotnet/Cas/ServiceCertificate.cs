// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cas
{
    /// <summary>
    /// Provides a SSL Certificates Certificate resource.
    /// 
    /// For information about SSL Certificates Certificate and how to use it, see [What is Certificate](https://www.alibabacloud.com/help/product/28533.html).
    /// 
    /// &gt; **NOTE:** Available since v1.129.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var defaultServiceCertificate = new AliCloud.Cas.ServiceCertificate("defaultServiceCertificate", new()
    ///     {
    ///         Cert = @"-----BEGIN CERTIFICATE-----
    /// MIIDeDCCAmCgAwIBAgIEN3ZT6zANBgkqhkiG9w0BAQsFADBVMQswCQYDVQQGEwJD
    /// TjEVMBMGA1UEAwwMKi50ZnRlc3QudG9wMRAwDgYDVQQIDAdCZWlKaW5nMRAwDgYD
    /// VQQHDAdCZWlKaW5nMQswCQYDVQQKDAJBQTAeFw0yMzA4MjgwNjQ5NDNaFw0yNTA4
    /// MjcwNjQ5NDNaMFUxCzAJBgNVBAYTAkNOMRUwEwYDVQQDDAwqLnRmdGVzdC50b3Ax
    /// EDAOBgNVBAgMB0JlaUppbmcxEDAOBgNVBAcMB0JlaUppbmcxCzAJBgNVBAoMAkFB
    /// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzkk9NJUH7PLSQK4RRrGQ
    /// Y5dVsftkhnKh9HhI6yrnlowWIDPS1PZHOU/5gQ7xPUPGdKQV5S7x8wROnAaXEimx
    /// N4GdQw25pGhRJvlwme9fzJJiSe6lG49NCxmuBiEdJAzPKaTPpK1cG1f1TqdgCfHR
    /// HAL6Jxb3ylHG2LlTNFLXikubUi5RT6/9C8psr713Zm4HveCI/cx0WdgZ+fmsc9ft
    /// rkIB1DdyV1kQ51m8r2rLi3J7aC5ggGOiex/VlGSd4e6SOQLpdQEdDbodtOJ4LgVk
    /// +arFNCMinUWIOPGFzXhdm6lssPbh4MOBrz8c/M9TcF4hoMn5/o/9johZIZ/DOvXt
    /// ZQIDAQABo1AwTjAdBgNVHQ4EFgQUOnWiddgeZj17IeysatqhE361o5YwHwYDVR0j
    /// BBgwFoAUOnWiddgeZj17IeysatqhE361o5YwDAYDVR0TBAUwAwEB/zANBgkqhkiG
    /// 9w0BAQsFAAOCAQEAfh3cnOszHM/5wXjY7BIkmgDOReksS+87ibhBz7T2ddZj+yCF
    /// 9GdIBzXCiHpQFXpW8a3kc3I7l3nGfMTkmF6ld3ot/6SXP17QKJwxtvUA4ib8QkWD
    /// S7FT+UcHCUHv42Sh1e5uAlQ5pMSul7iKcR7jwlwZGZ0905HOqrmdyUGJ+Ud2uZWD
    /// AC0dJF6Bv9VhNtci8Imp05PaPH09deXLZu8LRrKRZFy9qLW5R6Swv7nzxckOAqDk
    /// TTc40xwvQROekWUyxeJL7xaHuylUHE0bxsiIfx5bZsBizRjprIwGlj85CSPuTZyP
    /// DPfaiZAN/61h5HNAnxLltOZfqabKYYw7l9LBDg==
    /// -----END CERTIFICATE-----
    /// 
    /// ",
    ///         CertificateName = defaultRandomInteger.Result.Apply(result =&gt; $"tf-example-{result}"),
    ///         Key = @"-----BEGIN PRIVATE KEY-----
    /// MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDOST00lQfs8tJA
    /// rhFGsZBjl1Wx+2SGcqH0eEjrKueWjBYgM9LU9kc5T/mBDvE9Q8Z0pBXlLvHzBE6c
    /// BpcSKbE3gZ1DDbmkaFEm+XCZ71/MkmJJ7qUbj00LGa4GIR0kDM8ppM+krVwbV/VO
    /// p2AJ8dEcAvonFvfKUcbYuVM0UteKS5tSLlFPr/0LymyvvXdmbge94Ij9zHRZ2Bn5
    /// +axz1+2uQgHUN3JXWRDnWbyvasuLcntoLmCAY6J7H9WUZJ3h7pI5Aul1AR0Nuh20
    /// 4nguBWT5qsU0IyKdRYg48YXNeF2bqWyw9uHgw4GvPxz8z1NwXiGgyfn+j/2OiFkh
    /// n8M69e1lAgMBAAECggEAevPgTTT+0lYwx2h416ACJboP09O5KQGuUl5XaAPcoTjB
    /// /1OkOFbKQPjQCAJ1+0QoR2F9w2plv6kziX/MD4FWJXVV3J+TpNCgfhBy8u1gNjiR
    /// 6Osa8gBJtXIK7ZBTJCeWWoXnVYoWuh2FEupkLck6D+4eV6oy6x4u3QIo+6jc24n9
    /// dIXQG6/v/Iao34kB0LUdp/4WNaUDvfI6NDhEwchpKE95dtWIDlIN/YhfiYAdjrnl
    /// YmH2VDbAGgsdEiHP4wLZfjgsGPPDGS0+qBHoSiJGH0E6wWEZdAE4TsYGRFsO86n3
    /// LfjEPFGfPlcnZe2cTTe3kmyKq/DTjxtu2rh3I8o2CQKBgQD/5Xe7cenaOBefzPlx
    /// GOEsB+qv49UmzEPOXDNZe9hmAawuuuxPUM+xlE++P+mEgQm1LPT4WWgtFLPVuwmx
    /// ncxt4CJNZh+ZGFyAZ4dm4M4ZhIBXNonyIP+yGyAJUUVF9Iy3TYcJNiGzv2Rx9JRQ
    /// XWJMQnTDILmZbmU+ltTea7/zqwKBgQDOXqCqb17MuLt7OcKWSgthm79OlaOdzZvl
    /// i9qU6VzZKG7Axc5gA9yq6tHp3vWPI4bNdvwqIIa/nzVILjGA5fcYFbRN+7gHwo8s
    /// rNAgi5PAoKWqQRovyJRY9Eq/sn6l1jbJZAOUAMZMWDm8z89OqK7PNQSIAtfFSneo
    /// 2QxJkGeTLwKBgGJkafBB8af9b1/7YWISLepPNPbihH/BhMThAMGEdAVs2TaymtA4
    /// g1OFck/1pSVUtFXcbmjbf8ntruQcYbLQuNz6lFXsUXP9QPwCUrbE85ouL2bZSps2
    /// AvsJoPzUKe2nBUAp6CUrkjPaAJYsc6ae8X/fAaRRfeu33ef9+OV4yrq3AoGAYFZo
    /// ZmfrN2Kdkt7Z6dLTEVPlsMfGQ6pyNmxdM9rkzzNC0JcGymfDIb7RE35T3+hTy6La
    /// AMiCXv3xn6qAzY2NFh87tpPlyymWMOLTnf3Kkcfszlfp45idOBGCu46V9NDVbppT
    /// 2UmrSIR/H5dbTXsNcAlt/hhlpeInjhkU1VqmH10CgYEA7Kk+QhWq705SczpWjm5J
    /// 9kHqfFzJLwAWNBduiia0WypgPhLe/4wT1rYQkBtKMVKrgFo7Cvi4YKlrtlDnXyeU
    /// CIFqfEL5NriQelqrFsvgHsmD+MpvDoSWm5C8IrTubtlNyWUzXSVT4OIwzPobzPqG
    /// LILJ+e7bLw8RrM0HfgFnl8c=
    /// -----END PRIVATE KEY-----
    /// 
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SSL Certificates Certificate can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cas/serviceCertificate:ServiceCertificate example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cas/serviceCertificate:ServiceCertificate")]
    public partial class ServiceCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cert of the Certificate in which the Certificate will add.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// Name of the Certificate. 
        /// This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-",
        /// and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time.
        /// Suffix .sh and .tel are not supported.
        /// **NOTE:** One of `certificate_name` and `name` must be specified.
        /// </summary>
        [Output("certificateName")]
        public Output<string> CertificateName { get; private set; } = null!;

        /// <summary>
        /// Key of the Certificate in which the Certificate will add.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The lang.
        /// </summary>
        [Output("lang")]
        public Output<string?> Lang { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.129.0 and using `certificate_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceCertificate(string name, ServiceCertificateArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cas/serviceCertificate:ServiceCertificate", name, args ?? new ServiceCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceCertificate(string name, Input<string> id, ServiceCertificateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cas/serviceCertificate:ServiceCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceCertificate Get(string name, Input<string> id, ServiceCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceCertificate(name, id, state, options);
        }
    }

    public sealed class ServiceCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cert of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("cert", required: true)]
        public Input<string> Cert { get; set; } = null!;

        /// <summary>
        /// Name of the Certificate. 
        /// This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-",
        /// and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time.
        /// Suffix .sh and .tel are not supported.
        /// **NOTE:** One of `certificate_name` and `name` must be specified.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// Key of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The lang.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.129.0 and using `certificate_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceCertificateArgs()
        {
        }
        public static new ServiceCertificateArgs Empty => new ServiceCertificateArgs();
    }

    public sealed class ServiceCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cert of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Name of the Certificate. 
        /// This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-",
        /// and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time.
        /// Suffix .sh and .tel are not supported.
        /// **NOTE:** One of `certificate_name` and `name` must be specified.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// Key of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The lang.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.129.0 and using `certificate_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceCertificateState()
        {
        }
        public static new ServiceCertificateState Empty => new ServiceCertificateState();
    }
}
