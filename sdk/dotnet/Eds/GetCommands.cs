// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds
{
    public static class GetCommands
    {
        /// <summary>
        /// This data source provides the Ecd Commands of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.146.0+.
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
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultSimpleOfficeSite = new AliCloud.Eds.SimpleOfficeSite("default", new()
        ///     {
        ///         CidrBlock = "172.16.0.0/12",
        ///         DesktopAccessType = "Internet",
        ///         OfficeSiteName = "your_office_site_name",
        ///     });
        /// 
        ///     var @default = AliCloud.Eds.GetBundles.Invoke(new()
        ///     {
        ///         BundleType = "SYSTEM",
        ///         NameRegex = "windows",
        ///     });
        /// 
        ///     var defaultEcdPolicyGroup = new AliCloud.Eds.EcdPolicyGroup("default", new()
        ///     {
        ///         PolicyGroupName = "your_policy_group_name",
        ///         Clipboard = "readwrite",
        ///         LocalDrive = "read",
        ///         AuthorizeAccessPolicyRules = new[]
        ///         {
        ///             new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs
        ///             {
        ///                 Description = "example_value",
        ///                 CidrIp = "1.2.3.4/24",
        ///             },
        ///         },
        ///         AuthorizeSecurityPolicyRules = new[]
        ///         {
        ///             new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs
        ///             {
        ///                 Type = "inflow",
        ///                 Policy = "accept",
        ///                 Description = "example_value",
        ///                 PortRange = "80/80",
        ///                 IpProtocol = "TCP",
        ///                 Priority = "1",
        ///                 CidrIp = "0.0.0.0/0",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var defaultDesktop = new AliCloud.Eds.Desktop("default", new()
        ///     {
        ///         OfficeSiteId = defaultSimpleOfficeSite.Id,
        ///         PolicyGroupId = defaultEcdPolicyGroup.Id,
        ///         BundleId = @default.Apply(@default =&gt; @default.Apply(getBundlesResult =&gt; getBundlesResult.Bundles[0]?.Id)),
        ///         DesktopName = name,
        ///     });
        /// 
        ///     var defaultCommand = new AliCloud.Eds.Command("default", new()
        ///     {
        ///         CommandContent = "ipconfig",
        ///         CommandType = "RunPowerShellScript",
        ///         DesktopId = defaultDesktop.Id,
        ///     });
        /// 
        ///     var ids = AliCloud.Eds.GetCommands.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecdCommandId1"] = ids.Apply(getCommandsResult =&gt; getCommandsResult.Commands[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCommandsResult> InvokeAsync(GetCommandsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCommandsResult>("alicloud:eds/getCommands:getCommands", args ?? new GetCommandsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Ecd Commands of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.146.0+.
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
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultSimpleOfficeSite = new AliCloud.Eds.SimpleOfficeSite("default", new()
        ///     {
        ///         CidrBlock = "172.16.0.0/12",
        ///         DesktopAccessType = "Internet",
        ///         OfficeSiteName = "your_office_site_name",
        ///     });
        /// 
        ///     var @default = AliCloud.Eds.GetBundles.Invoke(new()
        ///     {
        ///         BundleType = "SYSTEM",
        ///         NameRegex = "windows",
        ///     });
        /// 
        ///     var defaultEcdPolicyGroup = new AliCloud.Eds.EcdPolicyGroup("default", new()
        ///     {
        ///         PolicyGroupName = "your_policy_group_name",
        ///         Clipboard = "readwrite",
        ///         LocalDrive = "read",
        ///         AuthorizeAccessPolicyRules = new[]
        ///         {
        ///             new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs
        ///             {
        ///                 Description = "example_value",
        ///                 CidrIp = "1.2.3.4/24",
        ///             },
        ///         },
        ///         AuthorizeSecurityPolicyRules = new[]
        ///         {
        ///             new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs
        ///             {
        ///                 Type = "inflow",
        ///                 Policy = "accept",
        ///                 Description = "example_value",
        ///                 PortRange = "80/80",
        ///                 IpProtocol = "TCP",
        ///                 Priority = "1",
        ///                 CidrIp = "0.0.0.0/0",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var defaultDesktop = new AliCloud.Eds.Desktop("default", new()
        ///     {
        ///         OfficeSiteId = defaultSimpleOfficeSite.Id,
        ///         PolicyGroupId = defaultEcdPolicyGroup.Id,
        ///         BundleId = @default.Apply(@default =&gt; @default.Apply(getBundlesResult =&gt; getBundlesResult.Bundles[0]?.Id)),
        ///         DesktopName = name,
        ///     });
        /// 
        ///     var defaultCommand = new AliCloud.Eds.Command("default", new()
        ///     {
        ///         CommandContent = "ipconfig",
        ///         CommandType = "RunPowerShellScript",
        ///         DesktopId = defaultDesktop.Id,
        ///     });
        /// 
        ///     var ids = AliCloud.Eds.GetCommands.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ecdCommandId1"] = ids.Apply(getCommandsResult =&gt; getCommandsResult.Commands[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCommandsResult> Invoke(GetCommandsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCommandsResult>("alicloud:eds/getCommands:getCommands", args ?? new GetCommandsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCommandsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
        /// </summary>
        [Input("commandType")]
        public string? CommandType { get; set; }

        /// <summary>
        /// That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
        /// </summary>
        [Input("contentEncoding")]
        public string? ContentEncoding { get; set; }

        /// <summary>
        /// The desktop id of the Desktop.
        /// </summary>
        [Input("desktopId")]
        public string? DesktopId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Command IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetCommandsArgs()
        {
        }
        public static new GetCommandsArgs Empty => new GetCommandsArgs();
    }

    public sealed class GetCommandsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
        /// </summary>
        [Input("commandType")]
        public Input<string>? CommandType { get; set; }

        /// <summary>
        /// That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// The desktop id of the Desktop.
        /// </summary>
        [Input("desktopId")]
        public Input<string>? DesktopId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Command IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetCommandsInvokeArgs()
        {
        }
        public static new GetCommandsInvokeArgs Empty => new GetCommandsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCommandsResult
    {
        public readonly string? CommandType;
        public readonly ImmutableArray<Outputs.GetCommandsCommandResult> Commands;
        public readonly string? ContentEncoding;
        public readonly string? DesktopId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetCommandsResult(
            string? commandType,

            ImmutableArray<Outputs.GetCommandsCommandResult> commands,

            string? contentEncoding,

            string? desktopId,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? status)
        {
            CommandType = commandType;
            Commands = commands;
            ContentEncoding = contentEncoding;
            DesktopId = desktopId;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
