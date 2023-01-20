// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetQueuesArgs, GetQueuesResult, GetQueuesOutputArgs } from "./getQueues";
export const getQueues: typeof import("./getQueues").getQueues = null as any;
export const getQueuesOutput: typeof import("./getQueues").getQueuesOutput = null as any;
utilities.lazyLoad(exports, ["getQueues","getQueuesOutput"], () => require("./getQueues"));

export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { GetTopicSubscriptionsArgs, GetTopicSubscriptionsResult, GetTopicSubscriptionsOutputArgs } from "./getTopicSubscriptions";
export const getTopicSubscriptions: typeof import("./getTopicSubscriptions").getTopicSubscriptions = null as any;
export const getTopicSubscriptionsOutput: typeof import("./getTopicSubscriptions").getTopicSubscriptionsOutput = null as any;
utilities.lazyLoad(exports, ["getTopicSubscriptions","getTopicSubscriptionsOutput"], () => require("./getTopicSubscriptions"));

export { GetTopicsArgs, GetTopicsResult, GetTopicsOutputArgs } from "./getTopics";
export const getTopics: typeof import("./getTopics").getTopics = null as any;
export const getTopicsOutput: typeof import("./getTopics").getTopicsOutput = null as any;
utilities.lazyLoad(exports, ["getTopics","getTopicsOutput"], () => require("./getTopics"));

export { QueueArgs, QueueState } from "./queue";
export type Queue = import("./queue").Queue;
export const Queue: typeof import("./queue").Queue = null as any;
utilities.lazyLoad(exports, ["Queue"], () => require("./queue"));

export { TopicArgs, TopicState } from "./topic";
export type Topic = import("./topic").Topic;
export const Topic: typeof import("./topic").Topic = null as any;
utilities.lazyLoad(exports, ["Topic"], () => require("./topic"));

export { TopicSubscriptionArgs, TopicSubscriptionState } from "./topicSubscription";
export type TopicSubscription = import("./topicSubscription").TopicSubscription;
export const TopicSubscription: typeof import("./topicSubscription").TopicSubscription = null as any;
utilities.lazyLoad(exports, ["TopicSubscription"], () => require("./topicSubscription"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:mns/queue:Queue":
                return new Queue(name, <any>undefined, { urn })
            case "alicloud:mns/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            case "alicloud:mns/topicSubscription:TopicSubscription":
                return new TopicSubscription(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "mns/queue", _module)
pulumi.runtime.registerResourceModule("alicloud", "mns/topic", _module)
pulumi.runtime.registerResourceModule("alicloud", "mns/topicSubscription", _module)
