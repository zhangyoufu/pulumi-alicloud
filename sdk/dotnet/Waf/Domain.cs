// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Waf
{
    /// <summary>
    /// &gt; **DEPRECATED:**  This resource has been deprecated and using alicloud.wafv3.Domain instead.
    /// 
    /// Provides a WAF Domain resource to create domain in the Web Application Firewall.
    /// 
    /// For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.82.0+ .
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var domain = new AliCloud.Waf.Domain("domain", new()
    ///     {
    ///         DomainName = "alicloud-provider.cn",
    ///         InstanceId = "waf-123455",
    ///         IsAccessProduct = "On",
    ///         SourceIps = new[]
    ///         {
    ///             "1.1.1.1",
    ///         },
    ///         ClusterType = "PhysicalCluster",
    ///         Http2Ports = new[]
    ///         {
    ///             "443",
    ///         },
    ///         HttpPorts = new[]
    ///         {
    ///             "80",
    ///         },
    ///         HttpsPorts = new[]
    ///         {
    ///             "443",
    ///         },
    ///         HttpToUserIp = "Off",
    ///         HttpsRedirect = "Off",
    ///         LoadBalancing = "IpHash",
    ///         LogHeaders = new[]
    ///         {
    ///             new AliCloud.Waf.Inputs.DomainLogHeaderArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "http",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// WAF domain can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:waf/domain:Domain domain waf-132435:www.domain.com
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:waf/domain:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
        /// </summary>
        [Output("clusterType")]
        public Output<string?> ClusterType { get; private set; } = null!;

        /// <summary>
        /// The CNAME record assigned by the WAF instance to the specified domain.
        /// </summary>
        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        /// <summary>
        /// The connection timeout for WAF exclusive clusters. Unit: seconds.
        /// </summary>
        [Output("connectionTime")]
        public Output<int?> ConnectionTime { get; private set; } = null!;

        /// <summary>
        /// Field `domain` has been deprecated from version 1.94.0. Use `domain_name` instead.
        /// </summary>
        [Output("domain")]
        public Output<string> DomainDeprecated { get; private set; } = null!;

        /// <summary>
        /// The domain that you want to add to WAF. The `domain_name` is required when the value of the `domain`  is Empty.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// List of the HTTP 2.0 ports.
        /// </summary>
        [Output("http2Ports")]
        public Output<ImmutableArray<string>> Http2Ports { get; private set; } = null!;

        /// <summary>
        /// List of the HTTP ports.
        /// </summary>
        [Output("httpPorts")]
        public Output<ImmutableArray<string>> HttpPorts { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server. 
        /// By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
        /// </summary>
        [Output("httpToUserIp")]
        public Output<string?> HttpToUserIp { get; private set; } = null!;

        /// <summary>
        /// List of the HTTPS ports.
        /// </summary>
        [Output("httpsPorts")]
        public Output<ImmutableArray<string>> HttpsPorts { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: "On" and `Off`. Default to `Off`.
        /// </summary>
        [Output("httpsRedirect")]
        public Output<string?> HttpsRedirect { get; private set; } = null!;

        /// <summary>
        /// The ID of the WAF instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
        /// </summary>
        [Output("isAccessProduct")]
        public Output<string> IsAccessProduct { get; private set; } = null!;

        /// <summary>
        /// The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
        /// </summary>
        [Output("loadBalancing")]
        public Output<string?> LoadBalancing { get; private set; } = null!;

        /// <summary>
        /// The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
        /// * key: The key of label
        /// * value: The value of label
        /// </summary>
        [Output("logHeaders")]
        public Output<ImmutableArray<Outputs.DomainLogHeader>> LogHeaders { get; private set; } = null!;

        /// <summary>
        /// The read timeout of a WAF exclusive cluster. Unit: seconds.
        /// </summary>
        [Output("readTime")]
        public Output<int?> ReadTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// List of the IP address or domain of the origin server to which the specified domain points.
        /// </summary>
        [Output("sourceIps")]
        public Output<ImmutableArray<string>> SourceIps { get; private set; } = null!;

        /// <summary>
        /// The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
        /// </summary>
        [Output("writeTime")]
        public Output<int?> WriteTime { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("alicloud:waf/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:waf/domain:Domain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// The connection timeout for WAF exclusive clusters. Unit: seconds.
        /// </summary>
        [Input("connectionTime")]
        public Input<int>? ConnectionTime { get; set; }

        /// <summary>
        /// Field `domain` has been deprecated from version 1.94.0. Use `domain_name` instead.
        /// </summary>
        [Input("domain")]
        public Input<string>? DomainDeprecated { get; set; }

        /// <summary>
        /// The domain that you want to add to WAF. The `domain_name` is required when the value of the `domain`  is Empty.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("http2Ports")]
        private InputList<string>? _http2Ports;

        /// <summary>
        /// List of the HTTP 2.0 ports.
        /// </summary>
        public InputList<string> Http2Ports
        {
            get => _http2Ports ?? (_http2Ports = new InputList<string>());
            set => _http2Ports = value;
        }

        [Input("httpPorts")]
        private InputList<string>? _httpPorts;

        /// <summary>
        /// List of the HTTP ports.
        /// </summary>
        public InputList<string> HttpPorts
        {
            get => _httpPorts ?? (_httpPorts = new InputList<string>());
            set => _httpPorts = value;
        }

        /// <summary>
        /// Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server. 
        /// By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
        /// </summary>
        [Input("httpToUserIp")]
        public Input<string>? HttpToUserIp { get; set; }

        [Input("httpsPorts")]
        private InputList<string>? _httpsPorts;

        /// <summary>
        /// List of the HTTPS ports.
        /// </summary>
        public InputList<string> HttpsPorts
        {
            get => _httpsPorts ?? (_httpsPorts = new InputList<string>());
            set => _httpsPorts = value;
        }

        /// <summary>
        /// Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: "On" and `Off`. Default to `Off`.
        /// </summary>
        [Input("httpsRedirect")]
        public Input<string>? HttpsRedirect { get; set; }

        /// <summary>
        /// The ID of the WAF instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
        /// </summary>
        [Input("isAccessProduct", required: true)]
        public Input<string> IsAccessProduct { get; set; } = null!;

        /// <summary>
        /// The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
        /// </summary>
        [Input("loadBalancing")]
        public Input<string>? LoadBalancing { get; set; }

        [Input("logHeaders")]
        private InputList<Inputs.DomainLogHeaderArgs>? _logHeaders;

        /// <summary>
        /// The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
        /// * key: The key of label
        /// * value: The value of label
        /// </summary>
        public InputList<Inputs.DomainLogHeaderArgs> LogHeaders
        {
            get => _logHeaders ?? (_logHeaders = new InputList<Inputs.DomainLogHeaderArgs>());
            set => _logHeaders = value;
        }

        /// <summary>
        /// The read timeout of a WAF exclusive cluster. Unit: seconds.
        /// </summary>
        [Input("readTime")]
        public Input<int>? ReadTime { get; set; }

        /// <summary>
        /// The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("sourceIps")]
        private InputList<string>? _sourceIps;

        /// <summary>
        /// List of the IP address or domain of the origin server to which the specified domain points.
        /// </summary>
        public InputList<string> SourceIps
        {
            get => _sourceIps ?? (_sourceIps = new InputList<string>());
            set => _sourceIps = value;
        }

        /// <summary>
        /// The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
        /// </summary>
        [Input("writeTime")]
        public Input<int>? WriteTime { get; set; }

        public DomainArgs()
        {
        }
        public static new DomainArgs Empty => new DomainArgs();
    }

    public sealed class DomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// The CNAME record assigned by the WAF instance to the specified domain.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// The connection timeout for WAF exclusive clusters. Unit: seconds.
        /// </summary>
        [Input("connectionTime")]
        public Input<int>? ConnectionTime { get; set; }

        /// <summary>
        /// Field `domain` has been deprecated from version 1.94.0. Use `domain_name` instead.
        /// </summary>
        [Input("domain")]
        public Input<string>? DomainDeprecated { get; set; }

        /// <summary>
        /// The domain that you want to add to WAF. The `domain_name` is required when the value of the `domain`  is Empty.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("http2Ports")]
        private InputList<string>? _http2Ports;

        /// <summary>
        /// List of the HTTP 2.0 ports.
        /// </summary>
        public InputList<string> Http2Ports
        {
            get => _http2Ports ?? (_http2Ports = new InputList<string>());
            set => _http2Ports = value;
        }

        [Input("httpPorts")]
        private InputList<string>? _httpPorts;

        /// <summary>
        /// List of the HTTP ports.
        /// </summary>
        public InputList<string> HttpPorts
        {
            get => _httpPorts ?? (_httpPorts = new InputList<string>());
            set => _httpPorts = value;
        }

        /// <summary>
        /// Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server. 
        /// By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
        /// </summary>
        [Input("httpToUserIp")]
        public Input<string>? HttpToUserIp { get; set; }

        [Input("httpsPorts")]
        private InputList<string>? _httpsPorts;

        /// <summary>
        /// List of the HTTPS ports.
        /// </summary>
        public InputList<string> HttpsPorts
        {
            get => _httpsPorts ?? (_httpsPorts = new InputList<string>());
            set => _httpsPorts = value;
        }

        /// <summary>
        /// Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: "On" and `Off`. Default to `Off`.
        /// </summary>
        [Input("httpsRedirect")]
        public Input<string>? HttpsRedirect { get; set; }

        /// <summary>
        /// The ID of the WAF instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
        /// </summary>
        [Input("isAccessProduct")]
        public Input<string>? IsAccessProduct { get; set; }

        /// <summary>
        /// The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
        /// </summary>
        [Input("loadBalancing")]
        public Input<string>? LoadBalancing { get; set; }

        [Input("logHeaders")]
        private InputList<Inputs.DomainLogHeaderGetArgs>? _logHeaders;

        /// <summary>
        /// The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
        /// * key: The key of label
        /// * value: The value of label
        /// </summary>
        public InputList<Inputs.DomainLogHeaderGetArgs> LogHeaders
        {
            get => _logHeaders ?? (_logHeaders = new InputList<Inputs.DomainLogHeaderGetArgs>());
            set => _logHeaders = value;
        }

        /// <summary>
        /// The read timeout of a WAF exclusive cluster. Unit: seconds.
        /// </summary>
        [Input("readTime")]
        public Input<int>? ReadTime { get; set; }

        /// <summary>
        /// The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("sourceIps")]
        private InputList<string>? _sourceIps;

        /// <summary>
        /// List of the IP address or domain of the origin server to which the specified domain points.
        /// </summary>
        public InputList<string> SourceIps
        {
            get => _sourceIps ?? (_sourceIps = new InputList<string>());
            set => _sourceIps = value;
        }

        /// <summary>
        /// The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
        /// </summary>
        [Input("writeTime")]
        public Input<int>? WriteTime { get; set; }

        public DomainState()
        {
        }
        public static new DomainState Empty => new DomainState();
    }
}
