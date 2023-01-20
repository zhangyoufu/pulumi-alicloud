// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas
{
    /// <summary>
    /// Provides a Nas Smb Acl resource.
    /// 
    /// Alibaba Cloud SMB protocol file storage service supports user authentication based on AD domain system and permission access control at the file system level. Connecting and accessing the SMB file system as a domain user can implement the requirements for access control at the file and directory level in the SMB protocol file system. The current Alibaba Cloud SMB protocol file storage service does not support multi-user file and directory-level permission access control, and only provides file system-level authentication and access based on the whitelist mechanism that supports cloud accounts and source IP permission groups control.
    /// &gt; **NOTE:** Available in 1.186.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleFileSystem = new AliCloud.Nas.FileSystem("exampleFileSystem", new()
    ///     {
    ///         ProtocolType = "SMB",
    ///         StorageType = "Performance",
    ///         FileSystemType = "standard",
    ///         ZoneId = "cn-hangzhou-g",
    ///     });
    /// 
    ///     var exampleSmbAclAttachment = new AliCloud.Nas.SmbAclAttachment("exampleSmbAclAttachment", new()
    ///     {
    ///         FileSystemId = exampleFileSystem.Id,
    ///         Keytab = "BQIAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAQAIqIx6v7p11oUAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAwAIqIx6v7p11oUAAABPAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAFwAQnQZWB3RAPHU7PMIJyBWePAAAAF8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQASACAGJ7F0s+bcBjf6jD5HlvlRLmPSOW+qDZe0Qk0lQcf8WwAAAE8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQARABDdFmanrSIatnDDhxxxxx",
    ///         KeytabMd5 = "E3CCF7E2416DF04FA958AA4513*****",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nas/smbAclAttachment:SmbAclAttachment")]
    public partial class SmbAclAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The method that is used to authenticate network identities.
        /// </summary>
        [Output("authMethod")]
        public Output<string> AuthMethod { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to allow anonymous access. Valid values:
        /// true: The file system allows anonymous access.
        /// false: The file system denies anonymous access. Default value: false.
        /// </summary>
        [Output("enableAnonymousAccess")]
        public Output<bool?> EnableAnonymousAccess { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable the ACL feature.
        /// true: enables the ACL feature.
        /// false: disables the ACL feature.
        /// </summary>
        [Output("enabled")]
        public Output<string> Enabled { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable encryption in transit. Valid values:
        /// true: enables encryption in transit.
        /// false: disables encryption in transit. Default value: false.
        /// </summary>
        [Output("encryptData")]
        public Output<bool?> EncryptData { get; private set; } = null!;

        /// <summary>
        /// The ID of the file system.
        /// </summary>
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// The home directory of each user. Each user-specific home directory must meet the following requirements:    
        /// Each segment starts with a forward slash (/) or a backslash (\).
        /// Each segment does not contain the following special characters: &lt;&gt;":?*.
        /// Each segment is 0 to 255 characters in length.
        /// The total length is 0 to 32,767 characters.
        /// For example, if you create a user named A and the home directory is /home, the file system automatically creates a directory named /home/A when User A logs on to the file system. If the /home/A directory already exists, the file system does not create the directory.
        /// </summary>
        [Output("homeDirPath")]
        public Output<string?> HomeDirPath { get; private set; } = null!;

        /// <summary>
        /// The string that is generated after the system encodes the keytab file by using Base64.
        /// </summary>
        [Output("keytab")]
        public Output<string> Keytab { get; private set; } = null!;

        /// <summary>
        /// RThe string that is generated after the system encodes the keytab file by using MD5.
        /// </summary>
        [Output("keytabMd5")]
        public Output<string> KeytabMd5 { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to deny access from non-encrypted clients. Valid values:
        /// true: The file system denies access from non-encrypted clients.
        /// false: The file system allows access from non-encrypted clients. Default value: false.
        /// </summary>
        [Output("rejectUnencryptedAccess")]
        public Output<bool?> RejectUnencryptedAccess { get; private set; } = null!;

        /// <summary>
        /// The ID of a super admin. The ID must meet the following requirements:
        /// The ID starts with S and does not contain letters except S.
        /// The ID contains at least three hyphens (-) as delimiters.
        /// Example: S-1-5-22 and S-1-5-22-23.
        /// </summary>
        [Output("superAdminSid")]
        public Output<string?> SuperAdminSid { get; private set; } = null!;


        /// <summary>
        /// Create a SmbAclAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmbAclAttachment(string name, SmbAclAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nas/smbAclAttachment:SmbAclAttachment", name, args ?? new SmbAclAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmbAclAttachment(string name, Input<string> id, SmbAclAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nas/smbAclAttachment:SmbAclAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SmbAclAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmbAclAttachment Get(string name, Input<string> id, SmbAclAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new SmbAclAttachment(name, id, state, options);
        }
    }

    public sealed class SmbAclAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to allow anonymous access. Valid values:
        /// true: The file system allows anonymous access.
        /// false: The file system denies anonymous access. Default value: false.
        /// </summary>
        [Input("enableAnonymousAccess")]
        public Input<bool>? EnableAnonymousAccess { get; set; }

        /// <summary>
        /// Specifies whether to enable encryption in transit. Valid values:
        /// true: enables encryption in transit.
        /// false: disables encryption in transit. Default value: false.
        /// </summary>
        [Input("encryptData")]
        public Input<bool>? EncryptData { get; set; }

        /// <summary>
        /// The ID of the file system.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// The home directory of each user. Each user-specific home directory must meet the following requirements:    
        /// Each segment starts with a forward slash (/) or a backslash (\).
        /// Each segment does not contain the following special characters: &lt;&gt;":?*.
        /// Each segment is 0 to 255 characters in length.
        /// The total length is 0 to 32,767 characters.
        /// For example, if you create a user named A and the home directory is /home, the file system automatically creates a directory named /home/A when User A logs on to the file system. If the /home/A directory already exists, the file system does not create the directory.
        /// </summary>
        [Input("homeDirPath")]
        public Input<string>? HomeDirPath { get; set; }

        /// <summary>
        /// The string that is generated after the system encodes the keytab file by using Base64.
        /// </summary>
        [Input("keytab", required: true)]
        public Input<string> Keytab { get; set; } = null!;

        /// <summary>
        /// RThe string that is generated after the system encodes the keytab file by using MD5.
        /// </summary>
        [Input("keytabMd5", required: true)]
        public Input<string> KeytabMd5 { get; set; } = null!;

        /// <summary>
        /// Specifies whether to deny access from non-encrypted clients. Valid values:
        /// true: The file system denies access from non-encrypted clients.
        /// false: The file system allows access from non-encrypted clients. Default value: false.
        /// </summary>
        [Input("rejectUnencryptedAccess")]
        public Input<bool>? RejectUnencryptedAccess { get; set; }

        /// <summary>
        /// The ID of a super admin. The ID must meet the following requirements:
        /// The ID starts with S and does not contain letters except S.
        /// The ID contains at least three hyphens (-) as delimiters.
        /// Example: S-1-5-22 and S-1-5-22-23.
        /// </summary>
        [Input("superAdminSid")]
        public Input<string>? SuperAdminSid { get; set; }

        public SmbAclAttachmentArgs()
        {
        }
        public static new SmbAclAttachmentArgs Empty => new SmbAclAttachmentArgs();
    }

    public sealed class SmbAclAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method that is used to authenticate network identities.
        /// </summary>
        [Input("authMethod")]
        public Input<string>? AuthMethod { get; set; }

        /// <summary>
        /// Specifies whether to allow anonymous access. Valid values:
        /// true: The file system allows anonymous access.
        /// false: The file system denies anonymous access. Default value: false.
        /// </summary>
        [Input("enableAnonymousAccess")]
        public Input<bool>? EnableAnonymousAccess { get; set; }

        /// <summary>
        /// Specifies whether to enable the ACL feature.
        /// true: enables the ACL feature.
        /// false: disables the ACL feature.
        /// </summary>
        [Input("enabled")]
        public Input<string>? Enabled { get; set; }

        /// <summary>
        /// Specifies whether to enable encryption in transit. Valid values:
        /// true: enables encryption in transit.
        /// false: disables encryption in transit. Default value: false.
        /// </summary>
        [Input("encryptData")]
        public Input<bool>? EncryptData { get; set; }

        /// <summary>
        /// The ID of the file system.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// The home directory of each user. Each user-specific home directory must meet the following requirements:    
        /// Each segment starts with a forward slash (/) or a backslash (\).
        /// Each segment does not contain the following special characters: &lt;&gt;":?*.
        /// Each segment is 0 to 255 characters in length.
        /// The total length is 0 to 32,767 characters.
        /// For example, if you create a user named A and the home directory is /home, the file system automatically creates a directory named /home/A when User A logs on to the file system. If the /home/A directory already exists, the file system does not create the directory.
        /// </summary>
        [Input("homeDirPath")]
        public Input<string>? HomeDirPath { get; set; }

        /// <summary>
        /// The string that is generated after the system encodes the keytab file by using Base64.
        /// </summary>
        [Input("keytab")]
        public Input<string>? Keytab { get; set; }

        /// <summary>
        /// RThe string that is generated after the system encodes the keytab file by using MD5.
        /// </summary>
        [Input("keytabMd5")]
        public Input<string>? KeytabMd5 { get; set; }

        /// <summary>
        /// Specifies whether to deny access from non-encrypted clients. Valid values:
        /// true: The file system denies access from non-encrypted clients.
        /// false: The file system allows access from non-encrypted clients. Default value: false.
        /// </summary>
        [Input("rejectUnencryptedAccess")]
        public Input<bool>? RejectUnencryptedAccess { get; set; }

        /// <summary>
        /// The ID of a super admin. The ID must meet the following requirements:
        /// The ID starts with S and does not contain letters except S.
        /// The ID contains at least three hyphens (-) as delimiters.
        /// Example: S-1-5-22 and S-1-5-22-23.
        /// </summary>
        [Input("superAdminSid")]
        public Input<string>? SuperAdminSid { get; set; }

        public SmbAclAttachmentState()
        {
        }
        public static new SmbAclAttachmentState Empty => new SmbAclAttachmentState();
    }
}
