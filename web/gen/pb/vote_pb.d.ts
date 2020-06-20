import * as jspb from "google-protobuf"

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';

export class Topic extends jspb.Message {
  getId(): number;
  setId(value: number): Topic;

  getMessage(): string;
  setMessage(value: string): Topic;

  getVotecount(): number;
  setVotecount(value: number): Topic;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Topic.AsObject;
  static toObject(includeInstance: boolean, msg: Topic): Topic.AsObject;
  static serializeBinaryToWriter(message: Topic, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Topic;
  static deserializeBinaryFromReader(message: Topic, reader: jspb.BinaryReader): Topic;
}

export namespace Topic {
  export type AsObject = {
    id: number,
    message: string,
    votecount: number,
  }
}

export class Topics extends jspb.Message {
  getTopicsList(): Array<Topic>;
  setTopicsList(value: Array<Topic>): Topics;
  clearTopicsList(): Topics;
  addTopics(value?: Topic, index?: number): Topic;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Topics.AsObject;
  static toObject(includeInstance: boolean, msg: Topics): Topics.AsObject;
  static serializeBinaryToWriter(message: Topics, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Topics;
  static deserializeBinaryFromReader(message: Topics, reader: jspb.BinaryReader): Topics;
}

export namespace Topics {
  export type AsObject = {
    topicsList: Array<Topic.AsObject>,
  }
}

export class RequestByID extends jspb.Message {
  getId(): number;
  setId(value: number): RequestByID;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RequestByID.AsObject;
  static toObject(includeInstance: boolean, msg: RequestByID): RequestByID.AsObject;
  static serializeBinaryToWriter(message: RequestByID, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RequestByID;
  static deserializeBinaryFromReader(message: RequestByID, reader: jspb.BinaryReader): RequestByID;
}

export namespace RequestByID {
  export type AsObject = {
    id: number,
  }
}

