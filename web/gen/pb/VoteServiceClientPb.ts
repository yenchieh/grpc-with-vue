/**
 * @fileoverview gRPC-Web generated client stub for vote
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';

import {
  RequestByID,
  Topic,
  Topics} from './vote_pb';

export class VoteClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetTopics = new grpcWeb.AbstractClientBase.MethodInfo(
    Topics,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    Topics.deserializeBinary
  );

  getTopics(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<Topics>;

  getTopics(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Topics) => void): grpcWeb.ClientReadableStream<Topics>;

  getTopics(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: Topics) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/vote.Vote/GetTopics',
        request,
        metadata || {},
        this.methodInfoGetTopics,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/vote.Vote/GetTopics',
    request,
    metadata || {},
    this.methodInfoGetTopics);
  }

  methodInfoVoteTopic = new grpcWeb.AbstractClientBase.MethodInfo(
    Topic,
    (request: RequestByID) => {
      return request.serializeBinary();
    },
    Topic.deserializeBinary
  );

  voteTopic(
    request: RequestByID,
    metadata: grpcWeb.Metadata | null): Promise<Topic>;

  voteTopic(
    request: RequestByID,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Topic) => void): grpcWeb.ClientReadableStream<Topic>;

  voteTopic(
    request: RequestByID,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: Topic) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/vote.Vote/VoteTopic',
        request,
        metadata || {},
        this.methodInfoVoteTopic,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/vote.Vote/VoteTopic',
    request,
    metadata || {},
    this.methodInfoVoteTopic);
  }

}

