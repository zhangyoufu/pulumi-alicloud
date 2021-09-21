// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Outputs
{

    [OutputType]
    public sealed class GetMscSubSubscriptionsSubscriptionResult
    {
        /// <summary>
        /// The channel the Subscription.
        /// </summary>
        public readonly string Channel;
        /// <summary>
        /// The ids of subscribed contacts.
        /// </summary>
        public readonly ImmutableArray<int> ContactIds;
        /// <summary>
        /// The description of the Subscription.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        /// </summary>
        public readonly int EmailStatus;
        /// <summary>
        /// The ID of the Subscription.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the Subscription.
        /// </summary>
        public readonly string ItemId;
        /// <summary>
        /// The name of the Subscription.
        /// </summary>
        public readonly string ItemName;
        /// <summary>
        /// The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        /// </summary>
        public readonly int PmsgStatus;
        /// <summary>
        /// The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        /// </summary>
        public readonly int SmsStatus;
        /// <summary>
        /// The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        /// </summary>
        public readonly int TtsStatus;
        /// <summary>
        /// The ids of subscribed webhooks.
        /// </summary>
        public readonly ImmutableArray<int> WebhookIds;
        /// <summary>
        /// The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
        /// </summary>
        public readonly int WebhookStatus;

        [OutputConstructor]
        private GetMscSubSubscriptionsSubscriptionResult(
            string channel,

            ImmutableArray<int> contactIds,

            string description,

            int emailStatus,

            string id,

            string itemId,

            string itemName,

            int pmsgStatus,

            int smsStatus,

            int ttsStatus,

            ImmutableArray<int> webhookIds,

            int webhookStatus)
        {
            Channel = channel;
            ContactIds = contactIds;
            Description = description;
            EmailStatus = emailStatus;
            Id = id;
            ItemId = itemId;
            ItemName = itemName;
            PmsgStatus = pmsgStatus;
            SmsStatus = smsStatus;
            TtsStatus = ttsStatus;
            WebhookIds = webhookIds;
            WebhookStatus = webhookStatus;
        }
    }
}
