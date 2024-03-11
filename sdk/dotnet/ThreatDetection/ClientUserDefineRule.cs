// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    /// <summary>
    /// Provides a Threat Detection Client User Define Rule resource. Malicious Behavior Defense Custom Rules.
    /// 
    /// For information about Threat Detection Client User Define Rule and how to use it, see [What is Client User Define Rule](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-addclientuserdefinerule).
    /// 
    /// &gt; **NOTE:** Available since v1.212.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var @default = new AliCloud.ThreatDetection.ClientUserDefineRule("default", new()
    ///     {
    ///         ActionType = 0,
    ///         Platform = "windows",
    ///         RegistryContent = "123",
    ///         ClientUserDefineRuleName = name,
    ///         ParentProcPath = "/root/bash",
    ///         Type = 5,
    ///         Cmdline = "bash",
    ///         ProcPath = "/root/bash",
    ///         ParentCmdline = "bash",
    ///         RegistryKey = "123",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Threat Detection Client User Define Rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:threatdetection/clientUserDefineRule:ClientUserDefineRule example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:threatdetection/clientUserDefineRule:ClientUserDefineRule")]
    public partial class ClientUserDefineRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The operation type. Value:
        /// - **0**: plus White
        /// - **1**: Plus Black.
        /// </summary>
        [Output("actionType")]
        public Output<int> ActionType { get; private set; } = null!;

        /// <summary>
        /// The custom rule name.
        /// </summary>
        [Output("clientUserDefineRuleName")]
        public Output<string> ClientUserDefineRuleName { get; private set; } = null!;

        /// <summary>
        /// Command line. When the value of the Type attribute is 2, 3, 4, 5, 6, or 7, the command line field is required.
        /// </summary>
        [Output("cmdline")]
        public Output<string?> Cmdline { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<int> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The file path. When the value of the Type attribute is 4 or 6, 7, the FilePath field is required.
        /// </summary>
        [Output("filePath")]
        public Output<string?> FilePath { get; private set; } = null!;

        /// <summary>
        /// Process hash list. When the value of the Type attribute is 1, the Hash attribute is required.
        /// </summary>
        [Output("hash")]
        public Output<string?> Hash { get; private set; } = null!;

        /// <summary>
        /// IP address. When the value of the Type attribute is 3, the Ip attribute is required.
        /// </summary>
        [Output("ip")]
        public Output<string?> Ip { get; private set; } = null!;

        /// <summary>
        /// The new file path to rename the file. When the value of the Type attribute is 7, the NewFilePath attribute is required.
        /// </summary>
        [Output("newFilePath")]
        public Output<string?> NewFilePath { get; private set; } = null!;

        /// <summary>
        /// The parent command line.
        /// </summary>
        [Output("parentCmdline")]
        public Output<string?> ParentCmdline { get; private set; } = null!;

        /// <summary>
        /// Parent process path.
        /// </summary>
        [Output("parentProcPath")]
        public Output<string?> ParentProcPath { get; private set; } = null!;

        /// <summary>
        /// The operating system type. Value:
        /// - **windows**:widows
        /// - **linux**:linux
        /// - **all**: all.
        /// </summary>
        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        /// <summary>
        /// The port number. When the value of the Type attribute is 3, the PortStr attribute is required. Value range: **1-65535**.
        /// </summary>
        [Output("portStr")]
        public Output<string> PortStr { get; private set; } = null!;

        /// <summary>
        /// The process path. When the Type attribute is set to 2, 3, 4, 5, 6, or 7, the ProcPath attribute is required.
        /// </summary>
        [Output("procPath")]
        public Output<string?> ProcPath { get; private set; } = null!;

        /// <summary>
        /// The registry value. When the value of the Type attribute is 5, the RegistryKey attribute is required.
        /// </summary>
        [Output("registryContent")]
        public Output<string?> RegistryContent { get; private set; } = null!;

        /// <summary>
        /// The registry key. When the value of the Type attribute is 5, the RegistryKey attribute is required.
        /// </summary>
        [Output("registryKey")]
        public Output<string?> RegistryKey { get; private set; } = null!;

        /// <summary>
        /// The rule type. Value:
        /// - **1**: Process hash
        /// - **2**: command line
        /// - **3**: Process network
        /// - **4**: File reading and writing
        /// - **5**: Operate the registry
        /// - **6**: Load Dynamic Link Library
        /// - **7**: File Rename.
        /// </summary>
        [Output("type")]
        public Output<int> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ClientUserDefineRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientUserDefineRule(string name, ClientUserDefineRuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/clientUserDefineRule:ClientUserDefineRule", name, args ?? new ClientUserDefineRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientUserDefineRule(string name, Input<string> id, ClientUserDefineRuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/clientUserDefineRule:ClientUserDefineRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientUserDefineRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientUserDefineRule Get(string name, Input<string> id, ClientUserDefineRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientUserDefineRule(name, id, state, options);
        }
    }

    public sealed class ClientUserDefineRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operation type. Value:
        /// - **0**: plus White
        /// - **1**: Plus Black.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<int> ActionType { get; set; } = null!;

        /// <summary>
        /// The custom rule name.
        /// </summary>
        [Input("clientUserDefineRuleName", required: true)]
        public Input<string> ClientUserDefineRuleName { get; set; } = null!;

        /// <summary>
        /// Command line. When the value of the Type attribute is 2, 3, 4, 5, 6, or 7, the command line field is required.
        /// </summary>
        [Input("cmdline")]
        public Input<string>? Cmdline { get; set; }

        /// <summary>
        /// The file path. When the value of the Type attribute is 4 or 6, 7, the FilePath field is required.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// Process hash list. When the value of the Type attribute is 1, the Hash attribute is required.
        /// </summary>
        [Input("hash")]
        public Input<string>? Hash { get; set; }

        /// <summary>
        /// IP address. When the value of the Type attribute is 3, the Ip attribute is required.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The new file path to rename the file. When the value of the Type attribute is 7, the NewFilePath attribute is required.
        /// </summary>
        [Input("newFilePath")]
        public Input<string>? NewFilePath { get; set; }

        /// <summary>
        /// The parent command line.
        /// </summary>
        [Input("parentCmdline")]
        public Input<string>? ParentCmdline { get; set; }

        /// <summary>
        /// Parent process path.
        /// </summary>
        [Input("parentProcPath")]
        public Input<string>? ParentProcPath { get; set; }

        /// <summary>
        /// The operating system type. Value:
        /// - **windows**:widows
        /// - **linux**:linux
        /// - **all**: all.
        /// </summary>
        [Input("platform", required: true)]
        public Input<string> Platform { get; set; } = null!;

        /// <summary>
        /// The port number. When the value of the Type attribute is 3, the PortStr attribute is required. Value range: **1-65535**.
        /// </summary>
        [Input("portStr")]
        public Input<string>? PortStr { get; set; }

        /// <summary>
        /// The process path. When the Type attribute is set to 2, 3, 4, 5, 6, or 7, the ProcPath attribute is required.
        /// </summary>
        [Input("procPath")]
        public Input<string>? ProcPath { get; set; }

        /// <summary>
        /// The registry value. When the value of the Type attribute is 5, the RegistryKey attribute is required.
        /// </summary>
        [Input("registryContent")]
        public Input<string>? RegistryContent { get; set; }

        /// <summary>
        /// The registry key. When the value of the Type attribute is 5, the RegistryKey attribute is required.
        /// </summary>
        [Input("registryKey")]
        public Input<string>? RegistryKey { get; set; }

        /// <summary>
        /// The rule type. Value:
        /// - **1**: Process hash
        /// - **2**: command line
        /// - **3**: Process network
        /// - **4**: File reading and writing
        /// - **5**: Operate the registry
        /// - **6**: Load Dynamic Link Library
        /// - **7**: File Rename.
        /// </summary>
        [Input("type", required: true)]
        public Input<int> Type { get; set; } = null!;

        public ClientUserDefineRuleArgs()
        {
        }
        public static new ClientUserDefineRuleArgs Empty => new ClientUserDefineRuleArgs();
    }

    public sealed class ClientUserDefineRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operation type. Value:
        /// - **0**: plus White
        /// - **1**: Plus Black.
        /// </summary>
        [Input("actionType")]
        public Input<int>? ActionType { get; set; }

        /// <summary>
        /// The custom rule name.
        /// </summary>
        [Input("clientUserDefineRuleName")]
        public Input<string>? ClientUserDefineRuleName { get; set; }

        /// <summary>
        /// Command line. When the value of the Type attribute is 2, 3, 4, 5, 6, or 7, the command line field is required.
        /// </summary>
        [Input("cmdline")]
        public Input<string>? Cmdline { get; set; }

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<int>? CreateTime { get; set; }

        /// <summary>
        /// The file path. When the value of the Type attribute is 4 or 6, 7, the FilePath field is required.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// Process hash list. When the value of the Type attribute is 1, the Hash attribute is required.
        /// </summary>
        [Input("hash")]
        public Input<string>? Hash { get; set; }

        /// <summary>
        /// IP address. When the value of the Type attribute is 3, the Ip attribute is required.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The new file path to rename the file. When the value of the Type attribute is 7, the NewFilePath attribute is required.
        /// </summary>
        [Input("newFilePath")]
        public Input<string>? NewFilePath { get; set; }

        /// <summary>
        /// The parent command line.
        /// </summary>
        [Input("parentCmdline")]
        public Input<string>? ParentCmdline { get; set; }

        /// <summary>
        /// Parent process path.
        /// </summary>
        [Input("parentProcPath")]
        public Input<string>? ParentProcPath { get; set; }

        /// <summary>
        /// The operating system type. Value:
        /// - **windows**:widows
        /// - **linux**:linux
        /// - **all**: all.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The port number. When the value of the Type attribute is 3, the PortStr attribute is required. Value range: **1-65535**.
        /// </summary>
        [Input("portStr")]
        public Input<string>? PortStr { get; set; }

        /// <summary>
        /// The process path. When the Type attribute is set to 2, 3, 4, 5, 6, or 7, the ProcPath attribute is required.
        /// </summary>
        [Input("procPath")]
        public Input<string>? ProcPath { get; set; }

        /// <summary>
        /// The registry value. When the value of the Type attribute is 5, the RegistryKey attribute is required.
        /// </summary>
        [Input("registryContent")]
        public Input<string>? RegistryContent { get; set; }

        /// <summary>
        /// The registry key. When the value of the Type attribute is 5, the RegistryKey attribute is required.
        /// </summary>
        [Input("registryKey")]
        public Input<string>? RegistryKey { get; set; }

        /// <summary>
        /// The rule type. Value:
        /// - **1**: Process hash
        /// - **2**: command line
        /// - **3**: Process network
        /// - **4**: File reading and writing
        /// - **5**: Operate the registry
        /// - **6**: Load Dynamic Link Library
        /// - **7**: File Rename.
        /// </summary>
        [Input("type")]
        public Input<int>? Type { get; set; }

        public ClientUserDefineRuleState()
        {
        }
        public static new ClientUserDefineRuleState Empty => new ClientUserDefineRuleState();
    }
}
