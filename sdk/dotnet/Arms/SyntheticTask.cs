// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms
{
    /// <summary>
    /// Provides a ARMS Synthetic Task resource. Cloud Synthetic task resources.
    /// 
    /// For information about ARMS Synthetic Task and how to use it, see [What is Synthetic Task](https://next.api.alibabacloud.com/document/ARMS/2019-08-08/CreateTimingSyntheticTask).
    /// 
    /// &gt; **NOTE:** Available since v1.215.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultSyntheticTask = new AliCloud.Arms.SyntheticTask("default", new()
    ///     {
    ///         Monitors = new[]
    ///         {
    ///             new AliCloud.Arms.Inputs.SyntheticTaskMonitorArgs
    ///             {
    ///                 CityCode = "1200101",
    ///                 OperatorCode = "246",
    ///                 ClientType = 4,
    ///             },
    ///         },
    ///         SyntheticTaskName = name,
    ///         CustomPeriod = new AliCloud.Arms.Inputs.SyntheticTaskCustomPeriodArgs
    ///         {
    ///             EndHour = 12,
    ///             StartHour = 11,
    ///         },
    ///         AvailableAssertions = new[]
    ///         {
    ///             new AliCloud.Arms.Inputs.SyntheticTaskAvailableAssertionArgs
    ///             {
    ///                 Type = "IcmpPackLoss",
    ///                 Operator = "neq",
    ///                 Expect = "200",
    ///                 Target = "example",
    ///             },
    ///             new AliCloud.Arms.Inputs.SyntheticTaskAvailableAssertionArgs
    ///             {
    ///                 Type = "IcmpPackAvgLatency",
    ///                 Operator = "lte",
    ///                 Expect = "1000",
    ///             },
    ///             new AliCloud.Arms.Inputs.SyntheticTaskAvailableAssertionArgs
    ///             {
    ///                 Type = "IcmpPackMaxLatency",
    ///                 Operator = "lte",
    ///                 Expect = "10000",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "example" },
    ///         },
    ///         Status = "RUNNING",
    ///         MonitorConf = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfArgs
    ///         {
    ///             NetTcp = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfNetTcpArgs
    ///             {
    ///                 TracertTimeout = 1050,
    ///                 TargetUrl = "www.aliyun.com",
    ///                 ConnectTimes = 6,
    ///                 Interval = 300,
    ///                 Timeout = 3000,
    ///                 TracertNumMax = 2,
    ///             },
    ///             NetDns = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfNetDnsArgs
    ///             {
    ///                 QueryMethod = 1,
    ///                 Timeout = 5050,
    ///                 TargetUrl = "www.aliyun.com",
    ///                 DnsServerIpType = 1,
    ///                 NsServer = "61.128.114.167",
    ///             },
    ///             ApiHttp = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfApiHttpArgs
    ///             {
    ///                 Timeout = 10050,
    ///                 TargetUrl = "https://www.aliyun.com",
    ///                 Method = "POST",
    ///                 RequestHeaders = 
    ///                 {
    ///                     { "key1", "value1" },
    ///                 },
    ///                 RequestBody = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfApiHttpRequestBodyArgs
    ///                 {
    ///                     Content = "example2",
    ///                     Type = "text/html",
    ///                 },
    ///                 ConnectTimeout = 6000,
    ///             },
    ///             Website = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfWebsiteArgs
    ///             {
    ///                 SlowElementThreshold = 5005,
    ///                 VerifyStringBlacklist = "Failed",
    ///                 ElementBlacklist = "a.jpg",
    ///                 DisableCompression = 1,
    ///                 IgnoreCertificateError = 0,
    ///                 MonitorTimeout = 20000,
    ///                 Redirection = 0,
    ///                 DnsHijackWhitelist = "www.aliyun.com:203.0.3.55",
    ///                 PageTamper = "www.aliyun.com:|/cc/bb/a.gif",
    ///                 FlowHijackJumpTimes = 10,
    ///                 CustomHeader = 1,
    ///                 DisableCache = 1,
    ///                 VerifyStringWhitelist = "Senyuan",
    ///                 TargetUrl = "http://www.aliyun.com",
    ///                 AutomaticScrolling = 1,
    ///                 WaitCompletionTime = 5005,
    ///                 FlowHijackLogo = "senyuan1",
    ///                 CustomHeaderContent = 
    ///                 {
    ///                     { "key1", "value1" },
    ///                 },
    ///                 FilterInvalidIp = 0,
    ///             },
    ///             FileDownload = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfFileDownloadArgs
    ///             {
    ///                 WhiteList = "www.aliyun.com:203.0.3.55",
    ///                 MonitorTimeout = 1050,
    ///                 IgnoreCertificateUntrustworthyError = 0,
    ///                 Redirection = 0,
    ///                 IgnoreCertificateCanceledError = 0,
    ///                 IgnoreCertificateAuthError = 0,
    ///                 IgnoreCertificateOutOfDateError = 0,
    ///                 IgnoreCertificateUsingError = 0,
    ///                 ConnectionTimeout = 6090,
    ///                 IgnoreInvalidHostError = 0,
    ///                 VerifyWay = 0,
    ///                 CustomHeaderContent = 
    ///                 {
    ///                     { "key1", "value1" },
    ///                 },
    ///                 TargetUrl = "https://www.aliyun.com",
    ///                 DownloadKernel = 0,
    ///                 QuickProtocol = 2,
    ///                 IgnoreCertificateStatusError = 1,
    ///                 TransmissionSize = 128,
    ///                 ValidateKeywords = "senyuan1",
    ///             },
    ///             Stream = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfStreamArgs
    ///             {
    ///                 StreamMonitorTimeout = 10,
    ///                 StreamAddressType = 0,
    ///                 PlayerType = 2,
    ///                 CustomHeaderContent = 
    ///                 {
    ///                     { "key1", "value1" },
    ///                 },
    ///                 WhiteList = "www.aliyun.com:203.0.3.55",
    ///                 TargetUrl = "https://acd-assets.alicdn.com:443/2021productweek/week1_s.mp4",
    ///                 StreamType = 1,
    ///             },
    ///             NetIcmp = new AliCloud.Arms.Inputs.SyntheticTaskMonitorConfNetIcmpArgs
    ///             {
    ///                 TargetUrl = "www.aliyun.com",
    ///                 Interval = 200,
    ///                 PackageNum = 36,
    ///                 PackageSize = 512,
    ///                 Timeout = 1000,
    ///                 TracertEnable = true,
    ///                 TracertNumMax = 1,
    ///                 TracertTimeout = 1200,
    ///             },
    ///         },
    ///         TaskType = 1,
    ///         Frequency = "1h",
    ///         MonitorCategory = 1,
    ///         CommonSetting = new AliCloud.Arms.Inputs.SyntheticTaskCommonSettingArgs
    ///         {
    ///             XtraceRegion = "cn-beijing",
    ///             CustomHost = new AliCloud.Arms.Inputs.SyntheticTaskCommonSettingCustomHostArgs
    ///             {
    ///                 Hosts = new[]
    ///                 {
    ///                     new AliCloud.Arms.Inputs.SyntheticTaskCommonSettingCustomHostHostArgs
    ///                     {
    ///                         Domain = "www.a.aliyun.com",
    ///                         Ips = new[]
    ///                         {
    ///                             "153.3.238.102",
    ///                         },
    ///                         IpType = 0,
    ///                     },
    ///                     new AliCloud.Arms.Inputs.SyntheticTaskCommonSettingCustomHostHostArgs
    ///                     {
    ///                         Domain = "www.shifen.com",
    ///                         Ips = new[]
    ///                         {
    ///                             "153.3.238.110",
    ///                             "114.114.114.114",
    ///                             "127.0.0.1",
    ///                         },
    ///                         IpType = 1,
    ///                     },
    ///                     new AliCloud.Arms.Inputs.SyntheticTaskCommonSettingCustomHostHostArgs
    ///                     {
    ///                         Domain = "www.aliyun.com",
    ///                         Ips = new[]
    ///                         {
    ///                             "153.3.238.110",
    ///                             "180.101.50.242",
    ///                             "180.101.50.188",
    ///                         },
    ///                         IpType = 0,
    ///                     },
    ///                 },
    ///                 SelectType = 1,
    ///             },
    ///             MonitorSamples = 1,
    ///             IpType = 1,
    ///             IsOpenTrace = true,
    ///             TraceClientType = 1,
    ///         },
    ///         ResourceGroupId = @default.Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0])),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ARMS Synthetic Task can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:arms/syntheticTask:SyntheticTask example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:arms/syntheticTask:SyntheticTask")]
    public partial class SyntheticTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Assertion List. See `available_assertions` below.
        /// </summary>
        [Output("availableAssertions")]
        public Output<ImmutableArray<Outputs.SyntheticTaskAvailableAssertion>> AvailableAssertions { get; private set; } = null!;

        /// <summary>
        /// Common settings. See `common_setting` below.
        /// </summary>
        [Output("commonSetting")]
        public Output<Outputs.SyntheticTaskCommonSetting> CommonSetting { get; private set; } = null!;

        /// <summary>
        /// Custom Cycle. See `custom_period` below.
        /// </summary>
        [Output("customPeriod")]
        public Output<Outputs.SyntheticTaskCustomPeriod?> CustomPeriod { get; private set; } = null!;

        /// <summary>
        /// Frequency.
        /// </summary>
        [Output("frequency")]
        public Output<string> Frequency { get; private set; } = null!;

        /// <summary>
        /// Classification of selected monitors.
        /// </summary>
        [Output("monitorCategory")]
        public Output<int> MonitorCategory { get; private set; } = null!;

        /// <summary>
        /// Monitoring configuration. See `monitor_conf` below.
        /// </summary>
        [Output("monitorConf")]
        public Output<Outputs.SyntheticTaskMonitorConf> MonitorConf { get; private set; } = null!;

        /// <summary>
        /// List of selected monitors. See `monitors` below.
        /// </summary>
        [Output("monitors")]
        public Output<ImmutableArray<Outputs.SyntheticTaskMonitor>> Monitors { get; private set; } = null!;

        /// <summary>
        /// Describes which resource group the resource belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// task status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The name of synthetic task.
        /// </summary>
        [Output("syntheticTaskName")]
        public Output<string> SyntheticTaskName { get; private set; } = null!;

        /// <summary>
        /// The list of tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of synthetic task.
        /// </summary>
        [Output("taskType")]
        public Output<int> TaskType { get; private set; } = null!;


        /// <summary>
        /// Create a SyntheticTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyntheticTask(string name, SyntheticTaskArgs args, CustomResourceOptions? options = null)
            : base("alicloud:arms/syntheticTask:SyntheticTask", name, args ?? new SyntheticTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyntheticTask(string name, Input<string> id, SyntheticTaskState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:arms/syntheticTask:SyntheticTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SyntheticTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyntheticTask Get(string name, Input<string> id, SyntheticTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new SyntheticTask(name, id, state, options);
        }
    }

    public sealed class SyntheticTaskArgs : global::Pulumi.ResourceArgs
    {
        [Input("availableAssertions")]
        private InputList<Inputs.SyntheticTaskAvailableAssertionArgs>? _availableAssertions;

        /// <summary>
        /// Assertion List. See `available_assertions` below.
        /// </summary>
        public InputList<Inputs.SyntheticTaskAvailableAssertionArgs> AvailableAssertions
        {
            get => _availableAssertions ?? (_availableAssertions = new InputList<Inputs.SyntheticTaskAvailableAssertionArgs>());
            set => _availableAssertions = value;
        }

        /// <summary>
        /// Common settings. See `common_setting` below.
        /// </summary>
        [Input("commonSetting")]
        public Input<Inputs.SyntheticTaskCommonSettingArgs>? CommonSetting { get; set; }

        /// <summary>
        /// Custom Cycle. See `custom_period` below.
        /// </summary>
        [Input("customPeriod")]
        public Input<Inputs.SyntheticTaskCustomPeriodArgs>? CustomPeriod { get; set; }

        /// <summary>
        /// Frequency.
        /// </summary>
        [Input("frequency", required: true)]
        public Input<string> Frequency { get; set; } = null!;

        /// <summary>
        /// Classification of selected monitors.
        /// </summary>
        [Input("monitorCategory", required: true)]
        public Input<int> MonitorCategory { get; set; } = null!;

        /// <summary>
        /// Monitoring configuration. See `monitor_conf` below.
        /// </summary>
        [Input("monitorConf", required: true)]
        public Input<Inputs.SyntheticTaskMonitorConfArgs> MonitorConf { get; set; } = null!;

        [Input("monitors", required: true)]
        private InputList<Inputs.SyntheticTaskMonitorArgs>? _monitors;

        /// <summary>
        /// List of selected monitors. See `monitors` below.
        /// </summary>
        public InputList<Inputs.SyntheticTaskMonitorArgs> Monitors
        {
            get => _monitors ?? (_monitors = new InputList<Inputs.SyntheticTaskMonitorArgs>());
            set => _monitors = value;
        }

        /// <summary>
        /// Describes which resource group the resource belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// task status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of synthetic task.
        /// </summary>
        [Input("syntheticTaskName", required: true)]
        public Input<string> SyntheticTaskName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The list of tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of synthetic task.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<int> TaskType { get; set; } = null!;

        public SyntheticTaskArgs()
        {
        }
        public static new SyntheticTaskArgs Empty => new SyntheticTaskArgs();
    }

    public sealed class SyntheticTaskState : global::Pulumi.ResourceArgs
    {
        [Input("availableAssertions")]
        private InputList<Inputs.SyntheticTaskAvailableAssertionGetArgs>? _availableAssertions;

        /// <summary>
        /// Assertion List. See `available_assertions` below.
        /// </summary>
        public InputList<Inputs.SyntheticTaskAvailableAssertionGetArgs> AvailableAssertions
        {
            get => _availableAssertions ?? (_availableAssertions = new InputList<Inputs.SyntheticTaskAvailableAssertionGetArgs>());
            set => _availableAssertions = value;
        }

        /// <summary>
        /// Common settings. See `common_setting` below.
        /// </summary>
        [Input("commonSetting")]
        public Input<Inputs.SyntheticTaskCommonSettingGetArgs>? CommonSetting { get; set; }

        /// <summary>
        /// Custom Cycle. See `custom_period` below.
        /// </summary>
        [Input("customPeriod")]
        public Input<Inputs.SyntheticTaskCustomPeriodGetArgs>? CustomPeriod { get; set; }

        /// <summary>
        /// Frequency.
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        /// <summary>
        /// Classification of selected monitors.
        /// </summary>
        [Input("monitorCategory")]
        public Input<int>? MonitorCategory { get; set; }

        /// <summary>
        /// Monitoring configuration. See `monitor_conf` below.
        /// </summary>
        [Input("monitorConf")]
        public Input<Inputs.SyntheticTaskMonitorConfGetArgs>? MonitorConf { get; set; }

        [Input("monitors")]
        private InputList<Inputs.SyntheticTaskMonitorGetArgs>? _monitors;

        /// <summary>
        /// List of selected monitors. See `monitors` below.
        /// </summary>
        public InputList<Inputs.SyntheticTaskMonitorGetArgs> Monitors
        {
            get => _monitors ?? (_monitors = new InputList<Inputs.SyntheticTaskMonitorGetArgs>());
            set => _monitors = value;
        }

        /// <summary>
        /// Describes which resource group the resource belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// task status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of synthetic task.
        /// </summary>
        [Input("syntheticTaskName")]
        public Input<string>? SyntheticTaskName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The list of tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of synthetic task.
        /// </summary>
        [Input("taskType")]
        public Input<int>? TaskType { get; set; }

        public SyntheticTaskState()
        {
        }
        public static new SyntheticTaskState Empty => new SyntheticTaskState();
    }
}
