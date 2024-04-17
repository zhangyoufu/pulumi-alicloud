// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    /// <summary>
    /// Provides a RDS Account resource.
    /// 
    /// For information about RDS Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createaccount).
    /// 
    /// &gt; **NOTE:** Available since v1.120.0.
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
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var @default = AliCloud.Rds.GetZones.Invoke(new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "5.6",
    ///     });
    /// 
    ///     var defaultGetInstanceClasses = AliCloud.Rds.GetInstanceClasses.Invoke(new()
    ///     {
    ///         ZoneId = @default.Apply(getZonesResult =&gt; getZonesResult.Ids[0]),
    ///         Engine = "MySQL",
    ///         EngineVersion = "5.6",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Ids[0])),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Rds.Instance("default", new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "5.6",
    ///         InstanceType = defaultGetInstanceClasses.Apply(getInstanceClassesResult =&gt; getInstanceClassesResult.InstanceClasses[0]?.InstanceClass),
    ///         InstanceStorage = 10,
    ///         VswitchId = defaultSwitch.Id,
    ///         InstanceName = name,
    ///     });
    /// 
    ///     var defaultRdsAccount = new AliCloud.Rds.RdsAccount("default", new()
    ///     {
    ///         DbInstanceId = defaultInstance.Id,
    ///         AccountName = name,
    ///         AccountPassword = "Example1234",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// RDS Account can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:rds/rdsAccount:RdsAccount example &lt;db_instance_id&gt;:&lt;account_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rds/rdsAccount:RdsAccount")]
    public partial class RdsAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        /// </summary>
        [Output("accountDescription")]
        public Output<string> AccountDescription { get; private set; } = null!;

        /// <summary>
        /// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Output("accountPassword")]
        public Output<string> AccountPassword { get; private set; } = null!;

        /// <summary>
        /// Privilege type of account. Default to `Normal`.
        /// `Normal`: Common privilege.
        /// `Super`: High privilege.
        /// </summary>
        [Output("accountType")]
        public Output<string> AccountType { get; private set; } = null!;

        /// <summary>
        /// The Id of instance in which account belongs.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_description` instead.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `db_instance_id` instead.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        /// </summary>
        [Output("kmsEncryptedPassword")]
        public Output<string?> KmsEncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        [Output("kmsEncryptionContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_password` instead.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
        /// </summary>
        [Output("resetPermissionFlag")]
        public Output<bool?> ResetPermissionFlag { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. Valid values: `Available`, `Unavailable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_type` instead.
        /// 
        /// &gt; **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RdsAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsAccount(string name, RdsAccountArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/rdsAccount:RdsAccount", name, args ?? new RdsAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsAccount(string name, Input<string> id, RdsAccountState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/rdsAccount:RdsAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accountPassword",
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RdsAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsAccount Get(string name, Input<string> id, RdsAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsAccount(name, id, state, options);
        }
    }

    public sealed class RdsAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("accountPassword")]
        private Input<string>? _accountPassword;

        /// <summary>
        /// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Privilege type of account. Default to `Normal`.
        /// `Normal`: Common privilege.
        /// `Super`: High privilege.
        /// </summary>
        [Input("accountType")]
        public Input<string>? AccountType { get; set; }

        /// <summary>
        /// The Id of instance in which account belongs.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_description` instead.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `db_instance_id` instead.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_password` instead.
        /// </summary>
        [Obsolete(@"Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.")]
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
        /// </summary>
        [Input("resetPermissionFlag")]
        public Input<bool>? ResetPermissionFlag { get; set; }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_type` instead.
        /// 
        /// &gt; **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public RdsAccountArgs()
        {
        }
        public static new RdsAccountArgs Empty => new RdsAccountArgs();
    }

    public sealed class RdsAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("accountPassword")]
        private Input<string>? _accountPassword;

        /// <summary>
        /// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        public Input<string>? AccountPassword
        {
            get => _accountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Privilege type of account. Default to `Normal`.
        /// `Normal`: Common privilege.
        /// `Super`: High privilege.
        /// </summary>
        [Input("accountType")]
        public Input<string>? AccountType { get; set; }

        /// <summary>
        /// The Id of instance in which account belongs.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_description` instead.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `db_instance_id` instead.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a db account. If the `account_password` is filled in, this field will be ignored.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a db account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_password` instead.
        /// </summary>
        [Obsolete(@"Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.")]
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
        /// </summary>
        [Input("resetPermissionFlag")]
        public Input<bool>? ResetPermissionFlag { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `Available`, `Unavailable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The attribute has been deprecated from 1.120.0 and using `account_type` instead.
        /// 
        /// &gt; **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public RdsAccountState()
        {
        }
        public static new RdsAccountState Empty => new RdsAccountState();
    }
}
