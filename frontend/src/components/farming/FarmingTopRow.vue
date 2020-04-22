<template>
  <v-container>
    <v-row align-center>
      <v-col cols="2" style="padding-left: 30px">
        <router-link :to="{name:previosPage}">
          <v-btn dark fab color="white">
            <v-icon color="primary">mdi-arrow-left</v-icon>
          </v-btn>
        </router-link>
      </v-col>
      <v-col align="center" align-self="center" cols="8">
        <h3 style="color:#828282" v-text="headtext"></h3>
      </v-col>
      <v-col cols="2" style="text-align: right; padding-right: 30px">
        <router-link :to="{name:'Home'}">
          <v-btn dark fab color="white">
            <v-icon color="primary">mdi-close</v-icon>
          </v-btn>
        </router-link>
      </v-col>
    </v-row>
    <v-row justify="center">
      <img :src="getImgUrl()" />
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "FarmingTopRow",

  components: {},
  methods: {
    getImgUrl() {
      return require("@/assets/ProgressBar/ProgressBar" + this.step + ".svg");
    }
  },
  data() {
    return {};
  },
  props: {
    step: Number,
    headtext: String,
    previosPage: String
  }
};
</script>

<style scoped>
#logo-text {
  font-family: Montserrat, sans-serif;
  font-style: normal;
  font-weight: 600;
  font-size: 36px;
  color: var(--v-primary-base);
}
</style>