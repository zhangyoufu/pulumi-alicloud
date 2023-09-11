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
    public sealed class DomainNewCertificateConfig
    {
        /// <summary>
        /// The ID of the certificate. It takes effect only when CertType = cas.
        /// </summary>
        public readonly string? CertId;
        /// <summary>
        /// Certificate name, only flyer names are supported.
        /// </summary>
        public readonly string? CertName;
        /// <summary>
        /// The certificate region, which takes effect only when CertType = cas, supports cn-hangzhou (domestic) and ap-southeast-1 (International), and is cn-hangzhou by default.
        /// </summary>
        public readonly string? CertRegion;
        /// <summary>
        /// Certificate type. Value:
        /// - **upload**: upload certificate.
        /// - **cas**: Cloud Shield certificate.
        /// - **free**: free certificate.
        /// &gt; If the certificate type is **cas**, **PrivateKey** does not need to pass parameters.
        /// </summary>
        public readonly string? CertType;
        /// <summary>
        /// The force set of the security certificate.
        /// </summary>
        public readonly string? ForceSet;
        /// <summary>
        /// The content of the private key. If the certificate is not enabled, you do not need to enter the content of the private key. To configure the certificate, enter the content of the private key.
        /// </summary>
        public readonly string? PrivateKey;
        /// <summary>
        /// The content of the security certificate. If the certificate is not enabled, you do not need to enter the content of the security certificate. Please enter the content of the certificate to configure the certificate.
        /// </summary>
        public readonly string? ServerCertificate;
        /// <summary>
        /// Whether the HTTPS certificate is enabled. Value:
        /// - **on**(default): enabled.
        /// - **off** : not enabled.
        /// </summary>
        public readonly string? ServerCertificateStatus;

        [OutputConstructor]
        private DomainNewCertificateConfig(
            string? certId,

            string? certName,

            string? certRegion,

            string? certType,

            string? forceSet,

            string? privateKey,

            string? serverCertificate,

            string? serverCertificateStatus)
        {
            CertId = certId;
            CertName = certName;
            CertRegion = certRegion;
            CertType = certType;
            ForceSet = forceSet;
            PrivateKey = privateKey;
            ServerCertificate = serverCertificate;
            ServerCertificateStatus = serverCertificateStatus;
        }
    }
}
