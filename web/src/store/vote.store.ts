import {ActionTree, GetterTree, Module, MutationTree} from "vuex"
import {RootState} from "@/store/rootState"

import Vue from 'vue'
import {Topic, RequestByID, Topics} from "../../gen/pb/vote_pb"
import {VoteClient} from "../../gen/pb/VoteServiceClientPb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

const service = new VoteClient('http://localhost:9090', null, null)

export interface LocalTopic {
  id: number;
  message: string;
  voteCount: number;
}

export interface VoteState {
  topics: LocalTopic[];
}

const state: VoteState = {
  topics: []
}

const getters: GetterTree<VoteState, RootState> = {
  topics: state => state.topics
}

const actions: ActionTree<VoteState, RootState> = {

  async fetchTopics({ commit }) {
    const empty = new google_protobuf_empty_pb.Empty()
    const topics = await service.getTopics(empty, {})
    commit('SET_TOPICS', topics)
  },

  async voteTopic({ commit }, topicId: number) {
    const r = new RequestByID().setId(topicId)
    const topic = await service.voteTopic(r, {})
    commit('UPDATE_TOPIC', topic)
  },

}

const mutations: MutationTree<VoteState> = {
  SET_TOPICS(state, topics: Topics) {
    state.topics = []
    for (const t of topics.getTopicsList()) {
      state.topics.push({
        id: t.getId(),
        message: t.getMessage(),
        voteCount: t.getVotecount(),
      })
    }
  },
  UPDATE_TOPIC(state, topic: Topic) {
    const index = state.topics.findIndex(t => t.id === topic.getId())

    if(index !== -1) {
      Vue.set(state.topics, index, {
        id: topic.getId(),
        message: topic.getMessage(),
        voteCount: topic.getVotecount(),
      })
    }
  },
}

const voteStore: Module<VoteState, RootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
}

export default voteStore
