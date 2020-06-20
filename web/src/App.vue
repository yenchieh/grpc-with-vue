<template>
  <div id="app">
    <h1>gRPC + Go + Vue + Typescript demo</h1>
    <span v-if="loading">
      Loading...
    </span>
    <template v-else>
      <div
          v-for="topic in topics"
          :key="topic.id"
      >
        <div class="topic">
          {{ topic.message }} ({{ topic.voteCount }})
          <button
              @click="vote(topic.id)"
              class="voteButton"
          >
            Vote
          </button>
        </div>
      </div>
    </template>
  </div>
</template>

<script lang="ts">
  import Vue from 'vue'
  import Component from "vue-class-component"
  import {Action, Getter} from "vuex-class";
  import {LocalTopic} from "@/store/vote.store";

  @Component
  export default class App extends Vue {
    @Action('voteStore/fetchTopics') fetchTopics!: () => Promise<void>
    @Action('voteStore/voteTopic') voteTopic!: (id: number) => Promise<void>
    @Getter('voteStore/topics') topics!: LocalTopic[]

    loading = true
    created() {
      this.fetchTopics().then(() => {
        this.loading = false
      })
    }

    vote(topicId: number) {
      this.voteTopic(topicId)
    }

  }
</script>

<style lang="scss" scoped>
  .topic {
    margin: 10px;
  }
</style>
