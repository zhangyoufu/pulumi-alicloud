// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Amqp.Outputs
{

    [OutputType]
    public sealed class GetExchangesExchangeResult
    {
        /// <summary>
        /// The attributes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Attributes;
        /// <summary>
        /// Indicates whether the Auto Delete attribute is configured.
        /// </summary>
        public readonly bool AutoDeleteState;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The name of the exchange.
        /// </summary>
        public readonly string ExchangeName;
        /// <summary>
        /// The type of the exchange.
        /// </summary>
        public readonly string ExchangeType;
        /// <summary>
        /// The ID of the Exchange. Its value is same as Queue Name.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The name of virtual host where an exchange resides.
        /// </summary>
        public readonly string VirtualHostName;

        [OutputConstructor]
        private GetExchangesExchangeResult(
            ImmutableDictionary<string, string> attributes,

            bool autoDeleteState,

            string createTime,

            string exchangeName,

            string exchangeType,

            string id,

            string instanceId,

            string virtualHostName)
        {
            Attributes = attributes;
            AutoDeleteState = autoDeleteState;
            CreateTime = createTime;
            ExchangeName = exchangeName;
            ExchangeType = exchangeType;
            Id = id;
            InstanceId = instanceId;
            VirtualHostName = virtualHostName;
        }
    }
}
