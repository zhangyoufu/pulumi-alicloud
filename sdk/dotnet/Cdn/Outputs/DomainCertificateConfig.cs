// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainCertificateConfig
    {
        public readonly string? PrivateKey;
        public readonly string? ServerCertificate;
        public readonly string? ServerCertificateStatus;

        [OutputConstructor]
        private DomainCertificateConfig(
            string? privateKey,

            string? serverCertificate,

            string? serverCertificateStatus)
        {
            PrivateKey = privateKey;
            ServerCertificate = serverCertificate;
            ServerCertificateStatus = serverCertificateStatus;
        }
    }
}
