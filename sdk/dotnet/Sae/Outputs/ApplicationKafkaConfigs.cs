// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Outputs
{

    [OutputType]
    public sealed class ApplicationKafkaConfigs
    {
        /// <summary>
        /// One or more logging configurations of ApsaraMQ for Kafka. See `kafka_configs` below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationKafkaConfigsKafkaConfig> KafkaConfigs;
        /// <summary>
        /// The endpoint of the ApsaraMQ for Kafka API.
        /// </summary>
        public readonly string? KafkaEndpoint;
        /// <summary>
        /// The  ID of the ApsaraMQ for Kafka instance.
        /// </summary>
        public readonly string? KafkaInstanceId;

        [OutputConstructor]
        private ApplicationKafkaConfigs(
            ImmutableArray<Outputs.ApplicationKafkaConfigsKafkaConfig> kafkaConfigs,

            string? kafkaEndpoint,

            string? kafkaInstanceId)
        {
            KafkaConfigs = kafkaConfigs;
            KafkaEndpoint = kafkaEndpoint;
            KafkaInstanceId = kafkaInstanceId;
        }
    }
}
